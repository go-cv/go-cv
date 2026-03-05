package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-pdf/fpdf"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/text"
)

// generatePDF generates a PDF file from markdown content
func generatePDF(markdownContent, name string) error {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetAutoPageBreak(true, 10)
	pdf.SetMargins(15, 15, 15)

	// Parse markdown to AST
	reader := text.NewReader([]byte(markdownContent))
	doc := goldmark.DefaultParser().Parse(reader)

	// Walk the AST and render to PDF
	err := ast.Walk(doc, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		switch node := n.(type) {
		case *ast.Heading:
			renderHeading(pdf, node, reader)
		case *ast.Paragraph:
			renderParagraph(pdf, node, reader)
		case *ast.List:
			// Lists are handled by their items
			return ast.WalkContinue, nil
		case *ast.ListItem:
			renderListItem(pdf, node, reader)
		case *ast.CodeBlock:
			renderCodeBlock(pdf, node, reader)
		case *ast.CodeSpan:
			renderCodeSpan(pdf, node, reader)
		}

		return ast.WalkContinue, nil
	})

	if err != nil {
		return fmt.Errorf("failed to render PDF: %w", err)
	}

	// Ensure output directory exists
	if err := os.MkdirAll(outputPath, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write PDF to file
	outputFile := filepath.Join(outputPath, name+".pdf")
	if err := pdf.OutputFileAndClose(outputFile); err != nil {
		return fmt.Errorf("failed to write PDF: %w", err)
	}

	fmt.Printf("  -> Written: %s\n", outputFile)
	return nil
}

// getNodeText extracts text content from a node
func getNodeText(n ast.Node, reader text.Reader) string {
	var sb strings.Builder
	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if textNode, ok := c.(*ast.Text); ok {
			segment := textNode.Segment
			sb.Write(reader.Value(segment))
		} else {
			sb.WriteString(getNodeText(c, reader))
		}
	}
	return sb.String()
}

// renderHeading renders a heading to the PDF
func renderHeading(pdf *fpdf.Fpdf, node *ast.Heading, reader text.Reader) {
	txt := getNodeText(node, reader)
	switch node.Level {
	case 1:
		pdf.SetFont("Helvetica", "B", 24)
		pdf.Ln(5)
	case 2:
		pdf.SetFont("Helvetica", "B", 18)
		pdf.Ln(3)
	case 3:
		pdf.SetFont("Helvetica", "B", 14)
		pdf.Ln(2)
	default:
		pdf.SetFont("Helvetica", "B", 12)
		pdf.Ln(1)
	}
	pdf.MultiCell(0, 8, txt, "", "", false)
	pdf.Ln(2)
}

// renderParagraph renders a paragraph to the PDF
func renderParagraph(pdf *fpdf.Fpdf, node *ast.Paragraph, reader text.Reader) {
	txt := getNodeText(node, reader)
	if txt == "" {
		return
	}
	pdf.SetFont("Helvetica", "", 12)
	pdf.MultiCell(0, 6, txt, "", "", false)
	pdf.Ln(2)
}

// renderListItem renders a list item to the PDF
func renderListItem(pdf *fpdf.Fpdf, node *ast.ListItem, reader text.Reader) {
	txt := getNodeText(node, reader)
	if txt == "" {
		return
	}
	pdf.SetFont("Helvetica", "", 12)
	pdf.Cell(5, 6, "- ")
	pdf.MultiCell(0, 6, txt, "", "", false)
}

// renderCodeBlock renders a code block to the PDF
func renderCodeBlock(pdf *fpdf.Fpdf, node *ast.CodeBlock, reader text.Reader) {
	segments := node.Lines()
	if segments.Len() == 0 {
		return
	}

	pdf.SetFont("Courier", "", 10)
	pdf.SetFillColor(240, 240, 240)
	pdf.Ln(2)

	for i := 0; i < segments.Len(); i++ {
		segment := segments.At(i)
		line := string(reader.Value(segment))
		line = strings.TrimRight(line, "\n\r")
		pdf.Cell(0, 5, "  "+line)
		pdf.Ln(-1)
	}

	pdf.Ln(2)
	pdf.SetFont("Helvetica", "", 12)
}

// renderCodeSpan renders inline code to the PDF
func renderCodeSpan(pdf *fpdf.Fpdf, node *ast.CodeSpan, reader text.Reader) {
	txt := getNodeText(node, reader)
	pdf.SetFont("Courier", "", 12)
	pdf.Cell(0, 6, txt)
	pdf.SetFont("Helvetica", "", 12)
}
