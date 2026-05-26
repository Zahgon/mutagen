// Subset of k8s.io/apimachinery vendored to avoid pulling in the full
// Kubernetes dependency tree (including klog and its init-time
// goroutines). Originally extracted from the following revision:
// https://github.com/kubernetes/apimachinery/tree/f916759cb6b8547418dc7708876ecab5c1961448
//
// The original code license:
//
// Copyright 2014 The Kubernetes Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// The original license header inside the code itself:
//

/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package labels

import (
	"github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/selection"
	"github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/util/sets"
	"github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/util/validation/field"
)

var (
	validRequirementOperators = []string{
		string(selection.In), string(selection.NotIn),
		string(selection.Equals), string(selection.DoubleEquals), string(selection.NotEquals),
		string(selection.Exists), string(selection.DoesNotExist),
		string(selection.GreaterThan), string(selection.LessThan),
	}
)

// Requirements is AND of all requirements.
type Requirements []Requirement

// Selector represents a label selector.
type Selector interface {
	// Matches returns true if this selector matches the given set of labels.
	Matches(Labels) bool

	// Empty returns true if this selector does not restrict the selection space.
	Empty() bool

	// String returns a human readable string that represents this selector.
	String() string

	// Add adds requirements to the Selector
	Add(r ...Requirement) Selector

	// Requirements converts this interface into Requirements to expose
	// more detailed selection information.
	// If there are querying parameters, it will return converted requirements and selectable=true.
	// If this selector doesn't want to select anything, it will return selectable=false.
	Requirements() (requirements Requirements, selectable bool)

	// Make a deep copy of the selector.
	DeepCopySelector() Selector

	// RequiresExactMatch allows a caller to introspect whether a given selector
	// requires a single specific label to be set, and if so returns the value it
	// requires.
	RequiresExactMatch(label string) (value string, found bool)
}

// Everything returns a selector that matches all labels.
func Everything() Selector { _ = "STUB: not implemented"; return *new(Selector) }

type nothingSelector struct{}

func (n nothingSelector) Matches(_ Labels) bool { _ = "STUB: not implemented"; return false }
func (n nothingSelector) Empty() bool           { _ = "STUB: not implemented"; return false }
func (n nothingSelector) String() string        { _ = "STUB: not implemented"; return "" }
func (n nothingSelector) Add(_ ...Requirement) Selector {
	_ = "STUB: not implemented"
	return *new(Selector)
}
func (n nothingSelector) Requirements() (Requirements, bool) {
	_ = "STUB: not implemented"
	return *new(Requirements), false
}
func (n nothingSelector) DeepCopySelector() Selector {
	_ = "STUB: not implemented"
	return *new(Selector)
}
func (n nothingSelector) RequiresExactMatch(label string) (value string, found bool) {
	_ = "STUB: not implemented"

	// Nothing returns a selector that matches no labels
	return "", false
}

func Nothing() Selector {
	_ = "STUB: not implemented"
	return *

	// NewSelector returns a nil selector
	new(Selector)
}

func NewSelector() Selector { _ = "STUB: not implemented"; return *new(Selector) }

type internalSelector []Requirement

func (s internalSelector) DeepCopy() internalSelector {
	_ = "STUB: not implemented"
	return *new(internalSelector)
}

func (s internalSelector) DeepCopySelector() Selector {
	_ = "STUB: not implemented"
	return *

	// ByKey sorts requirements by key to obtain deterministic parser
	new(Selector)
}

type ByKey []Requirement

func (a ByKey) Len() int { _ = "STUB: not implemented"; return 0 }

func (a ByKey) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (a ByKey) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Requirement contains values, a key, and an operator that relates the key and values.
// The zero value of Requirement is invalid.
// Requirement implements both set based match and exact match
// Requirement should be initialized via NewRequirement constructor for creating a valid Requirement.
// +k8s:deepcopy-gen=true
type Requirement struct {
	key      string
	operator selection.Operator
	// In huge majority of cases we have at most one value here.
	// It is generally faster to operate on a single-element slice
	// than on a single-element map, so we have a slice here.
	strValues []string
}

