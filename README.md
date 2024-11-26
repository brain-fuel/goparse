# GoParse

## Interface

### Rune Match (goforge.dev/tools/goparse/runematch)

## TODO

### Rune Match

- [x] `Single(r rune) datastructures.Matcher` [ ] code/test [ ] example
- [ ] `Single(r rune) datastructures.Matcher` example
- [x] `Not(r rune) datastructures.Matcher` code/test
- [ ] `Not(r rune) datastructures.Matcher` example

- [ ] `Any() datastructures.Matcher` code/test
- [ ] `Any() datastructures.Matcher` example
- [ ] `EOF() datastructures.Matcher` code/test
- [ ] `EOF() datastructures.Matcher` example

- [ ] `InRange(low rune, hi rune) datastructures.Matcher` code/test
- [ ] `InRange(low rune, hi rune) datastructures.Matcher` example
- [ ] `NotInRange(low rune, hi rune) datastructures.Matcher` code/test
- [ ] `NotInRange(low rune, hi rune) datastructures.Matcher` example

- [ ] `Lower() datastructures.Matcher` code/test
- [ ] `Lower() datastructures.Matcher` example
- [ ] `NotLower() datastructures.Matcher` code/test
- [ ] `NotLower() datastructures.Matcher` example

- [ ] `Upper() datastructures.Matcher` code/test
- [ ] `Upper() datastructures.Matcher` example
- [ ] `NotUpper() datastructures.Matcher` code/test
- [ ] `NotUpper() datastructures.Matcher` example

- [ ] `Digit() datastructures.Matcher` code/test
- [ ] `Digit() datastructures.Matcher` example
- [ ] `NotDigit() datastructures.Matcher` code/test
- [ ] `NotDigit() datastructures.Matcher` example

- [ ] `Space() datastructures.Matcher` code/test
- [ ] `Space() datastructures.Matcher` example
- [ ] `NotSpace() datastructures.Matcher` code/test
- [ ] `NotSpace() datastructures.Matcher` example

- [ ] `Tab() datastructures.Matcher` code/test
- [ ] `Tab() datastructures.Matcher` example
- [ ] `NotTab() datastructures.Matcher` code/test
- [ ] `NotTab() datastructures.Matcher` example

- [ ] `CarriageReturn() datastructures.Matcher` code/test
- [ ] `CarriageReturn() datastructures.Matcher` example
- [ ] `NotCarriageReturn() datastructures.Matcher` code/test
- [ ] `NotCarriageReturn() datastructures.Matcher` example

- [ ] `Newline() datastructures.Matcher` code/test
- [ ] `Newline() datastructures.Matcher` example
- [ ] `NotNewline() datastructures.Matcher` code/test
- [ ] `NotNewline() datastructures.Matcher` example

- [ ] `CR() datastructures.Matcher `-- optional code/test
- [ ] `CR() datastructures.Matcher `-- optional example
- [ ] `NotCR() datastructures.Matcher `-- optional code/test
- [ ] `NotCR() datastructures.Matcher `-- optional example

- [ ] `WS() datastructures.Matcher` code/test
- [ ] `WS() datastructures.Matcher` example
- [ ] `NotWS() datastructures.Matcher` code/test
- [ ] `NotWS() datastructures.Matcher` example

- [ ] `Alphanumeric() datastructures.Matcher` code/test
- [ ] `Alphanumeric() datastructures.Matcher` example
- [ ] `NotAlphanumeric() datastructures.Matcher` code/test
- [ ] `NotAlphanumeric() datastructures.Matcher` example

- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher` code/test
- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher` example
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher` code/test
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher` example

- [ ] `Underscore() datastructures.Matcher` code/test
- [ ] `Underscore() datastructures.Matcher` example
- [ ] `NotUnderscore() datastructures.Matcher` code/test
- [ ] `NotUnderscore() datastructures.Matcher` example

- [ ] `Hyphen() datastructures.Matcher` code/test
- [ ] `Hyphen() datastructures.Matcher` example
- [ ] `NotHyphen() datastructures.Matcher` code/test
- [ ] `NotHyphen() datastructures.Matcher` example

- [ ] `AnyOf(rs ...rune) datastructures.Matcher` code/test
- [ ] `AnyOf(rs ...rune) datastructures.Matcher` example
- [ ] `NoneOf(rs ...rune) datastructures.Matcher` code/test
- [ ] `NoneOf(rs ...rune) datastructures.Matcher` example

- [ ] `Exact(r rune) datastructures.Matcher` code/test
- [ ] `Exact(r rune) datastructures.Matcher` example
- [ ] `NotExact(r rune) datastructures.Matcher `-- if not <r> <EOF>, then fails code/test
- [ ] `NotExact(r rune) datastructures.Matcher `-- if not <r> <EOF>, then fails example

