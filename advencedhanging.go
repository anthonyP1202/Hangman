package Hangman

import (
	"bufio"
	"fmt"
	"os"
)

type HangManData struct {
	Word             []rune
	WordToFind       []string
	GivenLetter      []rune
	GivenWord        []string
	NbrOfAttempt     int
	VictoryCondition int
}

func HangmanADV(wordslist string) *HangManData { //return the struct
	var hangadv HangManData
	listOfWord := []string{}
	wordListFile, err := os.Open(wordslist)
	if err != nil {
		fmt.Println("error with the file")
		os.Exit(1)
	} else {
		fileScanner := bufio.NewScanner(wordListFile)

		fileScanner.Split(bufio.ScanLines)

		for fileScanner.Scan() {
			listOfWord = append(listOfWord, fileScanner.Text())
		}
		hangadv.Word = RndWord(&hangadv, listOfWord)
		shownLetter := len(hangadv.Word)/2 - 1
		hangadv.WordToFind, hangadv.GivenLetter = ShowWordp(hangadv.Word, shownLetter)
		hangadv.NbrOfAttempt = 10
		hangadv.VictoryCondition = 0
	}
	return &hangadv
}

func Playingadv(hangadv *HangManData, letter []rune) *HangManData {
	hangadv.VictoryCondition = 0
	for loop := 0; loop < len(hangadv.WordToFind); loop++ {
		fmt.Print(hangadv.WordToFind[loop])
	}
	fmt.Print("\n")
	if len(letter) > 1 {
		condition := 0
		alreadyGot := 0
		if 1 < len(letter) && len(letter) == len(hangadv.Word) {
			for b := 0; b < len(hangadv.GivenWord); b++ {
				if string(letter) == hangadv.GivenWord[b] {
					alreadyGot = 1
				}
			}
			if alreadyGot == 1 {
				fmt.Println("already attempted word")
				fmt.Println(hangadv.GivenWord)
			} else {
				hangadv.GivenWord = append(hangadv.GivenWord, string(letter))
				for i := 0; i < len(hangadv.Word); i++ {
					if letter[i] != hangadv.Word[i] {
						condition = 1
					}
				}
				if condition == 1 {
					hangadv.NbrOfAttempt = hangadv.NbrOfAttempt - 2
				} else {
					hangadv.VictoryCondition = 1
				}
			}
		} else {
			fmt.Println("attempted word cannot be seeked word (different lenght)")
		}
	} else {
		condition := 0
		for o := 0; o < len(hangadv.GivenLetter); o++ {
			if letter[0] == hangadv.GivenLetter[o] {
				condition = 1
			}
		}
		if condition == 0 {
			hangadv.GivenLetter = append(hangadv.GivenLetter, letter[0])
			wrong := 1
			for i := 0; i < len(hangadv.Word); i++ {
				if letter[0] == hangadv.Word[i] {
					hangadv.WordToFind[i] = string(letter[0])
					wrong = 0
				}
			}
			if wrong == 1 {
				hangadv.NbrOfAttempt = hangadv.NbrOfAttempt - 1
			}
		} else {
			fmt.Println("letter already tested")
		}
	}
	if hangadv.VictoryCondition == 0 {
		// fmt.Println("number of attempt left : ", nbrOfAttempt)
		hangadv.VictoryCondition = 1 // check if there still is a _ however if a word contain a _ it get block it isn't an issue yet
		for f := 0; f < len(hangadv.WordToFind); f++ {
			if hangadv.WordToFind[f] == "_" {
				hangadv.VictoryCondition = 0
			}
		}
	}
	return hangadv
	// if hangadv.NbrOfAttempt == 0 && hangadv.VictoryCondition == 0 {
	// 	fmt.Println("you lost the word was", string(hangadv.Word))
	// } else {
	// 	fmt.Print("you hangadv.VictoryCondition the word was ")
	// 	for v := 0; v < len(hangadv.Word); v++ {
	// 		fmt.Print(string(hangadv.Word[v]))
	// 	}
	// 	fmt.Print("\n")
	// }
}
