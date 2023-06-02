package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("all.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	initialTokens := equalizeWordCount(lines, 30)
	var chunks []string
	for _, token := range initialTokens {
		chunks = append(chunks, chunkToken(token)...)
	}
	addToUnsortedWordCloud(chunks)
	sortCloud()
	for _, token := range wordCloud {
		if token.Rank > 1 && len(token.Token) > 10 && !(strings.Contains(token.Token, "gutenberg")) {
			fmt.Println(token.Rank, token.Token)
		}
	}
}

func washTokens(inTokens []string) (outTokens []string) {
	outTokens = replacePunctuation(inTokens)
	outTokens = filterNonAlpha(outTokens)
	outTokens = removeShortStrings(outTokens, 3)
	return
}

func equalizeWordCount(inTokens []string, tokenLength int) (outTokens []string) {
	rinsed := washTokens(inTokens)
	rinsedString := standardizeSpaces(strings.Join(rinsed, " "))
	spaceCount := strings.Count(rinsedString, " ")
	var token string
	var ss []string
	for i := 0; i < spaceCount; i++ {
		spl := strings.SplitN(rinsedString, " ", 2)
		ss = append(ss, spl[0])
		rinsedString = spl[1]
		if i%tokenLength == 0 {
			token = strings.Join(ss, " ")
			outTokens = append(outTokens, token)
			ss = []string{}
		}
	}
	return
}
