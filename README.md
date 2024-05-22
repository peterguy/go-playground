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