/*
Copyright © 2025 Lex
*/
package cmd

import (
	"Goinator/loader"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "Goinator",
	Short: "Think of anything you want and I'll try to guess it",
	Long: `Goinator is a text-based guessing game that uses a decision tree to try 
	to figure out which character you're thinking of. From fictional heroes to historical 
	figures, the possibilities are endless! Just answer my questions with a simple 'yes' or 'no', 
	and I'll do my best to guess who's on your mind.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: func(cmd *cobra.Command, args []string) {
		recs, err := loader.LoadEntities("data/entities.json")
		if err != nil {
			panic(err)
		}

		root := loader.BuildTree(recs)
		node := root
		reader := bufio.NewReader(os.Stdin)

		for node.Entity == nil {
			fmt.Println(node.Question + "?")
			fmt.Print("y/n > ")
			ans, _ := reader.ReadString('\n')
			ans = strings.ToLower(strings.TrimSpace(ans))

			if ans == "y" && node.Yes != nil {
				node = node.Yes
			} else if node.No != nil {
				node = node.No
			} else {
				fmt.Println("I don’t know enough yet.")
				return
			}
		}
		fmt.Printf("I guess: %s\n", node.Entity.Name)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Goinator.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
