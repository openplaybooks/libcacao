// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package start implements the CACAO 2.0 workflow start step object.
//
// The Start Step workflow step is the starting point of a playbook and
// represents an explicit entry in the workflow to signify the start of a
// playbook. While this type inherits all of the common properties of a
// workflow step it does not define any additional properties. This workflow
// step MUST NOT use the on_success or on_failure properties.
package start

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// StartStep - This type implmenets the CACAO 2.0 workflow start step
// and defines all of the properties associated with the start step. Some
// properties are inherited from the workflow.CommonProperties type.
type StartStep struct {
	workflow.CommonProperties
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow start step object and return
// it as a pointer. It will also initialize the object by setting all of the
// basic properties.
func New() *StartStep {
	var w StartStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow start step object with the
// correct defaults.
func (w *StartStep) Init() {
	w.ObjectType = "start"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *StartStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}
