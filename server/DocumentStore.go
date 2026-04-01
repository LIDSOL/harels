package server

import (
	"sync"

	lsp "go.lsp.dev/protocol"
	"go.lsp.dev/uri"
)


type DocumentStore struct {
	documents sync.Map
}

func NewDocumentStore() *DocumentStore {
	return &DocumentStore{
		documents: sync.Map{},
	}
}

func (s *DocumentStore) didOpen(params *lsp.DidOpenTextDocumentParams) (*HareSourceDocument, error) {
	uri := params.TextDocument.URI
	path := uri.Filename()
	doc := NewHareSourceDocument(uri, []byte(params.TextDocument.Text), true)

	s.documents.Store(path, doc)

	return doc, nil
}

func (s *DocumentStore) didClose(params *lsp.DidCloseTextDocumentParams) error {
	uri := params.TextDocument.URI
	path := uri.Filename()

	s.documents.Delete(path)

	return nil
}

func (s *DocumentStore) CloseAll() {
	s.documents.Clear()
}

func (s *DocumentStore) GetSyncDocument(uri uri.URI) (DocumentInterface, bool) {
	path := uri.Filename()
	d, ok := s.documents.Load(path)
	if !ok {
		return nil, false
	}
	doc, ok := d.(DocumentInterface)
	return doc, ok
}
