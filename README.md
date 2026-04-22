# harels
Language Server Implementation for the Hare programming language

Esta rama contiene una implementación en desarrollo de un Language Server Protocol (LSP) para Harels, con el objetivo de integrar el intérprete dentro de Visual Studio Code.

## Objetivo de la rama

Proveer soporte básico de lenguaje para Harels dentro de VS Code, incluyendo:

* Comunicación cliente-servidor vía LSP
* Procesamiento de mensajes desde el editor
* Base para features como:

  * Hover
  * Autocompletado
  * Validación sintáctica

---

## Componentes principales

vscode/
├── client/        # Extensión VS Code (TypeScript)
├── server/        # Servidor LSP en Go
└── src/           # Código fuente de la extensión


### Cliente VS Code

Implementado en TypeScript, maneja:

* Conexión con el servidor LSP
* Integración con el editor

### Servidor (Go)

Encargado de:

* Procesar mensajes del cliente
* Interpretar el código Harels
* Enviar respuestas (diagnostics, etc.)

## Requisitos

* Go >= 1.24
* Node.js (para la extensión)
* npm

## Pasos

### 1. Clonar repositorio


### 2. Instalar dependencias Go

go mod tidy
go mod vendor

---

### 3. Instalar dependencias del cliente
cd vscode
npm install

## Ejecución

### Ejecutar servidor LSP

cd vscode/server
go build


### Ejecutar extensión en VS Code

go run harels.go


## Funciones

* Comunicación cliente-servidor básica
* Estructura inicial del LSP


##  Relación con el intérprete

El servidor LSP utiliza el parser generado con ANTLR para procesar código Harels.


## Desarrollo

Para extender el LSP:

1. Implementar handlers en el servidor (server/)
2. Conectar eventos en el cliente (client/)
3. Integrar con el parser de Harels

---

##  Estado del proyecto

En desarrollo activo — enfocado en integración con VS Code y soporte básico de lenguaje.

---

## License
MIT — ver LICENSE 

```