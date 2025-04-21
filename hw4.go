package main

import (
	"bufio"        // library to scan file word by word
	"encoding/csv" // library to read in csv files
	"fmt"          // standard go library to print stuff
	"io"           // library for basic i/o funcion
	"os"           // library so i can use command line arguments
	"strconv"      // library to convert strings to float65, used for the sentiment score
	"strings"      // library to format and manipulate strings
)

func main() {

	filename := "review.txt" // default name for file

	// switch statement to check command line arguments, print an error if there are too many, and use default name if none are provided
	switch len(os.Args) {
	case 1:

	case 2:
		filename = os.Args[1]
	default:
		fmt.Println("Error: Too many arguments")
		fmt.Println("Usage: program <filename>")
		fmt.Println("	If a filename is not provided, 'review.txt' will be used")
		os.Exit(1) // exit program
	}

	// calls my function that reads a csv and creates a dictionary based on the word and sentiment score
	socialSentimentMap, err := buildSocialSentimentTable("socialsent.csv")
	if err != nil { // error checking if function fails
		fmt.Println("Error:", err)
		return
	}

	// calling my function that uses the dictionary and the target file to calculate the sentiment score of the file
	socialSentimentScore, err := getSocialSentimentScore(filename, socialSentimentMap)
	if err != nil { // error checking
		fmt.Println("Error:", err)
		return
	}

	// calling my function that uses the sentiment score and calculates a star rating for the target file
	starRating := getStarRating(socialSentimentScore)

	fmt.Printf("%s Stars: %d\n", filename, starRating) // print star rating

}

/* function uses the filename and sentimentMap dictionary as parameters, opens the file, reads it word by word,
   and finds the score of each word by looking it up in the dictionary, and keeps a running total of the total score
   of the file. prints the score of each word and the accumulated score, and at the end prints the total score
*/

/*
go is cool because you can return multiple variables, and declare the datatype of the variables you are returning

	returning the sentimentscore, and an error datatype which should be nil if everything works, and if not will return
	the error.
*/
func getSocialSentimentScore(filename string, sentimentMap map[string]float64) (float64, error) {

	var totalScore float64 = 0 // initializing total score

	file, err := os.Open(filename) //open file
	if err != nil {                // error checking
		return 0, fmt.Errorf("Error opening file: %w", err)
	}

	defer file.Close() //defer waits until the whole function is completed, then executes. this makes sure the file is closed no matter what

	scanner := bufio.NewScanner(file) //NewScanner is a function that scans a file
	scanner.Split(bufio.ScanWords)    // telling the scanner to scan by word not by line

	fmt.Println("[word: current_score, accumulated score]")

	// this loop will scan a word, and will continue scanning new wordsuntil EOF
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(strings.Trim(word, ".,!?:;\"()[]{}-")) // nice function Claude told me about to clean up any punctuation and set word to lowercase

		//if formatting results in an empty word, skip it
		if word == "" {
			continue
		}

		/* another cool go function is that it's built in dictionaries, maps, can return 2 values, the value, and a boolean that
		   checks if the given key exists in the dictionary. i used it here.
		   gets the score for the current word by looking it up in the dictionary.
		*/
		score, exists := sentimentMap[word]

		// if the word exists in the dictionary, then add the score of the word to the totalScore,
		// print the score of the word, and the accumulating score.
		if exists {
			totalScore += score
			fmt.Printf("%s: %.2f, %.2f\n", word, score, totalScore)
		}

	}

	// error handling
	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("Error reading file: %w", err)
	}

	// print the total score of the file
	fmt.Printf("\n%s score: %.2f\n", filename, totalScore)

	return totalScore, nil
}

/*
Function for building the dictionary using the .csv file containing the word and sentiment score.

	Takes a string in as a parameter, and returns the dictionary, known as a map in Go, and a error datatype that
	defaults to nil if everything works well.
*/
func buildSocialSentimentTable(filename string) (map[string]float64, error) {

	// initialize dictionary, known as map in go
	socialSentimentTable := make(map[string]float64)

	// open the file
	file, err := os.Open(filename)
	if err != nil { // error checking
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close() // once function exits, either successfully or with an error, close file

	// NewReader is a function that reads a CSV into memory
	reader := csv.NewReader(file)

	// read in header file, and do nothing, basically skips it
	_, err = reader.Read()
	if err != nil { // error checking
		return nil, fmt.Errorf("Error reading header: %w", err)
	}

	// loop that reads the .csv file until the EOF is reached, then it exits
	for {
		record, err := reader.Read() // read in next row
		if err == io.EOF {           // if EOF exit
			break
		}
		if err != nil { // error checking
			return nil, fmt.Errorf("Error reading record: %w", err)
		}

		// checks if current row has word and sentiment score, if not it skips it
		if len(record) < 2 {
			continue
		}

		// sets word to the word value in the csv
		word := record[0]

		// converts the score value to a floating point value, and sets it equal to score
		score, err := strconv.ParseFloat(record[1], 64)
		if err != nil { // if it fails, skip it
			continue
		}

		//sets the map key to word, and the value to score
		socialSentimentTable[word] = score
	}

	// returns the map, and nil if successful
	return socialSentimentTable, nil
}

// function that uses the sentiment score and calculates the star rating. returns the star rating
func getStarRating(totalScore float64) int {
	var stars int // initialize stars variable

	// switch statement that checks the value of the totalScore and find the corresponding star value
	switch {
	case totalScore < -5.0:
		stars = 1
	case totalScore >= -5.0 && totalScore < -1.0:
		stars = 2
	case totalScore >= -1.0 && totalScore < 1.0:
		stars = 3
	case totalScore >= 1.0 && totalScore < 5.0:
		stars = 4
	case totalScore >= 5.0:
		stars = 5
	}

	return stars
}
