package cli

import (
	"github.com/scmn-dev/secman/v6/pkg/generator"
	"github.com/spf13/cobra"
)

func GenerateCMD() *cobra.Command{
	cmd := &cobra.Command{
		Use:   "generate",
		Aliases: []string{"gen"},
		Short: "Generate a new password.",
		Long:  "Generate a new password. you can set the password length with the --length flag.",
		RunE: func(cmd *cobra.Command, args []string) error {
			generator.Generator(&GenOpts)

			return nil
		},
	}

	cmd.Flags().IntVarP(&GenOpts.Length, "length", "l", 10, "Set the length of the password.")
	cmd.Flags().BoolVarP(&GenOpts.Raw, "raw", "r", false, "Generate a password and print it.")

	return cmd
}
