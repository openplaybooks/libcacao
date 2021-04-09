// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package single implements the CACAO 1.0 workflow single action step object.
//
// This workflow step contains the actual commands to be executed on one or more
// targets. These commands are intended to be processed sequentially one at a
// time. In addition to the inherited properties, this section defines five more
// specific properties that are valid for this type.
package single

import (
	"github.com/jordan2175/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// CommandData - This type implement the CACAO command data type. The CACAO
// command object (command-data) contains detailed information about the
// commands that are to be executed or processed automatically or manually as
// part of a workflow step (see section 4). Each command listed in a step may be
// of a different command type, however, all commands listed in a single step
// MUST be processed or executed by all of the targets defined in that step.
//
// Commands can use and refer to variables just like other parts of the
// playbook. For each command either the command property or the command_b64
// property MUST be present.
//
// The individual commands MAY be defined in other specifications, and when
// possible will be mapped to the JSON structure of this specification. When
// that is not possible, they will be base64 encoded.
type CommandData struct {
	ObjectType string `json:"type,omitempty"`
	Command    string `json:"command,omitempty"`
	CommandB64 string `json:"command_b64,omitempty"`
	Version    string `json:"version,omitempty"`
}

// WorkflowSingleActionStep - This type implmenets the CACAO 1.0 workflow single
// action step and defines all of the properties associated with the single
// action step. Some properties are inherited from the workflow.CommonProperties
// type.
type WorkflowSingleActionStep struct {
	workflow.CommonProperties
	Commands  []CommandData `json:"commands,omitempty"`
	TargetIDs []string      `json:"target_ids,omitempty"`
	InArgs    []string      `json:"in_args,omitempty"`
	OutArgs   []string      `json:"out_args,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow single action step object and
// return it as a pointer. It will also initialize the object by setting all of
// the basic properties.
func New() *WorkflowSingleActionStep {
	var w WorkflowSingleActionStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow single action step object with the
// correct defaults.
func (w *WorkflowSingleActionStep) Init() {
	w.ObjectType = "single"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods - CommandData
// ----------------------------------------------------------------------

// SetManual - This method will set the object type of the command to manual
func (c *CommandData) SetManual() {
	c.ObjectType = "manual"
}

// SetHTTPAPI - This method will set the object type of the command to http-api
func (c *CommandData) SetHTTPAPI() {
	c.ObjectType = "http-api"
}

// SetSSH - This method will set the object type of the command to ssh
func (c *CommandData) SetSSH() {
	c.ObjectType = "ssh"
}

// SetBash - This method will set the object type of the command to bash
func (c *CommandData) SetBash() {
	c.ObjectType = "bash"
}

// SetOpenC2JSON - This method will set the object type of the command to openc2-json
func (c *CommandData) SetOpenC2JSON() {
	c.ObjectType = "openc2-json"
}

// SetAttackCMD - This method will set the object type of the command to attack-cmd
func (c *CommandData) SetAttackCMD() {
	c.ObjectType = "attack-cmd"
}

// ----------------------------------------------------------------------
// Define Functions and Methods - WorkflowSingleActionStep
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *WorkflowSingleActionStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}

// NewCommand - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (w *WorkflowSingleActionStep) NewCommand(r ...CommandData) (*CommandData, error) {
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
