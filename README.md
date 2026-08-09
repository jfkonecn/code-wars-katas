# Code Wars Katas

https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&beta=false&order_by=sort_date%20desc

This repo is a place to practice Codewars katas in Go.

A kata is a small programming challenge. Each kata usually gives you a function to write and some tests that check whether your function works.

## Important Files

- `go.mod`: Tells Go the name of this project and which outside packages it needs.
- `go.sum`: Records exact versions of outside packages so everyone uses the same code.
- `katas/`: Holds one folder for each kata.
- `solution.go`: The file where you write your answer.
- `solution_test.go`: The file with tests that check your answer.

## Run The Tests

From the root of this repo, run:

```sh
go test ./...
```

What this means:

- `go test` tells Go to run tests.
- `./...` means "run tests in this folder and every folder inside it."

## Add A New Kata

These steps assume you are new to programming. Go slowly and check your work after each step.

### 1. Pick A Kata

[Open Codewars](https://www.codewars.com/kata/search/go?q=&r%5B%5D=-8&beta=false&order_by=sort_date%20desc) and choose a Go kata:

### 2. Create A Folder

Inside `katas/`, create a new folder for the kata.

Use a short name with lowercase letters and dashes. For example:

```text
katas/remove-string-spaces/
```

The folder name becomes part of the local package import path, so keep it simple.
I'll explain what an import path is later.

### 3. Copy The Solution Code From Codewars

Create this file in your new kata folder:

```text
solution.go
```

Copy and paste the starter solution code directly from Codewars into `solution.go`.

This file is where you will write the function that solves the kata.

### 4. Copy The Test Code From Codewars

Create this file in the same kata folder:

```text
solution_test.go
```

Copy and paste the sample test code directly from Codewars into `solution_test.go`.

This file checks whether your solution works.

### 5. Fix The Test Import

An import path tells Go where to find code from another package.

Think of it like an address for code. When a test file imports your solution package, Go uses the import path to find the folder that contains your `solution.go` file.

In this repo, each kata folder is its own package. The test file needs to import that package so it can call your solution function.

Codewars test files often contain this import:

```go
. "codewarrior/kata"
```

That import works on Codewars, but it does not work in this local repo.

Change it to this format:

```go
. "github.com/jfkonecn/code-wars-katas/katas/YOUR-KATA-FOLDER"
```

Replace `YOUR-KATA-FOLDER` with the folder you created.

For example, if your folder is:

```text
katas/remove-string-spaces/
```

Then the import should be:

```go
. "github.com/jfkonecn/code-wars-katas/katas/remove-string-spaces"
```

That path is made by joining the module name with the folder path.

The module name comes from `go.mod`:

```go
module github.com/jfkonecn/code-wars-katas
```

The folder path comes from where you put the kata:

```text
katas/remove-string-spaces
```

Put them together like this:

```text
github.com/jfkonecn/code-wars-katas + /katas/remove-string-spaces
```

That creates the full import path:

```go
. "github.com/jfkonecn/code-wars-katas/katas/remove-string-spaces"
```

The dot at the start of the import, `.`, means the test can call your function without writing the package name every time. Codewars uses this style, so keep it unless you know why you want to change it.

### 6. Add A Test Runner If Needed

Some Codewars Go tests use Ginkgo and Gomega. If the test file has imports like this:

```go
. "github.com/onsi/ginkgo"
. "github.com/onsi/gomega"
```

These are outside packages that help write tests.

`github.com/onsi/ginkgo` is a testing framework. It gives the test file words like `Describe` and `It`.

- `Describe` groups related tests together.
- `It` describes one specific thing the solution should do.

`github.com/onsi/gomega` is an assertion library. It gives the test file words like `Expect` and `Equal`.

- `Expect` says what value you are checking.
- `Equal` says what value you expected to get.

For example:

```go
Expect(Solution("world")).To(Equal("dlrow"))
```

This means: run `Solution("world")` and expect the answer to be `"dlrow"`.

The dot, `.`, before each import means the test file can use names like `Describe`, `It`, `Expect`, and `Equal` directly instead of writing the package name before them.

Make sure `solution_test.go` also has this test runner:

```go
func TestSolution(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Solution Suite")
}
```

You will also need `testing` in the import list:

```go
"testing"
```

The test runner is the bridge between Go's normal test command and the Ginkgo-style tests that Codewars provides.

### 7. Run The Tests

Run:

```sh
go test ./...
```

At first, tests may fail. That is normal. A failing test tells you what still
needs to be fixed.

## If You Get Stuck

Use this prompt with an AI assistant. It asks the assistant to help you think
through the problem instead of giving you the answer immediately.

```text
I am working on a Go Codewars kata. Please use the Socratic method to help me get
unstuck. Please keep in mind that I am a complete beginner at programming so
I will need extra context.

Make sure I'm using idiomatic go in my solution and always give an explanation
as to why it's a best practices so I can learn faster.

Do not give me the solution. Ask me one small question at a
time. Help me understand:

1. What the function is supposed to receive as input.
2. What the function is supposed to return as output.
3. What the examples are showing.
4. What smaller steps could solve the problem.
5. How I can test one small idea before writing the whole solution.

For any general go programming questions, please act like a stack overflow contributor
to give the answer on how to do that specific thing, but do not give away the
solution by answering the questions. For code snippets, please explicitly state
that you are giving me a more generic code snippet so that you do not give me
the answer.

Every time we work on a new problem please ask me for the following:

- The Codewars kata description
- The current solution code
- The test code
- Any error messages or warning before you start answering questions.
```
