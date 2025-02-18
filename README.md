# Build Your Own wc Tool

Inspired by John Crickett's [Coding Challenges](https://codingchallenges.fyi/challenges/challenge-wc).

From `man wc`:

> The wc utility displays the number of lines, words, and bytes contained in
     each input file, or standard input (if no file is specified) to the
     standard output.

## Requirements

This project assumes you're using MacOS. Other Unix-like systems may also work.

Check the `go.mod` file for the required version of Go.

## Run the application

Build the application using `go build -o ccwc ./cmd`.

Run the application using the flags described in the coding challenge. Basic example: `./ccwc -c test.txt`.