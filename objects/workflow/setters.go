// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

import (
	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Action Step Functions and Methods
// ----------------------------------------------------------------------

// NewCommand - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (s *ActionStep) NewCommand(r ...CommandData) (*CommandData, error) {
	positionThatAppendWillUse := len(s.Commands)

	if len(r) > 0 {
		for i := range r {
			// Update the value so we grab the last one entered
			positionThatAppendWillUse = len(s.Commands)
			s.Commands = append(s.Commands, r[i])
		}
		return &s.Commands[positionThatAppendWillUse], nil
	}

	// If one was not passed in, lets create one
	var c CommandData
	s.Commands = append(s.Commands, c)
	return &s.Commands[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Define Parallel Step Functions and Methods
// ----------------------------------------------------------------------

// AddNextSteps - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a next
// steps in the workflow and adds them to the next_steps property.
func (s *ParallelStep) AddNextSteps(values interface{}) error {
	return objects.AddValuesToList(&s.NextSteps, values)
}

// ----------------------------------------------------------------------
// Define Switch Step Functions and Methods
// ----------------------------------------------------------------------

// AddCase - This method takes in two values, the first is a string value
// representing the case of the switch statement. The second is a string value
// representing the identifier to be processed by the case condition.
func (s *SwitchStep) AddCase(k string, id string) error {
	if s.Cases == nil {
		m := make(map[string]string, 0)
		s.Cases = m
	}
	s.Cases[k] = id
	return nil
}

// ----------------------------------------------------------------------
// Define Functions and Methods - CommandData
// ----------------------------------------------------------------------

// SetManual - This method will set the object type of the command to manual
func (c *CommandData) SetManual() {
	c.ObjectType = "manual"
}

// SetBash - This method will set the object type of the command to bash
func (c *CommandData) SetBash() {
	c.ObjectType = "bash"
}

// SetHTTPAPI - This method will set the object type of the command to http-api
func (c *CommandData) SetHTTPAPI() {
	c.ObjectType = "http-api"
}

// SetSSH - This method will set the object type of the command to ssh
func (c *CommandData) SetSSH() {
	c.ObjectType = "ssh"
}

// SetCalderaCMD - This method will set the object type of the command to caldera-cmd
func (c *CommandData) SetCalderaCMD() {
	c.ObjectType = "caldera-cmd"
}

// SetElastic - This method will set the object type of the command to elastic
func (c *CommandData) SetElastic() {
	c.ObjectType = "elastic"
}

// SetJupyter - This method will set the object type of the command to jupyter
func (c *CommandData) SetJupyter() {
	c.ObjectType = "jupyter"
}

// SetKestrel - This method will set the object type of the command to kestrel
func (c *CommandData) SetKestrel() {
	c.ObjectType = "kestrel"
}

// SetOpenC2JSON - This method will set the object type of the command to openc2-json
func (c *CommandData) SetOpenC2JSON() {
	c.ObjectType = "openc2-json"
}

// SetSigma - This method will set the object type of the command to sigma
func (c *CommandData) SetSigma() {
	c.ObjectType = "sigma"
}

// SetYara - This method will set the object type of the command to yara
func (c *CommandData) SetYara() {
	c.ObjectType = "yara"
}
