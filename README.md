# WASM + Go

1. `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`
2. `GOOS=js GOARCH=wasm go build -o app.wasm`
3. `docker-compose up -d`
4. Go to http://localhost:8080