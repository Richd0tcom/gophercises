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
		exit("Failed to parse the provided csv")
	}
	problems := parseLines(lines)

	correctCounter := 0
	for i, prob := range problems {
		fmt.Printf("Quiz #%d: %s = \n", i+1, prob.q)
		var answer string

		fmt.Scanf("%s\n", &answer)
		if answer == prob.a {
			fmt.Println("Correct")
			correctCounter++
		}
	}
	fmt.Printf("You scored %d out of %d \n", correctCounter, len(problems))
}


type Problem struct {
	q string
	a string

}
func parseLines(lines [][]string) []Problem {
	res := make([]Problem, len(lines))

	for i, line := range lines {
		res[i] = Problem{
			q: line[0],
			a: line[1],
		}
	}
	return res
}


func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
