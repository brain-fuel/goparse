package runematch

import (
	"errors"
	"fmt"
	"io"
	"strings"

	ds "goforge.dev/tools/goparse/datastructures"
)

func EOF() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, err := in.CurrentCharString()
		if err == nil {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: fmt.Sprintf("expected EOF, got '%s'", actualString),
			}
		}
		newMatch := ds.Match{
			PosInfo:     in.PosInfo,
			Matched:     "",
			MatchedEOF:  true,
			Description: "EOF",
		}
		return newMatch, in, nil
	}
}

func eofMatchFailure(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
	return ds.Match{}, in, ds.MatchError{
		PosInfo: in.PosInfo,
		Message: io.EOF.Error(),
	}
}

func currentRuneMatch(
	in ds.MatcherInput,
	matchFn ds.MatchFn,
	expectation string,
) (ds.Match, ds.MatcherInput, error) {
	actualString, _, err := in.CurrentCharString()
	failureMessage := fmt.Sprintf("expected %s, got '%s'", expectation, actualString)
	if err != nil {
		return eofMatchFailure(in)
	}
	didMatch := matchFn(in)
	if !didMatch {
		return ds.Match{}, in, ds.MatchError{
			PosInfo: in.PosInfo,
			Message: failureMessage,
		}
	}
	newMatch := ds.Match{
		PosInfo:     in.PosInfo,
		Matched:     actualString,
		Description: expectation,
	}
	advancedInput, _ := in.AdvanceBy(1)
	return newMatch, advancedInput, nil
}

func Any() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := func(ds.MatcherInput) bool { return true }
		return currentRuneMatch(in, matchFn, "")
	}
}

func Single(expected rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := func(ds.MatcherInput) bool {
			actualString, _, _ := in.CurrentCharString()
			return string(expected) == actualString
		}
		expectation := fmt.Sprintf("'%c'", expected)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func Not(expected rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := func(ds.MatcherInput) bool {
			actualString, _, _ := in.CurrentCharString()
			return string(expected) != actualString
		}
		expectation := fmt.Sprintf("not '%c'", expected)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func matchesOneOf(rs ...rune) func(ds.MatcherInput) bool {
	return func(in ds.MatcherInput) bool {
		actualString, _, _ := in.CurrentCharString()
		for _, r := range rs {
			if string(r) == actualString {
				return true
			}
		}
		return false
	}
}

func doesNotMatchOneOf(rs ...rune) func(ds.MatcherInput) bool {
	return func(in ds.MatcherInput) bool {
		return !matchesOneOf(rs...)(in)
	}
}

func AnyOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := matchesOneOf(rs...)
		runeStrings := make([]string, len(rs))
		for idx, r := range rs {
			runeStrings[idx] = fmt.Sprintf("'%c'", r)
		}
		expectation := fmt.Sprintf(
			"any of [%s]",
			strings.Join(runeStrings, ", "),
		)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func NoneOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := doesNotMatchOneOf(rs...)
		runeStrings := make([]string, len(rs))
		for idx, r := range rs {
			runeStrings[idx] = fmt.Sprintf("'%c'", r)
		}
		expectation := fmt.Sprintf(
			"none of [%s]",
			strings.Join(runeStrings, ", "),
		)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func matchInRangeOf(r1 rune, r2 rune) func(ds.MatcherInput) bool {
	lo := min(r1, r2)
	hi := max(r1, r2)
	return func(in ds.MatcherInput) bool {
		actualRune, _, _ := in.CurrentRune()
		return lo <= actualRune && actualRune <= hi
	}
}

func matchNotInRangeOf(r1 rune, r2 rune) func(ds.MatcherInput) bool {
	lo := min(r1, r2)
	hi := max(r1, r2)
	return func(in ds.MatcherInput) bool {
		actualRune, _, _ := in.CurrentRune()
		return hi < actualRune || actualRune < lo
	}
}

func InRange(r1 rune, r2 rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := matchInRangeOf(r1, r2)
		lo := min(r1, r2)
		hi := max(r1, r2)
		expectation := fmt.Sprintf(
			"rune in range ['%c', '%c']",
			lo,
			hi,
		)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func NotInRange(r1 rune, r2 rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		matchFn := matchNotInRangeOf(r1, r2)
		lo := min(r1, r2)
		hi := max(r1, r2)
		expectation := fmt.Sprintf(
			"rune not in range ['%c', '%c']",
			lo,
			hi,
		)
		return currentRuneMatch(in, matchFn, expectation)
	}
}

func Lower() ds.Matcher {
	return InRange('a', 'z')
}

func NotLower() ds.Matcher {
	return NotInRange('a', 'z')
}

func Upper() ds.Matcher {
	return InRange('A', 'Z')
}

func NotUpper() ds.Matcher {
	return NotInRange('A', 'Z')
}

func Digit() ds.Matcher {
	return InRange('0', '9')
}

func NotDigit() ds.Matcher {
	return NotInRange('0', '9')
}

func Space() ds.Matcher {
	return Single(' ')
}

func NotSpace() ds.Matcher {
	return Not(' ')
}

func Tab() ds.Matcher {
	return Single('\t')
}

func NotTab() ds.Matcher {
	return Not('\t')
}

func CR() ds.Matcher {
	return Single('\r')
}

func NotCR() ds.Matcher {
	return Not('\r')
}

func Newline() ds.Matcher {
	return Single('\n')
}

func NotNewline() ds.Matcher {
	return Not('\n')
}

func WS() ds.Matcher {
	return AnyOf(' ', '\t', '\r', '\n')
}

func NotWS() ds.Matcher {
	return NoneOf(' ', '\t', '\r', '\n')
}

func Alphanumeric() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotAlphanumeric() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Underscore() ds.Matcher {
	return Single('_')
}

func NotUnderscore() ds.Matcher {
	return Not('_')
}

func AlphanumericPlusUnderscore() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotAlphanumericPlusUnderscore() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Hyphen() ds.Matcher {
	return Single('-')
}

func NotHyphen() ds.Matcher {
	return Not('-')
}

func Exact() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotExact() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}
