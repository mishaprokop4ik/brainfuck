# Brainfuck

Simple brainfuck interpreter for options to extensions

## Tests
#### To run tests run from cmd write:
```bash
make test
```

## Installation

Use the go get to install interpreter.

```bash
go get github.com/mishaprokop4ik/brainfuck
```

## Usage

Go to brainfuck folder

```bash
go run main.go -f brainfuck_example.b
```
or
```bash
go run main.go
```

## Make operations
lint project by golint 
```bash
make lint
```
fmt project
```bash
make fmt
```
Pre commit operations: fmt lint and test operations
```bash
make precommit
```
build binary file of project
```bash
make build
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.