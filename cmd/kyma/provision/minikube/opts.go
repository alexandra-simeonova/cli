package minikube

import "github.com/kyma-project/cli/internal/cli"

//options defines available options for the minikube provisioning command
type options struct {
	*cli.Options
	Domain              string
	VMDriver            string
	DiskSize            string
	Memory              string
	CPUS                string
	HypervVirtualSwitch string
}

//NewOptions creates options with default values
func NewOptions(o *cli.Options) *options {
	return &options{Options: o}
}
