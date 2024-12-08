package predicates

import (
	"fmt"
)

func ExampleEOF() {
	failureInput := "abacadaba"

	successfulInput := ""

	expectedFailMatch := EOF()(failureInput)
	expectedSuccessMatch := EOF()(successfulInput)

	fmt.Println(expectedFailMatch)
	fmt.Printf("%v\n", expectedSuccessMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {0 9}, match: "", rest: "abacadaba"}
	// MatchRes{type: SUCCESS_EOF, dldist: 0, odldist: {0 0}, match: "", rest: ""}
}

func ExampleAnyRune() {
	failureInput := ""

	successfulInput1 := "abacadaba"
	successfulInput2 := "123abc"
	successfulInput3 := "世界"

	expectedFailMatch := AnyRune()(failureInput)
	expectedSuccessMatch1 := AnyRune()(successfulInput1)
	expectedSuccessMatch2 := AnyRune()(successfulInput2)
	expectedSuccessMatch3 := AnyRune()(successfulInput3)

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
	successfulInput := "世界"

	expectedFailMatch1 := Rune(toMatch)(failureInput1)
	expectedFailMatch2 := Rune(toMatch)(failureInput2)
	expectedFailMatch3 := Rune(toMatch)(failureInput3)
	expectedSuccessMatch := Rune(toMatch)(successfulInput)

	fmt.Println(expectedFailMatch1)
	fmt.Println(expectedFailMatch2)
	fmt.Println(expectedFailMatch3)
	fmt.Println(expectedSuccessMatch)

	// Output: MatchRes{type: FAILURE_EOF, dldist: 1, odldist: {0 1}, match: "", rest: ""}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {1 8}, match: "", rest: "abacadaba"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 6, odldist: {1 5}, match: "", rest: "123abc"}
	// MatchRes{type: SUCCESS_RUNE, dldist: 0, odldist: {0 0}, match: "世", rest: "界"}
}

func ExampleStr_fail01() {
	toMatch := "世界"
	failureInput := ""
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
}

func ExampleStr_fail02() {
	toMatch := "世界"
	failureInput := "abacadaba"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "abacadaba"}
}

func ExampleStr_fail03() {
	toMatch := "世界"
	failureInput := "123abc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "123abc"}
}

func ExampleStr_fail04() {
	toMatch := "世界"
	failureInput := "世"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_MATCH_THEN_EOF, dldist: 1, odldist: {0 1}, match: "", rest: "世"}
}

func ExampleStr_fail05() {
	toMatch := "世界"
	failureInput := "界"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 1, odldist: {1 1}, match: "", rest: "界"}
}

func ExampleStr_succ01() {
	toMatch := "世界"
	successfulInput := "世界"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: ""}
}

func ExampleStr_succ02() {
	toMatch := "世界"
	successfulInput := "世界abacadaba"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "abacadaba"}
}

func ExampleStr_succ03() {
	toMatch := "世界"
	successfulInput := "世界123abc"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "123abc"}
}

func ExampleStr_fail06() {
	toMatch := "secondmatch"
	failureInput := ""
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
}

func ExampleStr_fail07() {
	toMatch := "secondmatch"
	failureInput := "abacadaba"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
}

func ExampleStr_fail08() {
	toMatch := "secondmatch"
	failureInput := "secondm@tc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
}

func ExampleStr_fail09() {
	toMatch := "secondmatch"
	failureInput := "secondm@tch"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
}

func ExampleStr_fail10() {
	toMatch := "secondmatch"
	failureInput := "secondm@tch123abc"
	expectedFailMatch := Str(toMatch)(failureInput)
	fmt.Printf("Failure: %v\n", expectedFailMatch)
	// Output: Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch123abc"}
}

func ExampleStr_succ04() {
	toMatch := "secondmatch"
	successfulInput := "secondmatch"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
}

func ExampleStr_succ05() {
	toMatch := "secondmatch"
	successfulInput := "secondmatch123"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "123"}
}

func ExampleStr_succ06() {
	toMatch := "secondmatch"
	successfulInput := "secondmatchabacadaba"
	expectedSuccessMatch := Str(toMatch)(successfulInput)
	fmt.Printf("Success: %v\n", expectedSuccessMatch)
	// Output: Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "abacadaba"}
}
