package main

import (
	"os"
	"testing"
)

func TestCountCharacters(t *testing.T) {
	input := `Project Gutenberg™
`

	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.CreateTemp(wd, "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(f.Name())

	_, err = f.WriteString(input)
	if err != nil {
		t.Fatal(err)
	}
	f.Sync()

	f, err = os.Open(f.Name())
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	count, err := countCharacters(f)
	if err != nil {
		t.Fatal(err)
	}

	if count != 19 {
		t.Errorf("want 19, got %d", count)
	}
}
