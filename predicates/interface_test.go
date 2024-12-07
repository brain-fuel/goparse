package predicates

import (
	"fmt"
)

func ExampleEOF() {
	toFail := "abacadaba"

	toSucceed := ""

	expectedFailMatch := EOF()(toFail)
	expectedSuccessMatch := EOF()(toSucceed)

	fmt.Println(expectedFailMatch)
	fmt.Printf("%v\n", expectedSuccessMatch)
	// Output: MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {0 9}, match: "", rest: "abacadaba"}
	// MatchRes{type: SUCCESS_EOF, dldist: 0, odldist: {0 0}, match: "", rest: ""}
}

func ExampleAnyRune() {
	toFail := ""

	toSucceed1 := "abacadaba"
	toSucceed2 := "123abc"
	toSucceed3 := "世界"

	expectedFailMatch := AnyRune()(toFail)
	expectedSuccessMatch1 := AnyRune()(toSucceed1)
	expectedSuccessMatch2 := AnyRune()(toSucceed2)
	expectedSuccessMatch3 := AnyRune()(toSucceed3)

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

	toFail1 := ""
	toFail2 := "abacadaba"
	toFail3 := "123abc"
	toSucceed := "世界"

	expectedFailMatch1 := Rune(toMatch)(toFail1)
	expectedFailMatch2 := Rune(toMatch)(toFail2)
	expectedFailMatch3 := Rune(toMatch)(toFail3)
	expectedSuccessMatch := Rune(toMatch)(toSucceed)

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

	toFail1_1 := ""
	toFail1_2 := "abacadaba"
	toFail1_3 := "123abc"
	toFail1_4 := "世"
	toFail1_5 := "界"
	expectedFailMatch1_1 := Str(toMatch1)(toFail1_1)
	expectedFailMatch1_2 := Str(toMatch1)(toFail1_2)
	expectedFailMatch1_3 := Str(toMatch1)(toFail1_3)
	expectedFailMatch1_4 := Str(toMatch1)(toFail1_4)
	expectedFailMatch1_5 := Str(toMatch1)(toFail1_5)
	fmt.Println(expectedFailMatch1_1)
	fmt.Println(expectedFailMatch1_2)
	fmt.Println(expectedFailMatch1_3)
	fmt.Println(expectedFailMatch1_4)
	fmt.Println(expectedFailMatch1_5)

	toSucceed1_1 := "世界"
	toSucceed1_2 := "世界abacadaba"
	toSucceed1_3 := "世界123abc"
	expectedSuccessMatch1_1 := Str(toMatch1)(toSucceed1_1)
	expectedSuccessMatch1_2 := Str(toMatch1)(toSucceed1_2)
	expectedSuccessMatch1_3 := Str(toMatch1)(toSucceed1_3)
	fmt.Println(expectedSuccessMatch1_1)
	fmt.Println(expectedSuccessMatch1_2)
	fmt.Println(expectedSuccessMatch1_3)

	toMatch2 := "secondmatch"

	toFail2_1 := ""
	toFail2_2 := "abacadaba"
	toFail2_3 := "secondm@tc"
	toFail2_4 := "secondm@tch"
	toFail2_5 := "secondm@tch123abc"
	expectedFailMatch2_1 := Str(toMatch2)(toFail2_1)
	expectedFailMatch2_2 := Str(toMatch2)(toFail2_2)
	expectedFailMatch2_3 := Str(toMatch2)(toFail2_3)
	expectedFailMatch2_4 := Str(toMatch2)(toFail2_4)
	expectedFailMatch2_5 := Str(toMatch2)(toFail2_5)
	fmt.Println(expectedFailMatch2_1)
	fmt.Println(expectedFailMatch2_2)
	fmt.Println(expectedFailMatch2_3)
	fmt.Println(expectedFailMatch2_4)
	fmt.Println(expectedFailMatch2_5)

	toSucceed2_1 := "secondmatch"
	toSucceed2_2 := "secondmatch123"
	toSucceed2_3 := "secondmatchabacadaba"
	expectedSuccessMatch2_1 := Str(toMatch2)(toSucceed2_1)
	expectedSuccessMatch2_2 := Str(toMatch2)(toSucceed2_2)
	expectedSuccessMatch2_3 := Str(toMatch2)(toSucceed2_3)
	fmt.Println(expectedSuccessMatch2_1)
	fmt.Println(expectedSuccessMatch2_2)
	fmt.Println(expectedSuccessMatch2_3)

	// Output: MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 2}, match: "", rest: ""}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {2 7}, match: "", rest: "abacadaba"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 6, odldist: {2 4}, match: "", rest: "123abc"}
	// MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {0 1}, match: "", rest: "世"}
	// MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "界"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 2, odldist: {0 0}, match: "", rest: "世界"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 11, odldist: {0 9}, match: "", rest: "世界abacadaba"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 8, odldist: {0 6}, match: "", rest: "世界123abc"}
	// MatchRes{type: FAILURE_EOF, dldist: 11, odldist: {0 11}, match: "", rest: ""}
	// MatchRes{type: FAILURE_EOF, dldist: 9, odldist: {8 2}, match: "", rest: "abacadaba"}
	// MatchRes{type: FAILURE_EOF, dldist: 2, odldist: {1 1}, match: "", rest: "secondm@tc"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 1, odldist: {1 0}, match: "", rest: "secondm@tch"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 7, odldist: {1 6}, match: "", rest: "secondm@tch123abc"}
	// MatchRes{type: SUCCESS_STRING, dldist: 0, odldist: {0 0}, match: "secondmatch", rest: ""}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 3, odldist: {0 3}, match: "", rest: "secondmatch123"}
	// MatchRes{type: FAILURE_NO_MATCH, dldist: 9, odldist: {0 9}, match: "", rest: "secondmatchabacadaba"}
}
