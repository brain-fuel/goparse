# TODO

## Data Structures

- `MatchRes` (instead of tuple)
- `ParseRes`

### RuneMatch
- `Single(r rune) MatchRes` DONE
- `Not(r rune) MatchRes`

- `Any() MatchRes`
- `EOF() MatchRes`

- `InRange(low rune, hi rune) MatchRes`
- `NotInRange(low rune, hi rune) MatchRes`

- `Lower() MatchRes`
- `NotLower() MatchRes`

- `Upper() MatchRes`
- `NotUpper() MatchRes`

- `Digit() MatchRes`
- `NotDigit() MatchRes`

- `Space() MatchRes`
- `NotSpace() MatchRes`

- `Tab() MatchRes`
- `NotTab() MatchRes`

- `Newline() MatchRes`
- `NotNewline() MatchRes`

- `CR() MatchRes `-- optional
- `NotCR() MatchRes `-- optional

- `WS() MatchRes`
- `NotWS() MatchRes`

- `Alphanumeric() MatchRes`
- `NotAlphanumeric() MatchRes`

- `AlphanumericPlusUnderscore() MatchRes`
- `NotAlphanumericPlusUnderscore() MatchRes`

- `Underscore() MatchRes`
- `NotUnderscore() MatchRes`

- `Hyphen() MatchRes`
- `NotHyphen() MatchRes`

- `AnyOf(rs ...rune) MatchRes`
- `NoneOf(rs ...rune) MatchRes`

- `Exact(r rune) MatchRes`
- `NotExact(r rune) MatchRes `-- if not <r> <EOF>, then fails

### StringMatch
- `Match(str string) MatchRes`
- `Not(str string) MatchRes`

- `AnyOf(strs ...string) MatchRes`
- `NoneOf(strs ...string) MatchRes`

- `Exact(str string) MatchRes `-- if not <str> <EOF>, then fails
- `NotExact(str string) MatchRes`

### Match
- `Succeed() MatchRes`
- `Fail() MatchRes`

- `Match(mat Matcher) MatchRes`
- `Not(mat Matcher) MatchRes`

- `AnyOf(mats ...Matcher) MatchRes`
- `NoneOf(mats ...Matcher) MatchRes`

- `AllOf(mats ...Matcher) MatchRes`
- `NotAllOf(mats ...Matcher) MatchRes`

- `NOf(n int, mat Matcher) MatchRes`
- `NotNOf(n int, mat Matcher) MatchRes`

- `OneOf(mat Matcher) MatchRes`
- `NotOneOf(mat Matcher) MatchRes`

- `NOrMoreMatchers(n int, mats ...Matcher) MatchRes`
- `NotNOrMoreMatchers(n int, mats ...Matcher) MatchRes`

- `NOrLessMatchers(n int, mats ...Matcher) MatchRes`
- `NotNOrLessMatchers(n int, mats ...Matcher) MatchRes`

- `NToMMatchers(n int, m int, mats ...Matcher) MatchRes`
- `NotNToMMatchers(n int, m int, mats ...Matcher) MatchRes`

- `InOrder(mats ...Matcher) MatchRes`
- `NotInOrder(mats ...Matcher) MatchRes`

- `NOrMore(n int, mat Matcher) MatchRes`
- `NotNOrMore(n int, mat Matcher) MatchRes`

- `ZeroOrMore(mat Matcher) MatchRes`
- `NotZeroOrMore(mat Matcher) MatchRes`

- `OneOrMore(mat Matcher) MatchRes`
- `NotOneOrMore(mat Matcher) MatchRes`

- `NOrLess(n int, mat Matcher) MatchRes`
- `NotNOrLess(n int, mat Matcher) MatchRes`

- `Optional(mat Matcher) MatchRes` = `ZeroOrOne(mat Matcher) MatchRes` = `OneOrLess(mat Matcher) MatchRes`
- `Mandatory(mat Matcher) MatchRes` = `Match(mat Matcher) MatchRes`

- `BetweenRunes(r1 rune, r2 rune) MatchRes`
- `NotBetweenRunes(r1 rune, r2 rune) MatchRes`

- `BetweenStrings(s1 string, s2 string) MatchRes`
- `NotBetweenStrings(s1 string, s2 string) MatchRes`

- `Between(mat1 Matcher, mat2 Matcher) MatchRes`
- `NotBetween(mat1 Matcher, mat2 Matcher) MatchRes`

- `Alphanumeric() MatchRes` -- captures while condition is true
- `NotAlphanumeric() MatchRes`

- `AlphanumericPlusUnderscore() MatchRes` -- captures while condition is true
- `NotAlphanumericPlusUnderscore() MatchRes`

- `Until(mat Matcher) MatchRes`
- `SkipUntil(mat Matcher) MatchRes`

- `Regex(re Regexp) MatchRes`

- `EOF() MatchRes`
- `UntilEOF() MatchRes`

- `WordBoundary() MatchRes`
- `NotWordBoundary() MatchRes`

- `Word() MatchRes` -- same as AlphanumericPlusUnderscore
- `NotWord() MatchRes`

- `Lookahead(mat Matcher) MatchRes` = `Before(mat Matcher) MatchRes`
- `NotLookahead(mat Matcher) MatchRes` = `NotBefore(mat Matcher) MatchRes`

- `Exact(mat Matcher) MatchRes `-- if not <mat> <EOF>, then fails
- `NotExact(mat Matcher) MatchRes`

- `Whitespace() MatchRes `-- consumes all whitespace until non-whitespace rune
- `NotWhitespace() MatchRes `-- consumes all non-whitespace until whitespace rune

### SemanticMatch

## Parser


