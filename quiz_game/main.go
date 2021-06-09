package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	normalAnswerTime = 10
	maxAnswerTime    = 15
)

type dbData struct {
	Question string    `db:"question"`
	Answer   string    `db:"answer"`
	Times    int64     `db:"times"`
	LastDate time.Time `db:"last_date"`
	NextDate time.Time `db:"next_date"`
}

func readInput(scanner bufio.Scanner, input chan<- string) {
	for {
		scanner.Scan()
		line := scanner.Text()
		input <- line
	}
}

// Find takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getDB(table string) []*dbData {
	var result []*dbData
	dbPass := os.Getenv("db_pass")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://localhost/quizes?user=postgres&password=%s&sslmode=disable", dbPass))
	if err != nil {
		fmt.Println("error:", err)
		return result
	}
	sql_query := fmt.Sprintf(`select question
									,answer
									,times
									,last_date
									,next_date 
									from public.%s 
									where times=1 or next_date<='%s'`,
		table, time.Now().Format("2006-01-02 15:04:05"))
	err = db.Select(&result, sql_query)
	return result
}

func updateDB(table string, payload *dbData, factor int) {
	var result dbData
	dbPass := os.Getenv("db_pass")
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgresql://localhost/quizes?user=postgres&password=%s&sslmode=disable", dbPass))
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = db.Get(&result, fmt.Sprintf(`select question
						,answer
						,times
						,last_date
						,next_date 
						from public.%s 
						where question=%s`, table, payload.Question))

	if factor == 1 {
		payload.Times = payload.Times * 2
	} else if factor == 2 {
		payload.Times = int64(float64(payload.Times) * 1.5)
	} else {
		payload.Times = 1
	}

	db.MustExec(fmt.Sprintf(`
		update public.%s
			set times=%v
			,last_date=current_date
			,next_date=current_date+ interval '%v' day
		where question='%s'`, table, payload.Times, payload.Times, payload.Question))
}

func main() {
	//parse command line
	var table string
	var dictSize int
	flag.StringVar(&table, "table", "english_idioms", "table in quizes database, schema public")
	flag.IntVar(&dictSize, "number", 20, "number of questions to answer")
	flag.Parse()
	//get data
	payload := getDB(table)
	//shuffle data
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(payload), func(i, j int) { payload[i], payload[j] = payload[j], payload[i] })

	//create arrays for swapping
	start_questions := make([]*dbData, 0)
	answered_questions := make([]string, 0)
	for _, row := range payload[:dictSize] {
		start_questions = append(start_questions, row)
	}
	wrong := 0
	right := 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	//create a channel with user input
	userInput := make(chan string)
	//create a reader to read input
	go readInput(*scanner, userInput)

	//Start quiz
	for len(answered_questions) < dictSize {
		for _, turn := range start_questions {
			if _, ok := Find(answered_questions, turn.Question); !ok {
				fmt.Print(turn.Question, "-> ")
				start := time.Now()
				factor := 1
				payload := dbData{
					Question: turn.Question,
					Answer:   turn.Answer,
					LastDate: turn.LastDate,
					NextDate: turn.NextDate,
				}
				select {
				case userAnswer := <-userInput:
					if userAnswer == turn.Answer {
						fmt.Println("Correct!")
						right++
						if time.Since(start) > time.Second*normalAnswerTime {
							factor = 2

						} else {
							factor = 1
						}
						updateDB(table, &payload, factor)
						answered_questions = append(answered_questions, turn.Question)
					} else {
						fmt.Println("Wrong! your answer is:", userAnswer, ", right is: ", turn.Answer)
						wrong++
						factor = 3
						updateDB(table, &payload, factor)
					}
				case <-time.After(maxAnswerTime * time.Second):
					fmt.Println("\n Time is over!")
				}

			}
		}

	}

	fmt.Println("Your score is:", float64(right)/float64(right+wrong)*100, ", right=", right, ", wrong=", wrong)
}
