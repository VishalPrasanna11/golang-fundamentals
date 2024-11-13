# GOLANG-FUNDAMENTALS

This repository is a structured collection of code samples, exercises, and projects aimed at teaching the fundamentals of the Go (Golang) programming language. Each folder covers a specific topic and contains practical examples to help you grasp Go's syntax and functionality.

## Table of Contents
- [GOLANG-FUNDAMENTALS](#golang-fundamentals)
  - [Table of Contents](#table-of-contents)
  - [About](#about)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Running Examples](#running-examples)
  - [Folder Structure](#folder-structure)
  - [Topics Covered](#topics-covered)
  - [Contributing](#contributing)
  - [Resources](#resources)
  - [License](#license)

---

## About

This repository provides hands-on examples covering fundamental Go programming concepts, from basic syntax to building web APIs. It is designed to be beginner-friendly while also introducing advanced topics for more experienced developers. 

## Getting Started

### Prerequisites

1. **Install Go**  
   Download and install Go from [Go's official website](https://golang.org/dl/).

   Verify your installation:

   ```bash
   go version
   ```

2. **Clone the Repository**  
   Clone this repository to your local environment.

   ```bash
   git clone https://github.com/your-username/go-lang-fundamentals.git
   cd go-lang-fundamentals
   ```

### Running Examples

To run an example file, navigate to the desired directory and use:

```bash
go run <filename>.go
```

For example:

```bash
cd 01-HelloWorld
go run main.go
```

## Folder Structure

The project is organized into topic-based folders, each containing examples and exercises related to a specific concept:

```
GOLANG-FUNDAMENTALS/
├── 01-HelloWorld
├── 02-Variables
├── 03-UserInput
├── 04-Conversions
├── 05-HandleTime
├── 06-Pointers
├── 07-Array
├── 08-Slices
├── 09-Maps
├── 10-Structs
├── 11-Conditions
├── 12-Switch
├── 13-Loops
├── 14-Functions
├── 15-Methods
├── 16-Defer
├── 17-Files
├── 18-WebRequest
├── 19-HandleUrl
├── 20-WebReqVerbs
├── 21-Json
├── 22-Modules
├── 23-BuildApi
├── 24-BuildAPI+DbConnection
└── README.md
```

## Topics Covered

1. **Basics** - Hello World, variables, user input, data type conversions.
2. **Data Types** - Arrays, slices, maps, structs.
3. **Control Flow** - Conditions, loops, switch statements.
4. **Functions & Methods** - Defining functions and methods in Go.
5. **Error Handling & Defer** - Handling errors and using `defer`.
6. **Working with Files** - File operations.
7. **Web Requests** - Making HTTP requests, handling URLs, and HTTP verbs.
8. **JSON Handling** - Working with JSON data.
9. **Modules & Packages** - Structuring Go projects with modules.
10. **API Development** - Building APIs and database connections.

## Contributing

Contributions are welcome! If you have additional examples or improvements, feel free to submit a pull request.

1. Fork this repository
2. Create a new branch (`git checkout -b feature-new-topic`)
3. Commit your changes (`git commit -am 'Add new topic example'`)
4. Push to your branch (`git push origin feature-new-topic`)
5. Open a Pull Request

## Resources

- [Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [A Tour of Go](https://tour.golang.org/)
- [Go by Example](https://gobyexample.com/)
- [Go Playground](https://play.golang.org/)

## License

This project is licensed under the MIT License.