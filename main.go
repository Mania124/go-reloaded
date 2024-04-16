package main

import (
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	// execute commands based on keyword
	JSlice := Processin(SplitWhiteSpaces(readFile()))
	// carry out the first punctuantion corrections
	jsLice := puncTuation(JSlice)
	// refine punctuations
	fsLice := finalPunc(jsLice)
	var sLice []string
	for _, word := range fsLice {
		if (word[0] == ',' || word[0] == '.') && len(word) > 3 {
			sLice = append(sLice, punc(word))
		} else {
			sLice = append(sLice, word)
		}
	}
	// finalize punctuations
	fnS := CheckQuotation(sLice)
	var mainString string
	for i, r := range fnS {
		if i < len(sLice) {
			mainString += r
		}
		mainString += " "
	}
	createNewFile(mainString)
}

func puncTuation(str []string) []string {
	// this function takes a slice of strings and iterates through its members
	// it finally returns the corrected slice
	// check and correct initial punctuations
	for i := 0; i < len(str); i++ {
		if i == 0 || str[i-1][len(str[i-1])-1] == '.' {
			// check need for change in "A" and "An" in the begining of slice members
			if (str[i] == "a" || str[i] == "An" || str[i] == "an" || str[i] == "A") && ((str[i+1][0] >= 'a' && str[i+1][0] <= 'z') || (str[i+1][0] >= 'A' && str[i+1][0] <= 'Z')) {
				if str[i+1][0] == 'a' || str[i+1][0] == 'e' || str[i+1][0] == 'i' || str[i+1][0] == 'o' || str[i+1][0] == 'u' || str[i+1][0] == 'h' {
					str[i] = "An"
				} else {
					str[i] = "A"
				}
			}
		}
		// check need for r or an in the middle of a string
		if i > 0 {
			if (str[i] == "a" || str[i] == "An" || str[i] == "an" || str[i] == "A") && ((str[i+1][0] >= 'a' && str[i+1][0] <= 'z') || (str[i+1][0] >= 'A' && str[i+1][0] <= 'Z')) {
				if str[i+1][0] == 'a' || str[i+1][0] == 'e' || str[i+1][0] == 'i' || str[i+1][0] == 'o' || str[i+1][0] == 'u' || str[i+1][0] == 'h' {
					str[i] = "an"
				} else {
					str[i] = "a"
				}
			}
		}
		// check for standalone punctuation mark and concatinate to previous member while eliminating the target
		// this part checks for occurences before the end of slice length
		if (str[i] == "," || str[i] == "." || str[i] == "!" || str[i] == "?" || str[i] == ":" || str[i] == ";" || str[i] == "'\"" || str[i] == "!?" || str[i] == "..." || str[i] == "!!") && i < len(str)-1 {
			str[i-1] = str[i-1] + str[i]
			str = append(str[:i], str[i+1:]...)
		}
		// this part specifically checks for the end of slice
		if (str[i] == "." || str[i] == "!" || str[i] == "?" || str[i] == ":" || str[i] == ";" || str[i] == "'\"" || str[i] == "!?" || str[i] == "..." || str[i] == "!!") && i == len(str)-1 {
			str[i-1] = str[i-1] + str[i]
			str = str[:i]
		}

	}
	return str
}

func CheckQuotation(str []string) []string {
	// this function checks the occurence of double and single quotation and corrects it position relative to proceeding and succeeding string
	var logic bool = true
	for i := 0; i < len(str); i++ {
		if str[i] == `'` || str[i] == `"` {
			if logic {
				str[i] = str[i] + str[i+1]
				str = append(str[:i+1], str[i+2:]...)
				logic = false
			} else {
				str[i-1] = str[i-1] + str[i]
				if i < len(str) {
					str = append(str[:i], str[i+1:]...)
				} else {
					str = str[:i]
					logic = true
				}
			}
		}
	}
	return str
}

func finalPunc(str []string) []string {
	// this is the function that eliminates the preceeding punctuation element within a string
	for j, char := range str {
		for i := 0; i < len(char); i++ {
			if (char[0] == ',' || char[0] == '.' || char[0] == '?' || char[0] == '!') && j != len(str)-1 {
				str[j-1] = str[j-1] + string(char[0])
				char = punc(char)
			}
		}
	}

	return str
}

func ConvIndex(str string) int {
	// the function singles out numerical index as a string and converts it to int
	var stri string
	for _, i := range str {
		if i >= '0' && i <= '9' {
			stri += string(i)
		}
	}
	num, _ := strconv.Atoi(stri)
	return num
}

func SplitWhiteSpaces(s string) []string {
	// function for spliting the strig data into slice of string
	neSli := []string{}
	var String string
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if String != "" {
				neSli = append(neSli, String)
				String = ""
			}
		} else {
			String += string(s[i])
		}
	}
	if String != "" {
		neSli = append(neSli, String)
		String = ""
	}
	return neSli
}

