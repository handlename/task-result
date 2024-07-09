# task-result

task-result (taskr) is output parser for [task](https://taskfile.dev/).
It outputs the result of the `task` in json format that is easy to use in other tools.

## Synopsis

```console
$ go run taskr -help
Usage of taskr:
taskr [flags...] source_path
  source_path
        Path to output of task. If set "-", it reads stdin.
  -log-level string
        Log level (trace, debug, info, warn, error, panic) (default "info")
  -out-raw
        Output raw input to stderr
  -version
        Print version

$ task -t testdata/app/parse/Taskfile.yaml 2>&1 | taskr -out-raw - | jq -S
task: [echo] echo "Hello, World!"
[echo] Hello, World!
task: [cal] cal
[cal]      July 2024
[cal] Su Mo Tu We Th Fr Sa
[cal]     1  2  3  4  5  6
[cal]  7  8  9 10 11 12 13
[cal] 14 15 16 17 18 19 20
[cal] 21 22 23 24 25 26 27
[cal] 28 29 30 31
[cal]
task: [ls] ls -la
[ls] total 16
[ls] drwxr-xr-x@ 4 nagata-hiroaki  staff  128 Jul  5 09:54 .
[ls] drwxr-xr-x@ 3 nagata-hiroaki  staff   96 Jul  5 09:54 ..
[ls] -rw-r--r--@ 1 nagata-hiroaki  staff  226 Jul  5 09:54 Taskfile.yaml
[ls] -rw-r--r--@ 1 nagata-hiroaki  staff  598 Jul  5 09:54 out.txt
[
  {
    "cmd": "echo \"Hello, World!\"",
    "name": "echo",
    "output": "Hello, World!\n"
  },
  {
    "cmd": "cal",
    "name": "cal",
    "output": "     July 2024        \nSu Mo Tu We Th Fr Sa  \n    1  2  3  4  5  6  \n 7  8  9 10 11 12 13  \n14 15 16 17 18 19 20  \n21 22 23 24 25 26 27  \n28 29 30 31           \n                      \n"
  },
  {
    "cmd": "ls -la",
    "name": "ls",
    "output": "total 16\ndrwxr-xr-x@ 4 nagata-hiroaki  staff  128 Jul  5 09:54 .\ndrwxr-xr-x@ 3 nagata-hiroaki  staff   96 Jul  5 09:54 ..\n-rw-r--r--@ 1 nagata-hiroaki  staff  226 Jul  5 09:54 Taskfile.yaml\n-rw-r--r--@ 1 nagata-hiroaki  staff  598 Jul  5 09:54 out.txt\n"
  }
]
```

## Installation

```console
$ go get github.com/handlename/task-result/cmd/task-result@latest
$ taskr ...
```

or

Download latest release binary from [Releases](https://github.com/handlename/task-result/releases)

## Lisence

MIT

## Author

@handlename
