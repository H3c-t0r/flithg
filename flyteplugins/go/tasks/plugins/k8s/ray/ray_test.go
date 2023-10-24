package ray

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/jsonpb"
	structpb "github.com/golang/protobuf/ptypes/struct"
	rayv1alpha1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/plugins"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/logs"
	pluginsCore "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/core/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/flytek8s"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/flytek8s/config"
	pluginIOMocks "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/io/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/k8s"
	mocks2 "github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/k8s/mocks"
	"github.com/flyteorg/flyte/flyteplugins/go/tasks/pluginmachinery/utils"
)

const testImage = "image://"
const serviceAccount = "ray_sa"
const serviceAccountOverride = "ray_sa_override"
const namespaceOverride = "ray_namespace_override"

var (
	dummyEnvVars = []*core.KeyValuePair{
		{Key: "Env_Var", Value: "Env_Val"},
	}

	testArgs = []string{
		"test-args",
	}

	headResourceOverride = core.Resources{
		Requests: []*core.Resources_ResourceEntry{
			{
				Name:  core.Resources_CPU,
				Value: "1000m",
			},
			{
				Name:  core.Resources_MEMORY,
				Value: "2Gi",
			},
		},
		Limits: []*core.Resources_ResourceEntry{
			{
				Name:  core.Resources_CPU,
				Value: "2000m",
			},
			{
				Name:  core.Resources_MEMORY,
				Value: "4Gi",
			},
		},
	}

	workerResourceOverride = core.Resources{
		Requests: []*core.Resources_ResourceEntry{
			{
				Name:  core.Resources_CPU,
				Value: "5",
			},
			{
				Name:  core.Resources_MEMORY,
				Value: "10G",
			},
		},
		Limits: []*core.Resources_ResourceEntry{
			{
				Name:  core.Resources_CPU,
				Value: "10",
			},
			{
				Name:  core.Resources_MEMORY,
				Value: "20G",
			},
		},
	}

	resourceRequirements = &corev1.ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:         resource.MustParse("1000m"),
			corev1.ResourceMemory:      resource.MustParse("1Gi"),
			flytek8s.ResourceNvidiaGPU: resource.MustParse("1"),
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:         resource.MustParse("100m"),
			corev1.ResourceMemory:      resource.MustParse("512Mi"),
			flytek8s.ResourceNvidiaGPU: resource.MustParse("1"),
		},
	}

	workerGroupName = "worker-group"
)

func dummyRayCustomObj() *plugins.RayJob {
	return &plugins.RayJob{
		RayCluster: &plugins.RayCluster{
			HeadGroupSpec:   &plugins.HeadGroupSpec{RayStartParams: map[string]string{"num-cpus": "1"}},
			WorkerGroupSpec: []*plugins.WorkerGroupSpec{{GroupName: workerGroupName, Replicas: 3}},
		},
	}
}

func dummyRayCustomObjWithOverrides() *plugins.RayJob {
	return &plugins.RayJob{
		RayCluster: &plugins.RayCluster{
			K8SServiceAccount: serviceAccountOverride,
			Namespace: namespaceOverride,
			HeadGroupSpec:   &plugins.HeadGroupSpec{RayStartParams: map[string]string{"num-cpus": "1"}, Resources: &headResourceOverride},
			WorkerGroupSpec: []*plugins.WorkerGroupSpec{{GroupName: workerGroupName, Replicas: 3, Resources: &workerResourceOverride}},
		},
	}
}

func dummyRayTaskTemplate(id string, rayJobObj *plugins.RayJob) *core.TaskTemplate {
	ptObjJSON, err := utils.MarshalToString(rayJobObj)
	if err != nil {
		panic(err)
	}

	structObj := structpb.Struct{}

	err = jsonpb.UnmarshalString(ptObjJSON, &structObj)
	if err != nil {
		panic(err)
	}

	return &core.TaskTemplate{
		Id:   &core.Identifier{Name: id},
		Type: "container",
		Target: &core.TaskTemplate_Container{
			Container: &core.Container{
				Image: testImage,
				Args:  testArgs,
				Env:   dummyEnvVars,
			},
		},
		Custom: &structObj,
	}
}

