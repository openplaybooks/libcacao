// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

import (
	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Start Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *StartStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// ----------------------------------------------------------------------
// Define End Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *EndStep) GetCommon() CommonProperties {
	return w.CommonProperties
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

// ----------------------------------------------------------------------
// Define Playbook Action Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *PlaybookActionStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// ----------------------------------------------------------------------
// Define Parallel Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *ParallelStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// AddNextSteps - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// next steps in the workflow and adds them to the next_steps property.
func (w *ParallelStep) AddNextSteps(values interface{}) error {
	return objects.AddValuesToList(&w.NextSteps, values)
}

// ----------------------------------------------------------------------
// Define If Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *IfStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// AddOnTrue - This method takes in a string value, a comma separated list of
// string values, or a slice of string values and add them to the on_true
// property. Each entry represents one or more identifiers to be processed if
// the condition returns "true".
func (w *IfStep) AddOnTrue(values interface{}) error {
	return objects.AddValuesToList(&w.OnTrue, values)
}

// AddOnFalse - This method takes in a string value, a comma separated list of
// string values, or a slice of string values and add them to the on_false
// property. Each entry represents one or more identifiers to be processed if
// the condition returns "false".
func (w *IfStep) AddOnFalse(values interface{}) error {
	return objects.AddValuesToList(&w.OnTrue, values)
}

// ----------------------------------------------------------------------
// Define While Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *WhileStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// AddOnTrue - This method takes in a string value, a comma separated list of
// string values, or a slice of string values and add them to the on_true
// property. Each entry represents one or more identifiers to be processed if
// the condition returns "true".
func (w *WhileStep) AddOnTrue(values interface{}) error {
	return objects.AddValuesToList(&w.OnTrue, values)
}

// ----------------------------------------------------------------------
// Define Switch Step Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *SwitchStep) GetCommon() CommonProperties {
	return w.CommonProperties
}

// AddCase - This method takes in two values, the first is a string value
// representing the case of the switch statement.  The second is a string value,
// a comma separated list of string values, or a slice of string values, each
// representing one or more identifiers to be processed by the case condition.
func (w *SwitchStep) AddCase(k string, v interface{}) error {
	if w.Cases == nil {
		m := make(map[string][]string, 0)
		w.Cases = m
	}
	temp := make([]string, 0)
	objects.AddValuesToList(&temp, v)
	w.Cases[k] = temp
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
