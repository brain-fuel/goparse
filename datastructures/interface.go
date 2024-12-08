package datastructures

import (
	"errors"
	"fmt"
	"io"
	"unicode/utf8"
)

const (
	EOF string = "EOF"
)

type PosInfo struct {
	OffsetBytes int
	OffsetRunes int
	Line        int
	ColBytes    int
	ColRunes    int
}

func NewPosInfo() PosInfo {
	return PosInfo{
		OffsetBytes: 0,
		OffsetRunes: 0,
		Line:        1,
		ColBytes:    1,
		ColRunes:    1,
	}
}

type MatcherInput struct {
	Source      string
	LengthBytes int
	LengthRunes int
	PosInfo     PosInfo
}

func NewMatcherInput(s string) MatcherInput {
	return MatcherInput{
		Source:      s,
		LengthBytes: len(s),
		LengthRunes: utf8.RuneCountInString(s),
		PosInfo:     NewPosInfo(),
	}
}

func (mi MatcherInput) CurrentPos() (int, int) {
	return (mi.PosInfo.OffsetBytes + 1), (mi.PosInfo.OffsetRunes + 1)
}

func (mi MatcherInput) CurrentRune() (rune, int, error) {
	if mi.LengthBytes <= mi.PosInfo.OffsetBytes {
		return rune(0), 1, io.EOF
	}
	r := rune(mi.Source[mi.PosInfo.OffsetBytes])
	runeLen := utf8.RuneLen(r)
	return r, runeLen, nil
}

func (mi MatcherInput) CurrentCharString() (string, int, error) {
	r, strLen, err := mi.CurrentRune()
	s := string(r)
	return s, strLen, err
}

func (mi MatcherInput) AdvanceBy(n int) (MatcherInput, error) {
	runesRemaining := mi.LengthRunes - mi.PosInfo.OffsetRunes
	if n <= 0 {
		return mi, errors.New(
			"AdvanceBy() requires a positive number less than the number of runes remaining, including EOF",
		)
	}
	if runesRemaining-n < 0 {
		return mi, io.EOF
	}
	offsetBytes := mi.PosInfo.OffsetBytes
	offsetRunes := mi.PosInfo.OffsetRunes
	line := mi.PosInfo.Line
	colBytes := mi.PosInfo.ColBytes
	colRunes := mi.PosInfo.ColRunes
	for idx := 0; idx < n; idx++ {
		r, runeLen := utf8.DecodeRuneInString(mi.Source[offsetBytes:])
		offsetBytes += runeLen
		colBytes += runeLen
		offsetRunes++
		colRunes++
		if r == '\n' {
			line++
			colBytes = 1
			colRunes = 1
		}

	}
	return MatcherInput{
		Source:      mi.Source,
		LengthBytes: mi.LengthBytes,
		LengthRunes: mi.LengthRunes,
		PosInfo: PosInfo{
			OffsetBytes: offsetBytes,
			OffsetRunes: offsetRunes,
			Line:        line,
			ColBytes:    colBytes,
			ColRunes:    colRunes,
		},
	}, nil
}

func (mi MatcherInput) Run(m Matcher) (Match, MatcherInput, error) {
	return m(mi)
}

type MatchError struct {
	InputIdentifier string
	PosInfo         PosInfo
	Message         string
}

func (e MatchError) Error() string {
	msg := fmt.Sprintf(
		"%d:%d: %s",
		e.PosInfo.Line,
		e.PosInfo.ColRunes,
		e.Message,
	)
	if e.InputIdentifier != "" {
		return fmt.Sprintf("%s:%s", e.InputIdentifier, msg)
	}
	return msg
}

type Match struct {
	PosInfo     PosInfo
	Matched     string
	MatchedEOF  bool
	Description string
}

type Matcher func(MatchStage) MatchStage

type MatchFn func(MatcherInput) bool

type M3 struct {
	Matcher Matcher
}

func (m M3) Run(ms MatchStage) MatchStage {
	return m.Matcher(ms)
}

// MatcherInput needs to be the triple so that the Run()s can be chained

type MatchStage struct {
	Match        Match
	MatcherInput MatcherInput
	Errs         []error
}

func InitMatchStage(input string) MatchStage {
	return MatchStage{
		Match:        Match{},
		MatcherInput: NewMatcherInput(input),
		Errs:         []error{},
	}
}

func (ms MatchStage) CurrentPos() (int, int) {
	return ms.MatcherInput.CurrentPos()
}

func (ms MatchStage) CurrentRune() (int, int) {
	return ms.MatcherInput.CurrentPos()
}

func (ms MatchStage) CurrentCharString() (string, int, error) {
	return ms.MatcherInput.CurrentCharString()
}

func (ms MatchStage) AdvanceBy(n int) MatchStage {
	newMi, err := ms.MatcherInput.AdvanceBy(n)
	if err != nil {
		return MatchStage{
			Match:        ms.Match,
			MatcherInput: ms.MatcherInput,
			Errs:         append(ms.Errs, err),
		}
	}
	return MatchStage{
		Match:        ms.Match,
		MatcherInput: newMi,
		Errs:         ms.Errs,
	}
}

func (ms MatchStage) Run(m Matcher) MatchStage {
	return m(ms)
}
