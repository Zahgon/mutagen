// Pattern matching infrastructure used to support .dockerignore files. Based on
// (but modified from)
// https://github.com/moby/patternmatcher/blob/c5e4b22c8cb290f9439a339c08bba6cb13aa296d/patternmatcher.go
//
// The original code license:
//
//                               Apache License
//                         Version 2.0, January 2004
//                      https://www.apache.org/licenses/
//
// TERMS AND CONDITIONS FOR USE, REPRODUCTION, AND DISTRIBUTION
//
// 1. Definitions.
//
//    "License" shall mean the terms and conditions for use, reproduction,
//    and distribution as defined by Sections 1 through 9 of this document.
//
//    "Licensor" shall mean the copyright owner or entity authorized by
//    the copyright owner that is granting the License.
//
//    "Legal Entity" shall mean the union of the acting entity and all
//    other entities that control, are controlled by, or are under common
//    control with that entity. For the purposes of this definition,
//    "control" means (i) the power, direct or indirect, to cause the
//    direction or management of such entity, whether by contract or
//    otherwise, or (ii) ownership of fifty percent (50%) or more of the
//    outstanding shares, or (iii) beneficial ownership of such entity.
//
//    "You" (or "Your") shall mean an individual or Legal Entity
//    exercising permissions granted by this License.
//
//    "Source" form shall mean the preferred form for making modifications,
//    including but not limited to software source code, documentation
//    source, and configuration files.
//
//    "Object" form shall mean any form resulting from mechanical
//    transformation or translation of a Source form, including but
//    not limited to compiled object code, generated documentation,
//    and conversions to other media types.
//
//    "Work" shall mean the work of authorship, whether in Source or
//    Object form, made available under the License, as indicated by a
//    copyright notice that is included in or attached to the work
//    (an example is provided in the Appendix below).
//
//    "Derivative Works" shall mean any work, whether in Source or Object
//    form, that is based on (or derived from) the Work and for which the
//    editorial revisions, annotations, elaborations, or other modifications
//    represent, as a whole, an original work of authorship. For the purposes
//    of this License, Derivative Works shall not include works that remain
//    separable from, or merely link (or bind by name) to the interfaces of,
//    the Work and Derivative Works thereof.
//
//    "Contribution" shall mean any work of authorship, including
//    the original version of the Work and any modifications or additions
//    to that Work or Derivative Works thereof, that is intentionally
//    submitted to Licensor for inclusion in the Work by the copyright owner
//    or by an individual or Legal Entity authorized to submit on behalf of
//    the copyright owner. For the purposes of this definition, "submitted"
//    means any form of electronic, verbal, or written communication sent
//    to the Licensor or its representatives, including but not limited to
//    communication on electronic mailing lists, source code control systems,
//    and issue tracking systems that are managed by, or on behalf of, the
//    Licensor for the purpose of discussing and improving the Work, but
//    excluding communication that is conspicuously marked or otherwise
//    designated in writing by the copyright owner as "Not a Contribution."
//
//    "Contributor" shall mean Licensor and any individual or Legal Entity
//    on behalf of whom a Contribution has been received by Licensor and
//    subsequently incorporated within the Work.
//
// 2. Grant of Copyright License. Subject to the terms and conditions of
//    this License, each Contributor hereby grants to You a perpetual,
//    worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//    copyright license to reproduce, prepare Derivative Works of,
//    publicly display, publicly perform, sublicense, and distribute the
//    Work and such Derivative Works in Source or Object form.
//
// 3. Grant of Patent License. Subject to the terms and conditions of
//    this License, each Contributor hereby grants to You a perpetual,
//    worldwide, non-exclusive, no-charge, royalty-free, irrevocable
//    (except as stated in this section) patent license to make, have made,
//    use, offer to sell, sell, import, and otherwise transfer the Work,
//    where such license applies only to those patent claims licensable
//    by such Contributor that are necessarily infringed by their
//    Contribution(s) alone or by combination of their Contribution(s)
//    with the Work to which such Contribution(s) was submitted. If You
//    institute patent litigation against any entity (including a
//    cross-claim or counterclaim in a lawsuit) alleging that the Work
//    or a Contribution incorporated within the Work constitutes direct
//    or contributory patent infringement, then any patent licenses
//    granted to You under this License for that Work shall terminate
//    as of the date such litigation is filed.
//
// 4. Redistribution. You may reproduce and distribute copies of the
//    Work or Derivative Works thereof in any medium, with or without
//    modifications, and in Source or Object form, provided that You
//    meet the following conditions:
//
//    (a) You must give any other recipients of the Work or
//        Derivative Works a copy of this License; and
//
//    (b) You must cause any modified files to carry prominent notices
//        stating that You changed the files; and
//
//    (c) You must retain, in the Source form of any Derivative Works
//        that You distribute, all copyright, patent, trademark, and
//        attribution notices from the Source form of the Work,
//        excluding those notices that do not pertain to any part of
//        the Derivative Works; and
//
//    (d) If the Work includes a "NOTICE" text file as part of its
//        distribution, then any Derivative Works that You distribute must
//        include a readable copy of the attribution notices contained
//        within such NOTICE file, excluding those notices that do not
//        pertain to any part of the Derivative Works, in at least one
//        of the following places: within a NOTICE text file distributed
//        as part of the Derivative Works; within the Source form or
//        documentation, if provided along with the Derivative Works; or,
//        within a display generated by the Derivative Works, if and
//        wherever such third-party notices normally appear. The contents
//        of the NOTICE file are for informational purposes only and
//        do not modify the License. You may add Your own attribution
//        notices within Derivative Works that You distribute, alongside
//        or as an addendum to the NOTICE text from the Work, provided
//        that such additional attribution notices cannot be construed
//        as modifying the License.
//
//    You may add Your own copyright statement to Your modifications and
//    may provide additional or different license terms and conditions
//    for use, reproduction, or distribution of Your modifications, or
//    for any such Derivative Works as a whole, provided Your use,
//    reproduction, and distribution of the Work otherwise complies with
//    the conditions stated in this License.
//
// 5. Submission of Contributions. Unless You explicitly state otherwise,
//    any Contribution intentionally submitted for inclusion in the Work
//    by You to the Licensor shall be under the terms and conditions of
//    this License, without any additional terms or conditions.
//    Notwithstanding the above, nothing herein shall supersede or modify
//    the terms of any separate license agreement you may have executed
//    with Licensor regarding such Contributions.
//
// 6. Trademarks. This License does not grant permission to use the trade
//    names, trademarks, service marks, or product names of the Licensor,
//    except as required for reasonable and customary use in describing the
//    origin of the Work and reproducing the content of the NOTICE file.
//
// 7. Disclaimer of Warranty. Unless required by applicable law or
//    agreed to in writing, Licensor provides the Work (and each
//    Contributor provides its Contributions) on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
//    implied, including, without limitation, any warranties or conditions
//    of TITLE, NON-INFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A
//    PARTICULAR PURPOSE. You are solely responsible for determining the
//    appropriateness of using or redistributing the Work and assume any
//    risks associated with Your exercise of permissions under this License.
//
// 8. Limitation of Liability. In no event and under no legal theory,
//    whether in tort (including negligence), contract, or otherwise,
//    unless required by applicable law (such as deliberate and grossly
//    negligent acts) or agreed to in writing, shall any Contributor be
//    liable to You for damages, including any direct, indirect, special,
//    incidental, or consequential damages of any character arising as a
//    result of this License or out of the use or inability to use the
//    Work (including but not limited to damages for loss of goodwill,
//    work stoppage, computer failure or malfunction, or any and all
//    other commercial damages or losses), even if such Contributor
//    has been advised of the possibility of such damages.
//
// 9. Accepting Warranty or Additional Liability. While redistributing
//    the Work or Derivative Works thereof, You may choose to offer,
//    and charge a fee for, acceptance of support, warranty, indemnity,
//    or other liability obligations and/or rights consistent with this
//    License. However, in accepting such obligations, You may act only
//    on Your own behalf and on Your sole responsibility, not on behalf
//    of any other Contributor, and only if You agree to indemnify,
//    defend, and hold each Contributor harmless for any liability
//    incurred by, or claims asserted against, such Contributor by reason
//    of your accepting any such warranty or additional liability.
//
// END OF TERMS AND CONDITIONS
//
// Copyright 2013-2018 Docker, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package patternmatcher

