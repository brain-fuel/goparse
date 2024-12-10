package predicates

import (
	"fmt"
)

// MatchesEOF()("any string") fails
func ExampleMatchesEOF_failingInput() {
	failingInput := "abacadaba"
	expectedFailingMatch := MatchesEOF()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dist: 9, oDist: {0 0 9}, match: "", rest: "abacadaba"}
}

// MatchesEOF()("") succeeds on an empty string
func ExampleMatchesEOF_succeedingInput() {
	succeedingInput := ""
	expectedSucceedingMatch := MatchesEOF()(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_EOF, dist: 0, oDist: {0 0 0}, match: "", rest: ""}
}

// MatchesAnyRune()("") fails on blank input
func ExampleMatchesAnyRune_failingInput() {
	failingInput := ""
	expectedFailingMatch := MatchesAnyRune()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dist: 1, oDist: {0 0 1}, match: "", rest: ""}
}

// MatchesAnyRune()("any string with length greater than zero")
// matches the first rune in the string, no matter what it is
func ExampleMatchesAnyRune_succeedingInput01() {
	succeedingInput1 := "abacadaba"
	expectedSucceedingMatch1 := MatchesAnyRune()(succeedingInput1)
	fmt.Println(expectedSucceedingMatch1)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -8}, match: "a", rest: "bacadaba"}
}

func ExampleMatchesAnyRune_succeedingInput02() {
	succeedingInput2 := "123abc"
	expectedSucceedingMatch2 := MatchesAnyRune()(succeedingInput2)
	fmt.Println(expectedSucceedingMatch2)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -5}, match: "1", rest: "23abc"}
}

func ExampleMatchesAnyRune_succeedingInput03() {
	succeedingInput3 := "世界"
	expectedSucceedingMatch3 := MatchesAnyRune()(succeedingInput3)
	fmt.Println(expectedSucceedingMatch3)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -1}, match: "世", rest: "界"}
}

func ExampleMatchesRune_failingInput01() {
	toMatch := '世'
	failingInput1 := ""
	expectedFailingMatch1 := MatchesRune(toMatch)(failingInput1)
	fmt.Println(expectedFailingMatch1)
	// Output: MatchRes{type: FAILURE_EOF, dist: 1, oDist: {0 0 1}, match: "", rest: ""}
}

func ExampleMatchesRune_failingInput02() {
	toMatch := '世'
	failingInput2 := "abacadaba"
	expectedFailingMatch2 := MatchesRune(toMatch)(failingInput2)
	fmt.Println(expectedFailingMatch2)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dist: 9, oDist: {1 1 8}, match: "", rest: "abacadaba"}
}

func ExampleMatchesRune_failingInput03() {
	toMatch := '世'
	failingInput3 := "123abc"
	expectedFailingMatch3 := MatchesRune(toMatch)(failingInput3)
	fmt.Println(expectedFailingMatch3)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dist: 6, oDist: {1 1 5}, match: "", rest: "123abc"}
}

func ExampleMatchesRune_succeedingInput01() {
	toMatch := '世'
	succeedingInput := "世界"
	expectedSucceedingMatch := MatchesRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -1}, match: "世", rest: "界"}
}

func ExampleMatchesRune_succeedingInput02() {
	toMatch := '世'
	succeedingInput := "世界abc"
	expectedSucceedingMatch := MatchesRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -4}, match: "世", rest: "界abc"}
}

func ExampleMatchesRune_succeedingInput03() {
	toMatch := '世'
	succeedingInput := "世界123"
	expectedSucceedingMatch := MatchesRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dist: 0, oDist: {0 1 -4}, match: "世", rest: "界123"}
}

func ExampleMatchesStr_failingInput01() {
	toMatch := "世界"
	failingInput := ""
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dist: 2, oDist: {0 0 2}, match: "", rest: ""}
}

func ExampleMatchesStr_failingInput02() {
	toMatch := "世界"
	failingInput := "abacadaba"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dist: 2, oDist: {2 2 0}, match: "", rest: "abacadaba"}
}

func ExampleMatchesStr_failingInput03() {
	toMatch := "世界"
	failingInput := "123abc"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dist: 2, oDist: {2 2 0}, match: "", rest: "123abc"}
}

func ExampleMatchesStr_failingInput04() {
	toMatch := "世界"
	failingInput := "世"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_MATCH_THEN_EOF, dist: 1, oDist: {0 1 1}, match: "", rest: "世"}
}

func ExampleMatchesStr_failingInput05() {
	toMatch := "世界"
	failingInput := "界"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_POTENTIAL_NEAR_MISS_THEN_EOF, dist: 1, oDist: {1 1 1}, match: "", rest: "界"}
}

func ExampleMatchesStr_succeedingInput01() {
	toMatch := "世界"
	succeedingInput := "世界"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 2 0}, match: "世界", rest: ""}
}

func ExampleMatchesStr_succeedingInput02() {
	toMatch := "世界"
	succeedingInput := "世界abacadaba"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 2 -9}, match: "世界", rest: "abacadaba"}
}

func ExampleMatchesStr_succeedingInput03() {
	toMatch := "世界"
	succeedingInput := "世界123abc"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 2 -6}, match: "世界", rest: "123abc"}
}

func ExampleMatchesStr_failingInput06() {
	toMatch := "secondmatch"
	failingInput := ""
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dist: 11, oDist: {0 0 11}, match: "", rest: ""}
}

func ExampleMatchesStr_failingInput07() {
	toMatch := "secondmatch"
	failingInput := "abacadaba"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dist: 9, oDist: {8 9 2}, match: "", rest: "abacadaba"}
}

func ExampleMatchesStr_failingInput08() {
	toMatch := "secondmatch"
	failingInput := "secondm@tc"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_POTENTIAL_NEAR_MISS_THEN_EOF, dist: 2, oDist: {1 10 1}, match: "", rest: "secondm@tc"}
}

func ExampleMatchesStr_failingInput09() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_POTENTIAL_NEAR_MISS, dist: 1, oDist: {1 11 0}, match: "", rest: "secondm@tch"}
}

func ExampleMatchesStr_failingInput10() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch123abc"
	expectedFailingMatch := MatchesStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_POTENTIAL_NEAR_MISS, dist: 1, oDist: {1 11 0}, match: "", rest: "secondm@tch123abc"}
}

func ExampleMatchesStr_succeedingInput04() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 11 0}, match: "secondmatch", rest: ""}
}

func ExampleMatchesStr_succeedingInput05() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch123"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 11 -3}, match: "secondmatch", rest: "123"}
}

func ExampleMatchesStr_succeedingInput06() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatchabacadaba"
	expectedSucceedingMatch := MatchesStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dist: 0, oDist: {0 11 -9}, match: "secondmatch", rest: "abacadaba"}
}
