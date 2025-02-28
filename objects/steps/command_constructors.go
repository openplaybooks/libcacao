// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package steps

import (
	"github.com/openplaybooks/libcacao/objects/commands"
)

// ----------------------------------------------------------------------
// Command Constructors
// ----------------------------------------------------------------------

// NewManualCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewManualCommand() (*commands.Manual, error) {
	var c commands.Manual
	c.ObjectType = "manual"
	s.AddCommand(&c)
	return &c, nil
}