### String Match
- [ ] `Match(str string) datastructures.Matcher` code/test
- [ ] `Match(str string) datastructures.Matcher` example
- [ ] `Not(str string) datastructures.Matcher` code/test
- [ ] `Not(str string) datastructures.Matcher` example

- [ ] `AnyOf(strs ...string) datastructures.Matcher` code/test
- [ ] `AnyOf(strs ...string) datastructures.Matcher` example
- [ ] `NoneOf(strs ...string) datastructures.Matcher` code/test
- [ ] `NoneOf(strs ...string) datastructures.Matcher` example

- [ ] `EOF() datastructures.Matcher` code/test
- [ ] `EOF() datastructures.Matcher` example
- [ ] `NotEOF() datastructures.Matcher` code/test
- [ ] `NotEOF() datastructures.Matcher` example

- [ ] `Exact(str string) datastructures.Matcher `-- if not <str> <EOF>, then fails code/test
- [ ] `Exact(str string) datastructures.Matcher `-- if not <str> <EOF>, then fails example
- [ ] `NotExact(str string) datastructures.Matcher` code/test
- [ ] `NotExact(str string) datastructures.Matcher` example

### Match
- [ ] `SucceedAndAdvanceBy(n int) datastructures.Matcher` code/test
- [ ] `SucceedAndAdvanceBy(n int) datastructures.Matcher` example
- [ ] `FailAndAdvanceBy(n int) datastructures.Matcher` code/test
- [ ] `FailAndAdvanceBy(n int) datastructures.Matcher` example

- [ ] `Match(mat Matcher) datastructures.Matcher` code/test
- [ ] `Match(mat Matcher) datastructures.Matcher` example
- [ ] `Not(mat Matcher) datastructures.Matcher` code/test
- [ ] `Not(mat Matcher) datastructures.Matcher` example

