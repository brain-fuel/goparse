package predicates

import (
	"fmt"
)

func ExampleEOF_failingInput() {
	failingInput := "abacadaba"
	expectedFailingMatch := EOF()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {0 9}, match: "", rest: "abacadaba"}
}

func ExampleEOF_succeedingInput() {
	succeedingInput := ""
	expectedSucceedingMatch := EOF()(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_EOF, dldist: 0, odldist: {0 0}, match: "", rest: ""}
}

func ExampleAnyRune_failingInput() {
	failingInput := ""
	expectedFailingMatch := AnyRune()(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
}

func ExampleAnyRune_succeedingInput01() {
	succeedingInput1 := "abacadaba"
	expectedSucceedingMatch1 := AnyRune()(succeedingInput1)
	fmt.Println(expectedSucceedingMatch1)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "a", rest: "bacadaba"}
}

func ExampleAnyRune_succeedingInput02() {
	succeedingInput2 := "123abc"
	expectedSucceedingMatch2 := AnyRune()(succeedingInput2)
	fmt.Println(expectedSucceedingMatch2)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "1", rest: "23abc"}
}

func ExampleAnyRune_succeedingInput03() {
	succeedingInput3 := "世界"
	expectedSucceedingMatch3 := AnyRune()(succeedingInput3)
	fmt.Println(expectedSucceedingMatch3)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleRune_failingInput01() {
	toMatch := '世'
	failingInput1 := ""
	expectedFailingMatch1 := Rune(toMatch)(failingInput1)
	fmt.Println(expectedFailingMatch1)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
}

func ExampleRune_failingInput02() {
	toMatch := '世'
	failingInput2 := "abacadaba"
	expectedFailingMatch2 := Rune(toMatch)(failingInput2)
	fmt.Println(expectedFailingMatch2)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {1 8}, match: "", rest: "abacadaba"}
}

func ExampleRune_failingInput03() {
	toMatch := '世'
	failingInput3 := "123abc"
	expectedFailingMatch3 := Rune(toMatch)(failingInput3)
	fmt.Println(expectedFailingMatch3)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 6, odldist: {1 5}, match: "", rest: "123abc"}
}

func ExampleRune_succeedingInput01() {
	toMatch := '世'
	succeedingInput := "世界"
	expectedSucceedingMatch := Rune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleRune_succeedingInput02() {
	toMatch := '世'
	succeedingInput := "世界abc"
	expectedSucceedingMatch := Rune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界abc"}
}

func ExampleRune_succeedingInput03() {
	toMatch := '世'
	succeedingInput := "世界123"
	expectedSucceedingMatch := Rune(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界123"}
}

func ExampleStr_failingInput01() {
	toMatch := "世界"
	failingInput := ""
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
}

func ExampleStr_failingInput02() {
	toMatch := "世界"
	failingInput := "abacadaba"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "abacadaba"}
}

func ExampleStr_failingInput03() {
	toMatch := "世界"
	failingInput := "123abc"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "123abc"}
}

func ExampleStr_failingInput04() {
	toMatch := "世界"
	failingInput := "世"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_MATCH_THEN_EOF, dldist: 1, odldist: {0 1}, match: "", rest: "世"}
}

func ExampleStr_failingInput05() {
	toMatch := "世界"
	failingInput := "界"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 1, odldist: {1 1}, match: "", rest: "界"}
}

func ExampleStr_succeedingInput01() {
	toMatch := "世界"
	succeedingInput := "世界"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: ""}
}

func ExampleStr_succeedingInput02() {
	toMatch := "世界"
	succeedingInput := "世界abacadaba"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "abacadaba"}
}

func ExampleStr_succeedingInput03() {
	toMatch := "世界"
	succeedingInput := "世界123abc"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "123abc"}
}

func ExampleStr_failingInput06() {
	toMatch := "secondmatch"
	failingInput := ""
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
}

func ExampleStr_failingInput07() {
	toMatch := "secondmatch"
	failingInput := "abacadaba"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
}

func ExampleStr_failingInput08() {
	toMatch := "secondmatch"
	failingInput := "secondm@tc"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
}

func ExampleStr_failingInput09() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
}

func ExampleStr_failingInput10() {
	toMatch := "secondmatch"
	failingInput := "secondm@tch123abc"
	expectedFailingMatch := Str(toMatch)(failingInput)
	fmt.Println(expectedFailingMatch)
	// Output: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch123abc"}
}

func ExampleStr_succeedingInput04() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
}

func ExampleStr_succeedingInput05() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch123"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "123"}
}

func ExampleStr_succeedingInput06() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatchabacadaba"
	expectedSucceedingMatch := Str(toMatch)(succeedingInput)
	fmt.Println(expectedSucceedingMatch)
	// Output: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "abacadaba"}
}
