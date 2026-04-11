package server

import (
	"bytes"
	"strings"

	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type Document struct {
	URI     lsp.DocumentURI
	Path    string
	Content []byte
	lines   []string
	IsOpen  bool
}

type TextDocument interface {
	ApplyChanges([]lsp.TextDocumentContentChangeEvent)
}

func NewDocument(fileURI uri.URI, content []byte, isOpen bool) *Document {
	return &Document{
		URI:     fileURI,
		Path:    fileURI.Filename(),
		Content: content,
		IsOpen:  isOpen,
	}
}

func PositionToIndex(pos lsp.Position, content []byte) int {
	index := 0
	for i := 0; i < int(pos.Line); i++ {
		if i < int(pos.Line) {
			index = index + strings.Index(string(content[index:]), "\n") + 1
		}
	}

	index = index + int(pos.Character)
	return index
}

// ApplyChanges updates the content of the document from LSP textDocument/didChange events.
func (d *Document) ApplyChanges(changes []lsp.TextDocumentContentChangeEvent) {

	content := d.Content

	for _, change := range changes {
		start, end := PositionToIndex(change.Range.Start, content), PositionToIndex(change.Range.End, content)

		var buf bytes.Buffer
		buf.Write(content[:start])
		buf.Write([]byte(change.Text))
		buf.Write(content[end:])
		content = buf.Bytes()
	}

	d.Content = content
	d.lines = nil
}
