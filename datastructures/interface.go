package datastructures

import (
	"errors"
	"fmt"
	"io"
	"strings"
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
		Source:      strings.ReplaceAll(s, "\r\n", "\n"),
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
			"MatcherInput.AdvanceBy() requires a positive number less than the number of runes remaining",
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
	PosInfo    PosInfo
	Matched    string
	MatchedEOF bool
}

type Matcher func(MatcherInput) (Match, MatcherInput, error)
