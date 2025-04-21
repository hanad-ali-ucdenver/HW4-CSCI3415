# Go Language Report - Sentiment Analysis Implementation

## Abstract

This report summarizes my experience with Go (Golang) during the implementation of a sentiment analysis program. Go is a statically typed, compiled language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It provides modern language features with a focus on simplicity, efficiency, and built-in support for concurrent programming. The language was designed to address the challenges of large-scale software development through easy-to-understand syntax, fast compilation, and robust standard libraries.

## Approach

### Problem Statement
I created a program that analyzes the sentiment of a text file based on a pre-defined dictionary of words and their associated sentiment scores. The program calculates an overall sentiment score and converts it to a star rating (1-5 stars).

### Design & Implementation
My implementation followed a modular approach with three main functions:

1. **buildSocialSentimentTable**: Reads a CSV file containing words and their sentiment scores, creating a map (dictionary) for quick lookups.

2. **getSocialSentimentScore**: Processes a text file word by word, looking up each word in the sentiment dictionary and accumulating a total score. Words not found in the dictionary are ignored.

3. **getStarRating**: Converts the numerical sentiment score to a star rating on a scale of 1 to 5.

### Algorithms
- The program uses a dictionary/hash map (Go's `map` type) for O(1) lookups of word sentiment scores
- Text processing involves tokenization (splitting into words) and normalization (converting to lowercase and removing punctuation)
- The star rating is determined using a simple range-based classification system

### Error Handling
I implemented comprehensive error handling throughout the code using Go's multiple return values pattern, returning both the expected result and an error value from functions. This approach ensures robustness and clear error reporting.

## New Things I Learned

### Unique/Special Features of Go

1. **Multiple Return Values**: Something I found really cool was that Go functions can return multiple values, which I found particularly useful for error handling. The pattern of returning both a result and an error is clean and effective.

   ```go
   func buildSocialSentimentTable(filename string) (map[string]float64, error) {
       // Function returns both a map (also known as a dictionary) and an error
   }
   ```

2. **Maps With Existence Checking**: Go's maps return both the value and a boolean indicating if the key exists, enabling cleaner code when checking for key presence.

   ```go
   score, exists := sentimentMap[word]
   if exists {
       // Use the score
   }
   ```

3. **Defer Statement**: The `defer` keyword allows scheduling function calls to execute after the surrounding function returns, which is perfect for resource cleanup like closing files.

   ```go
   defer file.Close() // File will be closed when the function exits
   ```

4. **Built-in Scanner Interface**: Go's `bufio` package provides elegant tools for reading input, with methods like `Split` to customize tokenization (e.g., reading word by word vs. reading line by line).

### Features I Liked

1. **Simple and Clear Syntax**: Go's syntax is straightforward and readable, making the code accessible even to those new to the language.

2. **Rich Standard Library**: The standard library contains comprehensive packages for common tasks like file I/O, CSV parsing, and string manipulation, reducing the need for external dependencies.

3. **Explicit Error Handling**: While verbose, Go's approach to error handling forces developers to consider error cases, leading to more robust code.

4. **Fast Compilation**:  Go's quick compilation makes the development cycle efficient.

5. **Strong Type System**: Static typing catches many errors at compile time rather than runtime.

### Features I Did Not Like

**Verbosity in Error Handling**: While explicit error handling is good for reliability, it can lead to repetitive code patterns and increased verbosity.

## Conclusion

Implementing the sentiment analysis program in Go was a positive experience. The language's simplicity and clarity made it easy to structure the program into logical components. Go's standard library provided all the necessary tools for file handling, text processing, and data structures. All in all, I think coding this in Go was a lot easier than if I had to code it in C.

The multiple return values and map functionality proved particularly valuable for this implementation, enabling clean error handling and efficient word-score lookups. While there are some trade-offs in terms of verbosity, the resulting code is robust and maintainable.

Go's focus on simplicity and practicality makes it well-suited for applications like this where performance, reliability, and maintainability are important considerations.

## Using Generative AI

I used Claude 3.7 Sonnet to help me with this assignment, as well as other QoL things like generating a readme.md file, and this .md for me based on my experiences and comments in my code. Claude was instrumental in helping me learn the syntax and general grammatical structure of Go, such as understanding their built in dictionary data strucure, maps, as well as Go's unique function structure that can return multiple values. Additionally, it recommended various libraries and functions that streamlined my development process and simplified my code drastically (I labeled these in my code). My views about AI from HW3 still hold, and I think actually strengthened. AI allowed me to learn, code, and develop much more efficiently than if I didn't use it at all.
