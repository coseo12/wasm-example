# WASM TEST

ls -alh $(go env GOROOT)/misc/wasm/wasm_exec.js

## Folder

- server: WebServer
- wasm: WebAssembly

## Commands

- WebAssembly build: GOOS=js GOARCH=wasm go build -o test.wasm
- Server start: go run server/server.go

## VSCode Preference Setting.json

```json
  "go.toolsEnvVars": {
    "GOARCH": "wasm",
    "GOOS": "js"
  },
  "go.testEnvVars": {
    "GOARCH": "wasm",
    "GOOS": "js"
  },
  "go.installDependenciesWhenBuilding": false
```
