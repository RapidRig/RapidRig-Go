package cmd

import (
	"github.com/RapidRig/rapidrig-go/internal/app"
	"github.com/RapidRig/rapidrig-go/internal/modules"
	"github.com/spf13/cobra"
)

func newUpdateCmd(app *app.App) {
	rootCmd.AddCommand(&cobra.Command{
		Use:   "update",
		Short: "update packages and configurations.",
		Long:  `This will run the update command against all modules.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := app.ModuleRegistry.ReadAndSetRegistryConfigsFromYAML()
			if err != nil {
				panic(err)
			}
			gpm := app.ModuleRegistry.GetModule("GenericPackageManager").(*modules.GenericPackageManager)
			call := app.ModuleRegistry.GetModule("SysCall").(*modules.SysCall)
			updateErr := gpm.Manager().Update(modules.ToBeInstalled[gpm.Manager()], call)
			if updateErr != nil {
				return updateErr
			}
			return app.ModuleRegistry.RunUpdate()
		},
	})
}
