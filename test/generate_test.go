package test

import (
	codechal "codechal/generate"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {
	filename := "testfile.txt"
	contents := "this is a test"
	codechal.WriteToFile(filename, contents)
	if _, err := os.Stat("testfile.txt"); err != nil {
		t.Errorf("File was not created")
	}

}

func TestReadFile(t *testing.T) {
	filename := "testfile.txt"
	contents := "this is a test"
	result := codechal.ReadFile(filename)

	if result != contents {
		t.Errorf("Contents didn't match")
	}

}

func TestIsValidInputLength(t *testing.T) {
	input := "your@email.com"

	if codechal.IsValidInputLength(input) != true {
		t.Errorf("Input is longer than 250 characters")
	}

	anotherInput := "testingtestingtesting@example.com"

	if codechal.IsValidInputLength(anotherInput) != true {
		t.Errorf("Input is longer than 250 characters")
	}

}

func TestGenerateNewKeys(t *testing.T) {
	input := "your@email.com"

	codechal.GenerateNewKeys(input)

	if _, err := os.Stat("your@email.com_signature.txt"); err != nil {
		t.Errorf("File was not created")
	}

	if _, err := os.Stat("your@email.com_public.pem"); err != nil {
		t.Errorf("File was not created")
	}

	if _, err := os.Stat("your@email.com_private.pem"); err != nil {
		t.Errorf("File was not created")
	}

}
