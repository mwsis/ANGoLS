package strings

import (
	"fmt"

	"github.com/synesissoftware/ANGoLS/_std_strings"
	debug "github.com/synesissoftware/Diagnosticism.Go/debug"
)

/*
 * C++ construct                            | Go construct                             | ANGoLS construct                                     |
 * ---------------------------------------- | ---------------------------------------- | ---------------------------------------------------- |
 * #find(string const& sf)                  | strings.Index(s, sf string)              | -                                                    |
 * #find(char c)                            | strings.IndexByte(s string, c byte),
 *                                            strings.IndexByte(s string, c rune)      | -                                                    |
 * #find_first_of(string const& chars)      | strings.IndexAny(s, chars string)        | -                                                    |
 * #rfind(string const& sf)                 | strings.LastIndex(s, sf string)          | -                                                    |
 * #rfind(char c)                           | strings.LastIndexByte(s string, c byte)  | -                                                    |
 * #find_last_of(string const& chars)       | strings.LastIndexAny(s, chars string)    | -                                                    |
 *
 *                                          |                                          | strings.IndexAfter(s, sf, string, ix int)            |
 *                                          |                                          | strings.IndexByteAfter(s string, c byte, ix int)     |
 *                                          |                                          | strings.IndexByteAfter(s string, c byte, ix int)     |
 *                                          |                                          | strings.IndexAnyAfter(s, chars, string, ix int)      |
 *                                          |                                          | strings.IndexNotAnyAfter(s, chars, string, ix int)   |
 */

// Finds the index of the given substring in the given string, starting from
// the position after the given index. -1 is returned if the find is not
// successful.
//
// To search from the start of the string, specify the value -1 for the
// index. Any index value less than -1 will be treated as if -1 specified.
// Any index value greater than the size of the string will result in a
// return value of -1.
//
// The returned value reflects the position of the found substring relative
// to the start of the string, not from the index.
func IndexAfter(s string, sf string, ix int) int {

	if ix < -1 {
		ix = -1
	}

	off := ix + 1

	if off > len(s) {
		return -1
	}

	r := _std_strings.PRIVATE_Index(s[off:], sf)

	if r < 0 {
		return r
	} else {
		return r + off
	}
}
