package logs

import (
	"context"
	"testing"

	"github.com/go-test/deep"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	pluginCore "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	coreMocks "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/tasklog"
)

const podName = "PodName"

func dummyTaskExecID() pluginCore.TaskExecutionID {
	tID := &coreMocks.TaskExecutionID{}
	tID.OnGetGeneratedName().Return("generated-name")
	tID.OnGetID().Return(core.TaskExecutionIdentifier{
		TaskId: &core.Identifier{
			ResourceType: core.ResourceType_TASK,
			Name:         "my-task-name",
			Project:      "my-task-project",
			Domain:       "my-task-domain",
			Version:      "1",
		},
		NodeExecutionId: &core.NodeExecutionIdentifier{
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Name:    "my-execution-name",
				Project: "my-execution-project",
				Domain:  "my-execution-domain",
			},
		},
		RetryAttempt: 1,
	})
	tID.OnGetUniqueNodeID().Return("n0-0-n0")
	return tID
}

func TestGetLogsForContainerInPod_NoPlugins(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{})
	assert.NoError(t, err)
	l, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), nil, 0, " Suffix", nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, l)
}

func TestGetLogsForContainerInPod_NoLogs(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsCloudwatchEnabled: true,
		CloudwatchRegion:    "us-east-1",
		CloudwatchLogGroup:  "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)
	p, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), nil, 0, " Suffix", nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, p)
}

func TestGetLogsForContainerInPod_BadIndex(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsCloudwatchEnabled: true,
		CloudwatchRegion:    "us-east-1",
		CloudwatchLogGroup:  "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	p, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 1, " Suffix", nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, p)
}

func TestGetLogsForContainerInPod_MissingStatus(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsCloudwatchEnabled: true,
		CloudwatchRegion:    "us-east-1",
		CloudwatchLogGroup:  "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{},
	}
	pod.Name = podName

	p, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 1, " Suffix", nil, nil)
	assert.NoError(t, err)
	assert.Nil(t, p)
}

func TestGetLogsForContainerInPod_Cloudwatch(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{IsCloudwatchEnabled: true,
		CloudwatchRegion:   "us-east-1",
		CloudwatchLogGroup: "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " Suffix", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, logs, 1)
}

func TestGetLogsForContainerInPod_K8s(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsKubernetesEnabled: true,
		KubernetesURL:       "k8s.com",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " Suffix", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, logs, 1)
}

func TestGetLogsForContainerInPod_All(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsKubernetesEnabled: true,
		KubernetesURL:       "k8s.com",
		IsCloudwatchEnabled: true,
		CloudwatchRegion:    "us-east-1",
		CloudwatchLogGroup:  "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " Suffix", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, logs, 2)
}

func TestGetLogsForContainerInPod_HostName(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsKubernetesEnabled: true,
		KubernetesURL:       "k8s.com",
		IsCloudwatchEnabled: true,
		CloudwatchRegion:    "us-east-1",
		CloudwatchLogGroup:  "/kubernetes/flyte-production",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
			Hostname: "my-hostname",
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " Suffix", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, logs, 2)
}

func TestGetLogsForContainerInPod_Stackdriver(t *testing.T) {
	logPlugin, err := InitializeLogPlugins(&LogConfig{
		IsStackDriverEnabled:       true,
		GCPProjectName:             "myGCPProject",
		StackdriverLogResourceName: "aws_ec2_instance",
	})
	assert.NoError(t, err)

	pod := &v1.Pod{
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}
	pod.Name = podName

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " Suffix", nil, nil)
	assert.Nil(t, err)
	assert.Len(t, logs, 1)
}

