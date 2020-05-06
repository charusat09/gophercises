package main

import(
	"fmt"
	"encoding/csv"
	"os"
	"io"
	"strconv"
)

func main() {
	fmt.Println("How many questions you want?")
	var no_of_que string
	fmt.Scanln(&no_of_que)
	no_of_q, err := strconv.Atoi(no_of_que)
	fmt.Printf("User Input %s \n", no_of_que)
	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	counter := 0
	for i := 0; i < no_of_q; i++ {
		var ans string

		line, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Printf("%s? \n", line[:1][0])
		fmt.Scanln(&ans)
		if ans == line[1:][0] {
			counter += 1
		}
	}
	fmt.Printf("Final score is: %d \n", counter)
}