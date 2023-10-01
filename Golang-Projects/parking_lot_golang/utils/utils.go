package utils

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"parking_lot/errors"
)

const (
	// Tab holds unicode value of tab
	Tab = "\t"
	// Space const holds the space string
	Space = " "
	// NewLineDelim holds unicode value of tab
	NewLineDelim = "\n"
	// EndLineDelim holds unicode value of tab
	EndLineDelim = '\n'
)

var regNoRegex = regexp.MustCompile(`^(([A-Za-z]){2}(|-)(?:[0-9]){1,2}(|-)(?:[A-Za-z]){1,2}(|-)([0-9]){1,4})$`)
var colourRegex = regexp.MustCompile(`^[A-Za-z]+$`)

// SplitCmdArguments attempts to split the input string by command and arguments
// Assuming that the string is seperated by space and the first instance is command
// and the rest followed by command is arguments.
func SplitCmdArguments(str string) (res []string, err error) {
	if strings.Contains(str, Tab) {
		err = errors.ErrInvalidTabSpace
		return
	}
	return strings.SplitN(str, Space, 2), nil
}

// FormatDateTime returns formatted date string
func FormatDateTime(t time.Time) string {
	return fmt.Sprintf("%d-%02d-%02d %s", t.Year(), t.Month(), t.Day(), t.Format(time.Kitchen))
}

// IsRegNoValid validates the regNo valid or not
func IsRegNoValid(regNo string) bool {
	return regNoRegex.MatchString(regNo)
}

// IsValidString validates the string
func IsValidString(str string) bool {
	return colourRegex.MatchString(str)
}
