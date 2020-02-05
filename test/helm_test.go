// Modified sample from https://github.com/gruntwork-io/terratest/blob/master/test/helm_basic_example_template_test.go

package helm_test

import (
	"path/filepath"
	"strings"
	"testing"

	_ "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	_ "github.com/smartystreets/goconvey/convey"
	_ "k8s.io/api/batch/v1beta1"
	certmanagerv1alpha2 "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	"github.com/stretchr/testify/require"
	appsv1 "k8s.io/api/apps/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/logger"
	"github.com/gruntwork-io/terratest/modules/random"
)

// An example of how to verify the rendered template object of a Helm Chart given various inputs.
func TestHelmBasicExampleTemplateRenderedDeployment(t *testing.T) {
	t.Parallel()

	// Path to the helm chart we will test
	helmChartPath, err := filepath.Abs("testdata/sample")
	releaseName := "helm-basic"
	require.NoError(t, err)

	// Since we aren't deploying any resources, there is no need to setup kubectl authentication or helm home.

	// Set up the namespace; confirm that the template renders the expected value for the namespace.
	namespaceName := "medieval-" + strings.ToLower(random.UniqueId())
	logger.Logf(t, "Namespace: %s\n", namespaceName)

	// Setup the args. For this test, we will set the following input values:
	// - containerImageRepo=nginx
	// - containerImageTag=1.15.8
	options := &helm.Options{
		SetValues: map[string]string{
			"image.repository": "nginx",
			"image.tag":        "1.15.8",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	// Run RenderTemplate to render the template and capture the output. Note that we use the version without `E`, since
	// we want to assert that the template renders without any errors.
	// Additionally, although we know there is only one yaml file in the template, we deliberately path a templateFiles
	// arg to demonstrate how to select individual templates to render.
	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/deployment.yaml"})

	// Now we use kubernetes/client-go library to render the template output into the Deployment struct. This will
	// ensure the Deployment resource is rendered correctly.
	var deployment appsv1.Deployment
	helm.UnmarshalK8SYaml(t, output, &deployment)

	// Verify the namespace matches the expected supplied namespace.
	require.Equal(t, namespaceName, deployment.Namespace)

	// Finally, we verify the deployment pod template spec is set to the expected container image value
	expectedContainerImage := "nginx:1.15.8"
	deploymentContainers := deployment.Spec.Template.Spec.Containers
	require.Equal(t, len(deploymentContainers), 1)
	require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}

func TestApplicationCertConfig(t *testing.T) {
	t.Parallel()
	helmChartPath, err := filepath.Abs("testdata/sample")
	releaseName := "argo-cd"
	require.NoError(t, err)

	options := &helm.Options{
		SetValues: map[string]string{
			"certificate.issuer.kind": "certKind",
			"certificate.issuer.name": "certName",
		},
		KubectlOptions: k8s.NewKubectlOptions("", "", "default"),
	}
	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{"templates/certificate.yaml"})
	var certificate certmanagerv1alpha2.Certificate
	helm.UnmarshalK8SYaml(t, output, &certificate)

	require.Equal(t, certificate.Spec.IssuerRef.Kind, "certKind")
	require.Equal(t, certificate.Spec.IssuerRef.Name, "certName")
	// require.Equal(t, deploymentContainers[0].Image, expectedContainerImage)
}
