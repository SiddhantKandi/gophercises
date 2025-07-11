package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main(){
	csvFilename := flag.String("csv","problems.csv","Reading the csv file in format : Question,Answer")

	timeLimit := flag.Int("limit",30,"time limit for the quiz in seconds")

	flag.Parse()

	file,err := os.Open(*csvFilename)

	if err != nil{
		exit(fmt.Sprintf("Unable to open the csv file %s\n",*csvFilename))
	}

	r := csv.NewReader(file)

	lines,err := r.ReadAll()

	if err != nil {
		exit("Failed to read the csv file")
	}

	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct:=0

	//Another way to do this is labels
	//problemloop:

	for i,problem := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1,problem.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)

			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou scored %d out of %d questions \n", correct,len(problems))
			return
			//break problemloop
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}
		
		
	}

	fmt.Printf("You scored %d out of %d questions \n", correct,len(problems))


	
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem,len(lines))

	for i,line := range lines {
		ret[i] = problem{
			q : line[0],
			a : strings.TrimSpace(line[1]),
		}
	}

	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}