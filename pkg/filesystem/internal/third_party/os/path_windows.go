// Windows path handling utilities based on (but modified from)
// https://github.com/golang/go/blob/da0d1a44bac379f5acedb1933f85400de08f4ac6/src/os/path_windows.go
//
// The original code license:
//
// Copyright (c) 2009 The Go Authors. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Google Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
//
// The original license header inside the code itself:
//
// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// isPathSeparator reports whether c is a directory separator character.
func isPathSeparator(c uint8) bool {
	_ = "STUB: not implemented"
	// NOTE: Windows accept / as path separator.
	return false
}

func isAbs(path string) (b bool) { _ = "STUB: not implemented"; return false }

func volumeName(path string) (v string) { _ = "STUB: not implemented"; return "" }

// with drive letter

// is it UNC

// first, leading `\\` and next shouldn't be `\`. its server name.

// second, next '\' shouldn't be repeated.

// third, following something characters. its share name.

// FixLongPath returns the extended-length (\\?\-prefixed) form of
// path when needed, in order to avoid the default 260 character file
// path limit imposed by Windows. If path is not easily converted to
// the extended-length form (for example, if path is a relative path
// or contains .. elements), or is short enough, fixLongPath returns
// path unmodified.
//
// See https://msdn.microsoft.com/en-us/library/windows/desktop/aa365247(v=vs.85).aspx#maxpath
func FixLongPath(path string) string {
	_ = "STUB: not implemented"
	// Do nothing (and don't allocate) if the path is "short".
	// Empirically (at least on the Windows Server 2013 builder),
	// the kernel is arbitrarily okay with < 248 bytes. That
	// matches what the docs above say:
	// "When using an API to create a directory, the specified
	// path cannot be so long that you cannot append an 8.3 file
	// name (that is, the directory name cannot exceed MAX_PATH
	// minus 12)." Since MAX_PATH is 260, 260 - 12 = 248.
	//
	// The MSDN docs appear to say that a normal path that is 248 bytes long
	// will work; empirically the path must be less then 248 bytes long.
	return ""
}

// Don't fix. (This is how Go 1.7 and earlier worked,
// not automatically generating the \\?\ form)

// The extended form begins with \\?\, as in
// \\?\c:\windows\foo.txt or \\?\UNC\server\share\foo.txt.
// The extended form disables evaluation of . and .. path
// elements and disables the interpretation of / as equivalent
// to \. The conversion here rewrites / to \ and elides
// . elements as well as trailing or duplicate separators. For
// simplicity it avoids the conversion entirely for relative
// paths or paths containing .. elements. For now,
// \\server\share paths are not converted to
// \\?\UNC\server\share paths because the rules for doing so
// are less well-specified.

// Don't canonicalize UNC paths.

// Relative path

// empty block

// /./

// /../ is currently unhandled

// A drive's root directory needs a trailing \
