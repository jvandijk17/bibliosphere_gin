package cmd

import (
	"bibliosphere_gin/adapters/database"
	"bibliosphere_gin/cmd/seed"
	"log"

	"github.com/spf13/cobra"
)

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with initial data",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.ConnectDatabase()
		if err != nil {
			log.Fatalf("Failed to connect to database: %v", err)
		}
		seed.SeedLibraries(db)
		seed.SeedUsers(db)
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
}
