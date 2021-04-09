// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package ifcondition implements the CACAO 1.0 workflow if condition step object.
//
// This section defines the 'if-then-else' conditional logic that can be used
// within the workflow of the playbook. In addition to the inherited properties,
// this section defines three additional specific properties that are valid for
// this type.
package ifcondition

import "github.com/jordan2175/libcacao/objects/workflow"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// WorkflowIfStep - This type implmenets the CACAO 1.0 workflow if condition
// step and defines all of the properties associated with the if condition step.
// Some properties are inherited from the workflow.CommonProperties type.
type WorkflowIfStep struct {
	workflow.CommonProperties
	Condition string   `json:"condition,omitempty"`
	OnTrue    []string `json:"on_true,omitempty"`
	OnFalse   []string `json:"on_false,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow if condition step object and
// return it as a pointer. It will also initialize the object by setting all of
// the basic properties.
func New() *WorkflowIfStep {
	var w WorkflowIfStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow if condition step object with
// the correct defaults.
func (w *WorkflowIfStep) Init() {
	w.ObjectType = "if-condition"
	w.SetNewID(w.ObjectType)
}