// NewRequirement is the constructor for a Requirement.
// If any of these rules is violated, an error is returned:
// (1) The operator can only be In, NotIn, Equals, DoubleEquals, NotEquals, Exists, or DoesNotExist.
// (2) If the operator is In or NotIn, the values set must be non-empty.
// (3) If the operator is Equals, DoubleEquals, or NotEquals, the values set must contain one value.
// (4) If the operator is Exists or DoesNotExist, the value set must be empty.
// (5) If the operator is Gt or Lt, the values set must contain only one value, which will be interpreted as an integer.
// (6) The key is invalid due to its length, or sequence
//
//	of characters. See validateLabelKey for more details.
//
// The empty string is a valid value in the input values set.
// Returned error, if not nil, is guaranteed to be an aggregated field.ErrorList
func NewRequirement(key string, op selection.Operator, vals []string, opts ...field.PathOption) (*Requirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Requirement) hasValue(value string) bool { _ = "STUB: not implemented"; return false }

// Matches returns true if the Requirement matches the input Labels.
// There is a match in the following cases:
// (1) The operator is Exists and Labels has the Requirement's key.
// (2) The operator is In, Labels has the Requirement's key and Labels'
//
//	value for that key is in Requirement's value set.
//
// (3) The operator is NotIn, Labels has the Requirement's key and
//
//	Labels' value for that key is not in Requirement's value set.
//
// (4) The operator is DoesNotExist or NotIn and Labels does not have the
//
//	Requirement's key.
//
// (5) The operator is GreaterThanOperator or LessThanOperator, and Labels has
//
//	the Requirement's key and the corresponding value satisfies mathematical inequality.
func (r *Requirement) Matches(ls Labels) bool { _ = "STUB: not implemented"; return false }

// There should be only one strValue in r.strValues, and can be converted to an integer.

// Key returns requirement key
func (r *Requirement) Key() string {
	_ = "STUB: not implemented"

	// Operator returns requirement operator
	return ""
}

func (r *Requirement) Operator() selection.Operator {
	_ = "STUB: not implemented"

	// Values returns requirement values
	return *new(selection.Operator)
}

func (r *Requirement) Values() sets.String { _ = "STUB: not implemented"; return *new(sets.String) }

// Equal checks the equality of requirement.
func (r Requirement) Equal(x Requirement) bool { _ = "STUB: not implemented"; return false }

// Empty returns true if the internalSelector doesn't restrict selection space
func (s internalSelector) Empty() bool { _ = "STUB: not implemented"; return false }

// String returns a human-readable string that represents this
// Requirement. If called on an invalid Requirement, an error is
// returned. See NewRequirement for creating a valid Requirement.
func (r *Requirement) String() string { _ = "STUB: not implemented"; return "" }

// length of r.key

// length of 'r.operator' + 2 spaces for the worst case ('in' and 'notin')

// length of 'r.strValues' slice times. Heuristically 5 chars per word

// only > 1 since == 0 prohibited by NewRequirement
// normalizes value order on output, without mutating the in-memory selector representation
// also avoids normalization when it is not required, and ensures we do not mutate shared data

// safeSort sorts input strings without modification
func safeSort(in []string) []string { _ = "STUB: not implemented"; return nil }

// Add adds requirements to the selector. It copies the current selector returning a new one
func (s internalSelector) Add(reqs ...Requirement) Selector {
	_ = "STUB: not implemented"
	return *new(Selector)
}

// Matches for a internalSelector returns true if all
// its Requirements match the input Labels. If any
// Requirement does not match, false is returned.
func (s internalSelector) Matches(l Labels) bool { _ = "STUB: not implemented"; return false }

func (s internalSelector) Requirements() (Requirements, bool) {
	_ = "STUB: not implemented"
	return *new(Requirements), false
}

// String returns a comma-separated string of all
// the internalSelector Requirements' human-readable strings.
func (s internalSelector) String() string { _ = "STUB: not implemented"; return "" }

// RequiresExactMatch introspects whether a given selector requires a single specific field
// to be set, and if so returns the value it requires.
func (s internalSelector) RequiresExactMatch(label string) (value string, found bool) {
	_ = "STUB: not implemented"
	return "", false
}

// Token represents constant definition for lexer token
type Token int

const (
	// ErrorToken represents scan error
	ErrorToken Token = iota
	// EndOfStringToken represents end of string
	EndOfStringToken
	// ClosedParToken represents close parenthesis
	ClosedParToken
	// CommaToken represents the comma
	CommaToken
	// DoesNotExistToken represents logic not
	DoesNotExistToken
	// DoubleEqualsToken represents double equals
	DoubleEqualsToken
	// EqualsToken represents equal
	EqualsToken
	// GreaterThanToken represents greater than
	GreaterThanToken
	// IdentifierToken represents identifier, e.g. keys and values
	IdentifierToken
	// InToken represents in
	InToken
	// LessThanToken represents less than
	LessThanToken
	// NotEqualsToken represents not equal
	NotEqualsToken
	// NotInToken represents not in
	NotInToken
	// OpenParToken represents open parenthesis
	OpenParToken
)

