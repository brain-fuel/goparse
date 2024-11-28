package runematch

import (
	"errors"
	"fmt"
	"io"

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

func Any() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, err := in.CurrentCharString()
		if err != nil {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: io.EOF.Error(),
			}
		}
		newMatch := ds.Match{
			PosInfo: in.PosInfo,
			Matched: actualString,
		}
		advancedInput, err := in.AdvanceBy(1)
		if err != nil {
			return newMatch, in, err
		}
		return newMatch, advancedInput, nil
	}
}

func Single(expected rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, err := in.CurrentCharString()
		if err != nil {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: io.EOF.Error(),
			}
		}
		if string(expected) != actualString {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: fmt.Sprintf("expected '%c', got '%s'", expected, actualString),
			}
		}
		newMatch := ds.Match{
			PosInfo: in.PosInfo,
			Matched: actualString,
		}
		advancedInput, err := in.AdvanceBy(1)
		if err != nil {
			return newMatch, in, err
		}
		return newMatch, advancedInput, nil
	}
}

func Not(expected rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		actualString, _, err := in.CurrentCharString()
		if err != nil {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: io.EOF.Error(),
			}
		}
		if string(expected) == actualString {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: fmt.Sprintf("expected not '%c', got '%s'", expected, actualString),
			}
		}
		newMatch := ds.Match{
			PosInfo: in.PosInfo,
			Matched: actualString,
		}
		advancedInput, err := in.AdvanceBy(1)
		if err != nil {
			return newMatch, in, err
		}
		return newMatch, advancedInput, nil
	}
}

func AnyOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotAnyOf(rs ...rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func InRange(low rune, hi rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotInRange(low rune, hi rune) ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
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
	return NotAnyOf(' ', '\t', '\r', '\n')
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
