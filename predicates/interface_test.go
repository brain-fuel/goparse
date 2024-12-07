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

func ExampleStr() {
	toMatch1 := "世界"

	failureInput1_1 := ""
	expectedFailMatch1_1 := Str(toMatch1)(failureInput1_1)
	fmt.Printf("Failure: %v\n", expectedFailMatch1_1)

	failureInput1_2 := "abacadaba"
	expectedFailMatch1_2 := Str(toMatch1)(failureInput1_2)
	fmt.Printf("Failure: %v\n", expectedFailMatch1_2)

	failureInput1_3 := "123abc"
	expectedFailMatch1_3 := Str(toMatch1)(failureInput1_3)
	fmt.Printf("Failure: %v\n", expectedFailMatch1_3)

	failureInput1_4 := "世"
	expectedFailMatch1_4 := Str(toMatch1)(failureInput1_4)
	fmt.Printf("Failure: %v\n", expectedFailMatch1_4)

	failureInput1_5 := "界"
	expectedFailMatch1_5 := Str(toMatch1)(failureInput1_5)
	fmt.Printf("Failure: %v\n", expectedFailMatch1_5)

	successfulInput1_1 := "世界"
	expectedSuccessMatch1_1 := Str(toMatch1)(successfulInput1_1)
	fmt.Printf("Success: %v\n", expectedSuccessMatch1_1)

	successfulInput1_2 := "世界abacadaba"
	expectedSuccessMatch1_2 := Str(toMatch1)(successfulInput1_2)
	fmt.Printf("Success: %v\n", expectedSuccessMatch1_2)

	successfulInput1_3 := "世界123abc"
	expectedSuccessMatch1_3 := Str(toMatch1)(successfulInput1_3)
	fmt.Printf("Success: %v\n", expectedSuccessMatch1_3)

	// --------------------
	toMatch2 := "secondmatch"

	failureInput2_1 := ""
	expectedFailMatch2_1 := Str(toMatch2)(failureInput2_1)
	fmt.Printf("Failure: %v\n", expectedFailMatch2_1)

	failureInput2_2 := "abacadaba"
	expectedFailMatch2_2 := Str(toMatch2)(failureInput2_2)
	fmt.Printf("Failure: %v\n", expectedFailMatch2_2)

	failureInput2_3 := "secondm@tc"
	expectedFailMatch2_3 := Str(toMatch2)(failureInput2_3)
	fmt.Printf("Failure: %v\n", expectedFailMatch2_3)

	failureInput2_4 := "secondm@tch"
	expectedFailMatch2_4 := Str(toMatch2)(failureInput2_4)
	fmt.Printf("Failure: %v\n", expectedFailMatch2_4)

	failureInput2_5 := "secondm@tch123abc"
	expectedFailMatch2_5 := Str(toMatch2)(failureInput2_5)
	fmt.Printf("Failure: %v\n", expectedFailMatch2_5)

	successfulInput2_1 := "secondmatch"
	expectedSuccessMatch2_1 := Str(toMatch2)(successfulInput2_1)
	fmt.Printf("Success: %v\n", expectedSuccessMatch2_1)

	successfulInput2_2 := "secondmatch123"
	expectedSuccessMatch2_2 := Str(toMatch2)(successfulInput2_2)
	fmt.Printf("Success: %v\n", expectedSuccessMatch2_2)

	successfulInput2_3 := "secondmatchabacadaba"
	expectedSuccessMatch2_3 := Str(toMatch2)(successfulInput2_3)
	fmt.Printf("Success: %v\n", expectedSuccessMatch2_3)

	// Output: Failure: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
	// Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "abacadaba"}
	// Failure: MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {2 0}, match: "", rest: "123abc"}
	// Failure: MatchRes{type: FAILURE_MATCH_THEN_EOF, dldist: 1, odldist: {0 1}, match: "", rest: "世"}
	// Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 1, odldist: {1 1}, match: "", rest: "界"}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: ""}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "abacadaba"}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "世界", rest: "123abc"}
	// Failure: MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
	// Failure: MatchRes{type: FAILURE_NO_MATCH_THEN_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
	// Failure: MatchRes{type: FAILURE_NEAR_MISS_THEN_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
	// Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
	// Failure: MatchRes{type: FAILURE_NEAR_MISS, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch123abc"}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "123"}
	// Success: MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: "abacadaba"}
}
