//go:build linux && mutagensspl

// Copyright (c) 2020-present Docker, Inc.
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

package linux

// Capabilities returns a bit mask of the process' effective capabilities.
func Capabilities() (uint32, error) {
	_ = "STUB: not implemented"
	// Grab the current process ID and ensure that we can fit it in the
	// capability API header.
	return 0, nil
}

// Construct the header. We request version 1 of the Linux capability API
// because it uses 32-bit capabilities, which are the only kind that the
// unix.CapUserData structure supports. Using a later API version can result
// in memory corruption when the kernel tries to stuff 64-bit values into
// 32-bit fields. Version 2 and 3 of the capability API also require more
// modern kernel versions (2.6.25 and 2.6.26, respectively) than the Go
// runtime requires (2.6.23), so using API version 1 keeps us in alignment
// with Go when it comes to version support.

// Perform the query.

// Success.
