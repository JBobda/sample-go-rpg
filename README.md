# Go RPG — Learning Project
 
A simple RPG built in Go, following the [rpg-in-golang project](https://github.com/vimichael/rpg-in-golang) as a hands-on way to learn the Go programming language.
 
## About
 
This project is a learning exercise. The goal isn't to ship a game — it's to get comfortable with Go by building something fun. The project uses [Ebiten](https://ebitengine.org/), a 2D game engine for Go, as the foundation.
 
## What I'm Learning
 
- Go syntax, types, and project structure
- Interfaces and how Go uses them (e.g. the `ebiten.Game` interface)
- The Ebiten game loop: `Update`, `Draw`, and `Layout`
- How Go handles packages and modules (`go.mod`, `go.sum`)
- Working with external dependencies in Go
## Prerequisites
 
- [Go](https://go.dev/dl/) 1.21 or later
- Ebiten dependencies — on Linux you may need a few system packages. See the [Ebiten install guide](https://ebitengine.org/en/documents/install.html) for your OS.
## Getting Started
 
```bash
# Clone the repo
git clone https://github.com/JBobda/sample-go-rpg.git
cd sample-go-rpg
 
# Install dependencies
go mod tidy
 
# Run the game
go run .
```
 
## Project Structure
 
```
sample-go-rpg/
├── main.go       # Entry point — game loop and window setup
├── go.mod        # Module definition and dependencies
├── go.sum        # Dependency checksums
└── .gitignore
```
 
## Reference
 
- [sample-go-rpg tutorial](https://github.com/JBobda/sample-go-rpg) — the tutorial this project follows
- [Ebiten documentation](https://ebitengine.org/)
- [A Tour of Go](https://go.dev/tour/) — useful alongside this project for Go fundamentals
- [Go by Example](https://gobyexample.com/) — quick reference for common Go patterns
 
