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
