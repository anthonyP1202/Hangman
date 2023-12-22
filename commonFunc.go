package Hangman

import (
	"math/rand"
)

func RndWord(hangpaused *HangManData, wordList []string) []rune { // pick a random word
	word := []rune(wordList[rune(rand.Intn(len(wordList)))])
	return word
}

func ShowWordp(word []rune, nbrOfLetterToShow int) ([]string, []rune) { // give the letter that are given with the word
	wordToFind := make([]string, len(word))
	for fillList := 0; fillList < len(word); fillList++ { // put _ for every letter
		wordToFind[fillList] = "_"
	}
	givenLetter := []rune{}
	for nbrOfLetterToShow != 0 {
		letterAlreadyRolld := 0
		letterChoosen := word[rand.Intn(len(word))]
		for i := 0; i < len(word); i++ {
			for checkThroughtAlready := 0; checkThroughtAlready < len(givenLetter); checkThroughtAlready++ {
				if letterChoosen == givenLetter[checkThroughtAlready] { // check that the letter wasn't already rolled as to not pick several times the same letter twice
					letterAlreadyRolld = 1
				}
			}
			if letterAlreadyRolld == 0 {
				f := 0
				nbrOfLetterToShow = nbrOfLetterToShow - 1
				for f < len(word) { // pass throught the list and replace the corresponding _ with the letter
					if word[f] == letterChoosen {
						wordToFind[f] = string(letterChoosen)
					}
					f++
				}
				givenLetter = append(givenLetter, letterChoosen)
			}
		}
	}
	return wordToFind, givenLetter
}
