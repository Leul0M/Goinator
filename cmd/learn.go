package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"Goinator/loader"

	"github.com/spf13/cobra"
)

// learnCmd represents the learn command
var learnCmd = &cobra.Command{
	Use:   "learn",
	Short: "Add a new entity to the knowledge base",
	Run: func(cmd *cobra.Command, args []string) {
		const file = "data/entities.json"

		// 1) Load existing records
		recs, err := loader.LoadEntities(file)
		if err != nil && !os.IsNotExist(err) {
			fmt.Println("Error loading:", err)
			return
		}

		// 2) Ask for new entity name
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Name of the new entity: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			fmt.Println("Aborted.")
			return
		}

		// 3) Ask each trait
		traits := map[string]interface{}{}
		keys := []string{
			"can_fly",
			"has_fur_feathers",
			"has_screen",
			"has_superpowers",
			"is_actor",
			"is_electronic",
			"is_famous",
			"is_footwear",
			"is_human",
			"is_large_country",
			"is_living",
			"is_male",
			"is_multiplayer",
			"is_musician",
			"is_personal",
			"is_photo_sharing",
			"is_politician",
			"is_portable",
			"is_real",
			"is_superhero",
			"is_sweet",
			"is_youtuber",
			"lives_in_water",
		}
		for _, k := range keys {
			fmt.Printf("%s (y/n): ", k)
			ans, _ := reader.ReadString('\n')
			ans = strings.ToLower(strings.TrimSpace(ans))
			switch ans {
			case "y":
				traits[k] = true
			case "n":
				traits[k] = false
			default:
				traits[k] = nil
			}
		}

		// 4) Build new record and append
		newID := strings.ReplaceAll(strings.ToLower(name), " ", "_")
		recs = append(recs, loader.Record{
			ID:       newID,
			Name:     name,
			Traits:   traits,
		})

		// 5) Save back
		b, err := json.MarshalIndent(recs, "", "  ")
		if err != nil {
			fmt.Println("Marshal error:", err)
			return
		}
		if err := os.WriteFile(file, b, 0644); err != nil {
			fmt.Println("Save error:", err)
			return
		}
		fmt.Println("Saved! Run the game again to see the new entity.")
	},
}

func init() {
	rootCmd.AddCommand(learnCmd)
}
