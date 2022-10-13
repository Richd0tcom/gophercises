package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

//were building a cli quiz app that reads a csv file and runs th quizes inside
func main(){
	csvFileName := flag.String("csv", "questions.csv", "a csv file with the format 'questions,answers'")
	//csvFilename is a pointer to a string (csv file name) that will be used
	flag.Parse()
	
	file, err := os.Open(*csvFileName)
	
	if err != nil {
		exit(fmt.Sprintf("Failed to open file: %v", *csvFileName,))
	}
	_ = file
	
	
	r := csv.NewReader(file)
	lines, errr := r.ReadAll()
	if errr != nil {
		exit("parse the provided csv")
	}
	fmt.Println(lines)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
