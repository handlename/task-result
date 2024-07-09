package taskr

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppParse(t *testing.T) {
	want := []Result{
		{
			Name:   "echo",
			Cmd:    `echo "Hello, World!"`,
			Output: "Hello, World!\n",
		},
		{
			Name: "cal",
			Cmd:  "cal",
			Output: strings.Join([]string{
				"     June 2024        ",
				"Su Mo Tu We Th Fr Sa  ",
				"                   1  ",
				" 2  3  4  5  6  7  8  ",
				" 9 10 11 12 13 14 15  ",
				"16 17 18 19 20 21 22  ",
				"23 24 25 26 27 28 29  ",
				"30                    ",
				"",
			}, "\n"),
		},
		{
			Name: "ls",
			Cmd:  "ls -la",
			Output: strings.Join([]string{
				"total 16",
				"drwxr-xr-x@ 4 nagata-hiroaki  staff  128 Jun 28 17:12 .",
				"drwxr-xr-x@ 3 nagata-hiroaki  staff   96 Jun 28 17:05 ..",
				"-rw-r--r--  1 nagata-hiroaki  staff  226 Jun 28 17:10 Taskfile.yaml",
				"-rw-r--r--@ 1 nagata-hiroaki  staff  321 Jun 28 17:12 out.txt",
				"",
			}, "\n"),
		},
	}

	out, err := os.Open("testdata/app/parse/out.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	ctx := context.Background()
	app := NewApp()
	b, err := app.Parse(ctx, out)
	if err != nil {
		t.Fatal(err)
	}

	results := []Result{}
	if err := json.Unmarshal(b, &results); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, want, results)
}
