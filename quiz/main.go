package main

import (
	"flag"
	"os"
	"fmt"
	"csv"
)

func main () {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse ()
	//_ = csvFilename

	file, err := os.Open(*csvFilename)
		if err != nil {
			exit(fmt.Sprintf("Failed to open the CSV file: %s/n", *csvFilename))
			os.Exit(1)
		}
		
		r := csv.NewReader(file)
		lines, err : = r.ReadAll()
}

func exit ( msg string ) {
	fmt.Println(msg)
	os.Exit(1)
}