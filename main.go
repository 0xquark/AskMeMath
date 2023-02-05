package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
	"bufio"
	"strings"
)

func main() {
	Filepath := flag.String("csv", "problems.csv", "csv file in format : 'question,answer'")
	timeout := flag.Int("Limit", 30, "Time limit : In seconds")
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

	timer := time.NewTimer(time.Duration(*timeout) * time.Second)
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, p.ques)
		answerCh := make(chan string)
		go func() {
			var answer string
			reader := bufio.NewReader(os.Stdin)
            answer, _ = reader.ReadString('\n')
            answer = strings.TrimSpace(answer)

			answerCh <- answer
		}()
		select {
		case <-timer.C: // Waiting for message from the channel
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return

		case answer := <-answerCh:
			if answer == p.ans {
				correct++
			}
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
	percentage := float64(correct) / float64(len(problems)) * 100
    if percentage > 70 {
	fmt.Println("Congratulations! You have scored more than 70%.")
	}
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
