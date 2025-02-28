// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

// ----------------------------------------------------------------------
// Command Constructors
// ----------------------------------------------------------------------

// NewManualCmd - Create and initialize a new manual command object and return it as
// a pointer.
func (s *Workflow) NewManualCmd() (*commands.Manual, error) {
	var c commands.Manual
	c.ObjectType = "maual"
	err := s.SetNewID(s.ObjectType)
	s.AddCommand(&s)
	return &s, err
}
