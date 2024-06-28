package cli

import (
	"flag"

	"github.com/morikuni/failure/v2"
)

type Flags struct {
	Version  bool
	LogLevel string
	OutRaw   bool
}

func parseFlags(appname string, args []string) (*Flags, error) {
	flags := &Flags{}

	fs := flag.NewFlagSet(appname, flag.ExitOnError)

	fs.BoolVar(&flags.Version, "version", false, "Print version")
	fs.StringVar(&flags.LogLevel, "log-level", "info", "Log level (trace, debug, info, warn, error, panic)")
	fs.BoolVar(&flags.OutRaw, "out-raw", false, "Output raw input to stderr")

	if err := fs.Parse(args); err != nil {
		return nil, failure.Wrap(err, failure.Message("failed to parse flags"))
	}

	return flags, nil
}
