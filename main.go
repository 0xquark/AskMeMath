package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	Filepath := flag.String("csv", "problems.csv", "csv file in format : 'question,answer'")
	flag.Parse()

	file, err := os.Open(*Filepath)
	if err != nil {
		fmt.Printf("Couldn't open csv file: %s\n", *Filepath)
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Printf("Not able to parse the csv file")
	}
	problems := parseLines(lines)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.ans {
			correct++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))

}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  line[1],
		}
	}
	return ret
}

type problem struct {
	ques string
	ans  string
}