func TestGetLogsForContainerInPod_LegacyTemplate(t *testing.T) {
	t.Run("All Templates available", func(t *testing.T) {
		assertTestSucceeded(t, &LogConfig{
			IsKubernetesEnabled:   true,
			KubernetesTemplateURI: "https://k8s-my-log-server/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",

			IsCloudwatchEnabled:   true,
			CloudwatchTemplateURI: "https://cw-my-log-server/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",

			IsStackDriverEnabled:   true,
			StackDriverTemplateURI: "https://sd-my-log-server/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
		}, nil, []*core.TaskLog{
			{
				Uri:           "https://k8s-my-log-server/my-namespace/my-pod/ContainerName/ContainerID",
				MessageFormat: core.TaskLog_JSON,
				Name:          "Kubernetes Logs my-Suffix",
			},
			{
				Uri:           "https://cw-my-log-server/my-namespace/my-pod/ContainerName/ContainerID",
				MessageFormat: core.TaskLog_JSON,
				Name:          "Cloudwatch Logs my-Suffix",
			},
			{
				Uri:           "https://sd-my-log-server/my-namespace/my-pod/ContainerName/ContainerID",
				MessageFormat: core.TaskLog_JSON,
				Name:          "Stackdriver Logs my-Suffix",
			},
		}, "")
	})

	t.Run("StackDriver", func(t *testing.T) {
		assertTestSucceeded(t, &LogConfig{
			IsStackDriverEnabled:   true,
			StackDriverTemplateURI: "https://sd-my-log-server/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
		}, nil, []*core.TaskLog{
			{
				Uri:           "https://sd-my-log-server/my-namespace/my-pod/ContainerName/ContainerID",
				MessageFormat: core.TaskLog_JSON,
				Name:          "Stackdriver Logs my-Suffix",
			},
		}, "")
	})
}

func assertTestSucceeded(tb testing.TB, config *LogConfig, taskTemplate *core.TaskTemplate, expectedTaskLogs []*core.TaskLog, hostname string) {
	logPlugin, err := InitializeLogPlugins(config)
	assert.NoError(tb, err)

	pod := &v1.Pod{
		ObjectMeta: v12.ObjectMeta{
			Namespace: "my-namespace",
			Name:      "my-pod",
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{
					Name: "ContainerName",
				},
			},
			Hostname: hostname,
		},
		Status: v1.PodStatus{
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
				},
			},
		},
	}

	logs, err := GetLogsForContainerInPod(context.TODO(), logPlugin, dummyTaskExecID(), pod, 0, " my-Suffix", nil, taskTemplate)
	assert.Nil(tb, err)
	assert.Len(tb, logs, len(expectedTaskLogs))
	if diff := deep.Equal(logs, expectedTaskLogs); len(diff) > 0 {
		assert.FailNowf(tb, "Not Equal.", "Diff: %v", diff)
	}
}

func TestGetLogsForContainerInPod_Templates(t *testing.T) {
	assertTestSucceeded(t, &LogConfig{
		Templates: []tasklog.TemplateLogPlugin{
			{
				DisplayName: "StackDriver",
				TemplateURIs: []string{
					"https://my-log-server/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
				},
				MessageFormat: core.TaskLog_JSON,
			},
			{
				DisplayName: "Internal",
				TemplateURIs: []string{
					"https://flyte.corp.net/console/projects/{{ .executionProject }}/domains/{{ .executionDomain }}/executions/{{ .executionName }}/nodeId/{{ .nodeID }}/taskId/{{ .taskID }}/attempt/{{ .taskRetryAttempt }}/view/logs",
				},
				MessageFormat: core.TaskLog_JSON,
			},
		},
	}, nil, []*core.TaskLog{
		{
			Uri:           "https://my-log-server/my-namespace/my-pod/ContainerName/ContainerID",
			MessageFormat: core.TaskLog_JSON,
			Name:          "StackDriver my-Suffix",
		},
		{
			Uri:           "https://flyte.corp.net/console/projects/my-execution-project/domains/my-execution-domain/executions/my-execution-name/nodeId/n0-0-n0/taskId/my-task-name/attempt/1/view/logs",
			MessageFormat: core.TaskLog_JSON,
			Name:          "Internal my-Suffix",
		},
	}, "")
}

func TestGetLogsForContainerInPodTemplates_Hostname(t *testing.T) {
	assertTestSucceeded(t, &LogConfig{
		Templates: []tasklog.TemplateLogPlugin{
			{
				DisplayName: "StackDriver",
				TemplateURIs: []string{
					"{{ .hostname }}/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
				},
				MessageFormat: core.TaskLog_JSON,
			},
		},
	}, nil, []*core.TaskLog{
		{
			Uri:           "my-hostname/my-namespace/my-pod/ContainerName/ContainerID",
			MessageFormat: core.TaskLog_JSON,
			Name:          "StackDriver my-Suffix",
		},
	}, "my-hostname")
}

