# Hare Language Server — Extensión de VSCode

Esta carpeta contiene la extensión de Visual Studio Code para brindar soporte
de desarrollo al lenguaje de programación Hare, conectándose con el servidor
de lenguaje `harels`.

---

## Requisitos Previos

Antes de instalar y ejecutar la extensión, asegúrate de tener instalado en tu sistema:

* [Node.js](https://nodejs.org/es/download/current) versión 18 o superior recomendada -> para distribuidor del usuario -> usando nvm -> con npm.

* Verificar version de Node.js node -v

* Verificar version de npm npm -v

---

## Instalación de Dependencias

La extensión tiene su propio `package.json` dentro de la carpeta `vscode/`.
Instala las dependencias **desde esta carpeta**, no desde la raíz del repositorio:

```bash
cd vscode
npm install
```

> Las advertencias `npm warn deprecated` son normales y no afectan el funcionamiento.

---

## Desarrollo y Pruebas Locales

1. Abre la carpeta `vscode` en Visual Studio Code.
2. Genera los archivos de salida de TypeScript necesarios para la compilación:

```bash
npx tsc -b client/tsconfig.json
```

> Este paso es necesario la primera vez que se clona el repositorio, ya que
> el directorio `client/out/` no existe hasta que se ejecuta este comando.

3. Compila el proyecto:

```bash
npm run compile
```

> Al finalizar correctamente verás: `Build successful! Binary created at: bin/harels-server`

4. Abre el archivo `client/src/extension.ts` en VSCode.
5. Listo para ejecutar con F5 o a Run -> Start Debugging para abrir una ventana nueva
   de VSCode con la extensión cargada.
6. Abre cualquier archivo `.hare` en esa ventana para verificar que el servidor
   de lenguaje esté activo.

---

## Scripts Disponibles

| Script            | Descripción |
| `npm run compile` | Ejecuta el chequeo de tipos, el linter, compila con esbuild y corre `./server/build.sh` |
| `npm run watch`   | Modo desarrollo: recompila automáticamente al detectar cambios en el código TypeScript|
| `npm run lint`    | Revisa la calidad y formato del código fuente |---|---|