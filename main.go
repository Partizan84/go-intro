package main

import (
	"go-intro/comment"
	"fmt"
	"os"
	"log"
	"go-intro/encodingcsv"
)

func openFile(fileName string) (file *os.File, err error) {
    file, err = os.Open(fileName)
        if err != nil {
           log.Fatal(err)
           return
        }
    return 
}
// getting size of file
func readFile(file *os.File) []byte {
		stat, err := file.Stat()
		if err != nil {
			  log.Fatal(err)
			  return nil
		}
		// reading file
		bs := make([]byte, stat.Size())
		_, err = file.Read(bs)
		if err != nil {
			log.Fatal(err)
			return nil 
		}
		return bs
	}
func fileContent (fileName string) string {
	file, _ := openFile(fileName)
		defer file.Close()
	    return string(readFile(file))
}
		
func main() {
	comment.Commenttaskone()
	fmt.Println(fileContent("main.go"))
	encodingcsv.EncodingCsv()
}
