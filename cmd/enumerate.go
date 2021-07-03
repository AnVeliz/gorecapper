package cmd

import (
	"fmt"
	"os"

	enumerator "github.com/AnVeliz/gorecapper/internal/searcher"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Sets what to search for",
	Long:  `You may need to search for interfaces or structs, etc`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println("Invalid arguments. You should define the object to find (interface or struct), its name and the root folder to search in.")
			os.Exit(1)
		}

		argmapper := map[string]enumerator.SearchObjectKind{
			"interface": enumerator.Interface,
			"struct":    enumerator.Struct,
		}
		if _, ok := argmapper[args[0]]; !ok {
			fmt.Printf("%s is not a kind of SearchObject", args[0])
			os.Exit(1)
		}
		if _, err := os.Stat(args[1]); os.IsNotExist(err) {
			fmt.Printf("%s path doesn't exist", args[1])
		}

		enumerator.Enumerate(argmapper[args[0]], args[1])
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
