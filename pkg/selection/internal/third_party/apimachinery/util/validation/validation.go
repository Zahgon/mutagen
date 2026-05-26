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

package validation

import (
	"math"
	"regexp"

	"github.com/mutagen-io/mutagen/pkg/selection/internal/third_party/apimachinery/util/validation/field"
)

const qnameCharFmt string = "[A-Za-z0-9]"
const qnameExtCharFmt string = "[-A-Za-z0-9_.]"
const qualifiedNameFmt string = "(" + qnameCharFmt + qnameExtCharFmt + "*)?" + qnameCharFmt
const qualifiedNameErrMsg string = "must consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character"
const qualifiedNameMaxLength int = 63

var qualifiedNameRegexp = regexp.MustCompile("^" + qualifiedNameFmt + "$")

// IsQualifiedName tests whether the value passed is what Kubernetes calls a
// "qualified name".  This is a format used in various places throughout the
// system.  If the value is not valid, a list of error strings is returned.
// Otherwise an empty list (or nil) is returned.
func IsQualifiedName(value string) []string { _ = "STUB: not implemented"; return nil }

// IsFullyQualifiedName checks if the name is fully qualified. This is similar
// to IsFullyQualifiedDomainName but requires a minimum of 3 segments instead of
// 2 and does not accept a trailing . as valid.
// TODO: This function is deprecated and preserved until all callers migrate to
// IsFullyQualifiedDomainName; please don't add new callers.
func IsFullyQualifiedName(fldPath *field.Path, name string) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// IsFullyQualifiedDomainName checks if the domain name is fully qualified. This
// is similar to IsFullyQualifiedName but only requires a minimum of 2 segments
// instead of 3 and accepts a trailing . as valid.
func IsFullyQualifiedDomainName(fldPath *field.Path, name string) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// Allowed characters in an HTTP Path as defined by RFC 3986. A HTTP path may
// contain:
// * unreserved characters (alphanumeric, '-', '.', '_', '~')
// * percent-encoded octets
// * sub-delims ("!", "$", "&", "'", "(", ")", "*", "+", ",", ";", "=")
// * a colon character (":")
const httpPathFmt string = `[A-Za-z0-9/\-._~%!$&'()*+,;=:]+`

var httpPathRegexp = regexp.MustCompile("^" + httpPathFmt + "$")

// IsDomainPrefixedPath checks if the given string is a domain-prefixed path
// (e.g. acme.io/foo). All characters before the first "/" must be a valid
// subdomain as defined by RFC 1123. All characters trailing the first "/" must
// be valid HTTP Path characters as defined by RFC 3986.
func IsDomainPrefixedPath(fldPath *field.Path, dpPath string) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

const labelValueFmt string = "(" + qualifiedNameFmt + ")?"
const labelValueErrMsg string = "a valid label must be an empty string or consist of alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character"

// LabelValueMaxLength is a label's max length
const LabelValueMaxLength int = 63

var labelValueRegexp = regexp.MustCompile("^" + labelValueFmt + "$")

// IsValidLabelValue tests whether the value passed is a valid label value.  If
// the value is not valid, a list of error strings is returned.  Otherwise an
// empty list (or nil) is returned.
func IsValidLabelValue(value string) []string { _ = "STUB: not implemented"; return nil }

const dns1123LabelFmt string = "[a-z0-9]([-a-z0-9]*[a-z0-9])?"
const dns1123LabelErrMsg string = "a lowercase RFC 1123 label must consist of lower case alphanumeric characters or '-', and must start and end with an alphanumeric character"

// DNS1123LabelMaxLength is a label's max length in DNS (RFC 1123)
const DNS1123LabelMaxLength int = 63

var dns1123LabelRegexp = regexp.MustCompile("^" + dns1123LabelFmt + "$")

// IsDNS1123Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1123).
func IsDNS1123Label(value string) []string { _ = "STUB: not implemented"; return nil }

const dns1123SubdomainFmt string = dns1123LabelFmt + "(\\." + dns1123LabelFmt + ")*"
const dns1123SubdomainErrorMsg string = "a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character"

// DNS1123SubdomainMaxLength is a subdomain's max length in DNS (RFC 1123)
const DNS1123SubdomainMaxLength int = 253

var dns1123SubdomainRegexp = regexp.MustCompile("^" + dns1123SubdomainFmt + "$")

// IsDNS1123Subdomain tests for a string that conforms to the definition of a
// subdomain in DNS (RFC 1123).
func IsDNS1123Subdomain(value string) []string { _ = "STUB: not implemented"; return nil }

const dns1035LabelFmt string = "[a-z]([-a-z0-9]*[a-z0-9])?"
const dns1035LabelErrMsg string = "a DNS-1035 label must consist of lower case alphanumeric characters or '-', start with an alphabetic character, and end with an alphanumeric character"

// DNS1035LabelMaxLength is a label's max length in DNS (RFC 1035)
const DNS1035LabelMaxLength int = 63

var dns1035LabelRegexp = regexp.MustCompile("^" + dns1035LabelFmt + "$")

// IsDNS1035Label tests for a string that conforms to the definition of a label in
// DNS (RFC 1035).
func IsDNS1035Label(value string) []string { _ = "STUB: not implemented"; return nil }

// wildcard definition - RFC 1034 section 4.3.3.
// examples:
// - valid: *.bar.com, *.foo.bar.com
// - invalid: *.*.bar.com, *.foo.*.com, *bar.com, f*.bar.com, *
const wildcardDNS1123SubdomainFmt = "\\*\\." + dns1123SubdomainFmt
const wildcardDNS1123SubdomainErrMsg = "a wildcard DNS-1123 subdomain must start with '*.', followed by a valid DNS subdomain, which must consist of lower case alphanumeric characters, '-' or '.' and end with an alphanumeric character"

