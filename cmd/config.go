package cmd

import (
	"github.com/RapidRig/rapidrig-go/internal/app"
	"github.com/spf13/cobra"
)

func newWriteCurrent(app *app.App) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "generate",
		Short: "Writes the current config to $HOME/.rapidrig/config.yml.",
		Long:  `This will create a default YAML file using the defaults provided by each module.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return app.ModuleRegistry.WriteRegistryConfigsToYAML()
		},
	})
}
