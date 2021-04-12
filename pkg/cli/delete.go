package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/replicate/cog/pkg/client"
)

func newDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Short:   "Delete a model",
		RunE:    deleteModel,
		Args:    cobra.MinimumNArgs(1),
		Aliases: []string{"rm"},
	}
	return cmd
}

func deleteModel(cmd *cobra.Command, args []string) error {
	repo, err := getRepo()
	if err != nil {
		return err
	}
	cli := client.NewClient()
	for _, id := range args {
		if err := cli.DeleteModel(repo, id); err != nil {
			return err
		}
		fmt.Printf("Deleted model %s:%s\n", repo, id)
	}
	return nil
}