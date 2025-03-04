# go-playground
A place to play with Go code

In the spirit of [the Go playground](https://go.dev/play/), this repo holds experiments and one-offs in Go programming.

To run each program, use `go run ./cmd/<program name>`

# programs
Arranged in order of addition (mostly)

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
This command requires the [kanaka/mal repo](https://github.com/kanaka/mal) to be cloned locally, and the path to that directory passed as the command line argument.

## file_read
Experiment with reading bytes from a file while reliably counting lines.
This command requires the [kanaka/mal repo](https://github.com/kanaka/mal) to be cloned locally, and the path to that directory passed as the second command line argument.

## pointers
More poking around with Go pointers and return by value vs return by reference.

## unique
Experiment with an algorithm to sort and uniqify a slice in place.

## file_name_from_language
More work with [enry](https://github.com/go-enry/go-enry), but this time the other way around: determine the file name from a given language. Generates a regular expression designed to encompass all of the file names/extensions that could be associated with the given language.

## language_from_file_name
The opposite of `file_name_from_language`.

## composition
Play around with [composition](https://www.codecademy.com/resources/docs/go/composition)

## execute
Illustrate how `os/exec` handles input, whether it's a slice or not.

## secure_storage
First pass at figuring out how to store secrets in the operating system's secure storage (aka keychains or keyrings)
There are two commands here:
1. `secure_store` - stores the secret passed as the second command line argument in the OS' secure storage
2. `secure_retrieve` - retrieves the secret stored in the OS' secure storage and prints it to stdout
Different operating systems use different secure storage mechanisms. Currently only macOS is supported.

I've experimented with various ways to make it work in Linux.

To run the secure storage under Linux, build a Docker container with GNOME keyring:
```
docker build -f secure_storage.dockerfile -t secure_storage_image .
docker run -it -v $(pwd):/src secure_storage_image
```

## led
A program to print out the seven-segment display for an integer

## histogram
Prints out a histogram of random number distribution segmented into buckets

## sqlf
Experiments with `sqlf`, query parameters, and `ILIKE`