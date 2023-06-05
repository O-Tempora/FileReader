package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"freader/logger"
	readers "freader/pkg"

	"github.com/spf13/cobra"
)

var (
	file         string
	infoRequired bool
)

var readCmd = &cobra.Command{
	Use:   "read",
	Short: "Reads file content",
	Long:  `Reads file content and displays it in terminal`,
	Run: func(cmd *cobra.Command, args []string) {
		switch filepath.Ext(file) {
		case ".csv":
			readers.ReadCsv(file)
		case ".pdf":
			readers.ReadPdf(file)
		default:
			readers.ReadPlane(file)
		}

		if infoRequired {
			abs, err := filepath.Abs(file)
			if err != nil {
				logger.Logger.Err(err).Msg("Absolute path error")
				return
			}
			f, err := os.Open(file)
			if err != nil {
				logger.Logger.Err(err).Msg("File open error")
			}
			defer f.Close()
			stat, err := f.Stat()
			if err != nil {
				logger.Logger.Err(err).Msg("File info error")
			}

			fmt.Println("\n\n\tFile info:",
				"\nFile name - ", filepath.Base(file),
				"\nAbsolute path - ", abs,
				"\nModification time - ", stat.ModTime(),
				"\nFile size - ", stat.Size(), "bytes",
			)
		}

		logger.Logger.Info().Msgf("File read: %s; detailed: %t", filepath.Base(file), infoRequired)
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().StringVarP(&file, "path", "p", "", "Path to file that needs to be read")
	if err := readCmd.MarkFlagRequired("path"); err != nil {
		logger.Logger.Fatal().Msg("Missing required flag \"p\"")
	}

	readCmd.Flags().BoolVarP(&infoRequired, "detailed", "d", false, "Additional information about file")
}