func dummyRayTaskContext(taskTemplate *core.TaskTemplate, resources *corev1.ResourceRequirements, extendedResources *core.ExtendedResources) pluginsCore.TaskExecutionContext {
	taskCtx := &mocks.TaskExecutionContext{}
	inputReader := &pluginIOMocks.InputReader{}
	inputReader.OnGetInputPrefixPath().Return("/input/prefix")
	inputReader.OnGetInputPath().Return("/input")
	inputReader.OnGetMatch(mock.Anything).Return(&core.LiteralMap{}, nil)
	taskCtx.OnInputReader().Return(inputReader)

	outputReader := &pluginIOMocks.OutputWriter{}
	outputReader.OnGetOutputPath().Return("/data/outputs.pb")
	outputReader.OnGetOutputPrefixPath().Return("/data/")
	outputReader.OnGetRawOutputPrefix().Return("")
	outputReader.OnGetCheckpointPrefix().Return("/checkpoint")
	outputReader.OnGetPreviousCheckpointsPrefix().Return("/prev")
	taskCtx.OnOutputWriter().Return(outputReader)

	taskReader := &mocks.TaskReader{}
	taskReader.OnReadMatch(mock.Anything).Return(taskTemplate, nil)
	taskCtx.OnTaskReader().Return(taskReader)

	tID := &mocks.TaskExecutionID{}
	tID.OnGetID().Return(core.TaskExecutionIdentifier{
		NodeExecutionId: &core.NodeExecutionIdentifier{
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Name:    "my_name",
				Project: "my_project",
				Domain:  "my_domain",
			},
		},
	})
	tID.OnGetGeneratedName().Return("some-acceptable-name")

	overrides := &mocks.TaskOverrides{}
	overrides.OnGetResources().Return(resources)
	overrides.OnGetExtendedResources().Return(extendedResources)

	taskExecutionMetadata := &mocks.TaskExecutionMetadata{}
	taskExecutionMetadata.OnGetTaskExecutionID().Return(tID)
	taskExecutionMetadata.OnGetNamespace().Return("test-namespace")
	taskExecutionMetadata.OnGetAnnotations().Return(map[string]string{"annotation-1": "val1"})
	taskExecutionMetadata.OnGetLabels().Return(map[string]string{"label-1": "val1"})
	taskExecutionMetadata.OnGetOwnerReference().Return(v1.OwnerReference{
		Kind: "node",
		Name: "blah",
	})
	taskExecutionMetadata.OnIsInterruptible().Return(true)
	taskExecutionMetadata.OnGetOverrides().Return(overrides)
	taskExecutionMetadata.OnGetK8sServiceAccount().Return(serviceAccount)
	taskExecutionMetadata.OnGetPlatformResources().Return(&corev1.ResourceRequirements{})
	taskExecutionMetadata.OnGetSecurityContext().Return(core.SecurityContext{
		RunAs: &core.Identity{K8SServiceAccount: serviceAccount},
	})
	taskExecutionMetadata.OnGetEnvironmentVariables().Return(nil)
	taskCtx.OnTaskExecutionMetadata().Return(taskExecutionMetadata)
	return taskCtx
}

