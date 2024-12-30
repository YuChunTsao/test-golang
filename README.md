# test golang

[![Linux](https://github.com/YuChunTsao/test-golang/actions/workflows/linux.yml/badge.svg)](https://github.com/YuChunTsao/test-golang/actions/workflows/linux.yml)
[![Mac](https://github.com/YuChunTsao/test-golang/actions/workflows/mac.yml/badge.svg)](https://github.com/YuChunTsao/test-golang/actions/workflows/mac.yml)
[![Windows](https://github.com/YuChunTsao/test-golang/actions/workflows/windows.yml/badge.svg)](https://github.com/YuChunTsao/test-golang/actions/workflows/windows.yml)

## How to use this repository

### Clone the repository

```bash
git@github.com:YuChunTsao/test-golang.git
```

### Update submodules

```bash
git submodule update --init --recursive
```

### Compile rust library

```bash
cargo build --manifest-path ./pkg/test-rust/lib/hello/Cargo.toml -r
```

### Execute go code

The go function will call the rust library function.

```bash
go run .
# go run main.go
```

## How to add submodules

```bash
git submodule add git@github.com:YuChunTsao/test-rust.git ./pkg/test-rust
```
