package cli

import (
	"flag"
	"fmt"

	"github.com/handlename/task-result/internal/errorcode"
	"github.com/morikuni/failure/v2"
)

type Flags struct {
	Version  bool
	LogLevel string
	OutRaw   bool

	// Source is path to source file or "-"
	Source string
}

func parseFlags(appname string, args []string) (*Flags, error) {
	flags := &Flags{}

	fs := flag.NewFlagSet(appname, flag.ExitOnError)
	fs.Usage = func() {
		u := `Usage of %s:
%s [flags...] source_path
  source_path
        Path to output of task. If set "-", it reads stdin.
`
		fmt.Fprintf(flag.CommandLine.Output(), u, appname, appname)
		fs.PrintDefaults()
	}

	fs.BoolVar(&flags.Version, "version", false, "Print version")
	fs.StringVar(&flags.LogLevel, "log-level", "info", "Log level (trace, debug, info, warn, error, panic)")
	fs.BoolVar(&flags.OutRaw, "out-raw", false, "Output raw input to stderr")

	if err := fs.Parse(args); err != nil {
		return nil, failure.Wrap(err, failure.Message("failed to parse flags"))
	}

	if flags.Version {
		return flags, nil
	}

	rest := fs.Args()
	if len(rest) != 1 {
		return nil, failure.New(errorcode.ErrInvalidArgument, failure.Message(`only one source path is required`))
	}

	flags.Source = rest[0]

	return flags, nil
}
