import * as path from 'path';
import { workspace, ExtensionContext } from 'vscode';

import {
  LanguageClient,
  LanguageClientOptions,
  ServerOptions,
  TransportKind,
  Executable
} from 'vscode-languageclient/node';

let client: LanguageClient | undefined;

export async function activate(context: ExtensionContext) {
  // The server is implemented in Go
  const serverExecutable: Executable = {
    command: context.asAbsolutePath(path.join('server', 'bin', 'harels-server')),
  };

  const serverOptions: ServerOptions = serverExecutable;
  
  // Options to control the language client
  let clientOptions: LanguageClientOptions = {
    // Register the server for hare files
    documentSelector: [{ scheme: 'file', language: 'hare' }],
    synchronize: {
      // Notify the server about file changes to '.clientrc files contained in the workspace
      fileEvents: workspace.createFileSystemWatcher('**/.clientrc')
    }
  };

  // Create the language client and start the client.
  client = new LanguageClient(
    'languageServerExample',
    'Language Server Example',
    serverOptions,
    clientOptions
  );

  // Start the client. This will also launch the server
  await client.start();
}

export async function deactivate() {
  await client?.dispose();
  client = undefined;
}
