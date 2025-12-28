package resources

import _ "embed"

//go:embed wasm_exec.js
var WasmExecJs []byte

//go:embed app.wasm
var AppWasm []byte
