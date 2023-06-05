package readers

import (
	"fmt"

	"github.com/dslipak/pdf"
)

func ReadPdf(file string) {
	r, err := pdf.Open(file)
	if err != nil {
		panic(err.Error())
	}
	total := r.NumPage()
	for i := 1; i <= total; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}
		rows, _ := page.GetTextByRow()
		for _, row := range rows {
			for _, word := range row.Content {
				fmt.Println(word.S)
			}
		}
	}
}
