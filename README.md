# Task Manager CLI

A simple tool for managing tasks using Go and the Cobra library.

## Installation

First, clone the repository to your local machine:

```bash
git clone https://github.com/lolverae/todo-cli.git
cd todo-cli
```

### Building the CLI

To build the CLI tool, run the following command from the root of the repository:

```bash
go build -o todo-cli
```

This will create an executable named `todo-cli` in the current directory.

### Adding the Executable to Your PATH

To make the `todo-cli` executable available from anywhere, move it to a directory that is included in your system's PATH. For example, you can move it to `~/.local/bin/`:

```bash
mv todo-cli $HOME/.local/bin/
```
### Creating an Alias

To create an alias for easier usage, add the following line to your `.bashrc`:

```bash
echo 'alias tasks="todo-cli"' >> ~/.bashrc
source ~/.bashrc
```

Now you can use `tasks` instead of `todo-cli` to run the CLI tool.

## Usage

To create a new task:
```bash
todo-cli new "Task Title" --list "list_name"
```

To get all tasks:
```bash
todo-cli get  --list "list_name"
```

To mark a task as completed:

```bash
todo-cli done "Task Title" --list list_name
```

