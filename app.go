package taskr

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/rs/zerolog/log"
)

var (
	taskHeaderReg = regexp.MustCompile(`^task: \[([^\]]+)\] (.+)$`)
	taskOutputReg = regexp.MustCompile(`^\[([^\]]+)\] (.+)$`)
)

type App struct {
	OutRaw bool
}

func NewApp() *App {
	return &App{}
}

func (app *App) Run(ctx context.Context, src io.Reader) error {
	b, err := app.Parse(ctx, src)
	if err != nil {
		return err
	}

	fmt.Printf("%s", b)

	return nil
}

func (app *App) ParseStdin(ctx context.Context) ([]byte, error) {
	return app.Parse(ctx, os.Stdin)
}

func (app *App) Parse(_ context.Context, r io.Reader) ([]byte, error) {
	scanner := bufio.NewScanner(r)
	results := []Result{}
	var currentResult *Result

	for scanner.Scan() {
		line := scanner.Text()

		if app.OutRaw {
			fmt.Fprintf(os.Stderr, "%s\n", line)
		}

		headers := taskHeaderReg.FindAllStringSubmatch(line, -1)
		if 0 < len(headers) {
			if currentResult != nil {
				results = append(results, *currentResult)
			}

			currentResult = &Result{
				Name:   headers[0][1],
				Cmd:    headers[0][2],
				Output: "",
			}

			log.Debug().Any("result", currentResult).Msg("new result detected")

			continue
		}

		if currentResult == nil {
			continue
		}

		outputs := taskOutputReg.FindAllStringSubmatch(line, -1)
		if len(outputs) == 0 {
			continue
		}

		outTaskName := outputs[0][1]
		if outTaskName != currentResult.Name {
			log.Warn().
				Str("task", currentResult.Name).
				Str("output", outTaskName).
				Msg("task name mismatch")
			continue
		}

		currentResult.Output += outputs[0][2] + "\n"
	}

	if currentResult != nil {
		results = append(results, *currentResult)
	}

	ret, err := json.Marshal(results)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal results")
		return nil, err
	}

	return ret, nil
}

type Result struct {
	Name   string `json:"name"`
	Cmd    string `json:"cmd"`
	Output string `json:"output"`
}
