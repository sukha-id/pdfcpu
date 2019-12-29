/*
Copyright 2019 The pdf Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"path/filepath"
	"testing"
)

func prepareForAttachmentTest() error {
	for _, fileName := range []string{"go.pdf", "golang.pdf", "T4.pdf", "go-lecture.pdf"} {
		inFile := filepath.Join(inDir, fileName)
		outFile := filepath.Join(outDir, fileName)
		if err := copyFile(inFile, outFile); err != nil {
			return err
		}
	}
	return copyFile(filepath.Join(resDir, "test.wav"), filepath.Join(outDir, "test.wav"))
}

func listAttachments(t *testing.T, msg, fileName string, want int) []string {
	t.Helper()
	list, err := ListAttachmentsFile(fileName, nil)
	if err != nil {
		t.Fatalf("%s list attachments: %v\n", msg, err)
	}
	// # of attachments must be want
	if len(list) != want {
		t.Fatalf("%s: list attachments %s: want %d got %d\n", msg, fileName, want, len(list))
	}
	return list
}

func TestAttachments(t *testing.T) {
	msg := "testAttachments"

	if err := prepareForAttachmentTest(); err != nil {
		t.Fatalf("%s prepare for attachments: %v\n", msg, err)
	}

	fileName := filepath.Join(outDir, "go.pdf")

	// # of attachments must be 0
	listAttachments(t, msg, fileName, 0)

	// attach add 4 files
	files := []string{
		outDir + "/golang.pdf",
		outDir + "/T4.pdf",
		outDir + "/go-lecture.pdf",
		outDir + "/test.wav"}

	if err := AddAttachmentsFile(fileName, "", files, false, nil); err != nil {
		t.Fatalf("%s add attachments: %v\n", msg, err)
	}
	list := listAttachments(t, msg, fileName, 4)
	for _, s := range list {
		t.Log(s)
	}

	// Extract all attachments.
	if err := ExtractAttachmentsFile(fileName, outDir, nil, nil); err != nil {
		t.Fatalf("%s extract all attachments: %v\n", msg, err)
	}

	// Extract 1 attachment.
	if err := ExtractAttachmentsFile(fileName, outDir, []string{"golang.pdf"}, nil); err != nil {
		t.Fatalf("%s extract one attachment: %v\n", msg, err)
	}

	// Remove 1 attachment.
	if err := RemoveAttachmentsFile(fileName, "", []string{"golang.pdf"}, nil); err != nil {
		t.Fatalf("%s remove one attachment: %v\n", msg, err)
	}
	listAttachments(t, msg, fileName, 3)

	// Remove all attachments.
	if err := RemoveAttachmentsFile(fileName, "", nil, nil); err != nil {
		t.Fatalf("%s remove all attachments: %v\n", msg, err)
	}
	listAttachments(t, msg, fileName, 0)

	// Validate the processed file.
	if err := ValidateFile(fileName, nil); err != nil {
		t.Fatalf("%s: validate: %v\n", msg, err)
	}
}

func TestPortfolio(t *testing.T) {
	msg := "testPortfolio"

	if err := prepareForAttachmentTest(); err != nil {
		t.Fatalf("%s prepare for portfolio: %v\n", msg, err)
	}

	fileName := filepath.Join(outDir, "go.pdf")

	// # of attachments must be 0
	listAttachments(t, msg, fileName, 0)

	// attach add 4 files including descriptions
	files := []string{
		outDir + "/golang.pdf",
		outDir + "/T4.pdf" + ", CCITT spec",
		outDir + "/go-lecture.pdf",
		outDir + "/test.wav" + ", test audio file"}

	if err := AddAttachmentsFile(fileName, "", files, false, nil); err != nil {
		t.Fatalf("%s add attachments: %v\n", msg, err)
	}
	list := listAttachments(t, msg, fileName, 4)
	for _, s := range list {
		t.Log(s)
	}

	// Extract all attachments.
	if err := ExtractAttachmentsFile(fileName, outDir, nil, nil); err != nil {
		t.Fatalf("%s extract all attachments: %v\n", msg, err)
	}

	// Extract 1 attachment.
	if err := ExtractAttachmentsFile(fileName, outDir, []string{"golang.pdf"}, nil); err != nil {
		t.Fatalf("%s extract one attachment: %v\n", msg, err)
	}

	// Remove 1 attachment.
	if err := RemoveAttachmentsFile(fileName, "", []string{"golang.pdf"}, nil); err != nil {
		t.Fatalf("%s remove one attachment: %v\n", msg, err)
	}
	listAttachments(t, msg, fileName, 3)

	// Remove all attachments.
	if err := RemoveAttachmentsFile(fileName, "", nil, nil); err != nil {
		t.Fatalf("%s remove all attachments: %v\n", msg, err)
	}
	listAttachments(t, msg, fileName, 0)

	// Validate the processed file.
	if err := ValidateFile(fileName, nil); err != nil {
		t.Fatalf("%s: validate: %v\n", msg, err)
	}
}