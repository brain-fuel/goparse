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

func TestEOFOnAnyCharacterExceptEOFShouldFail(t *testing.T) {
	blankMatcherInput := ds.NewMatcherInput("")
	m, newIn, err := EOF()(blankMatcherInput)
	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}
	if m.Matched != "" {
		t.Errorf("expected blank string match, got %s", m.Matched)
	}
	if !m.MatchedEOF {
		t.Error("expected MatchedEOF, got !MatchedEOF")
	}
	if newIn != blankMatcherInput {
		t.Error("expected new input to match blank input")
	}
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	for _, s := range lettersAndNumbersStrings {
		notEOFErr := fmt.Sprintf("1:1: expected EOF, got '%s'", s)
		matcherInput := ds.NewMatcherInput(s)
		_, _, err := EOF()(matcherInput)
		if err == nil {
			t.Errorf("expected %q, got no error", notEOFErr)
			continue
		}
		if err != nil {
			if err.Error() != notEOFErr {
				t.Errorf("expected %q, got %q", notEOFErr, err)
				continue
			}
		}
	}
}

func TestAnyOnAnyCharacterExceptEOFShouldSucceed(t *testing.T) {
	eofErr := errors.New("1:1: EOF").Error()
	blankMatcherInput := ds.NewMatcherInput("")
	_, _, err := Any()(blankMatcherInput)
	if err != nil {
		if err.Error() != eofErr {
			t.Errorf("expected %q, got %q", eofErr, err)
		}
	} else {
		t.Errorf("expected %q, got no error", eofErr)
	}
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	for _, s := range lettersAndNumbersStrings {
		matcherInput := ds.NewMatcherInput(s)
		match, nextInput, err := Any()(matcherInput)
		if err != nil {
			t.Errorf("expected no error; got %q", err)
		}
		expectedPosInfo := matcherInput.PosInfo
		actualPosInfo := match.PosInfo
		if actualPosInfo != expectedPosInfo {
			t.Errorf(
				"attempting to match any, expected PosInfo %v, got PosInfo %v",
				expectedPosInfo,
				actualPosInfo,
			)
		}
		increasedInput, _ := matcherInput.AdvanceBy(1)
		if nextInput.PosInfo != increasedInput.PosInfo {
			t.Errorf("attempted to advance matcher and failed")
		}
		cr, _, _ := matcherInput.CurrentCharString()
		if match.Matched != cr {
			t.Errorf("attempted to match any on '%s', got '%s'", cr, match.Matched)
		}
	}
}

func TestSingleOnEmptyStringShouldFail(t *testing.T) {
	lettersAndNumbers := generateLettersAndNumbersRunes()
	for _, r := range lettersAndNumbers {
		expectedErr := errors.New("1:1: EOF").Error()
		matcherInput := ds.NewMatcherInput("")
		_, _, err := Single(r)(matcherInput)
		if err == nil {
			t.Errorf("expected %q, got no error", expectedErr)
			continue
		}
		if err != nil {
			if err.Error() != expectedErr {
				t.Errorf("expected %q, got %q", expectedErr, err.Error())
			}
		}
	}
}

func TestFailingMatchShouldResultInSameMatcherInput(t *testing.T) {
	lettersAndNumbers := generateLettersAndNumbersRunes()
	for _, r := range lettersAndNumbers {
		expectedErr := errors.New("1:1: EOF").Error()
		matcherInput := ds.NewMatcherInput("")
		_, input, err := Single(r)(matcherInput)
		if err == nil {
			t.Errorf("expected %q, got no error", expectedErr)
			continue
		}
		if err != nil {
			if err.Error() != expectedErr {
				t.Errorf("expected %q, got %q", expectedErr, err.Error())
			}
		}
		if input != matcherInput {
			t.Errorf("expected new matcher input to be %v, got %v", matcherInput, input)
		}
	}
}

func TestSingleOnSingleCharacterStringsShouldFail(t *testing.T) {
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	for _, s := range lettersAndNumbersStrings {
		expectedErr := fmt.Sprintf("1:1: expected ':', got '%c'", s[0])
		matcherInput := ds.NewMatcherInput(s)
		_, _, err := Single(':')(matcherInput)
		if err == nil {
			t.Errorf("expected %q; got no error", expectedErr)
			continue
		}
		if err != nil {
			if err.Error() != expectedErr {
				t.Errorf("expected %q; got %q", expectedErr, err.Error())
			}
		}
	}
}

