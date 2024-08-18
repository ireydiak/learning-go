/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"github.com/ireydiak/learning-go/cli-todo-list/cli"
	"github.com/ireydiak/learning-go/cli-todo-list/fs"
	"github.com/ireydiak/learning-go/cli-todo-list/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [ID]",
	Short: "Delete an item from the todo list.",
	Long: `Delete an item from the todo list by its identifier.
	The todo item will be permanently removed from the system.
	Use "done" to mark an existing item as "done".`,
	Args: cli.RequiresExactIntArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}
		runDelete(i)
	},
}

func runDelete(toRemove int) {
	f, err := fs.OpenOrCreate("./todos.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer
	scanner := bufio.NewScanner(f)
	i := 0
	isFound := false
	for scanner.Scan() {
		if i == 0 {
			buf.WriteString(scanner.Text() + "\n")
		}
		if i != toRemove && i > 0 {
			txt := scanner.Text() + "\n"
			if i > toRemove {
				txt = utils.ReplaceUntilChar(scanner.Text(), strconv.Itoa(i-1), ',') + "\n"
			}
			buf.WriteString(txt)
		}
		if i == toRemove {
			isFound = true
		}
		i += 1
	}

	if err := f.Truncate(0); err != nil { // Erase file
		panic(err)
	}
	if _, err := f.Seek(0, 0); err != nil { // Go to the start of the file
		panic(err)
	}
	_, err = f.WriteString(buf.String())
	if err != nil {
		panic(err)
	}

	if isFound {
		fmt.Printf("Successfully deleted Todo item %d\n", toRemove)
	} else {
		fmt.Printf("Could not delete Todo item %d because it does not exist.\nPlease select an existing Todo item.\n", toRemove)
	}
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
