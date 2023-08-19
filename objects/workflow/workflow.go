// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// StepObject - This interface defines a workflow step object. I needed to add
// the ClearID() function to the interface to make sure I could call it on
// an object that is defined as fullfilling this interface.
type StepObject interface {
	GetCommon() CommonProperties
	ClearID()
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type
// and creates a new ID based on the specification format and updates the id
// property for the object.
func (w *CommonProperties) SetNewID(objType string) error {

	if !objects.IsVocabValueValid(objType, GetWorkflowStepTypesVocab()) {
		return fmt.Errorf("the object type %s is not a valid CACAO worflow step type", objType)
	}

	w.ID, _ = objects.CreateID(objType)
	return nil
}

// GetID - This method returns the ID of the step object
func (w *CommonProperties) GetID() string {
	return w.ID
}

// ClearID - This method will clear the ID from the object
func (w *CommonProperties) ClearID() {
	w.ID = ""
}

// AddVariable - This method takes in a Variable object and adds it to the
// workflow step object as a local step variable.
func (w *CommonProperties) AddVariable(v objects.Variables) error {
	if !objects.IsVocabValueValid(v.ObjectType, objects.GetVariableTypesVocab()) {
		return fmt.Errorf("the variable type %s is not valid", v.ObjectType)
	}

	if w.StepVariables == nil {
		m := make(map[string]objects.Variables, 0)
		w.StepVariables = m
	}
	name := v.Name
	v.Name = ""
	w.StepVariables[name] = v
	return nil
}
