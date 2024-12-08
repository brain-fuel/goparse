package predicates

import (
	"fmt"
)

func ExampleEOF() {
	failureInput := "abacadaba"

	succeedingInput := ""

	expectedFailMatch := EOF()(failureInput)
	expectedSuccessMatch := EOF()(succeedingInput)

	fmt.Println(expectedFailMatch)
	fmt.Printf("%v\n", expectedSuccessMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {0 9}, match: "", rest: "abacadaba"}
	// MatchRes{type: SUCCESS_EOF, dldist: 0, odldist: {0 0}, match: "", rest: ""}
}

func ExampleAnyRune() {
	failureInput := ""

	succeedingInput1 := "abacadaba"
	succeedingInput2 := "123abc"
	succeedingInput3 := "世界"

	expectedFailMatch := AnyRune()(failureInput)
	expectedSuccessMatch1 := AnyRune()(succeedingInput1)
	expectedSuccessMatch2 := AnyRune()(succeedingInput2)
	expectedSuccessMatch3 := AnyRune()(succeedingInput3)

	fmt.Println(expectedFailMatch)
	fmt.Println(expectedSuccessMatch1)
	fmt.Println(expectedSuccessMatch2)
	fmt.Println(expectedSuccessMatch3)

	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
	// MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "a", rest: "bacadaba"}
	// MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "1", rest: "23abc"}
	// MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleRune() {
	toMatch := '世'

	failureInput1 := ""
	failureInput2 := "abacadaba"
	failureInput3 := "123abc"
	succeedingInput := "世界"

	expectedFailMatch1 := Rune(toMatch)(failureInput1)
	expectedFailMatch2 := Rune(toMatch)(failureInput2)
	expectedFailMatch3 := Rune(toMatch)(failureInput3)
	expectedSuccessMatch := Rune(toMatch)(succeedingInput)

	fmt.Println(expectedFailMatch1)
	fmt.Println(expectedFailMatch2)
	fmt.Println(expectedFailMatch3)
	fmt.Println(expectedSuccessMatch)

	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {1 8}, match: "", rest: "abacadaba"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 6, odldist: {1 5}, match: "", rest: "123abc"}
	// MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleStr_matchWithFailingInput_01() {
	toMatch := "世界"
	failureInput := ""
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
}

func ExampleStr_matchWithFailingInput_02() {
	toMatch := "世界"
	failureInput := "abacadaba"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "abacadaba"}
}

func ExampleStr_matchWithFailingInput_03() {
	toMatch := "世界"
	failureInput := "123abc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "123abc"}
}

func ExampleStr_matchWithFailingInput_04() {
	toMatch := "世界"
	failureInput := "世"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_MATCH_THEN_EOF, dldist: 1, odldist: {0 1}, match: "", rest: "世"}
}

func ExampleStr_matchWithFailingInput_05() {
	toMatch := "世界"
	failureInput := "界"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 1, odldist: {1 1}, match: "", rest: "界"}
}

func ExampleStr_matchWithSucceedingInput_01() {
	toMatch := "世界"
	succeedingInput := "世界"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: ""}
}

func ExampleStr_matchWithSucceedingInput_02() {
	toMatch := "世界"
	succeedingInput := "世界abacadaba"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "abacadaba"}
}

func ExampleStr_matchWithSucceedingInput_03() {
	toMatch := "世界"
	succeedingInput := "世界123abc"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "123abc"}
}

func ExampleStr_matchWithFailingInput_06() {
	toMatch := "secondmatch"
	failureInput := ""
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
}

func ExampleStr_matchWithFailingInput_07() {
	toMatch := "secondmatch"
	failureInput := "abacadaba"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
}

func ExampleStr_matchWithFailingInput_08() {
	toMatch := "secondmatch"
	failureInput := "secondm@tc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
}

func ExampleStr_matchWithFailingInput_09() {
	toMatch := "secondmatch"
	failureInput := "secondm@tch"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
}

func ExampleStr_matchWithFailingInput_10() {
	toMatch := "secondmatch"
	failureInput := "secondm@tch123abc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch123abc"}
}

func ExampleStr_matchWithSucceedingInput_04() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
}

func ExampleStr_matchWithSucceedingInput_05() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatch123"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "123"}
}

func ExampleStr_matchWithSucceedingInput_06() {
	toMatch := "secondmatch"
	succeedingInput := "secondmatchabacadaba"
	expectedSuccessMatch := Str(toMatch)(succeedingInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "abacadaba"}
}
