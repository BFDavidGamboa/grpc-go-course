# grpc-go-course

## Setup

In order to configure the environment, we need to install previously Go and Protobuf for the implementation.

[Link to install Go](https://go.dev/)

To see the versions of Go use:

```bash
go version
```

Next install protobuf for go and gRPC

- Install protoc for go

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

- Install protoc for gRPC

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go-grpc@v1.1
```

To see the versions of protobuf use:

```bash
protoc --version
```

### Initialize module

For module initialization use:

```bash
go mod init github.com/{UserName}/{RepositoryName}
```

This will ensure that the application run with the propper provided packages

Check the `go.mod` file to ensure the `Go` version and the module name.

## Makefile

The make file will automatize the process of generating the binaries, clean dependencies, building, returning OS information and returning the commands fot it's ussage.

### Commands

Write make followed by the command

```bash
make all
```

- all: Generate Pbs and build
- greet: Generate Pbs and build for greet
- calculator:  Generate Pbs and build for calculator
- blog:  Generate Pbs and build for blog
- test: Launch tests
- clean: Clean generated files
- clean_greet: Clean generated files for greet
- clean_calculator: Clean generated files for Calculator
- clean_blog: Clean generated files for Blog
- rebuild: Rebuild the whole project
- bump: Update packages version
- about: Display info related to the build
- help: Show this help
