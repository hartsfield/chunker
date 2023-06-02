package main

import (
	"strings"

	"golang.org/x/net/html"
)

func chunkToken(inToken string) (outTokens []string) {
	words := strings.Split(inToken, " ")
	// log.Println(inToken)
	for groupingSize := range words {
		if groupingSize > 0 {
			for i := range words {
				if groupingSize+i < len(words) {
					token := strings.Join(words[i:groupingSize+i], " ")
					outTokens = append(outTokens, token)
				}
			}
		}
	}
	outTokens = append(outTokens, inToken)
	return
}

func replacePunctuation(inTokens []string) (outTokens []string) {
	punctuation := []string{
		",", ":", ";", "(", ")", "!", `"`, "?", ".", "'", "--", "- ",
		" -"}

	outTokens = inTokens
	for _, punc := range punctuation {
		outTokens = findAndReplace(outTokens, punc, "")
	}
	return
}

func standardizeSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func removeShortStrings(inTokens []string, minLength int) (outTokens []string) {
	for _, token := range inTokens {
		if len(token) > minLength {
			outTokens = append(outTokens, token)
		}
	}
	return
}

func trimNonAlphaLeftRight(token string) string {
	if len(token) >= 3 {
		if !containsLetters(token[:1]) {
			token = token[1:]
			trimNonAlphaLeftRight(token)
		}
		if !containsLetters(token[len(token)-1:]) {
			token = token[:len(token)-1]
			trimNonAlphaLeftRight(token)
		}
	}

	return token
}

func stripHTMLTags(hyperText string) (textContent []string) {
	domDocTest := html.NewTokenizer(strings.NewReader(hyperText))
	previousStartTokenTest := domDocTest.Token()

loopDomTest:
	for {
		tt := domDocTest.Next()
		switch {
		case tt == html.ErrorToken:
			break loopDomTest
		case tt == html.StartTagToken:
			previousStartTokenTest = domDocTest.Token()
		case tt == html.TextToken:
			if previousStartTokenTest.Data == "script" {
				continue
			}
			if previousStartTokenTest.Data == "style" {
				continue
			}
			texttoken := strings.TrimSpace(html.UnescapeString(string(domDocTest.Text())))
			if len(texttoken) > 2 {
				if len(strings.Split(texttoken, "</div>")) < 4 {
					if strings.Count(texttoken, "displaystyle ") < 1 {
						textContent = append(textContent, texttoken)
					}
				}
			}

		}
	}
	return
}

func filterStrings(inTokens []string, filters ...string) (outTokens []string) {
	for _, filter := range filters {
		for _, token := range inTokens {
			// doesn't contain filtered term
			if !strings.Contains(token, filter) {
				outTokens = append(outTokens, token)
			}
		}
	}
	if len(outTokens) > 0 {
		return
	}
	outTokens = append(outTokens, "")
	return
}

func filterNonAlpha(inTokens []string) (alpha []string) {
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, token := range inTokens {
	testAlpha:
		for _, char := range alphas {
			if strings.ContainsRune(token, char) {
				alpha = append(alpha, token)
				break testAlpha
			}
		}
	}
	return
}

func containsLetters(token string) bool {
	alphas := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for _, char := range alphas {
		if strings.ContainsRune(token, char) {
			return true
		}
	}
	return false
}

func filterByLength(dirty string, maxLength, minLength int) bool {
	for _, line := range strings.Split(dirty, " ") {
		if len(line) < maxLength && len(line) > minLength {
			return true
		}
	}
	return false
}

func findAndReplace(inTokens []string, find, replace string) (outTokens []string) {
	for _, token := range inTokens {
		outTokens = append(outTokens, strings.ReplaceAll(token, find, replace))
	}
	return
}