import (
	"regexp"
)

// escapeBytes is a bitmap used to check whether a character should be escaped when creating the regex.
var escapeBytes [8]byte

// shouldEscape reports whether a rune should be escaped as part of the regex.
//
// This only includes characters that require escaping in regex but are also NOT valid filepath pattern characters.
// Additionally, '\' is not excluded because there is specific logic to properly handle this, as it's a path separator
// on Windows.
//
// Adapted from regexp::QuoteMeta in go stdlib.
// See https://cs.opensource.google/go/go/+/refs/tags/go1.17.2:src/regexp/regexp.go;l=703-715;drc=refs%2Ftags%2Fgo1.17.2
func shouldEscape(b rune) bool { _ = "STUB: not implemented"; return false }

func init() {
	for _, b := range []byte(`.+()|{}$`) {
		escapeBytes[b%8] |= 1 << (b / 8)
	}
}

// PatternMatcher allows checking paths against a list of patterns
type PatternMatcher struct {
	patterns   []*Pattern
	exclusions bool
	// exclusionCount is the number of exclusion patterns.
	exclusionCount uint
}

// New creates a new matcher object for specific patterns that can
// be used later to match against patterns against paths
func New(patterns []string) (*PatternMatcher, error) { _ = "STUB: not implemented"; return nil, nil }

