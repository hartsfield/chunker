package main

import (
	"sort"
	"strings"
)

var (
	unsortedWordCloud map[string]int = make(map[string]int)
	wordCloud         []*rankedToken
)

type rankedToken struct {
	Token string
	Rank  int
}

func sortCloud() {
	for k, v := range unsortedWordCloud {
		wordCloud = append(wordCloud, &rankedToken{k, v})
	}

	sort.Slice(wordCloud, func(i, j int) bool {
		return wordCloud[i].Rank < wordCloud[j].Rank
	})
}

func addToUnsortedWordCloud(tokens []string) {
	tokens = filterStrings(tokens, "GMT+0000 (UTC)")
	tokens = filterNonAlpha(tokens)
	for _, token := range tokens {
		if len(token) > 2 {
			token = strings.ToLower(token)
			if unsortedWordCloud[token] == 0 {
				unsortedWordCloud[token] = 1
			} else {
				unsortedWordCloud[token] = unsortedWordCloud[token] + 1
			}
		}
	}
}
