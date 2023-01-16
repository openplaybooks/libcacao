// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package start implements the CACAO 1.0 workflow start step object.
//
// This workflow step is the starting point of a playbook or branch of steps.
// While this type inherits all of the common properties of a workflow step it
// does not define any additional properties.
package start

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// WorkflowStartStep - This type implmenets the CACAO 1.0 workflow start step
// and defines all of the properties associated with the start step. Some
// properties are inherited from the workflow.CommonProperties type.
type WorkflowStartStep struct {
	workflow.CommonProperties
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow start step object and return
// it as a pointer. It will also initialize the object by setting all of the
// basic properties.
func New() *WorkflowStartStep {
	var w WorkflowStartStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow start step object with the
// correct defaults.
func (w *WorkflowStartStep) Init() {
	w.ObjectType = "start"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *WorkflowStartStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}
