package shared

import (
	"os"
	"rest/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	if config.Mode.IsDevelopment() {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
