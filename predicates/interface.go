package predicates

import (
	"errors"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"

	"goforge.dev/tools/goparse/types"
)

const (
	POTENTIAL_NEAR_MISS_CHAR_THRESHOLD       int     = 1
	POTENTIAL_NEAR_MISS_PERCENTAGE_THRESHOLD float64 = 25.0
)

func NearMissThresholds() (int, float64) {
	lookupCharThreshold, ctExists := os.LookupEnv("POTENTIAL_NEAR_MISS_CHAR_THRESHOLD")
	lookupPercentThreshold, dtExists := os.LookupEnv("POTENTIAL_NEAR_MISS_PERCENTAGE_THRESHOLD")
	ct := POTENTIAL_NEAR_MISS_CHAR_THRESHOLD
	dt := POTENTIAL_NEAR_MISS_PERCENTAGE_THRESHOLD
	if ctExists {
		lct, err := strconv.Atoi(lookupCharThreshold)
		if err != nil {
			ct = POTENTIAL_NEAR_MISS_CHAR_THRESHOLD
		} else {
			ct = lct
		}
	}
	if dtExists {
		ldt, err := strconv.ParseFloat(lookupPercentThreshold, 64)
		if err != nil {
			dt = POTENTIAL_NEAR_MISS_PERCENTAGE_THRESHOLD
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
	return int(math.Floor((dt * float64(length) / 100)))
}

func MatchesEOF() types.MatchPred {
	return func(in string) types.MatchRes {
		if utf8.RuneCountInString(in) != 0 {
			return types.NewMatchFailure(types.FAILURE_EOF, DLDist(in, ""), ODLDist(in, ""), in)
		}
		return types.NewMatchSuccess(types.SUCCESS_EOF, "", "")
	}
}

func MatchesAnyRune() types.MatchPred {
	return func(in string) types.MatchRes {
		if utf8.RuneCountInString(in) == 0 {
			return types.NewMatchFailure(
				types.FAILURE_EOF,
				1,
				types.NewODist(types.Distance(0), types.Overlap(0), types.Leftover(1)),
				"",
			)
		}
		return types.NewMatchSuccess(
			types.SUCCESS_RUNE,
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

func MatchesRune(toMatch rune) types.MatchPred {
	return func(in string) types.MatchRes {
		firstRune, err := firstRuneInString(in)
		if err != nil {
			return types.NewMatchFailure(
				types.FAILURE_EOF,
				DLDist(in, string(toMatch)),
				ODLDist(in, string(toMatch)),
				in,
			)
		}
		if firstRune != toMatch {
			return types.NewMatchFailure(
				types.FAILURE_NO_MATCH,
				DLDist(in, string(toMatch)),
				ODLDist(in, string(toMatch)),
				in,
			)
		}
		return types.NewMatchSuccess(
			types.SUCCESS_RUNE,
			firstCharInString(in),
			butFirstCharInString(in),
		)
	}
}

func MatchesStr(toMatch string) types.MatchPred {
	return func(in string) types.MatchRes {
		inputLength := utf8.RuneCountInString(in)
		matchLength := utf8.RuneCountInString(toMatch)
		matchCut := firstNRunesInString(matchLength, toMatch)
		inputCut := firstNRunesInString(matchLength, in)
		dist := DLDist(matchCut, inputCut)
		oDist := ODLDistExpectedVsActual(matchCut, inputCut)
		eof := inputLength == 0
		success := dist == 0
		if eof {
			return types.NewMatchFailure(types.FAILURE_EOF, dist, oDist, in)
		}
		if !success {
			matchType := failMatchType(matchLength, dist, oDist)
			return types.NewMatchFailure(matchType, dist, oDist, in)
		}
		return types.NewMatchSuccess(
			types.SUCCESS_STRING,
			toMatch,
			butNCharInString(matchLength, in),
		)
	}
}

func failMatchType(
	length int,
	dist types.Distance,
	oDist types.OverlappingDistance,
) types.MatchType {
	disti := int(dist)
	nearMissThreshold := NearMissThreshold(length)
	count, _, leftover := oDist.ToTriple()
	counti := int(count)
	leftoveri := int(leftover)
	if 0 < leftoveri {
		if counti == 0 {
			return types.FAILURE_MATCH_THEN_EOF
		}
		if counti <= nearMissThreshold {
			return types.FAILURE_POTENTIAL_NEAR_MISS_THEN_EOF
		}
		return types.FAILURE_NO_MATCH_THEN_EOF
	}
	if disti <= nearMissThreshold {
		return types.FAILURE_POTENTIAL_NEAR_MISS
	}
	return types.FAILURE_NO_MATCH
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// https://en.wikipedia.org/wiki/Damerau-Levenshtein_distance
func DLDist(s1, s2 string) types.Distance {
	if len(s1) == 0 {
		return types.Distance(utf8.RuneCountInString(s2))
	}
	if len(s2) == 0 {
		return types.Distance(utf8.RuneCountInString(s1))
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
	return types.Distance(m[l1][l2])
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

func ODLDist(s1, s2 string) types.OverlappingDistance {
	l1 := utf8.RuneCountInString(s1)
	l2 := utf8.RuneCountInString(s2)
	minLen := min(l1, l2)
	leftover := absInt(l1 - l2)
	return types.NewODist(
		DLDist(s1[:minLen], s2[:minLen]),
		types.Overlap(minLen),
		types.Leftover(leftover),
	)
}

func ODLDistExpectedVsActual(expected, actual string) types.OverlappingDistance {
	lExpected := utf8.RuneCountInString(expected)
	lActual := utf8.RuneCountInString(actual)
	minLen := min(lExpected, lActual)
	leftover := lExpected - lActual
	return types.NewODist(
		DLDist(expected[:minLen], actual[:minLen]),
		types.Overlap(minLen),
		types.Leftover(leftover),
	)
}
