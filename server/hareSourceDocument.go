package server

import (
	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)

type HareSourceDocument struct {
	Document
}

func (d *HareSourceDocument) GetDocumentType() DocumentType {
	return HareSourceDocumentType
}

func NewHareSourceDocument(fileURI uri.URI, content []byte, isOpen bool) *HareSourceDocument {
	return &HareSourceDocument{
		Document:                *NewDocument(fileURI, content, isOpen),
	}
}

// ApplyChanges updates the content of the document from LSP textDocument/didChange events.
func (d *HareSourceDocument) ApplyChanges(changes []lsp.TextDocumentContentChangeEvent) {
	d.Document.ApplyChanges(changes)
}

func IsHareSourceDocumentLangID(langID lsp.LanguageIdentifier) bool {
	return langID == "hare"
}