func TestBuildResourceRay(t *testing.T) {
	rayJobResourceHandler := rayJobResourceHandler{}
	taskTemplate := dummyRayTaskTemplate("ray-id", dummyRayCustomObj())
	toleration := []corev1.Toleration{{
		Key:      "storage",
		Value:    "dedicated",
		Operator: corev1.TolerationOpExists,
		Effect:   corev1.TaintEffectNoSchedule,
	}}
	err := config.SetK8sPluginConfig(&config.K8sPluginConfig{DefaultTolerations: toleration})
	assert.Nil(t, err)

	RayResource, err := rayJobResourceHandler.BuildResource(context.TODO(), dummyRayTaskContext(taskTemplate, resourceRequirements, nil))
	assert.Nil(t, err)

	assert.NotNil(t, RayResource)
	ray, ok := RayResource.(*rayv1alpha1.RayJob)
	assert.True(t, ok)

	headReplica := int32(1)
	assert.Equal(t, &headReplica, ray.Spec.RayClusterSpec.HeadGroupSpec.Replicas)
	assert.Equal(t, serviceAccount, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.ServiceAccountName)
	assert.Equal(t, map[string]string{"dashboard-host": "0.0.0.0", "include-dashboard": "true", "node-ip-address": "$MY_POD_IP", "num-cpus": "1"},
		ray.Spec.RayClusterSpec.HeadGroupSpec.RayStartParams)
	assert.Equal(t, map[string]string{"annotation-1": "val1"}, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Annotations)
	assert.Equal(t, map[string]string{"label-1": "val1"}, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Labels)
	assert.Equal(t, toleration, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.Tolerations)

	workerReplica := int32(3)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Replicas)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MinReplicas)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MaxReplicas)
	assert.Equal(t, workerGroupName, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].GroupName)
	assert.Equal(t, serviceAccount, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.ServiceAccountName)
	assert.Equal(t, map[string]string{"node-ip-address": "$MY_POD_IP"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].RayStartParams)
	assert.Equal(t, map[string]string{"annotation-1": "val1"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Annotations)
	assert.Equal(t, map[string]string{"label-1": "val1"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Labels)
	assert.Equal(t, toleration, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.Tolerations)
}

func TestBuildResourceRayWithOverrides(t *testing.T) {
	rayJobResourceHandler := rayJobResourceHandler{}
	taskTemplate := dummyRayTaskTemplate("ray-id", dummyRayCustomObjWithOverrides())
	expectedHeadResources, _ := flytek8s.ToK8sResourceRequirements(&headResourceOverride)
	expectedWorkerResources, _ := flytek8s.ToK8sResourceRequirements(&workerResourceOverride)
	toleration := []corev1.Toleration{{
		Key:      "storage",
		Value:    "dedicated",
		Operator: corev1.TolerationOpExists,
		Effect:   corev1.TaintEffectNoSchedule,
	}}
	err := config.SetK8sPluginConfig(&config.K8sPluginConfig{DefaultTolerations: toleration})
	assert.Nil(t, err)

	RayResource, err := rayJobResourceHandler.BuildResource(context.TODO(), dummyRayTaskContext(taskTemplate))
	assert.Nil(t, err)

	assert.NotNil(t, RayResource)
	ray, ok := RayResource.(*rayv1alpha1.RayJob)
	assert.True(t, ok)

	headReplica := int32(1)
	assert.Equal(t, namespaceOverride, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.ObjectMeta.Namespace)
	assert.Equal(t, &headReplica, ray.Spec.RayClusterSpec.HeadGroupSpec.Replicas)
	assert.Equal(t, serviceAccountOverride, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.ServiceAccountName)
	assert.Equal(t, map[string]string{"dashboard-host": "0.0.0.0", "include-dashboard": "true", "node-ip-address": "$MY_POD_IP", "num-cpus": "1"},
		ray.Spec.RayClusterSpec.HeadGroupSpec.RayStartParams)
	assert.Equal(t, map[string]string{"annotation-1": "val1"}, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Annotations)
	assert.Equal(t, map[string]string{"label-1": "val1"}, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Labels)
	assert.Equal(t, toleration, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.Tolerations)
	assert.Equal(t, *expectedHeadResources, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.Containers[0].Resources)

	workerReplica := int32(3)
	assert.Equal(t, namespaceOverride, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.ObjectMeta.Namespace)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Replicas)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MinReplicas)
	assert.Equal(t, &workerReplica, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MaxReplicas)
	assert.Equal(t, workerGroupName, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].GroupName)
	assert.Equal(t, serviceAccountOverride, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.ServiceAccountName)
	assert.Equal(t, map[string]string{"node-ip-address": "$MY_POD_IP"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].RayStartParams)
	assert.Equal(t, map[string]string{"annotation-1": "val1"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Annotations)
	assert.Equal(t, map[string]string{"label-1": "val1"}, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Labels)
	assert.Equal(t, toleration, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.Tolerations)
	assert.Equal(t, *expectedWorkerResources, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.Containers[0].Resources)
}

