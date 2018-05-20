// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"unicode"
	"strings"

)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	reply := ""
	if isQuestion(remark) && allUpperCase(remark) {
		reply = "Calm down, I know what I'm doing!"
	} else if isQuestion(remark) {
		reply = "Sure."
	} else if any(remark, unicode.IsUpper) && !any(remark, unicode.IsLower) {
		reply = "Whoa, chill out!"
	} else {
		reply = "Whatever."
	}
	return reply
}

//"statement containing question mark",
//"Ending with ? means a question."

//"1, 2, 3"
//"Whoa, chill out"
func any(items string, test func(rune) bool) bool {
	for _, item := range items {
	  if test(item) {
		return true
	  }
	}
	return false
}

func allUpperCase(sentence string) bool {
	// return if the string contains number
	if isInt(sentence){
		return false
	} else if strings.ToUpper(sentence) == sentence {
		return true
	}
	return false
}	

func isInt(s string) bool {
    for _, c := range s {
        if unicode.IsDigit(c) {
            return true
        }
    }
    return false
}

func isNonLetter(s string) bool {
	for _, c := range s {
		if !unicode.IsLetter(c) && !unicode.IsDigit(c) {
            return true
        }
    }
    return false
}

// isQuestion 
func isQuestion(sentence string) bool {
	if sentence[len(sentence)-1] == '?' {
		return true
	} else {
		return false
	}
}