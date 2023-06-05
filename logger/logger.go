package logger

import (
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var Logger zerolog.Logger

func InitLogger(w io.Writer) {
	output := zerolog.ConsoleWriter{
		Out:        w,
		NoColor:    false,
		TimeFormat: time.ANSIC,
		FormatLevel: func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("[%s]", i))
		},
		FormatTimestamp: func(i interface{}) string {
			t, _ := time.Parse(time.RFC3339, fmt.Sprintf("%s", i))
			return fmt.Sprintf("%s", t.Format(time.RFC1123))
		},
	}
	Logger = zerolog.New(output).With().Timestamp().Logger()
}
