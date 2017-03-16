# http-route-audit

Summarize HTTP methods and response classes from access logs.

Small, dependency-light Go 1.8-era utility. Original implementation and history generated for a retrospective development exercise.

## Run

`go run . examples/input.txt --pretty`

Use `go run . -` for piped input. The program does not modify input files.