func TestBuildResourceRayExtendedResources(t *testing.T) {
	assert.NoError(t, config.SetK8sPluginConfig(&config.K8sPluginConfig{
		GpuDeviceNodeLabel:        "gpu-node-label",
		GpuPartitionSizeNodeLabel: "gpu-partition-size",
		GpuResourceName:           flytek8s.ResourceNvidiaGPU,
	}))

	fixtures := []struct {
		name                      string
		resources                 *corev1.ResourceRequirements
		extendedResourcesBase     *core.ExtendedResources
		extendedResourcesOverride *core.ExtendedResources
		expectedNsr               []corev1.NodeSelectorTerm
		expectedTol               []corev1.Toleration
	}{
		{
			"without overrides",
			&corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
					flytek8s.ResourceNvidiaGPU: resource.MustParse("1"),
				},
			},
			&core.ExtendedResources{
				GpuAccelerator: &core.GPUAccelerator{
					Device: "nvidia-tesla-t4",
				},
			},
			nil,
			[]corev1.NodeSelectorTerm{
				{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						corev1.NodeSelectorRequirement{
							Key:      "gpu-node-label",
							Operator: corev1.NodeSelectorOpIn,
							Values:   []string{"nvidia-tesla-t4"},
						},
					},
				},
			},
			[]corev1.Toleration{
				{
					Key:      "gpu-node-label",
					Value:    "nvidia-tesla-t4",
					Operator: corev1.TolerationOpEqual,
					Effect:   corev1.TaintEffectNoSchedule,
				},
			},
		},
		{
			"with overrides",
			&corev1.ResourceRequirements{
				Limits: corev1.ResourceList{
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
			[]corev1.NodeSelectorTerm{
				{
					MatchExpressions: []corev1.NodeSelectorRequirement{
						corev1.NodeSelectorRequirement{
							Key:      "gpu-node-label",
							Operator: corev1.NodeSelectorOpIn,
							Values:   []string{"nvidia-tesla-a100"},
						},
						corev1.NodeSelectorRequirement{
							Key:      "gpu-partition-size",
							Operator: corev1.NodeSelectorOpIn,
							Values:   []string{"1g.5gb"},
						},
					},
				},
			},
			[]corev1.Toleration{
				{
					Key:      "gpu-node-label",
					Value:    "nvidia-tesla-a100",
					Operator: corev1.TolerationOpEqual,
					Effect:   corev1.TaintEffectNoSchedule,
				},
				{
					Key:      "gpu-partition-size",
					Value:    "1g.5gb",
					Operator: corev1.TolerationOpEqual,
					Effect:   corev1.TaintEffectNoSchedule,
				},
			},
		},
	}

	for _, f := range fixtures {
		t.Run(f.name, func(t *testing.T) {
			taskTemplate := dummyRayTaskTemplate("ray-id", dummyRayCustomObj())
			taskTemplate.ExtendedResources = f.extendedResourcesBase
			taskContext := dummyRayTaskContext(taskTemplate, f.resources, f.extendedResourcesOverride)
			rayJobResourceHandler := rayJobResourceHandler{}
			r, err := rayJobResourceHandler.BuildResource(context.TODO(), taskContext)
			assert.Nil(t, err)
			assert.NotNil(t, r)
			rayJob, ok := r.(*rayv1alpha1.RayJob)
			assert.True(t, ok)

			// Head node
			headNodeSpec := rayJob.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec
			assert.EqualValues(
				t,
				f.expectedNsr,
				headNodeSpec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms,
			)
			assert.EqualValues(
				t,
				f.expectedTol,
				headNodeSpec.Tolerations,
			)

			// Worker node
			workerNodeSpec := rayJob.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec
			assert.EqualValues(
				t,
				f.expectedNsr,
				workerNodeSpec.Affinity.NodeAffinity.RequiredDuringSchedulingIgnoredDuringExecution.NodeSelectorTerms,
			)
			assert.EqualValues(
				t,
				f.expectedTol,
				workerNodeSpec.Tolerations,
			)
		})
	}
}

func TestDefaultStartParameters(t *testing.T) {
	rayJobResourceHandler := rayJobResourceHandler{}
	rayJob := &plugins.RayJob{
		RayCluster: &plugins.RayCluster{
			HeadGroupSpec:   &plugins.HeadGroupSpec{},
			WorkerGroupSpec: []*plugins.WorkerGroupSpec{{GroupName: workerGroupName, Replicas: 3}},
		},
	}

	taskTemplate := dummyRayTaskTemplate("ray-id", rayJob)
	toleration := []corev1.Toleration{{
		Key:      "storage",
		Value:    "dedicated",
		Operator: corev1.TolerationOpExists,
		Effect:   corev1.TaintEffectNoSchedule,
	}}
	err := config.SetK8sPluginConfig(&config.K8sPluginConfig{DefaultTolerations: toleration})
	assert.Nil(t, err)

	RayResource, err := rayJobResourceHandler.BuildResource(context.TODO(), dummyRayTaskContext(taskTemplate, resourceRequirements, nil))
	assert.Nil(t, err)

	assert.NotNil(t, RayResource)
	ray, ok := RayResource.(*rayv1alpha1.RayJob)
	assert.True(t, ok)

	headReplica := int32(1)
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.Replicas, &headReplica)
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.ServiceAccountName, serviceAccount)
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.RayStartParams,
		map[string]string{
			"dashboard-host": "0.0.0.0", "disable-usage-stats": "true", "include-dashboard": "true",
			"node-ip-address": "$MY_POD_IP"})
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Annotations, map[string]string{"annotation-1": "val1"})
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Labels, map[string]string{"label-1": "val1"})
	assert.Equal(t, ray.Spec.RayClusterSpec.HeadGroupSpec.Template.Spec.Tolerations, toleration)

	workerReplica := int32(3)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Replicas, &workerReplica)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MinReplicas, &workerReplica)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].MaxReplicas, &workerReplica)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].GroupName, workerGroupName)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.ServiceAccountName, serviceAccount)
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].RayStartParams, map[string]string{"disable-usage-stats": "true", "node-ip-address": "$MY_POD_IP"})
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Annotations, map[string]string{"annotation-1": "val1"})
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Labels, map[string]string{"label-1": "val1"})
	assert.Equal(t, ray.Spec.RayClusterSpec.WorkerGroupSpecs[0].Template.Spec.Tolerations, toleration)
}

