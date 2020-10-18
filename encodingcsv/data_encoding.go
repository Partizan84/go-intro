package encodingcsv

import (
	"encoding/csv"
	"fmt"
	"os"
)
//EncodingCsv функция кодировки.
func EncodingCsv() {
 
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
 
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'
 
	for {
		record, e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}