# cmit

A **CLI tool to standardize commit messages** following the Conventional Commits style, with predefined types, optional scopes, and descriptive emojis. It works as a wrapper around `git commit`, helping maintain consistency in the repository history.

---

## ✨ Features

* Select commit type (`feat`, `fix`, `docs`, `chore`, etc.)
* Optional **scope** support
* Automatic emojis based on commit type
* Interactive terminal interface (TUI)
* Generates the final commit message before executing `git commit`

---

## 📦 Installation

### Using Go

```sh
go install github.com/smokeeaasd/cmit/cmd/cmit@latest
```

Make sure `$GOPATH/bin` (or `%USERPROFILE%\go\bin` on Windows) is in your `PATH`.

### Using Make (Linux/Mac)

```sh
make build
```

The binary will be generated in `./bin/cmit`.

### Using Docker

```sh
docker build -t cmit .
docker run --rm -it -v $(pwd):/repo -w /repo cmit
```

---

## 🚀 Usage

Inside a Git repository:

```sh
# Start the interactive commit form
cmit

# Show a list of available commands and flags
cmit --help
```

Example of a generated commit message:

```
💡 feat(core): Add integration test
```

---

## 🧪 Testing

Run unit tests:

```sh
go test ./...
```

---

## 🛠 Development

Clone the repository:

```sh
git clone https://github.com/smokeeaasd/cmit.git
cd cmit
```

Useful commands:

```sh
make build   # Build binary
make test    # Run tests
make lint    # Run linter
```

---

## 📄 License

This project is licensed under the [MIT License](LICENSE).