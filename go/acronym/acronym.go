// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"strings"
)

// Abbreviate should have a comment documenting it.
func Abbreviate(sentence string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	var newSentence []string
	sentenceArray := strings.Split(sentence, "")
	for i := range sentenceArray {
		if i == 0 {
		  newSentence = append(newSentence, sentenceArray[i])
		} else if sentenceArray[i - 1] == " " {
		  newSentence = append(newSentence, strings.ToUpper(sentenceArray[i]))
		} else if sentenceArray[i - 1] == "-" {
		  newSentence = append(newSentence, strings.ToUpper(sentenceArray[i]))
		}
	}
	return strings.Join(newSentence, "")
}



// def acronym(sentence)
//     new_sentence = []
//     sentence.split("").each_with_index do |a , index|
//         if index == 0 
//             new_sentence << a           
//         elsif sentence[index - 1] == " "
//             new_sentence << sentence[index].capitalize
//         elsif sentence[index - 1] == "-"
//             new_sentence << sentence[index].capitalize
//         end    
//     end
//     new_sentence
// end

