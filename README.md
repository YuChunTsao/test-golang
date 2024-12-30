# test golang

```bash
go run .
```

## How to add submodules

```bash
git submodule add git@github.com:YuChunTsao/test-rust.git ./pkg/test-rust
```

## How to compile rust library

```bash
cargo build --manifest-path ./pkg/test-rust/lib/hello/Cargo.toml -r
```

## Execute go code

```bash
go run .
# go run main.go
```
