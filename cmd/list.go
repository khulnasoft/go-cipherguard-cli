package cmd

import (
	"github.com/khulnasoft/go-cipherguard-cli/folder"
	"github.com/khulnasoft/go-cipherguard-cli/group"
	"github.com/khulnasoft/go-cipherguard-cli/resource"
	"github.com/khulnasoft/go-cipherguard-cli/user"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists Cipherguard Entitys",
	Long:    `Lists Cipherguard Entitys`,
	Aliases: []string{"index", "ls", "filter", "search"},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().BoolP("json", "j", false, "Output JSON")
	listCmd.PersistentFlags().String("filter", "",
		"Define a CEl expression as filter for any list commands. In the expression, all available columns of subcommand can be used (see -c/--column).\n"+
			"See also CEl specifications under https://github.com/google/cel-spec.\n"+
			"Examples:\n"+
			"\t--filter '(Name == \"SomeName\" || matches(Name, \"RegExpr\")) && URI.startsWith(\"https://auth.\")'\n"+
			"\t--filter 'Username == \"User\" && CreatedTimestamp > timestamp(\"2022-06-10T00:00:00.000-00:00\")'")
	listCmd.AddCommand(resource.ResourceListCmd)
	listCmd.AddCommand(folder.FolderListCmd)
	listCmd.AddCommand(group.GroupListCmd)
	listCmd.AddCommand(user.UserListCmd)
}
