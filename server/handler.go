package server

import (
	"context"
	"os"

	lsp "go.lsp.dev/protocol"
	"go.uber.org/zap"
)

type ServerHandler struct {
	lsp.Server
	log	     *zap.Logger
	storedDocuments    *DocumentStore
}


func newHandler(ctx context.Context, server lsp.Server, logger *zap.Logger) (*ServerHandler, context.Context, error) {
	storedDocuments := NewDocumentStore()
	handler := &ServerHandler{
		Server:       server,     
		log:	      logger,
		storedDocuments:    storedDocuments,
	}

	_, err := os.Getwd()
	if err != nil {
		handler.log.Error("Error getting current directory")
	}

	return handler, ctx, nil
}

func (h ServerHandler) Initialize(ctx context.Context, params *lsp.InitializeParams) (*lsp.InitializeResult, error) {
	return &lsp.InitializeResult{
		Capabilities: lsp.ServerCapabilities{
			TextDocumentSync: lsp.TextDocumentSyncOptions {
				Change: lsp.TextDocumentSyncKindIncremental,
				OpenClose: true,
			},
		},
		ServerInfo: &lsp.ServerInfo{
			Name:    "hare-ls",
			Version: "0.1.0",
		},
	}, nil
}

func (h *ServerHandler) Exit(ctx context.Context) (err error) {
	return nil
}

func (h *ServerHandler) Shutdown(ctx context.Context) (err error) {
	h.storedDocuments.CloseAll()
	return lsp.Server.Shutdown(h, ctx)
}

