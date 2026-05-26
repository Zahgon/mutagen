//go:build mutagensspl

// Copyright (c) 2023-present Docker, Inc.
//
// This program is free software: you can redistribute it and/or modify it under
// the terms of the Server Side Public License, version 1, as published by
// MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
// FOR A PARTICULAR PURPOSE. See the Server Side Public License for more
// details.
//
// You should have received a copy of the Server Side Public License along with
// this program. If not, see
// <http://www.mongodb.com/licensing/server-side-public-license>.

package xxh128

import (
	"hash"

	"github.com/zeebo/xxh3"
)

// xxh128Hash implements hash.Hash using the XXH128 algorithm.
type xxh128Hash struct {
	// Hasher is the underlying hasher.
	*xxh3.Hasher
}

// New returns a new XXH128 hash.
func New() hash.Hash { _ = "STUB: not implemented"; return *new(hash.Hash) }

// Sum implements hash.Hash.Sum.
func (h *xxh128Hash) Sum(b []byte) []byte {
	_ = "STUB: not implemented"
	// Compute the sum and associated bytes.
	return nil
}

// If b is nil, then take the fast way out.

// Otherwise append the bytes to b.
