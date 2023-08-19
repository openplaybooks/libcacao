// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

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

// ----------------------------------------------------------------------
// Define Action Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *ActionStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// NewCommand - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (w *ActionStep) NewCommand(r ...CommandData) (*CommandData, error) {
	positionThatAppendWillUse := len(w.Commands)

	if len(r) > 0 {
		for i := range r {
			// Update the value so we grab the last one entered
			positionThatAppendWillUse = len(w.Commands)
			w.Commands = append(w.Commands, r[i])
		}
		return &w.Commands[positionThatAppendWillUse], nil
	}

	// If one was not passed in, lets create one
	var c CommandData
	w.Commands = append(w.Commands, c)
	return &w.Commands[positionThatAppendWillUse], nil
}
