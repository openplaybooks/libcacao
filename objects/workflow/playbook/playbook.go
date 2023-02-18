// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package playbookaction implements the CACAO 2.0 workflow playbook action step
// object.
//
// The Playbook Action Step workflow step executes a referenced playbook. In
// addition to the inherited properties, this section defines four more
// specific properties that are valid for this type.
package playbookaction

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// PlaybookActionStep - This type implmenets the CACAO 2.0 workflow
// playbook action step and defines all of the properties associated with the
// playbook action step. Some properties are inherited from the
// workflow.CommonProperties type.
type PlaybookActionStep struct {
	workflow.CommonProperties
	PlaybookID      string   `json:"playbook_id,omitempty"`
	PlaybookVersion string   `json:"playbook_version,omitempty"`
	InArgs          []string `json:"in_args,omitempty"`
	OutArgs         []string `json:"out_args,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow playbook action step object
// and return it as a pointer. It will also initialize the object by setting
// all of the basic properties.
func New() *PlaybookActionStep {
	var w PlaybookActionStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow playbook action step object
// with the correct defaults.
func (w *PlaybookActionStep) Init() {
	w.ObjectType = "playbook-action"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *PlaybookActionStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}
