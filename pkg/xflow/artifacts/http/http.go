package http

import (
	wfv1 "github.com/xcloudnative/xcloud/pkg/apis/xcloudnative.io/v1alpha1"
	"github.com/xcloudnative/xcloud/pkg/errors"
	"github.com/xcloudnative/xcloud/pkg/xflow/common"
)

// HTTPArtifactDriver is the artifact driver for a HTTP URL
type HTTPArtifactDriver struct{}

// Load download artifacts from an HTTP URL
func (h *HTTPArtifactDriver) Load(inputArtifact *wfv1.Artifact, path string) error {
	// Download the file to a local file path
	return common.RunCommand("curl", "-sS", "-L", "-o", path, inputArtifact.HTTP.URL)
}

func (h *HTTPArtifactDriver) Save(path string, outputArtifact *wfv1.Artifact) error {
	return errors.Errorf(errors.CodeBadRequest, "HTTP output artifacts unsupported")
}
