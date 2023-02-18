// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package action implements the CACAO 2.0 workflow action step object.
//
// The Action Step workflow step contains the actual commands to be executed on
// one or more agents. These commands are intended to be processed
// sequentially. In addition to the inherited properties, this section defines
// five more specific properties that are valid for this type.
package action

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// CommandData - This type implement the CACAO 2.0 command data type.
//
// The CACAO command object (command-data) contains detailed information about
// the commands that are to be executed or processed automatically or manually
// as part of an action step (see section 4.5). Each command listed in an
// action step may be of a different command type, however, all commands listed
// in a single step MUST be processed or executed by all of the agents defined
// in that step.
//
// Commands can use and refer to variables just like other parts of the
// playbook. For each command either the command property or the command_b64
// property MUST be present.
//
// The individual commands MAY be defined in other specifications, and when
// possible will be mapped to the JSON structure of this specification. When
// that is not possible, they will be base64 encoded.
type CommandData struct {
	ObjectType       string `json:"type,omitempty"`
	Description      string `json:"description,omitempty"`
	Command          string `json:"command,omitempty"`
	CommandB64       string `json:"command_b64,omitempty"`
	Version          string `json:"version,omitempty"`
	PlaybookActivity string `json:"playbook_activity,omitempty"`
}

// ActionStep - This type implmenets the CACAO 2.0 workflow action step
// and defines all of the properties associated with the action step. Some
// properties are inherited from the workflow.CommonProperties type.
type ActionStep struct {
	workflow.CommonProperties
	Commands []CommandData `json:"commands,omitempty"`
	Agent    string        `json:"agent,omitempty"`
	Targets  []string      `json:"targets,omitempty"`
	InArgs   []string      `json:"in_args,omitempty"`
	OutArgs  []string      `json:"out_args,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow action step object and return
// it as a pointer. It will also initialize the object by setting all of the
// basic properties.
func New() *ActionStep {
	var w ActionStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow action step object with the
// correct defaults.
func (w *ActionStep) Init() {
	w.ObjectType = "action"
	w.SetNewID(w.ObjectType)
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

// ----------------------------------------------------------------------
// Define Functions and Methods - ActionStep
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *ActionStep) GetCommon() workflow.CommonProperties {
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
