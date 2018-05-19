// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	reply := ""
	if (strings.Title(remark) == remark) && (strings.Contains(remark,"?")) {
		reply = "Calm down, I know what I'm doing!"
	} else if strings.Contains(remark, "?") {
		reply = "Sure."
	} else if strings.Title(remark) == remark {
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
