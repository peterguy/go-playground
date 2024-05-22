# go-playground
A place to play with Go code

In the spirit of [the Go playground](https://go.dev/play/), this repo holds experiments and one-offs in Go programming.

The current hacky way to run the desired program/module is to go through the `main` program, passing the desired program as the first command line argument:

```
go run . <program>
```

# programs

## pipe_commands
Experimenting with piping the output from one command into another.
The basic test is piping the output from an `echo` command to a `wc` command.

## lipsum
Play around with generating Lorem Ipsum in Go

## struct_pointers
Make sure I underswtand Go's pointers

## struct_types
Experiment with Go's type system and how it allocates memory

## stuff
Just a random collection of stuff

## untar
Developing and experimenting with using `git archive` and handling the resulting `tar` archive
This command requires a second command line paramete: a directory path that is a git repository.

## url_encoding
Playing around with URL encoding

## file_types
Experiment with using [enry](https://github.com/go-enry/go-enry) to determine the programming language of a file.
This command requires the [kanaka/mal repo](https://github.com/kanaka/mal) to be cloned locally, and the path to that directory passed as the second command line argument.

## file_read
Experiment with reading bytes from a file while reliably counting lines.
This command requires the [kanaka/mal repo](https://github.com/kanaka/mal) to be cloned locally, and the path to that directory passed as the second command line argument.
