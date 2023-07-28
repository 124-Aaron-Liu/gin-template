package cmd

import (
	"github.com/124-Aaron-Liu/gin-template/internal/app"
	"github.com/spf13/cobra"
	"log"
)

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "A brief description of your application",
	Long:  "A longer description that spans multiple lines and likely contains examples",
	Run: func(cmd *cobra.Command, args []string) {
		app, err := app.NewApplication()
		if err != nil {
			log.Fatalln(err)
		}
		app.Runserver()
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runserverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runserverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