func TestGetLogsForContainerInPod_Flyteinteractive(t *testing.T) {
	tests := []struct {
		name             string
		config           *LogConfig
		template         *core.TaskTemplate
		expectedTaskLogs []*core.TaskLog
	}{
		{
			"Flyteinteractive enabled but no task template",
			&LogConfig{
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://flyteinteractive.mydomain.com:{{ .taskConfig.port }}/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
						},
					},
				},
			},
			nil,
			nil,
		},
		{
			"Flyteinteractive enabled but config not found in task template",
			&LogConfig{
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://flyteinteractive.mydomain.com:{{ .taskConfig.port }}/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
						},
					},
				},
			},
			&core.TaskTemplate{},
			nil,
		},
		{
			"Flyteinteractive disabled but config present in TaskTemplate",
			&LogConfig{},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type": "vscode",
					"port":      "65535",
				},
			},
			nil,
		},
		{
			"Flyteinteractive - multiple dynamic options",
			&LogConfig{
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://abc.com:{{ .taskConfig.port }}/{{ .taskConfig.route }}",
						},
					},
				},
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type": "vscode",
					"port":      "65535",
					"route":     "a-route",
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://abc.com:65535/a-route",
					MessageFormat: core.TaskLog_JSON,
					Name:          "vscode link my-Suffix",
				},
			},
		},
		{
			"Flyteinteractive - multiple uses of the template (invalid use of ports in a URI)",
			&LogConfig{
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://abc.com:{{ .taskConfig.port }}:{{ .taskConfig.port}}",
						},
					},
				},
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type": "vscode",
					"port":      "65535",
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://abc.com:65535:65535",
					MessageFormat: core.TaskLog_JSON,
					Name:          "vscode link my-Suffix",
				},
			},
		},
		{
			"Flyteinteractive disabled and K8s enabled and flyteinteractive config present in TaskTemplate",
			&LogConfig{
				IsKubernetesEnabled:   true,
				KubernetesTemplateURI: "https://k8s.com/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type": "vscode",
					"port":      "65535",
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://k8s.com/my-namespace/my-pod/ContainerName/ContainerID",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Kubernetes Logs my-Suffix",
				},
			},
		},
		{
			"Flyteinteractive and K8s enabled",
			&LogConfig{
				IsKubernetesEnabled:   true,
				KubernetesTemplateURI: "https://k8s.com/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://flyteinteractive.mydomain.com:{{ .taskConfig.port }}/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
						},
					},
				},
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type": "vscode",
					"port":      "65535",
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://k8s.com/my-namespace/my-pod/ContainerName/ContainerID",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Kubernetes Logs my-Suffix",
				},
				{
					Uri:           "https://flyteinteractive.mydomain.com:65535/my-namespace/my-pod/ContainerName/ContainerID",
					MessageFormat: core.TaskLog_JSON,
					Name:          "vscode link my-Suffix",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTestSucceeded(t, tt.config, tt.template, tt.expectedTaskLogs, "")
		})
	}
}

