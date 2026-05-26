// Windows filesystem monitoring implementation based on
// golang.org/x/exp/winfsnotify
// (specifically
// https://github.com/golang/exp/tree/c84be7c6d1cd7b6a43fd7101daaf2dc35ded445f/winfsnotify),
// but modified to remove import path enforcement, increase
// ReadDirectoryChangesW buffer size, support recursive watching, use more
// idiomatic filesystem path joins, and remove test logging.
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

//go:build windows

// Package winfsnotify allows the user to receive
// file system event notifications on Windows.
package winfsnotify

import (
	"syscall"
)

// Event is the type of the notification messages
// received on the watcher's Event channel.
type Event struct {
	Mask   uint32 // Mask of events
	Cookie uint32 // Unique cookie associating related events (for rename)
	Name   string // File name (optional)
}

const (
	opAddWatch = iota
	opRemoveWatch
)

const (
	provisional uint64 = 1 << (32 + iota)
)

type input struct {
	op    int
	path  string
	flags uint32
	reply chan error
}

type inode struct {
	handle syscall.Handle
	volume uint32
	index  uint64
}

type watch struct {
	ov     syscall.Overlapped
	ino    *inode            // i-number
	path   string            // Directory path
	mask   uint64            // Directory itself is being watched with these notify flags
	names  map[string]uint64 // Map of names being watched and their notify flags
	rename string            // Remembers the old name while renaming a file
	buf    [65536]byte       // Maximum before we hit network packet limits
}

type indexMap map[uint64]*watch
type watchMap map[uint32]indexMap

// A Watcher waits for and receives event notifications
// for a specific set of files and directories.
type Watcher struct {
	port     syscall.Handle // Handle to completion port
	watches  watchMap       // Map of watches (key: i-number)
	input    chan *input    // Inputs to the reader are sent on this channel
	Event    chan *Event    // Events are returned on this channel
	Error    chan error     // Errors are sent on this channel
	isClosed bool           // Set to true when Close() is first called
	quit     chan chan<- error
	cookie   uint32
}

// NewWatcher creates and returns a Watcher.
func NewWatcher() (*Watcher, error) { _ = "STUB: not implemented"; return nil, nil }

// Close closes a Watcher.
// It sends a message to the reader goroutine to quit and removes all watches
// associated with the watcher.
func (w *Watcher) Close() error { _ = "STUB: not implemented"; return nil }

// Send "quit" message to the reader goroutine

// AddWatch adds path to the watched file set.
func (w *Watcher) AddWatch(path string, flags uint32) error { _ = "STUB: not implemented"; return nil }

// Watch adds path to the watched file set, watching all events.
func (w *Watcher) Watch(path string) error { _ = "STUB: not implemented"; return nil }

// RemoveWatch removes path from the watched file set.
func (w *Watcher) RemoveWatch(path string) error { _ = "STUB: not implemented"; return nil }

func (w *Watcher) wakeupReader() error { _ = "STUB: not implemented"; return nil }

func getDir(pathname string) (dir string, err error) { _ = "STUB: not implemented"; return "", nil }

func getIno(path string) (ino *inode, err error) { _ = "STUB: not implemented"; return nil, nil }

// Must run within the I/O thread.
func (m watchMap) get(ino *inode) *watch { _ = "STUB: not implemented"; return nil }

// Must run within the I/O thread.
func (m watchMap) set(ino *inode, watch *watch) { _ = "STUB: not implemented"; return }

// Must run within the I/O thread.
func (w *Watcher) addWatch(pathname string, flags uint64) error {
	_ = "STUB: not implemented"
	return nil
}

// Must run within the I/O thread.
func (w *Watcher) removeWatch(pathname string) error { _ = "STUB: not implemented"; return nil }

// We need the volume and index but not the handle itself.

// Must run within the I/O thread.
func (w *Watcher) deleteWatch(watch *watch) { _ = "STUB: not implemented"; return }

// Must run within the I/O thread.
func (w *Watcher) startRead(watch *watch) error { _ = "STUB: not implemented"; return nil }

// Watched directory was probably removed

// readEvents reads from the I/O completion port, converts the
// received events into Event objects and sends them via the Event channel.
// Entry point to the I/O thread.
func (w *Watcher) readEvents() { _ = "STUB: not implemented"; return }

// Watched directory was probably removed

// CancelIo was called on this handle

// Point "raw" to the event in the buffer

// Move to the next event in the buffer

func (w *Watcher) sendEvent(name string, mask uint64) bool { _ = "STUB: not implemented"; return false }

// String formats the event e in the form
// "filename: 0xEventMask = FS_ACCESS|FS_ATTRIB_|..."
func (e *Event) String() string { _ = "STUB: not implemented"; return "" }

func toWindowsFlags(mask uint64) uint32 { _ = "STUB: not implemented"; return 0 }

func toFSnotifyFlags(action uint32) uint64 { _ = "STUB: not implemented"; return 0 }

const (
	// Options for AddWatch
	FS_ONESHOT = 0x80000000
	FS_ONLYDIR = 0x1000000

	// Events
	FS_ACCESS      = 0x1
	FS_ALL_EVENTS  = 0xfff
	FS_ATTRIB      = 0x4
	FS_CLOSE       = 0x18
	FS_CREATE      = 0x100
	FS_DELETE      = 0x200
	FS_DELETE_SELF = 0x400
	FS_MODIFY      = 0x2
	FS_MOVE        = 0xc0
	FS_MOVED_FROM  = 0x40
	FS_MOVED_TO    = 0x80
	FS_MOVE_SELF   = 0x800

	// Special events
	FS_IGNORED    = 0x8000
	FS_Q_OVERFLOW = 0x4000
)

var eventBits = []struct {
	Value uint32
	Name  string
}{
	{FS_ACCESS, "FS_ACCESS"},
	{FS_ATTRIB, "FS_ATTRIB"},
	{FS_CREATE, "FS_CREATE"},
	{FS_DELETE, "FS_DELETE"},
	{FS_DELETE_SELF, "FS_DELETE_SELF"},
	{FS_MODIFY, "FS_MODIFY"},
	{FS_MOVED_FROM, "FS_MOVED_FROM"},
	{FS_MOVED_TO, "FS_MOVED_TO"},
	{FS_MOVE_SELF, "FS_MOVE_SELF"},
	{FS_IGNORED, "FS_IGNORED"},
	{FS_Q_OVERFLOW, "FS_Q_OVERFLOW"},
}
