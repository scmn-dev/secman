package clean

import (
	"fmt"

	survey "github.com/AlecAivazis/survey/v2"
	"github.com/abdfnx/shell"
	"github.com/scmn-dev/secman-v1/pkg/api/gh/pkg/prompt"
	commands "github.com/scmn-dev/secman-v1/tools/constants"
	checker "github.com/scmn-dev/version-checker"
	"github.com/spf13/cobra"
)

type CleanOptions struct {
	CleanAll bool
	CleanGit bool
}

func Clean(runF func(*CleanOptions) error) *cobra.Command {
	opts := CleanOptions{}

	cmd := &cobra.Command{
		Use:   "clean",
		Short: "Clean your ~/.secman (delete it).",
		RunE: func(c *cobra.Command, args []string) error {
			if runF != nil {
				return runF(&opts)
			}

			return _run(&opts)
		},
	}

	cmd.Flags().BoolVarP(&opts.CleanAll, "all", "a", false, "Delete ~/.secman para.")
	cmd.Flags().BoolVarP(&opts.CleanGit, "git", "g", false, "Delete .git in ~/.secman.")
	
	return cmd
}

func _run(opts *CleanOptions) error {
	if opts.CleanGit {
		shell.SHCore(commands.Clean_ml_git(), commands.Clean_w_git())
	} else if opts.CleanAll {
		shell.SHCore(commands.Clean_ml(), commands.Clean_w())
	} else {
		var err error

		err = _prompt()
		if err != nil {
			return err
		}
	}

	checker.Checker()

	return nil
}

func _prompt() (error) {
	var cleanType int
	err := prompt.SurveyAskOne(&survey.Select{
		Message: "What do you want to clean?",
		Options: []string{
			"Delete .git folder in ~/.secman",
			"Delete All ~/.secman",
		},
	}, &cleanType)

	if cleanType == 0 {
		shell.SHCore(commands.Clean_ml_git(), commands.Clean_w_git())
	} else if cleanType == 1 {
		shell.SHCore(commands.Clean_ml(), commands.Clean_w())
	}

	if err != nil {
		return fmt.Errorf("could not prompt: %w", err)
	}

	return nil
}
