# hello-go-wasm
wasm by golang

## Build

```sh
GOOS=js GOARCH=wasm go build -o main.wasm
```

## Run

```sh
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```
