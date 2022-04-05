package cli

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/abdfnx/tran/tools"
	"github.com/scmn-dev/tran/models"
	"github.com/scmn-dev/tran/constants"
	tranUI "github.com/scmn-dev/tran/tui"
)

func FilesCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "files",
		Short: "Securely transfer and send anything between computers.",
		Long: "Securely transfer and send anything between computers.",
	}

	cmd.AddCommand(FilesSendCMD())
	cmd.AddCommand(FilesReceiveCMD())

	return cmd
}

func FilesSendCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send",
		Short: "Send files/directories to remote.",
		Long: "Send files/directories to remote.",
		Aliases: []string{"s"},
		RunE: func(cmd *cobra.Command, args []string) error {
			tools.RandomSeed()

			err := tranUI.ValidateTranxAddress()

			if err != nil {
				log.Fatal(err)
			}

			tranUI.HandleSendCommand(models.TranOptions{
				TranxAddress: constants.DEFAULT_ADDRESS,
				TranxPort:    constants.DEFAULT_PORT,
			}, args)

			return nil
		},
	}

	return cmd
}

func FilesReceiveCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "receive",
		Short: "Receive files/directories from remote",
		Long:  "Receive files/directories from remote",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := tranUI.ValidateTranxAddress()

			if err != nil {
				return err
			}

			tranUI.HandleReceiveCommand(models.TranOptions{
				TranxAddress: constants.DEFAULT_ADDRESS,
				TranxPort:    constants.DEFAULT_PORT,
			}, args[0])

			return nil
		},
	}

	return cmd
}
