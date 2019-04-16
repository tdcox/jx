package cmd

import (
	"github.com/jenkins-x/jx/pkg/jx/cmd/opts"
	"github.com/jenkins-x/jx/pkg/jx/cmd/templates"
	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1"
)

// InstallDependenciesFlags flags for the install command
type InstallDependenciesFlags struct {
	Dependencies []string
}

// InstallDependenciesOptions options for install dependencies
type InstallDependenciesOptions struct {
	*opts.CommonOptions
	Flags InstallDependenciesFlags
}

var (
	installDependenciesLong = templates.LongDesc(`
		Installs required dependencies
`)

	installDependenciesExample = templates.Examples(`
		# Install dependencies
		jx install dependencies

		jx install dependencies -d gcloud
`)

	availableDependencies = []string{
		"az",
		"kubectl",
		"gcloud",
		"helm",
		"ibmcloud",
		"tiller",
		"helm3",
		"hyperkit",
		"kops",
		"kvm",
		"kvm2",
		"ksync",
		"minikube",
		"minishift",
		"oc",
		"virtualbox",
		"xhyve",
		"hyperv",
		"terraform",
		"oci",
		"aws",
		"eksctl",
		"heptio-authenticator-aws",
		"kustomize",
	}
)

// NewCmdInstallDependencies creates a command object to install any required dependencies
func NewCmdInstallDependencies(commonOpts *opts.CommonOptions) *cobra.Command {

	options := CreateInstallDependenciesOptions(commonOpts)

	cmd := &cobra.Command{
		Use:     "dependencies",
		Short:   "Install missing dependencies",
		Long:    installDependenciesLong,
		Example: installDependenciesExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			CheckErr(err)
		},
		SuggestFor: []string{"dependency"},
	}

	cmd.Flags().StringArrayVarP(&options.Flags.Dependencies, "dependencies", "d", []string{}, "The dependencies to install")

	return cmd
}

// CreateInstallDependenciesOptions creates the options for jx install dependencies
func CreateInstallDependenciesOptions(commonOpts *opts.CommonOptions) InstallDependenciesOptions {
	options := InstallDependenciesOptions{
		CommonOptions: commonOpts,
	}
	return options
}

// Run implements this command
func (options *InstallDependenciesOptions) Run() error {
	install := []string{}

	if len(options.Flags.Dependencies) == 0 {

		prompt := &survey.MultiSelect{
			Message: "What dependencies would you like to install:",
			Options: availableDependencies,
		}

		surveyOpts := survey.WithStdio(options.In, options.Out, options.Err)

		survey.AskOne(prompt, &install, nil, surveyOpts)
	} else {
		install = append(install, options.Flags.Dependencies...)
	}

	if len(install) > 0 {
		return options.DoInstallMissingDependencies(install)
	}

	options.Debugf("No dependencies selected to install\n")
	return nil
}
