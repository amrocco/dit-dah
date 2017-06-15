package ditdah

import (
	"bytes"
	"strings"
)

var charToMorse = map[string]string{
	"a": ".-",
	"b": "-...",
	"c": "-.-.",
	"d": "-..",
	"e": ".",
	"f": "..-.",
	"g": "--.",
	"h": "....",
	"i": "..",
	"j": ".---",
	"k": "-.-",
	"l": ".-..",
	"m": "--",
	"n": "-.",
	"o": "---",
	"p": ".--.",
	"q": "--.-",
	"r": ".-.",
	"s": "...",
	"t": "-",
	"u": "..-",
	"v": "...-",
	"w": ".--",
	"x": "-..-",
	"y": "-.--",
	"z": "--..",
	"1": ".----",
	"2": "..---",
	"3": "...--",
	"4": "....-",
	"5": ".....",
	"6": "-....",
	"7": "--...",
	"8": "---..",
	"9": "----.",
	"0": "-----",
	" ": "/",
}

var morseToChar = map[string]string{
	".-":    "a",
	"-...":  "b",
	"-.-.":  "c",
	"-..":   "d",
	".":     "e",
	"..-.":  "f",
	"--.":   "g",
	"....":  "h",
	"..":    "i",
	".---":  "j",
	"-.-":   "k",
	".-..":  "l",
	"--":    "m",
	"-.":    "n",
	"---":   "o",
	".--.":  "p",
	"--.-":  "q",
	".-.":   "r",
	"...":   "s",
	"-":     "t",
	"..-":   "u",
	"...-":  "v",
	".--":   "w",
	"-..-":  "x",
	"-.--":  "y",
	"--..":  "z",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
	"-----": "0",
	"/":     " ",
}

//Encode translates an alphanumeric string to morse code
func Encode(alphanumeric string) string {
	var morseCode bytes.Buffer
	alphanumericLength := len([]rune(alphanumeric))
	alphanumeric = strings.ToLower(alphanumeric)

	for pos, char := range alphanumeric {
		code := charToMorse[string(char)]
		morseCode.WriteString(code)
		if pos != alphanumericLength+1 {
			morseCode.WriteString(" ")
		}
	}

	return morseCode.String()
}

//Decode translates morse code to an alphanumeric string
func Decode(morseCode string) string {
	var alphanumeric bytes.Buffer
	splitMorseCode := strings.Split(morseCode, " ")

	for _, code := range splitMorseCode {
		alphanumeric.WriteString(morseToChar[code])
	}

	return alphanumeric.String()
}
