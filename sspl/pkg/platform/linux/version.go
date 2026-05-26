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

// Version returns the major and minor components of the Linux kernel version.
func Version() (uint64, uint64, error) {
	_ = "STUB: not implemented"
	// Grab system metadata using uname.
	return 0, 0, nil
}

// Extract the kernel version.

// Parse the kernel version.

// Success.
