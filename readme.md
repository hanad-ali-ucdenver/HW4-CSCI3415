# Sentiment Analysis Star Rating Generator

This Go program analyzes text files to generate star ratings based on sentiment scores. It uses a dictionary of words with associated sentiment values to calculate an overall sentiment score, which is then converted into a 1-5 star rating.

## Features

- Reads and processes text files word by word
- Uses a pre-defined sentiment dictionary from a CSV file
- Removes punctuation and normalizes text to lowercase
- Calculates cumulative sentiment scores
- Converts sentiment scores to star ratings (1-5 stars)
- Provides detailed word-by-word scoring output

## Prerequisites

- Go (installed and configured)
- Input files:
  - `socialsent.csv`: Dictionary file containing words and their sentiment scores
  - Text file to analyze (default: `review.txt`)

## Usage

```bash
# Using default review.txt
go run hw4.go

# Using a custom review file
go run hw4.go <filename>
```

## How It Works

1. **Dictionary Building**: The program first reads `socialsent.csv` to create a sentiment dictionary where:
   - Each word is a key
   - Each sentiment score is stored as a float64 value

2. **Text Analysis**: The program then:
   - Reads the target text file word by word
   - Cleans each word (removes punctuation, converts to lowercase)
   - Looks up each word in the sentiment dictionary
   - Accumulates the sentiment scores

3. **Star Rating**: The final sentiment score is converted to a star rating based on these ranges:
   - 1 star: < -5.0
   - 2 stars: -5.0 to -1.0
   - 3 stars: -1.0 to 1.0
   - 4 stars: 1.0 to 5.0
   - 5 stars: ≥ 5.0

## Output

The program provides:
- Word-by-word sentiment scoring
- Running total of accumulated scores
- Final sentiment score for the file
- Star rating (1-5 stars)

## Error Handling

The program includes comprehensive error handling for:
- File opening/reading errors
- CSV parsing errors
- Invalid command line arguments
- Malformed input data

## Example Output

```
[word: current_score, accumulated score]
great: 1.20, 1.20
excellent: 1.50, 2.70
...

review.txt score: 2.70
review.txt Stars: 4
```

## Notes

- If no filename is provided, the program defaults to `review.txt`
- The program expects the sentiment dictionary file to be named `socialsent.csv`
- Words not found in the sentiment dictionary are skipped in the scoring