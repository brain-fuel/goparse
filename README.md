# GoParse

## Gotchas

- The only matchers that match on EOF are:
  - runematch.EOF()
  - stringmatch.EOF()
  - match.EOF()
  - any semantic match using one of the aforementioned EOF matchers.

- Dependency Order amongst libraries should be:
  - runematch&lt;stringmatch&lt;match&lt;semanticmatch&lt;parse
  - Nothing on the left can have anything to its right as a dependency.

## Interface

### Data Structures (goforge.dev/tools/goparse/datastructures)

### Rune Match (goforge.dev/tools/goparse/runematch)

### String Match (goforge.dev/tools/goparse/stringmatch)

### Match (goforge.dev/tools/goparse/match)

### Semantic Match (goforge.dev/tools/goparse/semanticmatch)

### Parse (goforge.dev/tools/goparse/parse)

## TODO

### Rune Match

- [ ] `EOF() datastructures.Matcher`
  - [x] code/test
  - [ ] example
- [ ] `Any() datastructures.Matcher`
  - [x] code/test
  - [ ] example

- [ ] `Single(r rune) datastructures.Matcher`
  - [x] code/test
  - [ ] example
- [ ] `Not(r rune) datastructures.Matcher`
  - [x] code/test
  - [ ] example

- [ ] `InRange(low rune, hi rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotInRange(low rune, hi rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Lower() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotLower() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Upper() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotUpper() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Digit() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotDigit() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Space() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotSpace() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Tab() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotTab() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Newline() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNewline() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `CR() datastructures.Matcher `-- optional
  - [ ] code/test
  - [ ] example
- [ ] `NotCR() datastructures.Matcher `-- optional
  - [ ] code/test
  - [ ] example

- [ ] `WS() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotWS() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Alphanumeric() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotAlphanumeric() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Underscore() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotUnderscore() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Hyphen() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotHyphen() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AnyOf(rs ...rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NoneOf(rs ...rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Exact(r rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotExact(r rune) datastructures.Matcher `-- if not <r> <EOF>, then fails
  - [ ] code/test
  - [ ] example

### String Match
- [ ] `Match(str string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `Not(str string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AnyOf(strs ...string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NoneOf(strs ...string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `EOF() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotEOF() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Exact(str string) datastructures.Matcher `-- if not <str> <EOF>, then fails
  - [ ] code/test
  - [ ] example
- [ ] `NotExact(str string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

### Match
- [ ] `SucceedAndAdvanceBy(n int) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `FailAndAdvanceBy(n int) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Match(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `Not(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AnyOf(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NoneOf(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AllOf(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotAllOf(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NOf(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNOf(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `ExactlyOneOf(ms ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotExactlyOneOf(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `InOrder(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotInOrder(mats ...Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NOrMore(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNOrMore(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `ZeroOrMore(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotZeroOrMore(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `OneOrMore(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotOneOrMore(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `NOrLess(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotNOrLess(n int, mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Optional(mat Matcher) datastructures.Matcher` = `ZeroOrOne(mat Matcher) datastructures.Matcher` = `OneOrLess(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `Mandatory(mat Matcher) datastructures.Matcher` = `Match(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `BetweenRunes(r1 rune, r2 rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotBetweenRunes(r1 rune, r2 rune) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `BetweenStrings(s1 string, s2 string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotBetweenStrings(s1 string, s2 string) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Between(mat1 Matcher, mat2 Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotBetween(mat1 Matcher, mat2 Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Alphanumeric() datastructures.Matcher` -- captures while condition is true
  - [ ] code/test
  - [ ] example
- [ ] `NotAlphanumeric() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `AlphanumericPlusUnderscore() datastructures.Matcher` -- captures while condition is true
  - [ ] code/test
  - [ ] example
- [ ] `NotAlphanumericPlusUnderscore() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Until(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `SkipUntil(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Regex(re Regexp) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `EOF() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `UntilEOF() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `WordBoundary() datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotWordBoundary() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Word() datastructures.Matcher` -- same as AlphanumericPlusUnderscore
  - [ ] code/test
  - [ ] example
- [ ] `NotWord() datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Lookahead(mat Matcher) datastructures.Matcher` = `Before(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example
- [ ] `NotLookahead(mat Matcher) datastructures.Matcher` = `NotBefore(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Exact(mat Matcher) datastructures.Matcher `-- if not <mat> <EOF>, then fails
  - [ ] code/test
  - [ ] example
- [ ] `NotExact(mat Matcher) datastructures.Matcher`
  - [ ] code/test
  - [ ] example

- [ ] `Whitespace() datastructures.Matcher `-- consumes all whitespace until non-whitespace rune
  - [ ] code/test
  - [ ] example
- [ ] `NotWhitespace() datastructures.Matcher `-- consumes all non-whitespace until whitespace rune
  - [ ] code/test
  - [ ] example

### Semantic Match

### Parser