// IsWildcardDNS1123Subdomain tests for a string that conforms to the definition of a
// wildcard subdomain in DNS (RFC 1034 section 4.3.3).
func IsWildcardDNS1123Subdomain(value string) []string { _ = "STUB: not implemented"; return nil }

const cIdentifierFmt string = "[A-Za-z_][A-Za-z0-9_]*"
const identifierErrMsg string = "a valid C identifier must start with alphabetic character or '_', followed by a string of alphanumeric characters or '_'"

var cIdentifierRegexp = regexp.MustCompile("^" + cIdentifierFmt + "$")

// IsCIdentifier tests for a string that conforms the definition of an identifier
// in C. This checks the format, but not the length.
func IsCIdentifier(value string) []string { _ = "STUB: not implemented"; return nil }

// IsValidPortNum tests that the argument is a valid, non-zero port number.
func IsValidPortNum(port int) []string { _ = "STUB: not implemented"; return nil }

// IsInRange tests that the argument is in an inclusive range.
func IsInRange(value int, min int, max int) []string { _ = "STUB: not implemented"; return nil }

// Now in libcontainer UID/GID limits is 0 ~ 1<<31 - 1
// TODO: once we have a type for UID/GID we should make these that type.
const (
	minUserID  = 0
	maxUserID  = math.MaxInt32
	minGroupID = 0
	maxGroupID = math.MaxInt32
)

// IsValidGroupID tests that the argument is a valid Unix GID.
func IsValidGroupID(gid int64) []string { _ = "STUB: not implemented"; return nil }

// IsValidUserID tests that the argument is a valid Unix UID.
func IsValidUserID(uid int64) []string { _ = "STUB: not implemented"; return nil }

var portNameCharsetRegex = regexp.MustCompile("^[-a-z0-9]+$")
var portNameOneLetterRegexp = regexp.MustCompile("[a-z]")

// IsValidPortName check that the argument is valid syntax. It must be
// non-empty and no more than 15 characters long. It may contain only [-a-z0-9]
// and must contain at least one letter [a-z]. It must not start or end with a
// hyphen, nor contain adjacent hyphens.
//
// Note: We only allow lower-case characters, even though RFC 6335 is case
// insensitive.
func IsValidPortName(port string) []string { _ = "STUB: not implemented"; return nil }

// IsValidIP tests that the argument is a valid IP address.
func IsValidIP(value string) []string { _ = "STUB: not implemented"; return nil }

// IsValidIPv4Address tests that the argument is a valid IPv4 address.
func IsValidIPv4Address(fldPath *field.Path, value string) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

// IsValidIPv6Address tests that the argument is a valid IPv6 address.
func IsValidIPv6Address(fldPath *field.Path, value string) field.ErrorList {
	_ = "STUB: not implemented"
	return *new(field.ErrorList)
}

const percentFmt string = "[0-9]+%"
const percentErrMsg string = "a valid percent string must be a numeric string followed by an ending '%'"

var percentRegexp = regexp.MustCompile("^" + percentFmt + "$")

// IsValidPercent checks that string is in the form of a percentage
func IsValidPercent(percent string) []string { _ = "STUB: not implemented"; return nil }

const httpHeaderNameFmt string = "[-A-Za-z0-9]+"
const httpHeaderNameErrMsg string = "a valid HTTP header must consist of alphanumeric characters or '-'"

var httpHeaderNameRegexp = regexp.MustCompile("^" + httpHeaderNameFmt + "$")

// IsHTTPHeaderName checks that a string conforms to the Go HTTP library's
// definition of a valid header field name (a stricter subset than RFC7230).
func IsHTTPHeaderName(value string) []string { _ = "STUB: not implemented"; return nil }

const envVarNameFmt = "[-._a-zA-Z][-._a-zA-Z0-9]*"
const envVarNameFmtErrMsg string = "a valid environment variable name must consist of alphabetic characters, digits, '_', '-', or '.', and must not start with a digit"

var envVarNameRegexp = regexp.MustCompile("^" + envVarNameFmt + "$")

// IsEnvVarName tests if a string is a valid environment variable name.
func IsEnvVarName(value string) []string { _ = "STUB: not implemented"; return nil }

const configMapKeyFmt = `[-._a-zA-Z0-9]+`
const configMapKeyErrMsg string = "a valid config key must consist of alphanumeric characters, '-', '_' or '.'"

var configMapKeyRegexp = regexp.MustCompile("^" + configMapKeyFmt + "$")

// IsConfigMapKey tests for a string that is a valid key for a ConfigMap or Secret
func IsConfigMapKey(value string) []string { _ = "STUB: not implemented"; return nil }

// MaxLenError returns a string explanation of a "string too long" validation
// failure.
func MaxLenError(length int) string { _ = "STUB: not implemented"; return "" }

// RegexError returns a string explanation of a regex validation failure.
func RegexError(msg string, fmt string, examples ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// EmptyError returns a string explanation of a "must not be empty" validation
// failure.
func EmptyError() string { _ = "STUB: not implemented"; return "" }

func prefixEach(msgs []string, prefix string) []string { _ = "STUB: not implemented"; return nil }

// InclusiveRangeError returns a string explanation of a numeric "must be
// between" validation failure.
func InclusiveRangeError(lo, hi int) string { _ = "STUB: not implemented"; return "" }

func hasChDirPrefix(value string) []string { _ = "STUB: not implemented"; return nil }

// IsValidSocketAddr checks that string represents a valid socket address
// as defined in RFC 789. (e.g 0.0.0.0:10254 or [::]:10254))
func IsValidSocketAddr(value string) []string { _ = "STUB: not implemented"; return nil }
