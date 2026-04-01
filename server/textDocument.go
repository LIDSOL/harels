package harels

import (
	"context"
	"errors"
	"fmt"

	lsp "go.lsp.dev/protocol"
)

func (h *ServerHandler) didOpen(ctx context.Context, params *lsp.DidOpenTextDocumentParams) (err error) {
	_, err = h.storedDocuments.didOpen(params)

	if err != nil {
		return err
	}

	lsp.Server.DidOpen(ctx, params)

	return nil
}

func (h *ServerHandler) didClose(ctx context.Context, params *lsp.DidCloseTextDocumentParams) (err error) {
	err = h.storedDocuments.didClose(params)

	if err != nil {
		return err
	}

	lsp.Server.DidClose(ctx, params)

	return nil
}

func (h *ServerHandler) didChange(ctx context.Context, params *lsp.DidChangeTextDocumentParams) (err error) {
	doc, ok := h.storedDocuments.GetSyncDocument(params.TextDocument.URI)
	if !ok {
		return errors.New("Could not get document: " + params.TextDocument.URI.Filename())
	}

	doc.ApplyChanges(params.ContentChanges)

	return lsp.Server.DidChange(ctx, params)
}
