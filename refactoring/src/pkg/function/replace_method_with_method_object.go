package function

import "fmt"

type GuessStatisticsMessage struct {
	number         string
	verb           string
	pluralModifier string
}

func (gsm *GuessStatisticsMessage) noLetter() {
	gsm.number = "no"
	gsm.verb = "are"
	gsm.pluralModifier = "s"
}

func (gsm *GuessStatisticsMessage) oneLetter() {
	gsm.number = "l"
	gsm.verb = "is"
	gsm.pluralModifier = ""
}

func (gsm *GuessStatisticsMessage) manyLetters(count int) {
	gsm.number = string(count)
	gsm.verb = "are"
	gsm.pluralModifier = "s"
}

func (gsm *GuessStatisticsMessage) buildParams(count int) {
	if count == 0 {
		gsm.noLetter()
	} else if count == 1 {
		gsm.oneLetter()
	} else {
		gsm.manyLetters(count)
	}
}

func (gsm *GuessStatisticsMessage) printGuessStatistics(candidate string, count int) {
	gsm.buildParams(count)
	fmt.Printf("There %s %s %s %s\n", gsm.verb, gsm.number, candidate, gsm.pluralModifier)
}

// --------------原始函数-----------------------
func printGuessStatisticsOri(candidate string, count int) {
	// 重构下面这一段
	var number string
	var verb string
	var pluralModifier string
	if count == 0 {
		number = "no"
		verb = "are"
		pluralModifier = "s"
	} else if count == 1 {
		number = "l"
		verb = "is"
		pluralModifier = ""
	} else {
		number = string(count)
		verb = "are"
		pluralModifier = "s"
	}
	fmt.Printf("There %s %s %s %s\n", verb, number, candidate, pluralModifier)
}

func TestReplaceMethodWithMethodObj() {
	printGuessStatisticsOri("abc", 4)
	gsm := GuessStatisticsMessage{}
	gsm.printGuessStatistics("abc", 3)
}
