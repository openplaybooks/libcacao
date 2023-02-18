// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package switchcondition implements the CACAO 2.0 workflow switch condition step object.
//
// The Switch Condition Step workflow step defines the 'switch' condition logic
// that can be used within the workflow of the playbook. In addition to the
// inherited properties, this section defines two additional specific
// properties that are valid for this type.
package switchcondition

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// SwitchStep - This type implmenets the CACAO 2.0 workflow switch
// condition step and defines all of the properties associated with the switch
// condition step. Some properties are inherited from the
// workflow.CommonProperties type.
type SwitchStep struct {
	workflow.CommonProperties
	Switch string              `json:"switch,omitempty"`
	Cases  map[string][]string `json:"cases,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow switch condition step object
// and return it as a pointer. It will also initialize the object by setting all
// of the basic properties.
func New() *SwitchStep {
	var w SwitchStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow switch condition step
// object with the correct defaults.
func (w *SwitchStep) Init() {
	w.ObjectType = "switch-condition"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *SwitchStep) GetCommon() workflow.CommonProperties {
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
