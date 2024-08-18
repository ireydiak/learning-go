/*
Copyright © 2024 Jean-Charles Verdier <ireydiak@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/dustin/go-humanize"
	"github.com/gocarina/gocsv"
	"github.com/ireydiak/learning-go/cli-todo-list/cli"
	"github.com/ireydiak/learning-go/cli-todo-list/fs"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		runList()
	},
}

func runList() {
	fname := "./todos.csv"

	f, err := fs.OpenOrCreate(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records := []*cli.TodoItem{}
	err = gocsv.UnmarshalFile(f, &records)
	if err != nil {
		panic(err)
	}

	t := table.NewWriter()
	headerRow := table.Row{"Id", "Title", "Created At", "Status"}
	t.AppendHeader(headerRow)
	for _, record := range records {
		t.AppendRow(
			table.Row{record.Id, record.Title, humanize.Time(record.CreatedAt), record.Status},
		)
	}
	fmt.Println(t.Render())
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
