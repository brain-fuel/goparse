package runematch

import (
	"errors"
	"fmt"
	"testing"

	ds "goforge.dev/tools/goparse/datastructures"
)

// Helper Functions

func generateLettersAndNumbersRunes() []rune {
	lettersAndNumbers := []rune{}
	for idx := range 26 {
		lettersAndNumbers = append(lettersAndNumbers, rune('a'+idx))
	}
	for idx := range 26 {
		lettersAndNumbers = append(lettersAndNumbers, rune('A'+idx))
	}
	for idx := range 10 {
		lettersAndNumbers = append(lettersAndNumbers, rune('0'+idx))
	}
	return lettersAndNumbers
}

func generateLettersAndNumbersStrings() []string {
	lettersAndNumbers := []string{}
	for idx := range 26 {
		lettersAndNumbers = append(lettersAndNumbers, string(rune('a'+idx)))
	}
	for idx := range 26 {
		lettersAndNumbers = append(lettersAndNumbers, string(rune('A'+idx)))
	}
	for idx := range 10 {
		lettersAndNumbers = append(lettersAndNumbers, string(rune('0'+idx)))
	}
	return lettersAndNumbers
}

// Tests

func TestMatchOnEmptyStringShouldFail(t *testing.T) {
	lettersAndNumbers := generateLettersAndNumbersRunes()
	for _, r := range lettersAndNumbers {
		expectedErr := errors.New("1:1: EOF").Error()
		matcherInput := ds.NewMatcherInput("")
		_, _, err := Single(r, matcherInput)
		if err == nil {
			t.Errorf("expected %q, got no error", expectedErr)
			continue
		}
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
	}
}

func TestFailingMatchShouldResultInSameMatcherInput(t *testing.T) {
	lettersAndNumbers := generateLettersAndNumbersRunes()
	for _, r := range lettersAndNumbers {
		expectedErr := errors.New("1:1: EOF").Error()
		matcherInput := ds.NewMatcherInput("")
		_, input, err := Single(r, matcherInput)
		if err == nil {
			t.Errorf("expected %q, got no error", expectedErr)
			continue
		}
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err.Error())
		}
		if input != matcherInput {
			t.Errorf("expected new matcher input to be %v, got %v", matcherInput, input)
		}
	}
}

func TestMatchOnSingleCharacterStringsShouldFail(t *testing.T) {
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	for _, s := range lettersAndNumbersStrings {
		expectedErr := fmt.Sprintf("1:1: expected ':', got '%c'", s[0])
		matcherInput := ds.NewMatcherInput(s)
		_, _, err := Single(':', matcherInput)
		if err == nil {
			t.Errorf("expected %q; got no error", expectedErr)
			continue
		}
		if err.Error() != expectedErr {
			t.Errorf("expected %q; got %q", expectedErr, err.Error())
		}
	}
}

func TestMatchOnSingleCharacterStringsShouldSucceed(t *testing.T) {
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	lettersAndNumbersRunes := generateLettersAndNumbersRunes()
	for idx, s := range lettersAndNumbersStrings {
		matcherInput := ds.NewMatcherInput(s)
		toMatch := lettersAndNumbersRunes[idx]
		match, _, err := Single(toMatch, matcherInput)
		if err != nil {
			t.Errorf("expected no error; got %q", err)
			continue
		}
		expectedPosInfo := matcherInput.PosInfo
		actualPosInfo := match.PosInfo
		if actualPosInfo != expectedPosInfo {
			t.Errorf(
				"attempting to match '%c', expected PosInfo %v, got PosInfo %v",
				toMatch,
				expectedPosInfo,
				actualPosInfo,
			)
		}
		expectedMatchedString := string(toMatch)
		actualMatchedString := match.Matched
		if actualMatchedString != expectedMatchedString {
			t.Errorf(
				"attempting to match '%c', expected Matched %v, got Matched %v",
				toMatch,
				expectedMatchedString,
				actualMatchedString,
			)
		}
	}
}

func TestMatchOnSingleCharacterShouldAdvancePosInfo(t *testing.T) {
	inputStr := "ab\nc\n\nadaba\n"
	matcherInput := ds.NewMatcherInput(inputStr)
	charsToMatch := []rune{'a', 'b', '\n', 'c', '\n', '\n', 'a', 'd', 'a', 'b', 'a', '\n'}
	for _, charToMatch := range charsToMatch {
		_, inputRes, err := Single(charToMatch, matcherInput)
		if err != nil {
			t.Errorf("expected no error; got %q", err)
		}
		expectedNewOffset := matcherInput.PosInfo.OffsetRunes + 1
		actualNewOffset := inputRes.PosInfo.OffsetRunes
		if actualNewOffset != expectedNewOffset {
			t.Errorf(
				"matched '%c': expected new offset %d, actual %d",
				charToMatch,
				expectedNewOffset,
				actualNewOffset,
			)
		}
		expectedNewLine := matcherInput.PosInfo.Line
		expectedNewColumn := matcherInput.PosInfo.ColRunes + 1
		if charToMatch == '\n' {
			expectedNewLine++
			expectedNewColumn = 1
		}
		actualNewLine := inputRes.PosInfo.Line
		actualNewColumn := inputRes.PosInfo.ColRunes
		if actualNewLine != expectedNewLine {
			t.Errorf(
				"matched '%c': expected new line %d, actual %d",
				charToMatch,
				expectedNewLine,
				actualNewLine,
			)
		}
		if actualNewColumn != expectedNewColumn {
			t.Errorf(
				"matched '%c': expected new column %d, actual %d",
				charToMatch,
				expectedNewLine,
				actualNewLine,
			)
		}
		matcherInput = inputRes
	}
}
