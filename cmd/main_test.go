package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"testing"
	"time"
)

// Windows users need the ".exe" file extension
const BINARY_NAME = "ccwc"

func TestMain(m *testing.M) {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "build", "-o", BINARY_NAME)

	err := cmd.Run()
	if err != nil {
		os.Exit(1)
	}

	defer func() {
		err := os.Remove(BINARY_NAME)
		if err != nil {
			log.Fatalf("Error removing built binary: %v", err)
		}
	}()

	os.Exit(m.Run())
}

func TestApp(t *testing.T) {
	tt := []struct{ name, input, want, flag string }{
		{"count bytes", "ten bytes!", "10", "-c"},
		{"count lines", "line one\nline two\nline three", "3", "-l"},
		{"count words", "This sentence has five words.", "5", "-w"},
		{"whitespace does not add to word count", "     ", "0", "-w"},
		{"spaces, lines, and tabs, do not add to word count", "B.C.	\n514\tAccession of Ho Lu.\n", "6", "-w"},
		{"count characters", "Project Gutenberg™", "18", "-m"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			wd, err := os.Getwd()
			if err != nil {
				t.Fatal(err)
			}

			f, err := os.CreateTemp(wd, "test_*.txt")
			if err != nil {
				t.Fatal(err)
			}
			defer os.Remove(f.Name())

			_, err = f.WriteString(tc.input)
			if err != nil {
				t.Fatal(err)
			}

			binaryPath := path.Join(wd, BINARY_NAME)

			cmd := exec.CommandContext(ctx, binaryPath, tc.flag, f.Name())

			buf := new(bytes.Buffer)
			cmd.Stdout = buf

			err = cmd.Run()
			if err != nil {
				t.Fatalf("cannot run application, %s", err)
			}

			want := fmt.Sprintf("%s %s\n", tc.want, f.Name())
			got := buf.String()
			if want != got {
				t.Errorf("want '%s', got '%s'", want, got)
			}
		})
	}
}
