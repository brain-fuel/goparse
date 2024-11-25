# GoParse

## Interface

### Rune Match (goforge.dev/tools/goparse/runematch)
- `Single(r rune) datastructures.Matcher`

## TODO

### Rune Match
- `Not(r rune) datastructures.Matcher`

- `Any() datastructures.Matcher`
- `EOF() datastructures.Matcher`

- `InRange(low rune, hi rune) datastructures.Matcher`
- `NotInRange(low rune, hi rune) datastructures.Matcher`

- `Lower() datastructures.Matcher`
- `NotLower() datastructures.Matcher`

- `Upper() datastructures.Matcher`
- `NotUpper() datastructures.Matcher`

- `Digit() datastructures.Matcher`
- `NotDigit() datastructures.Matcher`

- `Space() datastructures.Matcher`
- `NotSpace() datastructures.Matcher`

- `Tab() datastructures.Matcher`
- `NotTab() datastructures.Matcher`

- `CarriageReturn() datastructures.Matcher`
- `NotCarriageReturn() datastructures.Matcher`

- `Newline() datastructures.Matcher`
- `NotNewline() datastructures.Matcher`

- `CR() datastructures.Matcher `-- optional
- `NotCR() datastructures.Matcher `-- optional

- `WS() datastructures.Matcher`
- `NotWS() datastructures.Matcher`

- `Alphanumeric() datastructures.Matcher`
- `NotAlphanumeric() datastructures.Matcher`

- `AlphanumericPlusUnderscore() datastructures.Matcher`
- `NotAlphanumericPlusUnderscore() datastructures.Matcher`

- `Underscore() datastructures.Matcher`
- `NotUnderscore() datastructures.Matcher`

- `Hyphen() datastructures.Matcher`
- `NotHyphen() datastructures.Matcher`

- `AnyOf(rs ...rune) datastructures.Matcher`
- `NoneOf(rs ...rune) datastructures.Matcher`

- `Exact(r rune) datastructures.Matcher`
- `NotExact(r rune) datastructures.Matcher `-- if not <r> <EOF>, then fails

### String Match
- `Match(str string) datastructures.Matcher`
- `Not(str string) datastructures.Matcher`

- `AnyOf(strs ...string) datastructures.Matcher`
- `NoneOf(strs ...string) datastructures.Matcher`

- `EOF() datastructures.Matcher`
- `NotEOF() datastructures.Matcher`

- `Exact(str string) datastructures.Matcher `-- if not <str> <EOF>, then fails
- `NotExact(str string) datastructures.Matcher`

### Match
- `Succeed() datastructures.Matcher`
- `Fail() datastructures.Matcher`

- `Match(mat Matcher) datastructures.Matcher`
- `Not(mat Matcher) datastructures.Matcher`

- `AnyOf(mats ...Matcher) datastructures.Matcher`
- `NoneOf(mats ...Matcher) datastructures.Matcher`

- `AllOf(mats ...Matcher) datastructures.Matcher`
- `NotAllOf(mats ...Matcher) datastructures.Matcher`

- `NOf(n int, mat Matcher) datastructures.Matcher`
- `NotNOf(n int, mat Matcher) datastructures.Matcher`

- `ExactlyOneOf(ms ...Matcher) datastructures.Matcher`
- `NotExactlyOneOf(mat Matcher) datastructures.Matcher`

- `NOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher`
- `NotNOrMoreMatchers(n int, mats ...Matcher) datastructures.Matcher`

- `NOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher`
- `NotNOrLessMatchers(n int, mats ...Matcher) datastructures.Matcher`

- `NToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher`
- `NotNToMMatchers(n int, m int, mats ...Matcher) datastructures.Matcher`

- `InOrder(mats ...Matcher) datastructures.Matcher`
- `NotInOrder(mats ...Matcher) datastructures.Matcher`

- `NOrMore(n int, mat Matcher) datastructures.Matcher`
- `NotNOrMore(n int, mat Matcher) datastructures.Matcher`

- `ZeroOrMore(mat Matcher) datastructures.Matcher`
- `NotZeroOrMore(mat Matcher) datastructures.Matcher`

- `OneOrMore(mat Matcher) datastructures.Matcher`
- `NotOneOrMore(mat Matcher) datastructures.Matcher`

- `NOrLess(n int, mat Matcher) datastructures.Matcher`
- `NotNOrLess(n int, mat Matcher) datastructures.Matcher`

- `Optional(mat Matcher) datastructures.Matcher` = `ZeroOrOne(mat Matcher) datastructures.Matcher` = `OneOrLess(mat Matcher) datastructures.Matcher`
- `Mandatory(mat Matcher) datastructures.Matcher` = `Match(mat Matcher) datastructures.Matcher`

- `BetweenRunes(r1 rune, r2 rune) datastructures.Matcher`
- `NotBetweenRunes(r1 rune, r2 rune) datastructures.Matcher`

- `BetweenStrings(s1 string, s2 string) datastructures.Matcher`
- `NotBetweenStrings(s1 string, s2 string) datastructures.Matcher`

- `Between(mat1 Matcher, mat2 Matcher) datastructures.Matcher`
- `NotBetween(mat1 Matcher, mat2 Matcher) datastructures.Matcher`

- `Alphanumeric() datastructures.Matcher` -- captures while condition is true
- `NotAlphanumeric() datastructures.Matcher`

- `AlphanumericPlusUnderscore() datastructures.Matcher` -- captures while condition is true
- `NotAlphanumericPlusUnderscore() datastructures.Matcher`

- `Until(mat Matcher) datastructures.Matcher`
- `SkipUntil(mat Matcher) datastructures.Matcher`

- `Regex(re Regexp) datastructures.Matcher`

- `EOF() datastructures.Matcher`
- `UntilEOF() datastructures.Matcher`

- `WordBoundary() datastructures.Matcher`
- `NotWordBoundary() datastructures.Matcher`

- `Word() datastructures.Matcher` -- same as AlphanumericPlusUnderscore
- `NotWord() datastructures.Matcher`

- `Lookahead(mat Matcher) datastructures.Matcher` = `Before(mat Matcher) datastructures.Matcher`
- `NotLookahead(mat Matcher) datastructures.Matcher` = `NotBefore(mat Matcher) datastructures.Matcher`

- `Exact(mat Matcher) datastructures.Matcher `-- if not <mat> <EOF>, then fails
- `NotExact(mat Matcher) datastructures.Matcher`

- `Whitespace() datastructures.Matcher `-- consumes all whitespace until non-whitespace rune
- `NotWhitespace() datastructures.Matcher `-- consumes all non-whitespace until whitespace rune

### Semantic Match

### Parser


