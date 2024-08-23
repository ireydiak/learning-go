/*
Copyright © 2024 Jean-Charles Verdier <ireydiak@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
	"github.com/ireydiak/learning-go/cli-todo-list/cli"
	"github.com/ireydiak/learning-go/cli-todo-list/fs"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [TITLE]",
	Short: "Create a new todo item",
	Long: `Create a new todo item with a [TITLE].
	The first positional argument [TITLE] is required.
	`,
	Args: cli.RequiresMinArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		runCreate(title, time.Now())
	},
}

func nextIdFromRecords(records []cli.TodoItem) int {
	return len(records) + 1
}

func runCreate(name string, createdAt time.Time) {
	fname := "./todos.csv"

	f, err := fs.OpenOrCreate(fname)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	records := []*cli.TodoItem{}
	if err = gocsv.UnmarshalFile(f, &records); err != nil {
		panic(err)
	}

	newRecord := &cli.TodoItem{
		Title:     name,
		Id:        0,
		CreatedAt: createdAt,
		Status:    cli.StatusPending,
	}
	records = append(records, newRecord)
	reorderRecords(records)

	if _, err := f.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	err = gocsv.MarshalFile(records, f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("New item %s has been added to the list\n", name)
}

func reorderRecords(records []*cli.TodoItem) {
	for i, record := range records {
		record.Id = i + 1
	}
}

func init() {
	rootCmd.AddCommand(createCmd)
}
