// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package whilecondition implements the CACAO 2.0 workflow if condition step object.
//
// The While Condition Step workflow step defines the 'while' conditional logic
// that can be used within the workflow of the playbook. In addition to the
// inherited properties, this section defines three additional specific
// properties that are valid for this type.
package whilecondition

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// WhileStep - This type implmenets the CACAO 2.0 workflow while condition
// step and defines all of the properties associated with the if condition step.
// Some properties are inherited from the workflow.CommonProperties type.
type WhileStep struct {
	workflow.CommonProperties
	Condition string   `json:"condition,omitempty"`
	OnTrue    []string `json:"on_true,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow if condition step object and
// return it as a pointer. It will also initialize the object by setting all of
// the basic properties.
func New() *WhileStep {
	var w WhileStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow while condition step object
// with the correct defaults.
func (w *WhileStep) Init() {
	w.ObjectType = "while-condition"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *WhileStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}

// AddOnTrue - This method takes in a string value, a comma separated list of
// string values, or a slice of string values and add them to the on_true
// property. Each entry represents one or more identifiers to be processed if
// the condition returns "true".
func (w *WhileStep) AddOnTrue(values interface{}) error {
	return objects.AddValuesToList(&w.OnTrue, values)
}
