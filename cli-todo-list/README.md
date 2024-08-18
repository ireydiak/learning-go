# Summary

This is another very simple TODO list application.
The main purpose of this project is to learn Golang through basic CLI applications.
It is not meant to be very robust or even user-friendly.

# Installation

Clone this project and build the binaries.
You can use the included [Makefile](./Makefile) to compile the binaries.

```bash
git clone git@github.com:ireydiak/learning-go.git
cd learning-go/cli-todo-list
make build
```

To test the installation:

```bash
./bin/go-todo list
```

You should see an empty table!
Read the following section to learn the available commands.

# Usage

## Create

To create a new TODO item:

```bash
todo add <description>
```

Replace `<description>` with a short sentence describing the task to be done.

## List

To list existing TODO items:

```bash
todo list
```

## Complete

To mark an item as completed:

```bash
todo complete <ID>
```

## Delete

Marking an item as comleted will not remove it from the list.
To permanently delete an existing TODO item:

```bash
todo delete <ID>
```

Replace `<ID>` with the TODO identifier. You can get a list of identifiers using the `list` command.
If the item does not exist, an error message is returned.

# Development

The CLI tool is developed using the well-established [Cobra library](https://github.com/spf13/cobra).
To add new commands, I recommend installing the Cobra CLI.
Instructions can be found in their [Github repository](https://github.com/spf13/cobra)
Once installed, run the following command from the root of this project to add a new command.

```bash
cobra-cli add <COMMAND_NAME>
```

Replace `<COMMAND_NAME>` with the desired name of the command.

