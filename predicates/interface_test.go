package predicates

import (
	"fmt"
)

func ExampleIsEOF_failingInput() {
	failingInput := "abacadaba"
	expectedFailingMatch := IsEOF()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {0 9}, match: "", rest: "abacadaba"}
}

func ExampleIsEOF_succeedingInput() {
	succeedingInput := ""
	expectedSucceedingMatch := IsEOF()(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_EOF, dldist: 0, odldist: {0 0}, match: "", rest: ""}
}

func ExampleIsAnyRune_failingInput() {
	failingInput := ""
	expectedFailingMatch := IsAnyRune()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
}

func ExampleIsAnyRune_succeedingInput01() {
	succeedingInput1 := "abacadaba"
	expectedSucceedingMatch1 := IsAnyRune()(succeedingInput1)
	fmt.Println(expectedSucceedingMatch1)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "a", rest: "bacadaba"}
}

func ExampleIsAnyRune_succeedingInput02() {
	succeedingInput2 := "123abc"
	expectedSucceedingMatch2 := IsAnyRune()(succeedingInput2)
	fmt.Println(expectedSucceedingMatch2)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "1", rest: "23abc"}
}

func ExampleIsAnyRune_succeedingInput03() {
	succeedingInput3 := "世界"
	expectedSucceedingMatch3 := IsAnyRune()(succeedingInput3)
	fmt.Println(expectedSucceedingMatch3)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleIsRune_failingInput01() {
	toMatch := '世'
	failingInput1 := ""
	expectedFailingMatch1 := IsRune(toMatch)(failingInput1)
	fmt.Println(expectedFailingMatch1)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
}

func ExampleIsRune_failingInput02() {
	toMatch := '世'
	failingInput2 := "abacadaba"
	expectedFailingMatch2 := IsRune(toMatch)(failingInput2)
	fmt.Println(expectedFailingMatch2)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {1 8}, match: "", rest: "abacadaba"}
}

func ExampleIsRune_failingInput03() {
	toMatch := '世'
	failingInput3 := "123abc"
	expectedFailingMatch3 := IsRune(toMatch)(failingInput3)
	fmt.Println(expectedFailingMatch3)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 6, odldist: {1 5}, match: "", rest: "123abc"}
}

func ExampleIsRune_succeedingInput01() {
	toMatch := '世'
	succeedingInput := "世界"
	expectedSucceedingMatch := IsRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleIsRune_succeedingInput02() {
	toMatch := '世'
	succeedingInput := "世界abc"
	expectedSucceedingMatch := IsRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界abc"}
}

func ExampleIsRune_succeedingInput03() {
	toMatch := '世'
	succeedingInput := "世界123"
	expectedSucceedingMatch := IsRune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界123"}
}

func ExampleIsStr_failingInput01() {
	toMatch := "世界"
	failingInput := ""
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
}

func ExampleIsStr_failingInput02() {
	toMatch := "世界"
	failingInput := "abacadaba"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "abacadaba"}
}

func ExampleIsStr_failingInput03() {
	toMatch := "世界"
	failingInput := "123abc"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "123abc"}
}

func ExampleIsStr_failingInput04() {
	toMatch := "世界"
	failingInput := "世"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_MATCH_THEN_EOF, dldist: 1, odldist: {0 1}, match: "", rest: "世"}
}

func ExampleIsStr_failingInput05() {
	toMatch := "世界"
	failingInput := "界"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 1, odldist: {1 1}, match: "", rest: "界"}
}

func ExampleIsStr_succeedingInput01() {
	toMatch := "世界"
	succeedingInput := "世界"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: ""}
}

func ExampleIsStr_succeedingInput02() {
	toMatch := "世界"
	succeedingInput := "世界abacadaba"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "abacadaba"}
}

func ExampleIsStr_succeedingInput03() {
	toMatch := "世界"
	succeedingInput := "世界123abc"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "123abc"}
}

func ExampleIsStr_failingInput06() {
	toMatch := "secondmatch"
	failingInput := ""
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
}

func ExampleIsStr_failingInput07() {
	toMatch := "secondmatch"
	failingInput := "abacadaba"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
}

func ExampleIsStr_failingInput08() {
	toMatch := "secondmatch"
	failingInput := "secondm@tc"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
}

func ExampleIsStr_failingInput09() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
}

func ExampleIsStr_failingInput10() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch123abc"
	expectedFailingMatch := IsStr(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch123abc"}
}

func ExampleIsStr_succeedingInput04() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
}

func ExampleIsStr_succeedingInput05() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch123"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "123"}
}

func ExampleIsStr_succeedingInput06() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatchabacadaba"
	expectedSucceedingMatch := IsStr(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "abacadaba"}
}
