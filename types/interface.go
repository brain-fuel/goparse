package types

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Pair[T any] struct {
	left  T
	right T
}

func NewPair[T any](a, b T) Pair[T] {
	return Pair[T]{
		left:  a,
		right: b,
	}
}

func (p Pair[T]) Left() T {
	return p.left
}

func (p Pair[T]) Right() T {
	return p.right
}

func (p Pair[T]) ToTuple() (T, T) {
	return p.left, p.right
}

type MatchType int

const (
	SUCCESS_EOF MatchType = iota
	SUCCESS_RUNE
	SUCCESS_STRING
	SUCCESS_REGEXP
	SUCCESS_SEMANTIC
	SUCCESS_OTHER

	FAILURE_EOF
	FAILURE_MATCH_THEN_EOF
	FAILURE_NEAR_MISS_THEN_EOF
	FAILURE_NO_MATCH_THEN_EOF

	FAILURE_NEAR_MISS
	FAILURE_NO_MATCH

	FAILURE_OTHER
)

func (mt MatchType) GoString() string {
	matchTypeNames := map[MatchType]string{
		SUCCESS_EOF:      "SUCCESS_EOF",
		SUCCESS_RUNE:     "SUCCESS_RUNE",
		SUCCESS_STRING:   "SUCCESS_STRING",
		SUCCESS_REGEXP:   "SUCCESS_REGEXP",
		SUCCESS_SEMANTIC: "SUCCESS_SEMANTIC",
		SUCCESS_OTHER:    "SUCCESS_OTHER",

		FAILURE_EOF:                "FAILURE_EOF",
		FAILURE_MATCH_THEN_EOF:     "FAILURE_MATCH_THEN_EOF",
		FAILURE_NEAR_MISS_THEN_EOF: "FAILURE_NEAR_MISS_THEN_EOF",
		FAILURE_NO_MATCH_THEN_EOF:  "FAILURE_NO_MATCH_THEN_EOF",

		FAILURE_NEAR_MISS: "FAILURE_NEAR_MISS",
		FAILURE_NO_MATCH:  "FAILURE_NO_MATCH",

		FAILURE_OTHER: "FAILURE_OTHER",
	}
	name, present := matchTypeNames[mt]
	if !present {
		return "FAILURE_UNK"
	}
	return name
}

func (mt MatchType) String() string {
	return mt.GoString()
}

type InputConsumption int

const (
	SLURP_EXACT InputConsumption = iota
	SLURP_RUNES
	SLURP_BYTES
	SLURP_STRING
	SPIT_EXACT
	SPIT_RUNES
	SPIT_BYTES
	SPIT_STRING
	DO_NOT_CONSUME
)

type ReadDirection int

const (
	HERE ReadDirection = iota
	FORWARD
	BACKWARD
)

type ParseErrType int

type MatchPred func(string) MatchRes

type MatchRes struct {
	matchType MatchType
	dldist    int
	odldist   Pair[int]
	match     string
	rest      string
}

func (mr MatchRes) GoString() string {
	restOfLine := strings.Split(mr.rest, "\n")[0]
	rest := strings.TrimRightFunc(restOfLine, unicode.IsSpace)
	if 64 < utf8.RuneCountInString(rest) {
		rs := []rune(rest)
		rest = string(rs[:64]) + "..."
	}
	return fmt.Sprintf(
		"MatchRes{type: %s, dldist: %d, odldist: %d, match: %q, rest: %q}",
		mr.matchType,
		mr.dldist,
		mr.odldist,
		mr.match,
		rest,
	)
}

func (mr MatchRes) String() string {
	return fmt.Sprintf("%#v", mr)
}

func (mr MatchRes) HowManyRunesMatched() int {
	return utf8.RuneCountInString(mr.match)
}

func (mr MatchRes) HowManyBytesMatched() int {
	return len(mr.match)
}

func NewMatchSuccess(mt MatchType, actual string, rest string) MatchRes {
	return MatchRes{
		matchType: mt,
		dldist:    0,
		odldist:   NewPair[int](0, 0),
		match:     actual,
		rest:      rest,
	}
}

func NewMatchFailure(mt MatchType, dldist int, odldist Pair[int], rest string) MatchRes {
	return MatchRes{
		matchType: mt,
		dldist:    dldist,
		odldist:   odldist,
		match:     "",
		rest:      rest,
	}
}

type Matcher struct {
	matchPred        MatchPred
	inputConsumption InputConsumption
}

type MatchStage struct {
	input   MatchInput
	posInfo PosInfo
	matches []Match
	errors  []MatchErr
}

type MatchInput struct {
	source      string
	lengthBytes int
	lengthRunes int
}

type PosInfo struct {
	offsetBytes int
	offsetRunes int
	line        int
	columnBytes int
	columnRunes int
}

type Match struct {
	matchType   MatchType
	posInfo     PosInfo
	matched     string
	matchedEOF  bool
	description string
}

type MatchErr struct {
	inputIdentifier string
	posInfo         PosInfo
	matchErrType    MatchType
	err             error
}

func (e MatchErr) Error() string {
	var matchErrType string
	switch e.matchErrType {
	case FAILURE_EOF:
		matchErrType = "EOF"
	case FAILURE_NO_MATCH:
		matchErrType = "in bounds, no match"
	default:
		matchErrType = "unknown error type"
	}
	err := e.err.Error()
	msg := fmt.Sprintf(
		"%d:%d: %s: %s",
		e.posInfo.line,
		e.posInfo.columnRunes,
		matchErrType,
		err,
	)
	if e.inputIdentifier != "" {
		return fmt.Sprintf("%s:%s", e.inputIdentifier, msg)
	}
	return msg
}

type ParseStage[T any] struct {
	matchStage   MatchStage
	parseResults []ParseResult[T]
	errors       []ParseErr
}

type ParseResult[T any] struct {
	contents    T
	description string
}

type Parser[T any] struct {
	matcher     Matcher
	onMatch     func([]ParseStage[T]) ParseStage[T]
	onFailMatch func([]ParseStage[T]) ParseStage[T]
}

type ParseErr struct {
	inputIdentifier string
	posInfo         PosInfo
	parseErrType    ParseErrType
	err             error
}

func (e ParseErr) Error() string {
	var parseErrType string
	switch e.parseErrType {
	default:
		parseErrType = "unknown parse error"
	}
	err := e.err.Error()
	msg := fmt.Sprintf(
		"%d:%d: %s: %s",
		e.posInfo.line,
		e.posInfo.columnRunes,
		parseErrType,
		err,
	)
	if e.inputIdentifier != "" {
		return fmt.Sprintf("%s:%s", e.inputIdentifier, msg)
	}
	return msg
}
