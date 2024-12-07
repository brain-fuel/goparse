package predicates

import (
	"errors"
	"math"
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
	lookupPercentThreshold, dtExists := os.LookupEnv("NEAR_MISS_DECIMAL_THRESHOLD")
	ct := NEAR_MISS_CHAR_THRESHOLD
	dt := NEAR_MISS_DECIMAL_THRESHOLD
	if ctExists {
		lct, err := strconv.Atoi(lookupCharThreshold)
		if err != nil {
			ct = NEAR_MISS_CHAR_THRESHOLD
		} else {
			ct = lct
		}
	}
	if dtExists {
		ldt, err := strconv.ParseFloat(lookupPercentThreshold, 64)
		if err != nil {
			dt = NEAR_MISS_DECIMAL_THRESHOLD
		} else {
			dt = ldt
		}
	}
	return ct, dt
}

func NearMissThreshold(length int) int {
	ct, dt := NearMissThresholds()
	isShort := length <= 4
	if isShort {
		return ct
	}
	return int(math.Floor((dt * float64(length))))
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

func firstNRunesInString(n int, s string) string {
	if n <= 0 {
		return ""
	}
	count := 0
	for idx := range s {
		count++
		if n < count {
			return s[:idx]
		}
	}
	return s
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
		matchCut := firstNRunesInString(matchLength, toMatch)
		inputCut := firstNRunesInString(matchLength, in)
		dldist := DLDist(matchCut, inputCut)
		odldist := ODLDistTest(matchCut, inputCut)
		eof := inputLength == 0
		success := dldist == 0
		if eof {
			return t.NewMatchFailure(t.FAILURE_EOF, dldist, odldist, in)
		}
		if !success {
			matchType := failMatchType(matchLength, dldist, odldist)
			return t.NewMatchFailure(matchType, dldist, odldist, in)
		}
		return t.NewMatchSuccess(t.SUCCESS_STRING, toMatch, butNCharInString(matchLength, in))
	}
}

func failMatchType(length int, dldist int, odldist t.Pair[int]) t.MatchType {
	nearMissThreshold := NearMissThreshold(length)
	ocount, lacking := odldist.ToTuple()
	if 0 < lacking {
		if ocount == 0 {
			return t.FAILURE_MATCH_THEN_EOF
		}
		if ocount <= nearMissThreshold {
			return t.FAILURE_NEAR_MISS_THEN_EOF
		}
		return t.FAILURE_NO_MATCH_THEN_EOF
	}
	if dldist <= nearMissThreshold {
		return t.FAILURE_NEAR_MISS
	}
	return t.FAILURE_NO_MATCH
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
	rs1 := strings.Split(s1, "")
	rs2 := strings.Split(s2, "")
	l1 := len(rs1)
	l2 := len(rs2)
	m := make([][]int, l1+1)
	var cost int
	for i := range m {
		m[i] = make([]int, l2+1)
	}
	for i := 0; i < l1+1; i++ {
		for j := 0; j < l2+1; j++ {
			if i == 0 {
				m[i][j] = j
			} else if j == 0 {
				m[i][j] = i
			} else {
				cost = 0
				if rs1[i-1] != rs2[j-1] {
					cost = 1
				}
				m[i][j] = multiMin(m[i-1][j]+1, m[i][j-1]+1, m[i-1][j-1]+cost)
				if i > 1 && j > 1 && rs1[i-1] == rs2[j-2] && rs1[i-2] == rs2[j-1] {
					m[i][j] = multiMin(m[i][j], m[i-2][j-2]+cost)
				}
			}
		}
	}
	return m[l1][l2]
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

func ODLDistTest(expected, actual string) t.Pair[int] {
	lExpected := utf8.RuneCountInString(expected)
	lActual := utf8.RuneCountInString(actual)
	minLen := min(lExpected, lActual)
	diff := lExpected - lActual
	return t.NewPair[int](DLDist(expected[:minLen], actual[:minLen]), diff)
}