func newPluginContext() k8s.PluginContext {
	plg := &mocks2.PluginContext{}

	taskExecID := &mocks.TaskExecutionID{}
	taskExecID.OnGetID().Return(core.TaskExecutionIdentifier{
		NodeExecutionId: &core.NodeExecutionIdentifier{
			ExecutionId: &core.WorkflowExecutionIdentifier{
				Name:    "my_name",
				Project: "my_project",
				Domain:  "my_domain",
			},
		},
	})

	tskCtx := &mocks.TaskExecutionMetadata{}
	tskCtx.OnGetTaskExecutionID().Return(taskExecID)
	plg.OnTaskExecutionMetadata().Return(tskCtx)
	return plg
}

func init() {
	f := defaultConfig
	f.Logs = logs.LogConfig{
		IsKubernetesEnabled: true,
	}

	if err := SetConfig(&f); err != nil {
		panic(err)
	}
}

func TestGetTaskPhase(t *testing.T) {
	ctx := context.Background()
	rayJobResourceHandler := rayJobResourceHandler{}
	pluginCtx := newPluginContext()

	testCases := []struct {
		rayJobPhase       rayv1alpha1.JobStatus
		rayClusterPhase   rayv1alpha1.JobDeploymentStatus
		expectedCorePhase pluginsCore.Phase
	}{
		{"", rayv1alpha1.JobDeploymentStatusInitializing, pluginsCore.PhaseInitializing},
		{rayv1alpha1.JobStatusPending, rayv1alpha1.JobDeploymentStatusFailedToGetOrCreateRayCluster, pluginsCore.PhasePermanentFailure},
		{rayv1alpha1.JobStatusPending, rayv1alpha1.JobDeploymentStatusWaitForDashboard, pluginsCore.PhaseRunning},
		{rayv1alpha1.JobStatusPending, rayv1alpha1.JobDeploymentStatusFailedJobDeploy, pluginsCore.PhasePermanentFailure},
		{rayv1alpha1.JobStatusPending, rayv1alpha1.JobDeploymentStatusRunning, pluginsCore.PhaseRunning},
		{rayv1alpha1.JobStatusPending, rayv1alpha1.JobDeploymentStatusFailedToGetJobStatus, pluginsCore.PhaseUndefined},
		{rayv1alpha1.JobStatusRunning, rayv1alpha1.JobDeploymentStatusRunning, pluginsCore.PhaseRunning},
		{rayv1alpha1.JobStatusFailed, rayv1alpha1.JobDeploymentStatusRunning, pluginsCore.PhasePermanentFailure},
		{rayv1alpha1.JobStatusSucceeded, rayv1alpha1.JobDeploymentStatusRunning, pluginsCore.PhaseSuccess},
		{rayv1alpha1.JobStatusSucceeded, rayv1alpha1.JobDeploymentStatusComplete, pluginsCore.PhaseSuccess},
	}

	for _, tc := range testCases {
		t.Run("TestGetTaskPhase_"+string(tc.rayJobPhase), func(t *testing.T) {
			rayObject := &rayv1alpha1.RayJob{}
			rayObject.Status.JobStatus = tc.rayJobPhase
			rayObject.Status.JobDeploymentStatus = tc.rayClusterPhase
			startTime := metav1.NewTime(time.Now())
			rayObject.Status.StartTime = &startTime
			phaseInfo, err := rayJobResourceHandler.GetTaskPhase(ctx, pluginCtx, rayObject)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedCorePhase.String(), phaseInfo.Phase().String())
		})
	}
}

func TestGetPropertiesRay(t *testing.T) {
	rayJobResourceHandler := rayJobResourceHandler{}
	expected := k8s.PluginProperties{}
	assert.Equal(t, expected, rayJobResourceHandler.GetProperties())
}
