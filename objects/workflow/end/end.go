// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package end implements the CACAO 2.0 workflow end step object.
//
// The End Step workflow step is the ending point of a playbook or branch of
// step (e.g., a list of steps that are part of a parallel processing branch)
// and represents an explicit point in the workflow to signify the end of a
// playbook or branch of steps. While this type inherits all of the common
// properties of a workflow step it does not define any additional properties.
// When a playbook or branch of a playbook terminates it MUST call an End Step.
// This workflow step MUST NOT use the on_completion, on_success, or on_failure
// properties.
package end

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// EndStep - This type implmenets the CACAO 2.0 workflow end step
// and defines all of the properties associated with the end step. Some
// properties are inherited from the workflow.CommonProperties type.
type EndStep struct {
	workflow.CommonProperties
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow end step object and return
// it as a pointer. It will also initialize the object by setting all of the
// basic properties.
func New() *EndStep {
	var w EndStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow end step object with the
// correct defaults.
func (w *EndStep) Init() {
	w.ObjectType = "end"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *EndStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}
