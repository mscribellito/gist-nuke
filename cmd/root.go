package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func NewRootCommand() *cobra.Command {

	var token string
	var noDryRun bool
	var removed, failed int

	command := &cobra.Command{
		Use:   "gist-nuke",
		Short: "gist-nuke removes all gists from a GitHub account",
		Long:  `A tool which removes all gists from a GitHub account. Use it with caution.`,
	}

	command.Flags().StringVarP(&token, "token", "t", "", "Personal Access Token")
	command.Flags().BoolVarP(&noDryRun, "no-dry-run", "", false, "If specified, will actually delete gists")

	command.MarkFlagRequired("token")

	command.RunE = func(cmd *cobra.Command, args []string) error {

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		gists, _, err := client.Gists.List(ctx, "", nil)
		if err != nil {
			return err
		}

		// Check if gists found
		if len(gists) == 0 {
			fmt.Println("No gists found.")
			return nil
		}

		// Iterate over gists
		for _, gist := range gists {

			id := fmt.Sprintf("%v", *gist.ID)
			desc := fmt.Sprintf("%v", *gist.Description)

			// we're actually doing it
			if noDryRun == true {

				status := "finished"

				_, err := client.Gists.Delete(ctx, id)

				if err != nil {
					failed++
					status = "failed"
				} else {
					removed++
				}

				fmt.Printf("%s - '%s' - %s\n", desc, id, status)

			} else {

				fmt.Printf("%s - '%s' - would remove\n", desc, id)

			}

		}

		if noDryRun == true {
			fmt.Printf("\nRemoval: %d removed, %d failed\n", removed, failed)
		}

		return nil

	}

	return command

}
