package cmd

import (
	// this is so that we load the auth plugins so we can connect to, say, GCP

	"github.com/xcloudnative/xcloud/pkg/client/clientset/versioned"
	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
)

// Factory is the interface defined for jx interactions via the cli
//go:generate pegomock generate github.com/xcloudnative/xcloud/pkg/jx/cmd Factory -o mocks/factory.go --generate-matchers
type Factory interface {
	//WithBearerToken(token string) Factory
	//
	//ImpersonateUser(user string) Factory
	//
	//CreateJenkinsClient(kubeClient kubernetes.Interface, ns string, in terminal.FileReader, out terminal.FileWriter, errOut io.Writer) (gojenkins.JenkinsClient, error)
	//
	//GetJenkinsURL(kubeClient kubernetes.Interface, ns string) (string, error)
	//
	//CreateAuthConfigService(fileName string) (auth.AuthConfigService, error)
	//
	//CreateJenkinsAuthConfigService(kubernetes.Interface, string) (auth.AuthConfigService, error)
	//
	//CreateChartmuseumAuthConfigService() (auth.AuthConfigService, error)
	//
	//CreateIssueTrackerAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)
	//
	//CreateChatAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)
	//
	//CreateAddonAuthConfigService(secrets *corev1.SecretList) (auth.AuthConfigService, error)
	//
	CreateClient() (kubernetes.Interface, string, error)
	//
	//CreateGitProvider(string, string, auth.AuthConfigService, string, bool, gits.Gitter, terminal.FileReader, terminal.FileWriter, io.Writer) (gits.GitProvider, error)
	//
	CreateKubeConfig() (*rest.Config, error)

	CreateJXClient() (versioned.Interface, string, error)
	//
	CreateApiExtensionsClient() (apiextensionsclientset.Interface, error)
	//
	//CreateDynamicClient() (*dynamic.APIHelper, string, error)
	//
	//CreateMetricsClient() (*metricsclient.Clientset, error)
	//
	//CreateComplianceClient() (*client.SonobuoyClient, error)
	//
	//CreateKnativeBuildClient() (buildclient.Interface, string, error)
	//
	//CreateTable(out io.Writer) table.Table
	//
	//SetBatch(batch bool)
	//
	//IsInCluster() bool
	//
	//IsInCDPipeline() bool
	//
	//AuthMergePipelineSecrets(config *auth.AuthConfig, secrets *corev1.SecretList, kind string, isCDPipeline bool) error
	//
	//CreateVaultOperatorClient() (vaultoperatorclient.Interface, error)
}