- [ ] `AnyOf(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `AnyOf(mats ...Matcher) datastructures.Matcher` example
- [ ] `NoneOf(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NoneOf(mats ...Matcher) datastructures.Matcher` example

- [ ] `AllOf(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `AllOf(mats ...Matcher) datastructures.Matcher` example
- [ ] `NotAllOf(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NotAllOf(mats ...Matcher) datastructures.Matcher` example

- [ ] `NOf(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NOf(n int, mat Matcher) datastructures.Matcher` example
- [ ] `NotNOf(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NotNOf(n int, mat Matcher) datastructures.Matcher` example

- [ ] `ExactlyOneOf(ms ...Matcher) datastructures.Matcher` code/test
- [ ] `ExactlyOneOf(ms ...Matcher) datastructures.Matcher` example
- [ ] `NotExactlyOneOf(mat Matcher) datastructures.Matcher` code/test
- [ ] `NotExactlyOneOf(mat Matcher) datastructures.Matcher` example

- [ ] `NOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher` example
- [ ] `NotNOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NotNOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher` example

- [ ] `NOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher` example
- [ ] `NotNOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NotNOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher` example

- [ ] `NToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher` example
- [ ] `NotNToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NotNToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher` example

- [ ] `InOrder(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `InOrder(mats ...Matcher) datastructures.Matcher` example
- [ ] `NotInOrder(mats ...Matcher) datastructures.Matcher` code/test
- [ ] `NotInOrder(mats ...Matcher) datastructures.Matcher` example

- [ ] `NOrMore(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NOrMore(n int, mat Matcher) datastructures.Matcher` example
- [ ] `NotNOrMore(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NotNOrMore(n int, mat Matcher) datastructures.Matcher` example

- [ ] `ZeroOrMore(mat Matcher) datastructures.Matcher` code/test
- [ ] `ZeroOrMore(mat Matcher) datastructures.Matcher` example
- [ ] `NotZeroOrMore(mat Matcher) datastructures.Matcher` code/test
- [ ] `NotZeroOrMore(mat Matcher) datastructures.Matcher` example

- [ ] `OneOrMore(mat Matcher) datastructures.Matcher` code/test
- [ ] `OneOrMore(mat Matcher) datastructures.Matcher` example
- [ ] `NotOneOrMore(mat Matcher) datastructures.Matcher` code/test
- [ ] `NotOneOrMore(mat Matcher) datastructures.Matcher` example

- [ ] `NOrLess(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NOrLess(n int, mat Matcher) datastructures.Matcher` example
- [ ] `NotNOrLess(n int, mat Matcher) datastructures.Matcher` code/test
- [ ] `NotNOrLess(n int, mat Matcher) datastructures.Matcher` example

- [ ] `Optional(mat Matcher) datastructures.Matcher` = `ZeroOrOne(mat Matcher) datastructures.Matcher` = `OneOrLess(mat Matcher) datastructures.Matcher` code/test
- [ ] `Optional(mat Matcher) datastructures.Matcher` = `ZeroOrOne(mat Matcher) datastructures.Matcher` = `OneOrLess(mat Matcher) datastructures.Matcher` example
- [ ] `Mandatory(mat Matcher) datastructures.Matcher` = `Match(mat Matcher) datastructures.Matcher` code/test
- [ ] `Mandatory(mat Matcher) datastructures.Matcher` = `Match(mat Matcher) datastructures.Matcher` example

- [ ] `BetweenRunes(r1 rune, r2 rune) datastructures.Matcher` code/test
- [ ] `BetweenRunes(r1 rune, r2 rune) datastructures.Matcher` example
- [ ] `NotBetweenRunes(r1 rune, r2 rune) datastructures.Matcher` code/test
- [ ] `NotBetweenRunes(r1 rune, r2 rune) datastructures.Matcher` example

- [ ] `BetweenStrings(s1 string, s2 string) datastructures.Matcher` code/test
- [ ] `BetweenStrings(s1 string, s2 string) datastructures.Matcher` example
- [ ] `NotBetweenStrings(s1 string, s2 string) datastructures.Matcher` code/test
- [ ] `NotBetweenStrings(s1 string, s2 string) datastructures.Matcher` example

- [ ] `Between(mat1 Matcher, mat2 Matcher) datastructures.Matcher` code/test
- [ ] `Between(mat1 Matcher, mat2 Matcher) datastructures.Matcher` example
- [ ] `NotBetween(mat1 Matcher, mat2 Matcher) datastructures.Matcher` code/test
- [ ] `NotBetween(mat1 Matcher, mat2 Matcher) datastructures.Matcher` example

- [ ] `Alphanumeric() datastructures.Matcher` -- captures while condition is true code/test
- [ ] `Alphanumeric() datastructures.Matcher` -- captures while condition is true example
- [ ] `NotAlphanumeric() datastructures.Matcher` code/test
- [ ] `NotAlphanumeric() datastructures.Matcher` example

- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher` -- captures while condition is true code/test
- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher` -- captures while condition is true example
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher` code/test
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher` example

- [ ] `Until(mat Matcher) datastructures.Matcher` code/test
- [ ] `Until(mat Matcher) datastructures.Matcher` example
- [ ] `SkipUntil(mat Matcher) datastructures.Matcher` code/test
- [ ] `SkipUntil(mat Matcher) datastructures.Matcher` example

- [ ] `Regex(re Regexp) datastructures.Matcher` code/test
- [ ] `Regex(re Regexp) datastructures.Matcher` example

- [ ] `EOF() datastructures.Matcher` code/test
- [ ] `EOF() datastructures.Matcher` example
- [ ] `UntilEOF() datastructures.Matcher` code/test
- [ ] `UntilEOF() datastructures.Matcher` example

- [ ] `WordBoundary() datastructures.Matcher` code/test
- [ ] `WordBoundary() datastructures.Matcher` example
- [ ] `NotWordBoundary() datastructures.Matcher` code/test
- [ ] `NotWordBoundary() datastructures.Matcher` example

- [ ] `Word() datastructures.Matcher` -- same as AlphanumericPlusUnderscore code/test
- [ ] `Word() datastructures.Matcher` -- same as AlphanumericPlusUnderscore example
- [ ] `NotWord() datastructures.Matcher` code/test
- [ ] `NotWord() datastructures.Matcher` example

- [ ] `Lookahead(mat Matcher) datastructures.Matcher` = `Before(mat Matcher) datastructures.Matcher` code/test
- [ ] `Lookahead(mat Matcher) datastructures.Matcher` = `Before(mat Matcher) datastructures.Matcher` example
- [ ] `NotLookahead(mat Matcher) datastructures.Matcher` = `NotBefore(mat Matcher) datastructures.Matcher` code/test
- [ ] `NotLookahead(mat Matcher) datastructures.Matcher` = `NotBefore(mat Matcher) datastructures.Matcher` example

- [ ] `Exact(mat Matcher) datastructures.Matcher `-- if not <mat> <EOF>, then fails code/test
- [ ] `Exact(mat Matcher) datastructures.Matcher `-- if not <mat> <EOF>, then fails example
- [ ] `NotExact(mat Matcher) datastructures.Matcher` code/test
- [ ] `NotExact(mat Matcher) datastructures.Matcher` example

- [ ] `Whitespace() datastructures.Matcher `-- consumes all whitespace until non-whitespace rune code/test
- [ ] `Whitespace() datastructures.Matcher `-- consumes all whitespace until non-whitespace rune example
- [ ] `NotWhitespace() datastructures.Matcher `-- consumes all non-whitespace until whitespace rune code/test
- [ ] `NotWhitespace() datastructures.Matcher `-- consumes all non-whitespace until whitespace rune example

### Semantic Match

### Parser


