package main

import (
	"bytes"
	"testing"
)

func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	expected := 4

	result := countWords(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1\n word2 word3\n word4")

	expected := 3

	result := countLines(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}

func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4")

	expected := 23

	result := countBytes(b)

	if result != expected {
		t.Errorf("Expected %d, got %d instead\n", expected, result)
	}
}
