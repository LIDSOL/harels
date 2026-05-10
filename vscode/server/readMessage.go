package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"go.lsp.dev/protocol"
)

// Document tracking
var documents = make(map[protocol.DocumentURI]string)

func readFromStdin() {
	//Read the header
	reader := bufio.NewReader(os.Stdin)

	for {
		// Read all header lines until empty line
		var headerLines []string
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error reading header: %v\n", err)
				return
			}
			line = strings.TrimRight(line, "\r\n")
			if line == "" {
				// Empty line marks end of headers
				break
			}
			headerLines = append(headerLines, line)
		}

		// Parse headers
		header := parseHeaders(headerLines)
		if header == nil {
			fmt.Fprintln(os.Stderr, "Error parsing headers")
			continue
		}

		//Read the body
		body, err := readBody(reader, header["Content-Length"])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading body: %v\n", err)
			continue
		}

		// Parse the message
		var req RequestMessage
		err = json.Unmarshal(body, &req)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing message: %v\n", err)
			continue
		}

		// Handle initialize request
		if req.Method == "initialize" {
			handleInitialize(req)
		}
		if req.Method == "textDocument/didOpen" {
			handleTextDocumentDidOpen(req)
		}
		if req.Method == "textDocument/didChange" {
			handleTextDocumentDidChange(req)
		}
		if req.Method == "textDocument/foldingRange" {
			handleFoldingRange(req)
		}
	}
}

func parseHeaders(headerLines []string) map[string]string {
	header := make(map[string]string)

	for _, line := range headerLines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		header[key] = value
	}

	if _, ok := header["Content-Length"]; !ok {
		return nil
	}
	// Content-Type is optional in LSP

	return header
}

func readBody(reader *bufio.Reader, contentLengthStr string) ([]byte, error) {
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return nil, fmt.Errorf("invalid Content-Length: %v", err)
	}
	body := make([]byte, contentLength)
	_, err = io.ReadFull(reader, body)
	if err != nil {
		return nil, fmt.Errorf("error reading body: %v", err)
	}
	return body, nil
}

// handleInitialize processes the initialize request and sends back a response
func handleInitialize(req RequestMessage) {
	// Parse initialize params (optional - not strictly needed for minimal response)
	var params protocol.InitializeParams
	json.Unmarshal(req.Params, &params)

	// Create minimal initialize result
	result := protocol.InitializeResult{
		Capabilities: protocol.ServerCapabilities{
			TextDocumentSync: protocol.TextDocumentSyncOptions{
				OpenClose: true,
				Change:    protocol.TextDocumentSyncKindFull,
			},
			FoldingRangeProvider: true,
		},

		ServerInfo: &protocol.ServerInfo{
			Name:    "Hare Analyzer Server",
			Version: "1.0.0",
		},
	}

	fmt.Fprintln(os.Stderr, "HandleInitialize called")

	// Send response
	sendResponse(req.ID, result, nil)
}

func handleFoldingRange(req RequestMessage) {
	fmt.Fprintln(os.Stderr, "handleFoldingRange called with params:", string(req.Params))

	foldingRanges := []protocol.FoldingRange{
		{
			StartLine: 0, // first line to fold
			EndLine:   4, // last line to fold
		},
	}

	fmt.Fprintln(os.Stderr, "Sending folding ranges:", foldingRanges)

	sendResponse(req.ID, foldingRanges, nil)
}

func handleTextDocumentDidOpen(req RequestMessage) {
	var params protocol.DidOpenTextDocumentParams
	err := json.Unmarshal(req.Params, &params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing didOpen params: %v\n", err)
		return
	}
	documents[params.TextDocument.URI] = params.TextDocument.Text
	fmt.Fprintf(os.Stderr, "Document opened: %s (len=%d)\n", params.TextDocument.URI, len(params.TextDocument.Text))
}

func handleTextDocumentDidChange(req RequestMessage) {
	var params protocol.DidChangeTextDocumentParams
	err := json.Unmarshal(req.Params, &params)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing didChange params: %v\n", err)
		return
	}
	if _, ok := documents[params.TextDocument.URI]; ok {
		documents[params.TextDocument.URI] = params.ContentChanges[len(params.ContentChanges)-1].Text
		fmt.Fprintf(os.Stderr, "Document changed: %s\n", params.TextDocument.URI)
	}
}

// sendResponse sends a JSON-RPC response to stdout
func sendResponse(id interface{}, result interface{}, err *ResponseError) {

	response := ResponseMessage{
		Message: Message{Jsonrpc: "2.0"},
		ID:      id,
		Result:  result,
		Error:   err,
	}

	// Marshal to JSON
	body, marshalErr := json.Marshal(response)
	if marshalErr != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling response: %v\n", marshalErr)
		return
	}

	fmt.Fprintf(os.Stderr, "Sending response body: %s\n", string(body))

	// Write response with LSP headers to stdout
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	fmt.Fprint(os.Stdout, header)
	fmt.Fprint(os.Stdout, string(body))

	os.Stdout.Sync()
}

func main() {
	readFromStdin()
}
