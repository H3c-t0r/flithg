package pod

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	pluginsCore "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	pluginsCoreMock "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/flytek8s"
	flytek8sConfig "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/flytek8s/config"
	pluginsIOMock "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/k8s"
)

var containerResourceRequirements = &v1.ResourceRequirements{
	Limits: v1.ResourceList{
		v1.ResourceCPU:     resource.MustParse("1024m"),
		v1.ResourceStorage: resource.MustParse("100M"),
	},
}

func dummyContainerTaskTemplate(command []string, args []string) *core.TaskTemplate {
	return &core.TaskTemplate{
		Type: "test",
		Target: &core.TaskTemplate_Container{
			Container: &core.Container{
				Command: command,
				Args:    args,
			},
		},
	}
}

func dummyContainerTaskMetadata(resources *v1.ResourceRequirements, extendedResources *core.ExtendedResources) pluginsCore.TaskExecutionMetadata {
	taskMetadata := &pluginsCoreMock.TaskExecutionMetadata{}
	taskMetadata.On("GetNamespace").Return("test-namespace")
	taskMetadata.On("GetAnnotations").Return(map[string]string{"annotation-1": "val1"})
	taskMetadata.On("GetLabels").Return(map[string]string{"label-1": "val1"})
	taskMetadata.On("GetOwnerReference").Return(metav1.OwnerReference{
		Kind: "node",
		Name: "blah",
	})
	taskMetadata.On("GetK8sServiceAccount").Return("service-account")
	taskMetadata.On("GetSecurityContext").Return(core.SecurityContext{
		RunAs: &core.Identity{K8SServiceAccount: "service-account"},
	})
	taskMetadata.On("GetOwnerID").Return(types.NamespacedName{
		Namespace: "test-namespace",
		Name:      "test-owner-name",
	})
	taskMetadata.OnGetPlatformResources().Return(&v1.ResourceRequirements{})

	tID := &pluginsCoreMock.TaskExecutionID{}
	tID.On("GetID").Return(core.TaskExecutionIdentifier{
		NodeExecutionId: &core.NodeExecutionIdentifier{
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Name:    "my_name",
				Project: "my_project",
				Domain:  "my_domain",
			},
		},
	})
	tID.On("GetGeneratedName").Return("my_project:my_domain:my_name")
	taskMetadata.On("GetTaskExecutionID").Return(tID)

	to := &pluginsCoreMock.TaskOverrides{}
	to.On("GetResources").Return(resources)
	to.On("GetExtendedResources").Return(extendedResources)
	taskMetadata.On("GetOverrides").Return(to)
	taskMetadata.On("IsInterruptible").Return(true)
	taskMetadata.On("GetEnvironmentVariables").Return(nil)
	return taskMetadata
}

func dummyContainerTaskContext(taskTemplate *core.TaskTemplate, resources *v1.ResourceRequirements, extendedResources *core.ExtendedResources) pluginsCore.TaskExecutionContext {
	dummyTaskMetadata := dummyContainerTaskMetadata(resources, extendedResources)
	taskCtx := &pluginsCoreMock.TaskExecutionContext{}
	inputReader := &pluginsIOMock.InputReader{}
	inputReader.OnGetInputPrefixPath().Return("test-data-reference")
	inputReader.OnGetInputPath().Return("test-data-reference")
	inputReader.OnGetMatch(mock.Anything).Return(&core.LiteralMap{}, nil)
	taskCtx.OnInputReader().Return(inputReader)

	outputReader := &pluginsIOMock.OutputWriter{}
	outputReader.OnGetOutputPath().Return("/data/outputs.pb")
	outputReader.OnGetOutputPrefixPath().Return("/data/")
	outputReader.OnGetRawOutputPrefix().Return("")
	outputReader.OnGetCheckpointPrefix().Return("/checkpoint")
	outputReader.OnGetPreviousCheckpointsPrefix().Return("/prev")

	taskCtx.OnOutputWriter().Return(outputReader)

	taskReader := &pluginsCoreMock.TaskReader{}
	taskReader.OnReadMatch(mock.Anything).Return(taskTemplate, nil)
	taskCtx.OnTaskReader().Return(taskReader)

	taskCtx.OnTaskExecutionMetadata().Return(dummyTaskMetadata)

	pluginStateReader := &pluginsCoreMock.PluginStateReader{}
	pluginStateReader.OnGetMatch(mock.Anything).Return(0, nil)
	taskCtx.OnPluginStateReader().Return(pluginStateReader)

	return taskCtx
}

func TestContainerTaskExecutor_BuildIdentityResource(t *testing.T) {
	taskMetadata := &pluginsCoreMock.TaskExecutionMetadata{}
	r, err := DefaultPodPlugin.BuildIdentityResource(context.TODO(), taskMetadata)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	_, ok := r.(*v1.Pod)
	assert.True(t, ok)
	assert.Equal(t, flytek8s.PodKind, r.GetObjectKind().GroupVersionKind().Kind)
}

