package runematch

import (
	"errors"
	"fmt"
	"io"

	ds "goforge.dev/tools/goparse/datastructures"
)

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

func EOF() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		s, _, err := in.CurrentCharString()
		if err == nil {
			return ds.Match{}, in, ds.MatchError{
				PosInfo: in.PosInfo,
				Message: fmt.Sprintf("expected EOF, got '%s'", s),
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
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Upper() ds.Matcher {
	return InRange('A', 'Z')
}

func NotUpper() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Digit() ds.Matcher {
	return InRange('0', '9')
}

func NotDigit() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Space() ds.Matcher {
	return Single(' ')
}

func NotSpace() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Tab() ds.Matcher {
	return Single('\t')
}

func NotTab() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func CR() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotCR() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func Newline() ds.Matcher {
	return Single('\n')
}

func NotNewline() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func WS() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotWS() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
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
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
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
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func AnyOf() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
}

func NotAnyOf() ds.Matcher {
	return func(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
		return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
	}
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
