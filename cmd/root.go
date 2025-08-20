/*
Copyright © 2025 Lex
*/
package cmd

import (
	"Goinator/loader"
	"Goinator/tui"
	"fmt"
	"os"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "Goinator",
	Short: "Think of anything you want and I'll try to guess it",
	Long: `Goinator is a text-based guessing game that uses a decision tree to try 
	to figure out which character you're thinking of. From fictional heroes to historical 
	figures, the possibilities are endless! Just answer my questions with a simple 'yes' or 'no', 
	and I'll do my best to guess who's on your mind.`,
	Run: func(cmd *cobra.Command, args []string) {
		recs, err := loader.LoadEntities("data/entities.json")
		if err != nil {
			panic(err)
		}
		root := loader.BuildTree(recs)

		p := tea.NewProgram(tui.New(root))
		if _, err := p.Run(); err != nil {
			fmt.Println("TUI error:", err)
		}
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
