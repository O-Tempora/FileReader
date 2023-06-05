package readers

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

func ReadCsv(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Failed reading file")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Failed reading file")
		return
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 4, 3, ' ', 0)
	for _, v := range records {
		fmt.Fprintln(writer, strings.Join(v, "\t")+"\t")
	}
	writer.Flush()
}