func TestSingleOnSingleCharacterStringsShouldSucceed(t *testing.T) {
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	lettersAndNumbersRunes := generateLettersAndNumbersRunes()
	for idx, s := range lettersAndNumbersStrings {
		matcherInput := ds.NewMatcherInput(s)
		toMatch := lettersAndNumbersRunes[idx]
		match, _, err := Single(toMatch)(matcherInput)
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

func TestSingleOnSingleCharacterShouldAdvancePosInfo(t *testing.T) {
	inputStr := "ab\nc\n\nadaba\n"
	matcherInput := ds.NewMatcherInput(inputStr)
	charsToMatch := []rune{'a', 'b', '\n', 'c', '\n', '\n', 'a', 'd', 'a', 'b', 'a', '\n'}
	for _, charToMatch := range charsToMatch {
		_, inputRes, err := Single(charToMatch)(matcherInput)
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

func TestNotOnMatchingCharacterShouldFail(t *testing.T) {
	expectedErr := "1:1: expected not 'a', got 'a'"
	matcherInput := ds.NewMatcherInput("a")
	_, _, err := Not('a')(matcherInput)
	if err == nil {
		t.Errorf("expected %q, got no error", expectedErr)
	}
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err)
		}
	}
}

func TestNotOnNotMatchingCharacterShouldSucceed(t *testing.T) {
	lettersAndNumbersStrings := generateLettersAndNumbersStrings()
	for idx, s := range lettersAndNumbersStrings {
		if idx == 0 {
			continue
		}
		matcherInput := ds.NewMatcherInput(s)
		toMatch := 'a'
		match, nextInput, err := Not(toMatch)(matcherInput)
		if err != nil {
			t.Errorf("expected no error; got %q", err)
			continue
		}
		expectedMatchedString := string(toMatch)
		actualMatchedString := match.Matched
		if actualMatchedString == expectedMatchedString {
			t.Errorf(
				"attempting to match not '%c', expected Matched %v, got Matched %v",
				toMatch,
				expectedMatchedString,
				actualMatchedString,
			)
		}
		increasedInput, _ := matcherInput.AdvanceBy(1)
		if nextInput.PosInfo != increasedInput.PosInfo {
			t.Errorf("attempted to advance matcher and failed")
		}
	}
}

func TestAnyOfOnNotMatchingCharacterShouldFail(t *testing.T) {
	expectedErr := "1:1: expected any of ['b', 'c', 'd', 'e', 'f'], got 'a'"
	matcherInput := ds.NewMatcherInput("abacadaba")
	_, _, err := AnyOf('b', 'c', 'd', 'e', 'f')(matcherInput)
	if err == nil {
		t.Errorf("expected %q, got no error", expectedErr)
	}
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err)
		}
	}
}

func TestAnyOfOnMatchingCharactersShouldSucceed(t *testing.T) {
	rs := []rune{'a', 'b', 'c', 'd', 'e'}
	for _, r := range rs {
		expectedMatchedString := string(r)
		matcherInput := ds.NewMatcherInput(expectedMatchedString)
		match, _, err := AnyOf(rs...)(matcherInput)
		if err != nil {
			t.Errorf("expected no error; got %q", err)
			continue
		}
		actualMatchedString := match.Matched
		if actualMatchedString != expectedMatchedString {
			t.Errorf(
				"attempting to match any of ['a', 'b', 'c', 'd', 'e'], expected Matched %v, got Matched %v",
				expectedMatchedString,
				actualMatchedString,
			)
		}
	}
}

func TestNoneOfOnMatchingCharactersShouldFail(t *testing.T) {
	expectedErr := "1:1: expected none of ['a', 'b', 'c', 'd', 'e'], got 'a'"
	matcherInput := ds.NewMatcherInput("abacadaba")
	_, _, err := NoneOf('a', 'b', 'c', 'd', 'e')(matcherInput)
	if err == nil {
		t.Errorf("expected %q, got no error", expectedErr)
	}
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err)
		}
	}
}

func TestNoneOfOnNotMatchingCharactersShouldSucceed(t *testing.T) {
	rs := []rune{'b', 'c', 'd', 'e', 'f'}
	matcherInput := ds.NewMatcherInput("abacadaba")
	match, _, err := NoneOf(rs...)(matcherInput)
	if err != nil {
		t.Errorf("expected no error; got %q", err)
	}
	actualMatchedString := match.Matched
	for _, r := range rs {
		expectedFailMatchString := string(r)
		if actualMatchedString == expectedFailMatchString {
			t.Errorf(
				"attempting to match none of ['a', 'b', 'c', 'd', 'e'], expected Matched %v, got Matched %v",
				expectedFailMatchString,
				actualMatchedString,
			)
		}
	}
}

func TestInRangeOnCharOutOfRangeShouldFail(t *testing.T) {
	expectedErr := "1:1: expected rune in range ['b', 'f'], got 'a'"
	matcherInput := ds.NewMatcherInput("abacadaba")
	_, _, err := InRange('b', 'f')(matcherInput)
	if err == nil {
		t.Errorf("expected %q, got no error", expectedErr)
	}
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err)
		}
	}
}

func TestInRangeOnCharInRangeShouldSucceed(t *testing.T) {
	matcherInput := ds.NewMatcherInput("abacadaba")
	expectedMatchedString := "a"
	match, _, err := InRange('a', 'e')(matcherInput)
	if err != nil {
		t.Errorf("expected no error; got %q", err)
	}
	actualMatchedString := match.Matched
	if actualMatchedString != expectedMatchedString {
		t.Errorf(
			"attempting to match range of ['a', 'e'], expected Matched %v, got Matched %v",
			expectedMatchedString,
			actualMatchedString,
		)
	}
}

func TestNotInRangeOnCharInRangeShouldFail(t *testing.T) {
	expectedErr := "1:1: expected rune not in range ['a', 'e'], got 'a'"
	matcherInput := ds.NewMatcherInput("abacadaba")
	_, _, err := NotInRange('a', 'e')(matcherInput)
	if err == nil {
		t.Errorf("expected %q, got no error", expectedErr)
	}
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("expected %q, got %q", expectedErr, err)
		}
	}
}

func TestNotInRangeOnCharNotInRangeShouldSucceed(t *testing.T) {
	matcherInput := ds.NewMatcherInput("fbacadaba")
	expectedFailMatchString := "a"
	match, _, err := NotInRange('a', 'e')(matcherInput)
	if err != nil {
		t.Errorf("expected no error; got %q", err)
	}
	actualMatchedString := match.Matched
	if actualMatchedString == expectedFailMatchString {
		t.Errorf(
			"attempting to match not in range ['a', 'e'], expected Matched %v, got Matched %v",
			expectedFailMatchString,
			actualMatchedString,
		)
	}
}
