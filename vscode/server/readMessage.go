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
			// Add minimal capabilities here - empty for now
		},
	}
	fmt.Fprintln(os.Stderr, "HandleInitialize called")
	// Send response
	sendResponse(req.ID, result, nil)
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

	// Write response with LSP headers to stdout
	header := fmt.Sprintf("Content-Length: %d\r\n\r\n", len(body))
	fmt.Fprint(os.Stdout, header)
	fmt.Fprint(os.Stdout, string(body))
	os.Stdout.Sync()
}

func main() {
	readFromStdin()
}
