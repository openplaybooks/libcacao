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

// NewAssignmentCommand - Create and initialize a new assignment command object
// and return it as a pointer.
func (s *ActionStep) NewAssignmentCommand() (*commands.Assignment, error) {
	var c commands.Assignment
	c.ObjectType = "assignment"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewManualCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewManualCommand() (*commands.Manual, error) {
	var c commands.Manual
	c.ObjectType = "manual"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewShellCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewShellCommand(t string) (*commands.Shell, error) {
	var c commands.Shell
	c.ObjectType = "shell"
	c.ShellType = t
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewHTTPCommand - Create and initialize a new http-api command object and
// return it as a pointer.
func (s *ActionStep) NewHTTPCommand() (*commands.HTTP, error) {
	var c commands.HTTP
	c.ObjectType = "http"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewPowerShellCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewPowerShellCommand() (*commands.PowerShell, error) {
	var c commands.PowerShell
	c.ObjectType = "powershell"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewSigmaCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewSigmaCommand() (*commands.Sigma, error) {
	var c commands.Sigma
	c.ObjectType = "sigma"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewYaraCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewYaraCommand() (*commands.Yara, error) {
	var c commands.Yara
	c.ObjectType = "yara"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewKestrelCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewKestrelCommand() (*commands.Kestrel, error) {
	var c commands.Kestrel
	c.ObjectType = "kestrel"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}

// NewElasticCommand - Create and initialize a new manual command object and return it as
// a pointer.
func (s *ActionStep) NewElasticCommand() (*commands.Elastic, error) {
	var c commands.Elastic
	c.ObjectType = "elastic"
	c.SetNewID(c.ObjectType)
	s.Commands = append(s.Commands, &c)
	return &c, nil
}
