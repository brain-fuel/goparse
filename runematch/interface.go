package runematch

import (
	"errors"
	"fmt"
	"io"

	ds "goforge.dev/tools/goparse/datastructures"
)

func Single(expected rune, in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
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

func MatchEOF(in ds.MatcherInput) (ds.Match, ds.MatcherInput, error) {
	return ds.Match{}, ds.MatcherInput{}, errors.New("Not implemented yet")
}