// Eliminate leading and trailing whitespace.

// Do some syntax checking on the pattern.
// filepath's Match() has some really weird rules that are inconsistent
// so instead of trying to dup their logic, just call Match() for its
// error state and if there is an error in the pattern return it.
// If this becomes an issue we can remove this since its really only
// needed in the error (syntax) case - which isn't really critical.

// PrecompileForMutagen is a utility function that will pre-compile patterns to
// watch for validation errors.
func (pm *PatternMatcher) PrecompileForMutagen() error {
	_ = "STUB: not implemented"
	// Pre-compile any as-of-yet uncompiled patterns.
	return nil
}

// Success.

// Matches returns true if "file" matches any of the patterns
// and isn't excluded by any of the subsequent patterns.
//
// The "file" argument should be a slash-delimited path.
//
// Matches is not safe to call concurrently.
//
// Deprecated: This implementation is buggy (it only checks a single parent dir
// against the pattern) and will be removed soon. Use either
// MatchesOrParentMatches or MatchesUsingParentResults instead.
func (pm *PatternMatcher) Matches(file string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Skip evaluation if this is an inclusion and the filename
// already matched the pattern, or it's an exclusion and it has
// not matched the pattern yet.

// Check to see if the pattern matches one of our parent dirs.

// MatchesOrParentMatches returns true if "file" matches any of the patterns
// and isn't excluded by any of the subsequent patterns.
//
// The "file" argument should be a slash-delimited path.
//
// Matches is not safe to call concurrently.
func (pm *PatternMatcher) MatchesOrParentMatches(file string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Skip evaluation if this is an inclusion and the filename
// already matched the pattern, or it's an exclusion and it has
// not matched the pattern yet.

// Check to see if the pattern matches one of our parent dirs.

// MatchStatus encodes the different potential match states of a file path.
type MatchStatus uint8

const (
	// MatchStatusNominal indicates that the path has been neither matched nor
	// exclude-matched.
	MatchStatusNominal MatchStatus = iota
	// MatchStatusMatched indicates that the path has been matched.
	MatchStatusMatched
	// MatchStatusInverted indicates that the path has been exclude-matched.
	MatchStatusInverted
)

// MatchesForMutagen is a variant of MatchesOrParentMatches that doesn't perform
// any parent comparison and which returns a trinary match state and traversal
// continuation information. Note that this method may panic if the constituent
// patterns haven't been validated prior to PatternMatcher construction.
func (pm *PatternMatcher) MatchesForMutagen(path string, directory bool) (MatchStatus, bool) {
	_ = "STUB: not implemented"
	// Start with a nominal match status.
	return *new(MatchStatus), false
}

// Convert to native path separators. This is a little expensive on Windows
// since all inbound paths from Mutagen will be forward-slash-separated, but
// adjusting this would require significant changes to this vendored code.

// Run through the ignore patterns, updating the match state as we reach
// more specific rules.

// See if we can skip the (relatively expensive) matching process. If
// we're already in a matched state and there aren't any exclusion
// patterns remaining, then we can't leave that state, and thus we can
// skip any further matching. If this is an exclusion pattern, then
// we'll decrement the remaining exclusion pattern count, and we can
// also skip matching for this particular pattern if we're already in an
// inverted state. Finally, if we're already in a matched state and this
// is a non-exclusion pattern, then we also won't change state as a
// result of this particular pattern and can skip matching.

// Perform a matching operation and adjust the status as appropriate. We
// panic on any error (which can only result from compilation) because
// the pattern should have already been externally validated.

// If we're dealing with a directory that is explicitly inverted, then
// traversal continuation should be false, because traversal continuation is
// implicit.

// If we're not dealing with a directory or we don't have any exclusion
// patterns, then we won't need to continue traversal and we're done.

// Determine whether or not filesystem traversal should continue based on
// whether or not any exclusion patterns have this path as a prefix. Note
// that we compute this in the case of both matched and nominal, because in
// the case of nominal we want to know whether or not to continue if an
// ignore mask is set by an ignored parent directory.
//
// Note that this behavior won't recommend continued traversal based on
// exclusions with wildcard prefixes or internal elements (because we're
// just using a simple prefix match), but this is aligned with what Moby
// does:
// https://github.com/moby/moby/blob/462d6ef826861fad021fb565c0481fb61d2db6bc/pkg/archive/archive.go#L1014-L1044
//
// And apparently this is a known issue with no plans to fix at the moment:
// https://github.com/moby/moby/issues/30018

// Done.

// MatchesUsingParentResult returns true if "file" matches any of the patterns
// and isn't excluded by any of the subsequent patterns. The functionality is
// the same as Matches, but as an optimization, the caller keeps track of
// whether the parent directory matched.
//
// The "file" argument should be a slash-delimited path.
//
// MatchesUsingParentResult is not safe to call concurrently.
//
// Deprecated: this function does behave correctly in some cases (see
// https://github.com/docker/buildx/issues/850).
//
// Use MatchesUsingParentResults instead.
func (pm *PatternMatcher) MatchesUsingParentResult(file string, parentMatched bool) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Skip evaluation if this is an inclusion and the filename
// already matched the pattern, or it's an exclusion and it has
// not matched the pattern yet.

// MatchInfo tracks information about parent dir matches while traversing a
// filesystem.
type MatchInfo struct {
	parentMatched []bool
}

// MatchesUsingParentResults returns true if "file" matches any of the patterns
// and isn't excluded by any of the subsequent patterns. The functionality is
// the same as Matches, but as an optimization, the caller passes in
// intermediate results from matching the parent directory.
//
// The "file" argument should be a slash-delimited path.
//
// MatchesUsingParentResults is not safe to call concurrently.
func (pm *PatternMatcher) MatchesUsingParentResults(file string, parentMatchInfo MatchInfo) (bool, MatchInfo, error) {
	_ = "STUB: not implemented"
	return false, *new(MatchInfo), nil
}

// If the parent matched this pattern, we don't need to recheck.

// Skip evaluation if this is an inclusion and the filename
// already matched the pattern, or it's an exclusion and it has
// not matched the pattern yet.

// If the zero value of MatchInfo was passed in, we don't have
// any information about the parent dir's match results, and we
// apply the same logic as MatchesOrParentMatches.

// Check to see if the pattern matches one of our parent dirs.

// Exclusions returns true if any of the patterns define exclusions
func (pm *PatternMatcher) Exclusions() bool { _ = "STUB: not implemented"; return false }

// Patterns returns array of active patterns
func (pm *PatternMatcher) Patterns() []*Pattern {
	_ = "STUB: not implemented"

	// Pattern defines a single regexp used to filter file paths.
	return nil
}

type Pattern struct {
	matchType      matchType
	cleanedPattern string
	dirs           []string
	regexp         *regexp.Regexp
	exclusion      bool
}

type matchType int

const (
	unknownMatch matchType = iota
	exactMatch
	prefixMatch
	suffixMatch
	regexpMatch
)

func (p *Pattern) String() string { _ = "STUB: not implemented"; return "" }

// Exclusion returns true if this pattern defines exclusion
func (p *Pattern) Exclusion() bool { _ = "STUB: not implemented"; return false }

func (p *Pattern) match(path string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

// strip trailing **

// strip leading **

// **/foo matches "foo"

func (p *Pattern) compile(sl string) error { _ = "STUB: not implemented"; return nil }

// Go through the pattern and convert it to a regexp.
// We use a scanner so we can support utf-8 chars.

// is some flavor of "**"

// Treat **/ as ** so eat the "/"

// is "**EOF" - to align with .gitignore just accept all

// is "**"
// Note that this allows for any # of /'s (even 0) because
// the .* will eat everything, even /'s

// is "*" so map it to anything but "/"

// "?" is any char except "/"

// Escape some regexp special chars that have no meaning
// in golang's filepath.Match

// escape next char. Note that a trailing \ in the pattern
// will be left alone (but need to escape it)

// On windows map "\" to "\\", meaning an escaped backslash,
// and then just continue because filepath.Match on
// Windows doesn't allow escaping at all

// Matches returns true if file matches any of the patterns
// and isn't excluded by any of the subsequent patterns.
//
// This implementation is buggy (it only checks a single parent dir against the
// pattern) and will be removed soon. Use MatchesOrParentMatches instead.
func Matches(file string, patterns []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Don't let them exclude everything, kind of silly.

// MatchesOrParentMatches returns true if file matches any of the patterns
// and isn't excluded by any of the subsequent patterns.
func MatchesOrParentMatches(file string, patterns []string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Don't let them exclude everything, kind of silly.
