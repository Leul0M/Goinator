/*
Copyright © 2025 Lex
*/
package cmd

import (
	"Goinator/data"
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
    reader := bufio.NewReader(os.Stdin)
    answers := make([]bool, 0, len(data.Questions))

    for _, q := range data.Questions {
        fmt.Println(q.Text)
        fmt.Print("y/n > ")

        ans, _ := reader.ReadString('\n')
        ans = strings.ToLower(strings.TrimSpace(ans))
        answers = append(answers, ans == "y")
    }

    // naive match: all 3 answers must be yes
    if len(answers) == 3 && answers[0] && answers[1] && answers[2] {
        fmt.Printf("I guess: %s\n", data.Secret.Name)
    } else {
        fmt.Println("I don't know yet.")
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
