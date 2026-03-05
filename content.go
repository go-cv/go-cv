package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
)

// ContentFile represents a parsed content file
type ContentFile struct {
	SourcePath string // Original .md file path
	Name       string // File name without extension
	Content    string // Raw markdown content
	HTML       string // Converted HTML
}

// generateOutput reads content from ./content, processes it, and writes to ./output
func generateOutput() error {
	// Ensure content directory exists
	if _, err := os.Stat(contentPath); os.IsNotExist(err) {
		return fmt.Errorf("content directory does not exist: %s", contentPath)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Read all markdown files from content directory
	files, err := readContentFiles()
	if err != nil {
		return fmt.Errorf("failed to read content files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("Warning: No markdown files found in content directory")
		return nil
	}

	// Process each file
	for _, file := range files {
		fmt.Printf("Processing: %s\n", file.SourcePath)

		// Convert markdown to HTML
		html, err := markdownToHTML(file.Content)
		if err != nil {
			return fmt.Errorf("failed to convert %s: %w", file.SourcePath, err)
		}
		file.HTML = html

		// Write output
		outputFile := filepath.Join(outputPath, file.Name+".html")
		if err := os.WriteFile(outputFile, []byte(html), 0644); err != nil {
			return fmt.Errorf("failed to write %s: %w", outputFile, err)
		}
		fmt.Printf("  -> Written: %s\n", outputPath)
	}

	return nil
}

// readContentFiles reads all .md files from the content directory
func readContentFiles() ([]ContentFile, error) {
	var files []ContentFile

	entries, err := os.ReadDir(contentPath)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		if !strings.HasSuffix(name, ".md") {
			continue
		}

		fullPath := filepath.Join(contentPath, name)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read %s: %w", fullPath, err)
		}

		files = append(files, ContentFile{
			SourcePath: fullPath,
			Name:       strings.TrimSuffix(name, ".md"),
			Content:    string(content),
		})
	}

	return files, nil
}

// markdownToHTML converts markdown content to HTML using goldmark
func markdownToHTML(markdown string) (string, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert([]byte(markdown), &buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}
