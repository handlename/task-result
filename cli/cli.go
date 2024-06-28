package cli

import (
	"context"
	"errors"
	"os"
	"os/signal"

	taskr "github.com/handlename/task-result"
	"github.com/rs/zerolog/log"
)

type ExitCode int

const (
	ExitCodeOK    ExitCode = 0
	ExitCodeError ExitCode = 1
)

func Run() ExitCode {
	flags, err := parseFlags(os.Args[0], os.Args[1:])
	if err != nil {
		log.Error().Err(err).Msg("failed to parse flags")
		return ExitCodeError
	}

	taskr.InitLogger(flags.LogLevel)

	if flags.Version {
		log.Info().Msgf("task-result v%s", taskr.Version)
		return ExitCodeOK
	}

	app := taskr.NewApp()
	app.OutRaw = flags.OutRaw

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	if err := app.Run(ctx); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Error().Msg("canceled")
		} else {
			log.Error().Stack().Err(err).Msg("")
		}

		return ExitCodeError
	}

	return ExitCodeOK
}
