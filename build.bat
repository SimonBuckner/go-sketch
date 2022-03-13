SET GOOS=js
SET GOARCH=wasm

cd app
go build -o ..\main.wasm
cd ..

SET GOOS=""
SET GOARCH=""
