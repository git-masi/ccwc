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
	t.Run("num bytes", func(t *testing.T) {
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
		t.Logf("test file name, '%s'", f.Name())

		_, err = f.WriteString("ten bytes!")
		if err != nil {
			t.Fatal(err)
		}

		binaryPath := path.Join(wd, BINARY_NAME)

		cmd := exec.CommandContext(ctx, binaryPath, "-c", f.Name())

		buf := new(bytes.Buffer)
		cmd.Stdout = buf

		err = cmd.Run()
		if err != nil {
			t.Fatalf("cannot run application, %s", err)
		}

		want := fmt.Sprintf("10 %s", f.Name())
		got := buf.String()
		if want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})
}
