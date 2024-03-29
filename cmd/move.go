package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/spf13/cobra"
)

// moveCmd represents the move command
var moveCmd = &cobra.Command{
	Use:   "move",
	Short: "Moves a Cipherguard Entity",
	Long:  `Moves a Cipherguard Entity`,
}

func init() {
	rootCmd.AddCommand(moveCmd)
	moveCmd.AddCommand(resource.ResourceMoveCmd)
	moveCmd.AddCommand(folder.FolderMoveCmd)
}
