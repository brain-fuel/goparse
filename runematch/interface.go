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
			PosInfo:    in.PosInfo,
			Matched:    "",
			MatchedEOF: true,
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
	failureMessage string,
) (ds.Match, ds.MatcherInput, error) {
	actualString, _, err := in.CurrentCharString()
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
		PosInfo: in.PosInfo,
		Matched: actualString,
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
		actualString, _, _ := in.CurrentCharString()
		matchFn := func(ds.MatcherInput) bool {
			return string(expected) == actualString
		}
		failureMessage := fmt.Sprintf(
			"expected '%c', got '%s'",
			expected,
			actualString)
		return currentRuneMatch(in, matchFn, failureMessage)
	}
}

func Not(expected rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, _ := in.CurrentCharString()
		matchFn := func(ds.MatcherInput) bool {
			return string(expected) != actualString
		}
		failureMessage := fmt.Sprintf("expected not '%c', got '%s'", expected, actualString)
		return currentRuneMatch(in, matchFn, failureMessage)
	}
}

func inRangeOfRunes(rs ...rune) func(ds.MatcherInput) bool {
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

func notInRangeOfRunes(rs ...rune) func(ds.MatcherInput) bool {
	return func(in ds.MatcherInput) bool {
		return !inRangeOfRunes(rs...)(in)
	}
}

func AnyOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, _ := in.CurrentCharString()
		matchFn := inRangeOfRunes(rs...)
		var runeStrings []string
		for _, r := range rs {
			runeStrings = append(runeStrings, fmt.Sprintf("'%c'", r))
		}
		failureMessage := fmt.Sprintf(
			"expected any of [%s], got '%s'",
			strings.Join(runeStrings, ", "),
			actualString,
		)
		return currentRuneMatch(in, matchFn, failureMessage)
	}
}

func NoneOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, _ := in.CurrentCharString()
		matchFn := notInRangeOfRunes(rs...)
		runeStrings := make([]string, len(rs))
		for idx, r := range rs {
			runeStrings[idx] = fmt.Sprintf("'%c'", r)
		}
		failureMessage := fmt.Sprintf(
			"expected none of [%s], got '%s'",
			strings.Join(runeStrings, ", "),
			actualString,
		)
		return currentRuneMatch(in, matchFn, failureMessage)
	}
}

func InRange(r1 rune, r2 rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualRune, _, _ := in.CurrentRune()
		lo := min(r1, r2)
		hi := max(r1, r2)
		matchFn := func(in ds.MatcherInput) bool {
			return lo <= actualRune && actualRune <= hi
		}
		failureMessage := fmt.Sprintf(
			"expected rune in range ['%c', '%c'], got '%c'",
			lo,
			hi,
			actualRune,
		)
		return currentRuneMatch(in, matchFn, failureMessage)
	}
}

func NotInRange(r1 rune, r2 rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualRune, _, _ := in.CurrentRune()
		lo := min(r1, r2)
		hi := max(r1, r2)
		matchFn := func(in ds.MatcherInput) bool {
			return actualRune < lo || hi < actualRune
		}
		failureMessage := fmt.Sprintf(
			"expected rune not in range ['%c', '%c'], got '%c'",
			lo,
			hi,
			actualRune,
		)
		return currentRuneMatch(in, matchFn, failureMessage)
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
