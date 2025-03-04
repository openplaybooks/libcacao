// Copyright 2019-2025 Bret Jordan, All rights reserved.
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
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewBashCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewBashCommand() (*commands.Bash, error) {
	var c commands.Bash
	c.ObjectType = "bash"
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewHTTPAPICommand - Create and initialize a new http-api command object and
// return it as a pointer.
func (s *ActionStep) NewHTTPAPICommand() (*commands.HTTPAPI, error) {
	var c commands.HTTPAPI
	c.ObjectType = "http-api"

	if c.ReturnedData == nil {
		m := make(map[string]string, 0)
		c.ReturnedData = m
	}
	s.Commands = append(s.Commands, &c)
	return &c, nil
}