func readFile() string {
	// open file with name provided as an argument
	// if no error continue to reading content and return the string text data
	file, _ := os.Open(os.Args[1])
	// defer file close until the end of function
	defer file.Close()

	// read content data and give out the byte value
	data, _ := io.ReadAll(file)
	// return a converted data in string format
	return string(data)
}

func BinConv2Dec(str string) string {
	// Convert binary string to decimal using big int
	// Create a big.Int to hold the result
	decimal := new(big.Int)
	decimal, success := decimal.SetString(str, 2)
	if !success {
		return " Error"
	}
	return decimal.String()
}

func ToUpper(s string) string {
	// convert lowerCase string to upperCase
	// create a empty string as a place holder for final string
	var UpperCase string
	for _, letter := range s {
		if letter >= 'a' && letter <= 'z' {
			UpperCase += string(letter - 'a' + 'A')
		} else {
			UpperCase += string(letter)
		}
	}
	return UpperCase
}

func ToLower(s string) string {
	// convert uppercase string to lowercase
	// create a empty string as a place holder for final string
	var loWer string
	for _, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			loWer = loWer + string(letter-'A'+'a')
		} else {
			loWer = loWer + string(letter)
		}
	}
	return loWer
}

func CapiTalize(str string) string {
	if len(str) == 0 {
		return str
	}
	// Convert the first character to uppercase and concatenate it with the rest of the string
	return strings.ToUpper(string(str[0])) + str[1:]
}

func HexConv2Dec(hex string) string {
	// convert to decimal big int using SetrString() funtion
	decimal, success := new(big.Int).SetString(hex, 16)
	if !success {
		return "Error"
	}
	return decimal.String()
}

func punc(str string) string {
	var sTr string
	if (strings.Contains(str, ",") || strings.Contains(str, ".")) && (str[0] == ',' || str[0] == '.') {
		for _, char := range str {
			if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '\'' {
				sTr += string(char)
			}
		}
	}
	return sTr
}

func Processin(wordSlice []string) []string {
	// execute commands based on "keyword" match
	for i := 0; i < len(wordSlice); i++ {
		// single instance of keyword followed by keyword+index
		// single instance calls for execution on the previous index
		// keyword+index respects the presence of the index by looing back based on th integral value of the index
		if wordSlice[i] == "(cap)" {

			wordSlice[i-1] = strings.Title(wordSlice[i-1])

			wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)

		}
		if wordSlice[i] == "(cap," {
			index := ConvIndex(wordSlice[i+1])
			for j := i; j > i-index; j-- {
				wordSlice[j-1] = CapiTalize(wordSlice[j-1])
			}
			wordSlice = append(wordSlice[:i], wordSlice[i+2:]...)
		}
		if wordSlice[i] == "(up)" {

			wordSlice[i-1] = ToUpper(wordSlice[i-1])

			wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)

		}
		if wordSlice[i] == "(up," {
			// carries out lowercase to uppercase conversion considering the number of times(index)
			index := ConvIndex(wordSlice[i+1])
			for j := i; j > i-index; j-- {
				wordSlice[j-1] = ToUpper(wordSlice[j-1])
			}
			wordSlice = append(wordSlice[:i], wordSlice[i+2:]...)
		}
		if wordSlice[i] == "(low)" {
			// carries out uppercase to lowercase conversion
			wordSlice[i-1] = ToLower(wordSlice[i-1])

			wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)

		}
		if wordSlice[i] == "(low," {
			// carries out uppercase to lowercase conversion considering the number of times(index)
			index := ConvIndex(wordSlice[i+1])
			for j := i; j > i-index; j-- {
				wordSlice[j-1] = ToLower(wordSlice[j-1])
			}
			wordSlice = append(wordSlice[:i], wordSlice[i+2:]...)
		}
		if wordSlice[i] == "(bin)" {
			// carries out binary to decimal conversion
			wordSlice[i-1] = BinConv2Dec(wordSlice[i-1])

			wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)

		}
		if wordSlice[i] == "(bin," {
			// carries out binary to decimal conversion considering the number of times(index)
			index := ConvIndex(wordSlice[i+1])
			for j := i; j >= i-index; j-- {
				wordSlice[j-1] = BinConv2Dec(wordSlice[j-1])
			}
			wordSlice = append(wordSlice[:i], wordSlice[i+2:]...)
		}
		if wordSlice[i] == "(hex)" {

			wordSlice[i-1] = HexConv2Dec(wordSlice[i-1])

			wordSlice = append(wordSlice[:i], wordSlice[i+1:]...)

		}
		if wordSlice[i] == "(hex," {
			index := ConvIndex(wordSlice[i+1])
			for j := i; j >= i-index; j-- {
				wordSlice[j-1] = HexConv2Dec(wordSlice[j-1])
			}
			wordSlice = append(wordSlice[:i], wordSlice[i+2:]...)
		}
	}

	return wordSlice
}

func createNewFile(str string) {
	file, _ := os.Create(os.Args[2])

	// Defer is used for purposes of cleanup like
	// closing a running file after the file has
	// been written and main //function has
	// completed execution
	defer file.Close()

	//  string written to the file.
	file.WriteString(str)
}
