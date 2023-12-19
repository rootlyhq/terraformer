package cmd

import (
	rootly_terraforming "github.com/GoogleCloudPlatform/terraformer/providers/rootly"

	"github.com/GoogleCloudPlatform/terraformer/terraformutils"
	"github.com/spf13/cobra"
)

func newCmdRootlyImporter(options ImportOptions) *cobra.Command {
	token := ""
	cmd := &cobra.Command{
		Use:   "rootly",
		Short: "Import current state to Terraform configuration from Rootly",
		Long:  "Import current state to Terraform configuration from Rootly",
		RunE: func(cmd *cobra.Command, args []string) error {
			provider := newRootlyProvider()
			err := Import(provider, options, []string{token})
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.AddCommand(listCmd(newRootlyProvider()))
	baseProviderFlags(cmd.PersistentFlags(), &options, "environment", "")
	cmd.PersistentFlags().StringVarP(&token, "token", "t", "", "env param ROOTLY_API_TOKEN")
	return cmd
}

func newRootlyProvider() terraformutils.ProviderGenerator {
	return &rootly_terraforming.RootlyProvider{}
}
