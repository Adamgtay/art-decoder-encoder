package art_interf

import (
	"fmt"
	"regexp"
	"strconv"
)

func sortBracketedAndNonBracketedStrings(inputSliceString []string) []string {

	var sliceString []string // sort into bracketed and non-bracketed strings

	for _, section := range inputSliceString {
		for i, char := range section {
			if i > 0 && char == '[' {
				sliceString = append(sliceString, section[:i])
				sliceString = append(sliceString, section[i:])
				break

			} else if char == '[' {
				sliceString = append(sliceString, section[i:])
				break
			}
		}
	}

	return sliceString

}

func useRegExToValidateData(sliceString []string) bool {
	// analyse each string in sliceString to validate structure
	// [5 #]  <-- regexp to match this where # can one or more of any character (including a space) and 5 can be one or more digits
	squareBracketRegExpPattern := `\[\d+\s.+?\]`

	squareBracketRegExpPatternCompile := regexp.MustCompile(squareBracketRegExpPattern)
	newLineCount := 1

	for _, data := range sliceString {
		if data[0] == '[' {
			validDataStructure := squareBracketRegExpPatternCompile.FindAllStringSubmatch(data, -1)
			if validDataStructure == nil {
				if len(sliceString) > 1 { // is multiline
					//errorData := data + "check line:" + strconv.Itoa(newLineCount)
					return true // isMalform is true
				} else {
					return true // isMalform is true
				}
			} else {
				continue
			}
		} else if data == "\n" {
			newLineCount += 1
		} else {
			continue
		}
	}
	return false
}

func readString(sliceString []string) (string, bool) {
	var output string
	// read each strin
	for _, data := range sliceString {
		if data[0] == '[' {
			// bracketed data
			var extractedDigits string
			var extractedSymbols string
			mandatorySpaceCount := false
			for _, char := range data {
				if char == '[' || char == ']' { // if brackets -> ignore
					continue
				} else if char >= '0' && char <= '9' && !mandatorySpaceCount {
					extractedDigits += string(char)
				} else if char >= '0' && char <= '9' && mandatorySpaceCount {
					extractedSymbols += string(char)
				} else if char == ' ' && !mandatorySpaceCount { // mandatory space -> ignore
					mandatorySpaceCount = true
					continue
				} else if char == ' ' && mandatorySpaceCount { // printed space
					extractedSymbols += " "
				} else {
					extractedSymbols += string(char)
				}
			}
			mandatorySpaceCount = false
			// method to convert extractedDigits into single integer
			extractedDigitsInteger, err := strconv.Atoi(extractedDigits)
			if err != nil {
				//PrintError(FORMAT_ERROR, extractedDigits)
				return output, true
			} else {
				for x := 0; x < extractedDigitsInteger; x++ {
					output += extractedSymbols
				}
			}
		} else {
			output += data // print unbracketed data
		}
	}
	return output, false
}

func isDuplicateSymbol(i int, line string) bool {
	if i == len(line)-1 {
		return line[i] == line[i-1]
	}
	return line[i] == line[i+1]
}

func ifDuplicateSymbol(i, lineNum, matchCount int, line string, splitStringFromArgs []string) (currentSymbol string, newMatchCount int) {
	currentSymbol = ""
	if i == len(line)-1 { // if end of current line
		if lineNum < len(splitStringFromArgs)-1 {
			matchCount += 1
			currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]\n")

			matchCount = 0
		} else { // if last line of input (exclude new line)
			matchCount += 1
			currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")

			matchCount = 0

		}
	} else {
		matchCount += 1
	}
	newMatchCount = matchCount

	return currentSymbol, newMatchCount
}

func endOfDuplicateSymbols(i, matchCount int, line string) string {
	currentSymbol := ""
	matchCount += 1
	currentSymbol = fmt.Sprint("[" + strconv.Itoa(matchCount) + " " + string(line[i]) + "]")

	return currentSymbol
}

func ifSingleSymbol(i, lineNum int, line string, splitStringFromArgs []string) string {
	currentSymbol := ""
	if i == len(line)-1 && lineNum < len(splitStringFromArgs)-1 { // if end of line and not last line of input
		currentSymbol = string(line[i]) + "\n" // <-- add newline
	} else {
		currentSymbol = string(line[i])
	}
	return currentSymbol
}