// string2token contains the mapping between lexer Token and token literal
// (except IdentifierToken, EndOfStringToken and ErrorToken since it makes no sense)
var string2token = map[string]Token{
	")":     ClosedParToken,
	",":     CommaToken,
	"!":     DoesNotExistToken,
	"==":    DoubleEqualsToken,
	"=":     EqualsToken,
	">":     GreaterThanToken,
	"in":    InToken,
	"<":     LessThanToken,
	"!=":    NotEqualsToken,
	"notin": NotInToken,
	"(":     OpenParToken,
}

// ScannedItem contains the Token and the literal produced by the lexer.
type ScannedItem struct {
	tok     Token
	literal string
}

// isWhitespace returns true if the rune is a space, tab, or newline.
func isWhitespace(ch byte) bool { _ = "STUB: not implemented"; return false }

// isSpecialSymbol detects if the character ch can be an operator
func isSpecialSymbol(ch byte) bool { _ = "STUB: not implemented"; return false }

// Lexer represents the Lexer struct for label selector.
// It contains necessary informationt to tokenize the input string
type Lexer struct {
	// s stores the string to be tokenized
	s string
	// pos is the position currently tokenized
	pos int
}

// read returns the character currently lexed
// increment the position and check the buffer overflow
func (l *Lexer) read() (b byte) { _ = "STUB: not implemented"; return 0 }

// unread 'undoes' the last read character
func (l *Lexer) unread() {
	_ = "STUB: not implemented"

	// scanIDOrKeyword scans string to recognize literal token (for example 'in') or an identifier.
	return
}

func (l *Lexer) scanIDOrKeyword() (tok Token, lit string) {
	_ = "STUB: not implemented"
	return *new(Token), ""
}

// is a literal token?

// otherwise is an identifier

// scanSpecialSymbol scans string starting with special symbol.
// special symbol identify non literal operators. "!=", "==", "="
func (l *Lexer) scanSpecialSymbol() (Token, string) {
	_ = "STUB: not implemented"
	return *new(Token), ""
}

// skipWhiteSpaces consumes all blank characters
// returning the first non blank character
func (l *Lexer) skipWhiteSpaces(ch byte) byte { _ = "STUB: not implemented"; return 0 }

// Lex returns a pair of Token and the literal
// literal is meaningfull only for IdentifierToken token
func (l *Lexer) Lex() (tok Token, lit string) { _ = "STUB: not implemented"; return *new(Token), "" }

// Parser data structure contains the label selector parser data structure
type Parser struct {
	l            *Lexer
	scannedItems []ScannedItem
	position     int
	path         *field.Path
}

// ParserContext represents context during parsing:
// some literal for example 'in' and 'notin' can be
// recognized as operator for example 'x in (a)' but
// it can be recognized as value for example 'value in (in)'
type ParserContext int

const (
	// KeyAndOperator represents key and operator
	KeyAndOperator ParserContext = iota
	// Values represents values
	Values
)

// lookahead func returns the current token and string. No increment of current position
func (p *Parser) lookahead(context ParserContext) (Token, string) {
	_ = "STUB: not implemented"
	return *new(Token), ""
}

// consume returns current token and string. Increments the position
func (p *Parser) consume(context ParserContext) (Token, string) {
	_ = "STUB: not implemented"
	return *new(Token), ""
}

// scan runs through the input string and stores the ScannedItem in an array
// Parser can now lookahead and consume the tokens
func (p *Parser) scan() { _ = "STUB: not implemented"; return }

// parse runs the left recursive descending algorithm
// on input string. It returns a list of Requirement objects.
func (p *Parser) parse() (internalSelector, error) {
	_ = "STUB: not implemented"
	// init scannedItems
	return *new(internalSelector), nil
}