func TestContainerTaskExecutor_BuildResource(t *testing.T) {
	command := []string{"command"}
	args := []string{"{{.Input}}"}
	taskTemplate := dummyContainerTaskTemplate(command, args)
	taskCtx := dummyContainerTaskContext(taskTemplate, containerResourceRequirements, nil)

	r, err := DefaultPodPlugin.BuildResource(context.TODO(), taskCtx)
	assert.NoError(t, err)
	assert.NotNil(t, r)
	j, ok := r.(*v1.Pod)
	assert.True(t, ok)

	assert.NotEmpty(t, j.Spec.Containers)
	assert.Equal(t, containerResourceRequirements.Limits[v1.ResourceCPU], j.Spec.Containers[0].Resources.Limits[v1.ResourceCPU])

	// TODO: Once configurable, test when setting storage is supported on the cluster vs not.
	storageRes := j.Spec.Containers[0].Resources.Limits[v1.ResourceStorage]
	assert.Equal(t, int64(0), (&storageRes).Value())

	assert.Equal(t, command, j.Spec.Containers[0].Command)
	assert.Equal(t, []string{"test-data-reference"}, j.Spec.Containers[0].Args)

	assert.Equal(t, "service-account", j.Spec.ServiceAccountName)
}

func TestContainerTaskExecutor_BuildResource_ExtendedResources(t *testing.T) {
	assert.NoError(t, flytek8sConfig.SetK8sPluginConfig(&flytek8sConfig.K8sPluginConfig{
		GpuDeviceNodeLabel:        "gpu-node-label",
		GpuPartitionSizeNodeLabel: "gpu-partition-size",
		GpuResourceName:           flytek8s.ResourceNvidiaGPU,
	}))

	fixtures := []struct {
		name                      string
		resources                 *v1.ResourceRequirements
		extendedResourcesBase     *core.ExtendedResources
		extendedResourcesOverride *core.ExtendedResources
		expectedNsr               []v1.NodeSelectorTerm
		expectedTol               []v1.Toleration
	}{
		{
			"without overrides",
			&v1.ResourceRequirements{
				Limits: v1.ResourceList{
					flytek8s.ResourceNvidiaGPU: resource.MustParse("1"),
				},
			},
			&core.ExtendedResources{
				GpuAccelerator: &core.GPUAccelerator{
					Device: "nvidia-tesla-t4",
				},
			},
			nil,
			[]v1.NodeSelectorTerm{
				{
					MatchExpressions: []v1.NodeSelectorRequirement{
						v1.NodeSelectorRequirement{
							Key:      "gpu-node-label",
							Operator: v1.NodeSelectorOpIn,
							Values:   []string{"nvidia-tesla-t4"},
						},
					},
				},
			},
			[]v1.Toleration{
				{
					Key:      "gpu-node-label",
					Value:    "nvidia-tesla-t4",
					Operator: v1.TolerationOpEqual,
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
		},
		{
			"with overrides",
			&v1.ResourceRequirements{
				Limits: v1.ResourceList{
					flytek8s.ResourceNvidiaGPU: resource.MustParse("1"),
				},
			},
			&core.ExtendedResources{
				GpuAccelerator: &core.GPUAccelerator{
					Device: "nvidia-tesla-t4",
				},
			},
			&core.ExtendedResources{
				GpuAccelerator: &core.GPUAccelerator{
					Device: "nvidia-tesla-a100",
					PartitionSizeValue: &core.GPUAccelerator_PartitionSize{
						PartitionSize: "1g.5gb",
					},
				},
			},
			[]v1.NodeSelectorTerm{
				{
					MatchExpressions: []v1.NodeSelectorRequirement{
						v1.NodeSelectorRequirement{
							Key:      "gpu-node-label",
							Operator: v1.NodeSelectorOpIn,
							Values:   []string{"nvidia-tesla-a100"},
						},
						v1.NodeSelectorRequirement{
							Key:      "gpu-partition-size",
							Operator: v1.NodeSelectorOpIn,
							Values:   []string{"1g.5gb"},
						},
					},
				},
			},
			[]v1.Toleration{
				{
					Key:      "gpu-node-label",
					Value:    "nvidia-tesla-a100",
					Operator: v1.TolerationOpEqual,
					Effect:   v1.TaintEffectNoSchedule,
				},
				{
					Key:      "gpu-partition-size",
					Value:    "1g.5gb",
					Operator: v1.TolerationOpEqual,
					Effect:   v1.TaintEffectNoSchedule,
				},
			},
		},
	}

	for _, f := range fixtures {
		t.Run(f.name, func(t *testing.T) {
			taskTemplate := dummyContainerTaskTemplate([]string{"command"}, []string{"{{.Input}}"})
			taskTemplate.ExtendedResources = f.extendedResourcesBase
			taskContext := dummyContainerTaskContext(taskTemplate, f.resources, f.extendedResourcesOverride)
			r, err := DefaultPodPlugin.BuildResource(context.TODO(), taskContext)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			pod, ok := r.(*v1.Pod)
			assert.True(t, ok)

			assert.EqualValues(
				t,
				f.expectedNsr,
				pod.Spec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms,
			)
			assert.EqualValues(
				t,
				f.expectedTol,
				pod.Spec.Tolerations,
			)
		})
	}
}

func TestContainerTaskExecutor_GetTaskStatus(t *testing.T) {
	command := []string{"command"}
	args := []string{"{{.Input}}"}
	taskTemplate := dummyContainerTaskTemplate(command, args)
	taskCtx := dummyContainerTaskContext(taskTemplate, containerResourceRequirements, nil)

	j := &v1.Pod{
		Status: v1.PodStatus{},
	}

	ctx := context.TODO()
	t.Run("running", func(t *testing.T) {
		j.Status.Phase = v1.PodRunning
		j.Status.ContainerStatuses = []v1.ContainerStatus{
			{
				State: v1.ContainerState{
					Running: &v1.ContainerStateRunning{},
				},
			},
		}

		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseRunning, phaseInfo.Phase())
	})

	t.Run("queued", func(t *testing.T) {
		j.Status.Phase = v1.PodPending
		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseQueued, phaseInfo.Phase())
	})

	t.Run("failNoCondition", func(t *testing.T) {
		j.Status.Phase = v1.PodFailed
		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseRetryableFailure, phaseInfo.Phase())
		ec := phaseInfo.Err().GetCode()
		assert.Equal(t, "UnknownError", ec)
	})

	t.Run("failConditionUnschedulable", func(t *testing.T) {
		j.Status.Phase = v1.PodFailed
		j.Status.Reason = "Unschedulable"
		j.Status.Message = "some message"
		j.Status.Conditions = []v1.PodCondition{
			{
				Type: v1.PodReasonUnschedulable,
			},
		}
		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseRetryableFailure, phaseInfo.Phase())
		ec := phaseInfo.Err().GetCode()
		assert.Equal(t, "Unschedulable", ec)
	})

	t.Run("successOptimized", func(t *testing.T) {
		j.Status.Phase = v1.PodRunning
		j.Status.ContainerStatuses = []v1.ContainerStatus{
			{
				State: v1.ContainerState{
					Terminated: &v1.ContainerStateTerminated{
						ExitCode: 0,
					},
				},
			},
		}

		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhaseSuccess, phaseInfo.Phase())
	})

	t.Run("success", func(t *testing.T) {
		j.Status.Phase = v1.PodSucceeded
		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, j)
		assert.NoError(t, err)
		assert.NotNil(t, phaseInfo)
		assert.Equal(t, pluginsCore.PhaseSuccess, phaseInfo.Phase())
	})
}