func TestGetLogsForContainerInPod_GenericDynamicLogLinks(t *testing.T) {
	tests := []struct {
		name             string
		config           *LogConfig
		template         *core.TaskTemplate
		expectedTaskLogs []*core.TaskLog
	}{
		{
			"Generic Dynamic Log Links enabled but no task template",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
			},
			nil,
			nil,
		},
		{
			"Generic Dynamic Log Links enabled and empty task template",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
			},
			&core.TaskTemplate{},
			nil,
		},
		{
			"Generic Dynamic Log Links enabled and simple config present in TaskTemplate",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "vscode",
					"port":                      "65535",
					"generic_dynamic_log_links": `{"links": [{"name": "vscode", "display_name": "vscode", "template_uri": "vscode://def.com:{{ .taskConfig.port }}/{{ .podName }}"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "vscode://def.com:65535/my-pod",
					MessageFormat: core.TaskLog_JSON,
					Name:          "vscode",
				},
			},
		},
		{
			"Generic Dynamic Log Links disabled and simple config present in TaskTemplate - Notice how we can mix and match parameters defined in the task template, i.e. using {{ .taskConfig.port }} also works",
			&LogConfig{},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "vscode",
					"port":                      "65535",
					"generic_dynamic_log_links": `{"links": [{"name": "vscode", "display_name": "vscode", "template_uri": "vscode://def.com:{{ .taskConfig.port }}/{{ .podName }}"}]}`,
				},
			},
			nil,
		},
		{
			"Generic Dynamic Log Links - multiple dynamic options",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"vscode": tasklog.TemplateLogPlugin{
						DisplayName: "vscode link",
						TemplateURIs: []tasklog.TemplateURI{
							"https://abc.com:{{ .taskConfig.port }}/{{ .taskConfig.route }}",
						},
					},
				},
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "vscode",
					"port":                      "65535",
					"route":                     "a-route",
					"generic_dynamic_log_links": `{"links": [{"name": "vscode", "display_name": "Other logs", "template_uri": "coming://from/other/place"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://abc.com:65535/a-route",
					MessageFormat: core.TaskLog_JSON,
					Name:          "vscode link my-Suffix",
				},
				{
					Uri:           "coming://from/other/place",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Other logs",
				},
			},
		},
		{
			"Generic Dynamic Log Links disabled and K8s enabled and generic dynamic log links config present in TaskTemplate",
			&LogConfig{
				IsKubernetesEnabled:           true,
				KubernetesTemplateURI:         "https://k8s.com/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
				GenericDynamicLogLinksEnabled: false,
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "abc",
					"generic_dynamic_log_links": `{"links": [{"display_name": "Other logs", "template_uri": "coming://from/other/place"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://k8s.com/my-namespace/my-pod/ContainerName/ContainerID",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Kubernetes Logs my-Suffix",
				},
			},
		},
		{
			"Generic Dynamic Log Links enabled and K8s enabled and generic dynamic log links config present in TaskTemplate",
			&LogConfig{
				IsKubernetesEnabled:           true,
				KubernetesTemplateURI:         "https://k8s.com/{{ .namespace }}/{{ .podName }}/{{ .containerName }}/{{ .containerId }}",
				GenericDynamicLogLinksEnabled: true,
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "abc",
					"generic_dynamic_log_links": `{"links": [{"name": "abc", "display_name": "Other logs", "template_uri": "coming://from/other/place"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "https://k8s.com/my-namespace/my-pod/ContainerName/ContainerID",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Kubernetes Logs my-Suffix",
				},
				{
					Uri:           "coming://from/other/place",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Other logs",
				},
			},
		},
		{
			"Multiple Generic Dynamic Log Links defined in task template",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
			},
			&core.TaskTemplate{
				Config: map[string]string{
					"link_type":                 "abc,defg",
					"generic_dynamic_log_links": `{"links": [{"name": "abc", "display_name": "Other logs", "template_uri": "coming://from/abc"},{"name": "defg", "display_name": "defg logs", "template_uri": "coming://from/defg"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "coming://from/abc",
					MessageFormat: core.TaskLog_JSON,
					Name:          "Other logs",
				},
				{
					Uri:           "coming://from/defg",
					MessageFormat: core.TaskLog_JSON,
					Name:          "defg logs",
				},
			},
		},
		{
			"Multiple Generic Dynamic Log Links defined in task template - one defined as a Generic Dynamic Log Link and a regular Dynamic Log Link",
			&LogConfig{
				GenericDynamicLogLinksEnabled: true,
				DynamicLogLinks: map[string]tasklog.TemplateLogPlugin{
					"abc": tasklog.TemplateLogPlugin{
						DisplayName: "abc link",
						TemplateURIs: []tasklog.TemplateURI{
							"coming://from/abc",
						},
					},
				},
			},
			&core.TaskTemplate{
				Config: map[string]string{
					// Notice that only defg is defined as a generic dynamic log link
					"link_type":                 "abc,defg",
					"generic_dynamic_log_links": `{"links": [{"name": "defg", "display_name": "defg logs", "template_uri": "coming://from/defg"}]}`,
				},
			},
			[]*core.TaskLog{
				{
					Uri:           "coming://from/abc",
					MessageFormat: core.TaskLog_JSON,
					Name:          "abc link my-Suffix",
				},
				{
					Uri:           "coming://from/defg",
					MessageFormat: core.TaskLog_JSON,
					Name:          "defg logs",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assertTestSucceeded(t, tt.config, tt.template, tt.expectedTaskLogs, "")
		})
	}
}