func (p *Parser) parseRequirement() (*Requirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// operator found lookahead set checked

// parseKeyAndInferOperator parses literals.
// in case of no operator '!, in, notin, ==, =, !=' are found
// the 'exists' operator is inferred
func (p *Parser) parseKeyAndInferOperator() (string, selection.Operator, error) {
	_ = "STUB: not implemented"
	return "", *new(selection.Operator), nil
}

// parseOperator returns operator and eventually matchType
// matchType can be exact
func (p *Parser) parseOperator() (op selection.Operator, err error) {
	_ = "STUB: not implemented"
	return *new(selection.Operator), nil
}

// DoesNotExistToken shouldn't be here because it's a unary operator, not a binary operator

// parseValues parses the values for set based matching (x,y,z)
func (p *Parser) parseValues() (sets.String, error) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// handles general cases

// handles "()"

// parseIdentifiersList parses a (possibly empty) list of
// of comma separated (possibly empty) identifiers
func (p *Parser) parseIdentifiersList() (sets.String, error) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// handled here since we can have "(,"

// to handle (,

// to handle ,)  Double "" removed by StringSet

// to handle ,, Double "" removed by StringSet

// it can be operator

// parseExactValue parses the only value for exact match style
func (p *Parser) parseExactValue() (sets.String, error) {
	_ = "STUB: not implemented"
	return *new(sets.String), nil
}

// Parse takes a string representing a selector and returns a selector
// object, or an error. This parsing function differs from ParseSelector
// as they parse different selectors with different syntaxes.
// The input will cause an error if it does not follow this form:
//
//	<selector-syntax>         ::= <requirement> | <requirement> "," <selector-syntax>
//	<requirement>             ::= [!] KEY [ <set-based-restriction> | <exact-match-restriction> ]
//	<set-based-restriction>   ::= "" | <inclusion-exclusion> <value-set>
//	<inclusion-exclusion>     ::= <inclusion> | <exclusion>
//	<exclusion>               ::= "notin"
//	<inclusion>               ::= "in"
//	<value-set>               ::= "(" <values> ")"
//	<values>                  ::= VALUE | VALUE "," <values>
//	<exact-match-restriction> ::= ["="|"=="|"!="] VALUE
//
// KEY is a sequence of one or more characters following [ DNS_SUBDOMAIN "/" ] DNS_LABEL. Max length is 63 characters.
// VALUE is a sequence of zero or more characters "([A-Za-z0-9_-\.])". Max length is 63 characters.
// Delimiter is white space: (' ', '\t')
// Example of valid syntax:
//
//	"x in (foo,,baz),y,z notin ()"
//
// Note:
//
//	(1) Inclusion - " in " - denotes that the KEY exists and is equal to any of the
//	    VALUEs in its requirement
//	(2) Exclusion - " notin " - denotes that the KEY is not equal to any
//	    of the VALUEs in its requirement or does not exist
//	(3) The empty string is a valid VALUE
//	(4) A requirement with just a KEY - as in "y" above - denotes that
//	    the KEY exists and can be any VALUE.
//	(5) A requirement with just !KEY requires that the KEY not exist.
func Parse(selector string, opts ...field.PathOption) (Selector, error) {
	_ = "STUB: not implemented"
	return *new(Selector), nil
}

// parse parses the string representation of the selector and returns the internalSelector struct.
// The callers of this method can then decide how to return the internalSelector struct to their
// callers. This function has two callers now, one returns a Selector interface and the other
// returns a list of requirements.
func parse(selector string, path *field.Path) (internalSelector, error) {
	_ = "STUB: not implemented"
	return *new(internalSelector), nil
}

// sort to grant determistic parsing

func validateLabelKey(k string, path *field.Path) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

func validateLabelValue(k, v string, path *field.Path) *field.Error {
	_ = "STUB: not implemented"
	return nil
}

// SelectorFromSet returns a Selector which will match exactly the given Set. A
// nil and empty Sets are considered equivalent to Everything().
// It does not perform any validation, which means the server will reject
// the request if the Set contains invalid values.
func SelectorFromSet(ls Set) Selector { _ = "STUB: not implemented"; return *new(Selector) }

// ValidatedSelectorFromSet returns a Selector which will match exactly the given Set. A
// nil and empty Sets are considered equivalent to Everything().
// The Set is validated client-side, which allows to catch errors early.
func ValidatedSelectorFromSet(ls Set) (Selector, error) {
	_ = "STUB: not implemented"
	return *new(Selector), nil
}

// sort to have deterministic string representation

// SelectorFromValidatedSet returns a Selector which will match exactly the given Set.
// A nil and empty Sets are considered equivalent to Everything().
// It assumes that Set is already validated and doesn't do any validation.
func SelectorFromValidatedSet(ls Set) Selector { _ = "STUB: not implemented"; return *new(Selector) }

// sort to have deterministic string representation

// ParseToRequirements takes a string representing a selector and returns a list of
// requirements. This function is suitable for those callers that perform additional
// processing on selector requirements.
// See the documentation for Parse() function for more details.
// TODO: Consider exporting the internalSelector type instead.
func ParseToRequirements(selector string, opts ...field.PathOption) ([]Requirement, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
