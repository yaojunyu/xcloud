package common

import (
	"time"

	"github.com/xcloudnative/xcloud/pkg/apis/xcloudnative.io"
)

const (
	// xcloudnativeioControllerConfigMapKey is the key in the configmap to retrieve xcloudnativeio configuration from.
	// Content encoding is expected to be YAML.
	xcloudnativeioControllerConfigMapKey = "config"

	// DefaultArchivePattern is the default pattern when storing artifacts in an archive repository
	DefaultArchivePattern = "{{xcloudnativeio.name}}/{{pod.name}}"

	// Container names used in the xcloudnativeio pod
	MainContainerName = "main"
	InitContainerName = "init"
	WaitContainerName = "wait"

	// PodMetadataVolumeName is the volume name defined in a xcloudnativeio pod spec to expose pod metadata via downward API
	PodMetadataVolumeName = "podmetadata"

	// PodMetadataAnnotationsVolumePath is volume path for metadata.annotations in the downward API
	PodMetadataAnnotationsVolumePath = "annotations"
	// PodMetadataMountPath is the directory mount location for DownwardAPI volume containing pod metadata
	PodMetadataMountPath = "/xflow-controller/" + PodMetadataVolumeName
	// PodMetadataAnnotationsPath is the file path containing pod metadata annotations. Examined by executor
	PodMetadataAnnotationsPath = PodMetadataMountPath + "/" + PodMetadataAnnotationsVolumePath

	// DockerLibVolumeName is the volume name for the /var/lib/docker host path volume
	DockerLibVolumeName = "docker-lib"
	// DockerLibHostPath is the host directory path containing docker runtime state
	DockerLibHostPath = "/var/lib/docker"
	// DockerSockVolumeName is the volume name for the /var/run/docker.sock host path volume
	DockerSockVolumeName = "docker-sock"

	// AnnotationKeyNodeName is the pod metadata annotation key containing the xcloudnativeio node name
	AnnotationKeyNodeName = xcloudnativeio.FullName + "/node-name"
	// AnnotationKeyNodeMessage is the pod metadata annotation key the executor will use to
	// communicate errors encountered by the executor during artifact load/save, etc...
	AnnotationKeyNodeMessage = xcloudnativeio.FullName + "/node-message"
	// AnnotationKeyTemplate is the pod metadata annotation key containing the container template as JSON
	AnnotationKeyTemplate = xcloudnativeio.FullName + "/template"
	// AnnotationKeyOutputs is the pod metadata annotation key containing the container outputs
	AnnotationKeyOutputs = xcloudnativeio.FullName + "/outputs"
	// AnnotationKeyExecutionControl is the pod metadata annotation key containing execution control parameters
	// set by the controller and obeyed by the executor. For example, the controller will use this annotation to
	// signal the executors of daemoned containers that it should terminate.
	AnnotationKeyExecutionControl = xcloudnativeio.FullName + "/execution"

	// LabelKeyControllerInstanceID is the label the controller will carry forward to xcloudnativeios/pod labels
	// for the purposes of xcloudnativeio segregation
	LabelKeyControllerInstanceID = xcloudnativeio.FullName + "/controller-instanceid"
	// LabelKeyCompleted is the metadata label applied on worfklows and xcloudnativeio pods to indicates if resource is completed
	// xcloudnativeios and pods with a completed=true label will be ignored by the controller
	LabelKeyCompleted = xcloudnativeio.FullName + "/completed"
	// LabelKeyxcloudnativeio is the pod metadata label to indicate the associated xcloudnativeio name
	LabelKeyxcloudnativeio = xcloudnativeio.FullName + "/xcloudnativeio"
	// LabelKeyPhase is a label applied to xcloudnativeios to indicate the current phase of the xcloudnativeio (for filtering purposes)
	LabelKeyPhase = xcloudnativeio.FullName + "/phase"

	// ExecutorArtifactBaseDir is the base directory in the init container in which artifacts will be copied to.
	// Each artifact will be named according to its input name (e.g: /argo/inputs/artifacts/CODE)
	ExecutorArtifactBaseDir = "/xflow-controller/inputs/artifacts"

	// InitContainerMainFilesystemDir is a path made available to the init container such that the init container
	// can access the same volume mounts used in the main container. This is used for the purposes of artifact loading
	// (when there is overlapping paths between artifacts and volume mounts)
	InitContainerMainFilesystemDir = "/mainctrfs"

	// ExecutorStagingEmptyDir is the path of the emptydir which is used as a staging area to transfer a file between init/main container for script/resource templates
	ExecutorStagingEmptyDir = "/xflow-controller/staging"
	// ExecutorScriptSourcePath is the path which init will write the script source file to for script templates
	ExecutorScriptSourcePath = "/xflow-controller/staging/script"
	// ExecutorResourceManifestPath is the path which init will write the a manifest file to for resource templates
	ExecutorResourceManifestPath = "/tmp/manifest.yaml"

	// Various environment variables containing pod information exposed to the executor container(s)

	// EnvVarPodName contains the name of the pod (currently unused)
	EnvVarPodName = "ARGO_POD_NAME"

	// EnvVarContainerRuntimeExecutor contains the name of the container runtime executor to use, empty is equal to "docker"
	EnvVarContainerRuntimeExecutor = "ARGO_CONTAINER_RUNTIME_EXECUTOR"
	// EnvVarDownwardAPINodeIP is the envvar used to get the `status.hostIP`
	EnvVarDownwardAPINodeIP = "ARGO_KUBELET_HOST"
	// EnvVarKubeletPort is used to configure the kubelet api port
	EnvVarKubeletPort = "ARGO_KUBELET_PORT"
	// EnvVarKubeletInsecure is used to disable the TLS verification
	EnvVarKubeletInsecure = "ARGO_KUBELET_INSECURE"

	// ContainerRuntimeExecutorDocker to use docker as container runtime executor
	ContainerRuntimeExecutorDocker = "docker"

	// ContainerRuntimeExecutorKubelet to use the kubelet as container runtime executor
	ContainerRuntimeExecutorKubelet = "kubelet"

	// ContainerRuntimeExecutorK8sAPI to use the Kubernetes API server as container runtime executor
	ContainerRuntimeExecutorK8sAPI = "k8sapi"

	// Variables that are added to the scope during template execution and can be referenced using {{}} syntax

	// GlobalVarWorkflowName is a global xcloudnativeio variable referencing the xcloudnativeio's metadata.name field
	GlobalVarWorkflowName = "xcloudnativeio.name"
	// GlobalVarWorkflowNamespace is a global xcloudnativeio variable referencing the xcloudnativeio's metadata.namespace field
	GlobalVarWorkflowNamespace = "xcloudnativeio.namespace"
	// GlobalVarWorkflowUID is a global xcloudnativeio variable referencing the xcloudnativeio's metadata.uid field
	GlobalVarWorkflowUID = "xcloudnativeio.uid"
	// GlobalVarWorkflowStatus is a global xcloudnativeio variable referencing the xcloudnativeio's status.phase field
	GlobalVarWorkflowStatus = "xcloudnativeio.status"
	// GlobalVarWorkflowCreationTimestamp is the xcloudnativeio variable referencing the xcloudnativeios metadata.creationTimestamp field
	GlobalVarWorkflowCreationTimestamp = "xcloudnativeio.creationTimestamp"
	// LocalVarPodName is a step level variable that references the name of the pod
	LocalVarPodName = "pod.name"
)

// ExecutionControl contains execution control parameters for executor to decide how to execute the container
type ExecutionControl struct {
	// Deadline is a max timestamp in which an executor can run the container before terminating it
	// It is used to signal the executor to terminate a daemoned container. In the future it will be
	// used to support xcloudnativeio or steps/dag level timeouts.
	Deadline *time.Time `json:"deadline,omitempty"`
}
