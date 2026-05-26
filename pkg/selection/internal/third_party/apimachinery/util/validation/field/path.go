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
Copyright 2015 The Kubernetes Authors.

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

package field

type pathOptions struct {
	path *Path
}

// PathOption modifies a pathOptions
type PathOption func(o *pathOptions)

// WithPath generates a PathOption
func WithPath(p *Path) PathOption { _ = "STUB: not implemented"; return *new(PathOption) }

// ToPath produces *Path from a set of PathOption
func ToPath(opts ...PathOption) *Path { _ = "STUB: not implemented"; return nil }

// Path represents the path from some root to a particular field.
type Path struct {
	name   string // the name of this field or "" if this is an index
	index  string // if name == "", this is a subscript (index or map key) of the previous element
	parent *Path  // nil if this is the root element
}

// NewPath creates a root Path object.
func NewPath(name string, moreNames ...string) *Path { _ = "STUB: not implemented"; return nil }

// Root returns the root element of this Path.
func (p *Path) Root() *Path { _ = "STUB: not implemented"; return nil }

// Do nothing.

// Child creates a new Path that is a child of the method receiver.
func (p *Path) Child(name string, moreNames ...string) *Path { _ = "STUB: not implemented"; return nil }

// Index indicates that the previous Path is to be subscripted by an int.
// This sets the same underlying value as Key.
func (p *Path) Index(index int) *Path { _ = "STUB: not implemented"; return nil }

// Key indicates that the previous Path is to be subscripted by a string.
// This sets the same underlying value as Index.
func (p *Path) Key(key string) *Path { _ = "STUB: not implemented"; return nil }

// String produces a string representation of the Path.
func (p *Path) String() string { _ = "STUB: not implemented"; return "" }

// make a slice to iterate

// iterate, but it has to be backwards

// This is either the root or it is a subscript.
