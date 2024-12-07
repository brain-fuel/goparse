package predicates

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	t "goforge.dev/tools/goparse/types"
)

const (
	NEAR_MISS_CHAR_THRESHOLD    int     = 1
	NEAR_MISS_DECIMAL_THRESHOLD float64 = 0.25
)

func NearMissThresholds() (int, float64) {
	lookupCharThreshold, ctExists := os.LookupEnv("NEAR_MISS_CHAR_THRESHOLD")
	lookupPercentThreshold, ptExists := os.LookupEnv("NEAR_MISS_DECIMAL_THRESHOLD")
	var ct int
	var pt float64
	if ctExists {
		lct, err := strconv.Atoi(lookupCharThreshold)
		if err != nil {
			ct = NEAR_MISS_CHAR_THRESHOLD
		} else {
			ct = lct
		}
	}
	if ptExists {
		lpt, err := strconv.ParseFloat(lookupPercentThreshold, 64)
		if err != nil {
			pt = NEAR_MISS_DECIMAL_THRESHOLD
		} else {
			pt = lpt
		}
	}
	return ct, pt
}

func EOF() t.MatchPred {
	return func(in string) t.MatchRes {
		if utf8.RuneCountInString(in) != 0 {
			return t.NewMatchFailure(t.FAILURE_EOF, DLDist(in, ""), ODLDist(in, ""), in)
		}
		return t.NewMatchSuccess(t.SUCCESS_EOF, "", "")
	}
}

func AnyRune() t.MatchPred {
	return func(in string) t.MatchRes {
		if utf8.RuneCountInString(in) == 0 {
			return t.NewMatchFailure(t.FAILURE_EOF, 1, t.NewPair[int](0, 1), "")
		}
		return t.NewMatchSuccess(
			t.SUCCESS_RUNE,
			firstCharInString(in),
			butFirstCharInString(in),
		)
	}
}

func firstRuneInString(cs string) (rune, error) {
	if utf8.RuneCountInString(cs) == 0 {
		return '0', errors.New("empty string contains no first rune")
	}
	var r rune
	for _, c := range cs {
		r = c
		break
	}
	return r, nil
}

func firstCharInString(cs string) string {
	if utf8.RuneCountInString(cs) == 0 {
		return ""
	}
	var c string
	for _, r := range cs {
		c = string(r)
		break
	}
	return c
}

func butNCharInString(n int, cs string) string {
	if n <= 0 {
		return cs
	}
	for idx := range cs {
		if n == 0 {
			return cs[idx:]
		}
		n--
	}
	return ""
}

func butFirstCharInString(cs string) string {
	return butNCharInString(1, cs)
}

func Rune(r rune) t.MatchPred {
	return func(in string) t.MatchRes {
		firstRune, err := firstRuneInString(in)
		if err != nil {
			return t.NewMatchFailure(
				t.FAILURE_EOF,
				DLDist(in, string(r)),
				ODLDist(in, string(r)),
				in,
			)
		}
		if firstRune != r {
			return t.NewMatchFailure(
				t.FAILURE_NO_MATCH,
				DLDist(in, string(r)),
				ODLDist(in, string(r)),
				in,
			)
		}
		return t.NewMatchSuccess(t.SUCCESS_RUNE, firstCharInString(in), butFirstCharInString(in))
	}
}

func Str(toMatch string) t.MatchPred {
	return func(in string) t.MatchRes {
		inputLength := utf8.RuneCountInString(in)
		matchLength := utf8.RuneCountInString(toMatch)
		dldistLength := matchLength
		dldist := DLDist(in, toMatch[:dldistLength])
		odldist := ODLDist(in, toMatch[:dldistLength])
		eof := inputLength < matchLength
		success := dldist == 0
		if eof {
			return t.NewMatchFailure(t.FAILURE_EOF, dldist, odldist, in)
		}
		if !success {
			return t.NewMatchFailure(t.FAILURE_NO_MATCH, dldist, odldist, in)
		}
		return t.NewMatchSuccess(t.SUCCESS_STRING, toMatch, butNCharInString(matchLength, in))
	}
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// https://en.wikipedia.org/wiki/Damerau-Levenshtein_distance
func DLDist(s1, s2 string) int {
	if len(s1) == 0 {
		return utf8.RuneCountInString(s2)
	}
	if len(s2) == 0 {
		return utf8.RuneCountInString(s1)
	}
	if s1 == s2 {
		return 0
	}
	s1Array := strings.Split(s1, "")
	s2Array := strings.Split(s2, "")
	lenS1Array := len(s1Array)
	lenS2Array := len(s2Array)
	m := make([][]int, lenS1Array+1)
	var cost int
	for i := range m {
		m[i] = make([]int, lenS2Array+1)
	}
	for i := 0; i < lenS1Array+1; i++ {
		for j := 0; j < lenS2Array+1; j++ {
			if i == 0 {
				m[i][j] = j
			} else if j == 0 {
				m[i][j] = i
			} else {
				cost = 0
				if s1Array[i-1] != s2Array[j-1] {
					cost = 1
				}
				m[i][j] = multiMin(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+cost)
				if i > 1 && j > 1 && s1Array[i-1] == s2Array[j-2] && s1Array[i-2] == s2Array[j-1] {
					m[i][j] = multiMin(m[i][j], m[i-2][j-2]+cost)
				}
			}
		}
	}
	return m[lenS1Array][lenS2Array]
}

func multiMin(ns ...int) int {
	if len(ns) == 0 {
		return 0
	}
	acc := ns[0]
	for idx, n := range ns {
		if idx == 0 {
			continue
		}
		acc = min(acc, n)
	}
	return acc
}

func ODLDist(s1, s2 string) t.Pair[int] {
	l1 := utf8.RuneCountInString(s1)
	l2 := utf8.RuneCountInString(s2)
	minLen := min(l1, l2)
	diff := absInt(l1 - l2)
	return t.NewPair[int](DLDist(s1[:minLen], s2[:minLen]), diff)
}
