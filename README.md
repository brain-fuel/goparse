# GoParse

## Interface

### Rune Match (goforge.dev/tools/goparse/runematch)
- `Single(r rune) (ds.Match, ds.MatcherInput, error)`

## TODO

### Data Structures

- `(ds.Match, ds.MatcherInput, error)` (instead of tuple)
- `ParseRes`

### Rune Match
- `Not(r rune) (ds.Match, ds.MatcherInput, error)`

- `Any() (ds.Match, ds.MatcherInput, error)`
- `EOF() (ds.Match, ds.MatcherInput, error)`

- `InRange(low rune, hi rune) (ds.Match, ds.MatcherInput, error)`
- `NotInRange(low rune, hi rune) (ds.Match, ds.MatcherInput, error)`

- `Lower() (ds.Match, ds.MatcherInput, error)`
- `NotLower() (ds.Match, ds.MatcherInput, error)`

- `Upper() (ds.Match, ds.MatcherInput, error)`
- `NotUpper() (ds.Match, ds.MatcherInput, error)`

- `Digit() (ds.Match, ds.MatcherInput, error)`
- `NotDigit() (ds.Match, ds.MatcherInput, error)`

- `Space() (ds.Match, ds.MatcherInput, error)`
- `NotSpace() (ds.Match, ds.MatcherInput, error)`

- `Tab() (ds.Match, ds.MatcherInput, error)`
- `NotTab() (ds.Match, ds.MatcherInput, error)`

- `Newline() (ds.Match, ds.MatcherInput, error)`
- `NotNewline() (ds.Match, ds.MatcherInput, error)`

- `CR() (ds.Match, ds.MatcherInput, error) `-- optional
- `NotCR() (ds.Match, ds.MatcherInput, error) `-- optional

- `WS() (ds.Match, ds.MatcherInput, error)`
- `NotWS() (ds.Match, ds.MatcherInput, error)`

- `Alphanumeric() (ds.Match, ds.MatcherInput, error)`
- `NotAlphanumeric() (ds.Match, ds.MatcherInput, error)`

- `AlphanumericPlusUnderscore() (ds.Match, ds.MatcherInput, error)`
- `NotAlphanumericPlusUnderscore() (ds.Match, ds.MatcherInput, error)`

- `Underscore() (ds.Match, ds.MatcherInput, error)`
- `NotUnderscore() (ds.Match, ds.MatcherInput, error)`

- `Hyphen() (ds.Match, ds.MatcherInput, error)`
- `NotHyphen() (ds.Match, ds.MatcherInput, error)`

- `AnyOf(rs ...rune) (ds.Match, ds.MatcherInput, error)`
- `NoneOf(rs ...rune) (ds.Match, ds.MatcherInput, error)`

- `Exact(r rune) (ds.Match, ds.MatcherInput, error)`
- `NotExact(r rune) (ds.Match, ds.MatcherInput, error) `-- if not <r> <EOF>, then fails

### String Match
- `Match(str string) (ds.Match, ds.MatcherInput, error)`
- `Not(str string) (ds.Match, ds.MatcherInput, error)`

- `AnyOf(strs ...string) (ds.Match, ds.MatcherInput, error)`
- `NoneOf(strs ...string) (ds.Match, ds.MatcherInput, error)`

- `Exact(str string) (ds.Match, ds.MatcherInput, error) `-- if not <str> <EOF>, then fails
- `NotExact(str string) (ds.Match, ds.MatcherInput, error)`

### Match
- `Succeed() (ds.Match, ds.MatcherInput, error)`
- `Fail() (ds.Match, ds.MatcherInput, error)`

- `Match(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `Not(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `AnyOf(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NoneOf(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `AllOf(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotAllOf(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `NOf(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNOf(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `OneOf(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotOneOf(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `NOrMoreMatchers(n int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNOrMoreMatchers(n int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `NOrLessMatchers(n int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNOrLessMatchers(n int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `NToMMatchers(n int, m int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNToMMatchers(n int, m int, mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `InOrder(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotInOrder(mats ...Matcher) (ds.Match, ds.MatcherInput, error)`

- `NOrMore(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNOrMore(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `ZeroOrMore(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotZeroOrMore(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `OneOrMore(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotOneOrMore(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `NOrLess(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotNOrLess(n int, mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `Optional(mat Matcher) (ds.Match, ds.MatcherInput, error)` = `ZeroOrOne(mat Matcher) (ds.Match, ds.MatcherInput, error)` = `OneOrLess(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `Mandatory(mat Matcher) (ds.Match, ds.MatcherInput, error)` = `Match(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `BetweenRunes(r1 rune, r2 rune) (ds.Match, ds.MatcherInput, error)`
- `NotBetweenRunes(r1 rune, r2 rune) (ds.Match, ds.MatcherInput, error)`

- `BetweenStrings(s1 string, s2 string) (ds.Match, ds.MatcherInput, error)`
- `NotBetweenStrings(s1 string, s2 string) (ds.Match, ds.MatcherInput, error)`

- `Between(mat1 Matcher, mat2 Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotBetween(mat1 Matcher, mat2 Matcher) (ds.Match, ds.MatcherInput, error)`

- `Alphanumeric() (ds.Match, ds.MatcherInput, error)` -- captures while condition is true
- `NotAlphanumeric() (ds.Match, ds.MatcherInput, error)`

- `AlphanumericPlusUnderscore() (ds.Match, ds.MatcherInput, error)` -- captures while condition is true
- `NotAlphanumericPlusUnderscore() (ds.Match, ds.MatcherInput, error)`

- `Until(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `SkipUntil(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `Regex(re Regexp) (ds.Match, ds.MatcherInput, error)`

- `EOF() (ds.Match, ds.MatcherInput, error)`
- `UntilEOF() (ds.Match, ds.MatcherInput, error)`

- `WordBoundary() (ds.Match, ds.MatcherInput, error)`
- `NotWordBoundary() (ds.Match, ds.MatcherInput, error)`

- `Word() (ds.Match, ds.MatcherInput, error)` -- same as AlphanumericPlusUnderscore
- `NotWord() (ds.Match, ds.MatcherInput, error)`

- `Lookahead(mat Matcher) (ds.Match, ds.MatcherInput, error)` = `Before(mat Matcher) (ds.Match, ds.MatcherInput, error)`
- `NotLookahead(mat Matcher) (ds.Match, ds.MatcherInput, error)` = `NotBefore(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `Exact(mat Matcher) (ds.Match, ds.MatcherInput, error) `-- if not <mat> <EOF>, then fails
- `NotExact(mat Matcher) (ds.Match, ds.MatcherInput, error)`

- `Whitespace() (ds.Match, ds.MatcherInput, error) `-- consumes all whitespace until non-whitespace rune
- `NotWhitespace() (ds.Match, ds.MatcherInput, error) `-- consumes all non-whitespace until whitespace rune

### Semantic Match

### Parser