func TestContainerTaskExecutor_GetProperties(t *testing.T) {
	expected := k8s.PluginProperties{}
	assert.Equal(t, expected, DefaultPodPlugin.GetProperties())
}

func TestContainerTaskExecutor_GetTaskStatus_InvalidImageName(t *testing.T) {
	command := []string{"command"}
	args := []string{"{{.Input}}"}
	taskTemplate := dummyContainerTaskTemplate(command, args)
	taskCtx := dummyContainerTaskContext(taskTemplate, containerResourceRequirements, nil)

	ctx := context.TODO()
	reason := "InvalidImageName"
	message := "Failed to apply default image tag \"TEST/flyteorg/myapp:latest\": couldn't parse image reference" +
		" \"TEST/flyteorg/myapp:latest\": invalid reference format: repository name must be lowercase"

	pendingPod := &v1.Pod{
		Status: v1.PodStatus{
			Phase: v1.PodPending,
			Conditions: []v1.PodCondition{
				{
					Type:   v1.PodReady,
					Status: v1.ConditionFalse,
				},
			},
			ContainerStatuses: []v1.ContainerStatus{
				{
					ContainerID: "ContainerID",
					Ready:       false,
					State: v1.ContainerState{
						Waiting: &v1.ContainerStateWaiting{
							Reason:  reason,
							Message: message,
						},
					},
				},
			},
		},
	}

	t.Run("failInvalidImageName", func(t *testing.T) {
		pendingPod.Status.Phase = v1.PodPending
		phaseInfo, err := DefaultPodPlugin.GetTaskPhase(ctx, taskCtx, pendingPod)
		finalReason := fmt.Sprintf("|%s", reason)
		finalMessage := fmt.Sprintf("|%s", message)
		assert.NoError(t, err)
		assert.Equal(t, pluginsCore.PhasePermanentFailure, phaseInfo.Phase())
		assert.Equal(t, &core.ExecutionError{Code: finalReason, Message: finalMessage, Kind: core.ExecutionError_USER}, phaseInfo.Err())
	})
}
