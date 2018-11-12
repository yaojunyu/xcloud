package cmd

type factory struct {
}

// NewFactory creates a factory with the default Kubernetes resources defined.
func NewFactory() Factory {
	f := &factory{}

	return f
}
