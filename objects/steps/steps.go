// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package steps

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type
// and creates a new ID based on the specification format and updates the id
// property for the object.
func (s *CommonProperties) SetNewID(objType string) error {

	if !objects.IsVocabValueValid(objType, GetWorkflowStepTypesVocab()) {
		return fmt.Errorf("the object type %s is not a valid CACAO worflow step type", objType)
	}

	s.ID, _ = objects.CreateID(objType)
	return nil
}

// GetID - This method returns the ID of the step object
func (s *CommonProperties) GetID() string {
	return s.ID
}

// ClearID - This method will clear the ID from the object
func (s *CommonProperties) ClearID() {
	s.ID = ""
}

// AddVariable - This method takes in a Variable object and adds it to the
// workflow step object as a local step variable.
func (s *CommonProperties) AddVariable(v objects.Variables) error {
	if !objects.IsVocabValueValid(v.ObjectType, objects.GetVariableTypesVocab()) {
		return fmt.Errorf("the variable type %s is not valid", v.ObjectType)
	}

	if s.StepVariables == nil {
		m := make(map[string]objects.Variables, 0)
		s.StepVariables = m
	}
	name := v.Name
	v.Name = ""
	s.StepVariables[name] = v
	return nil
}

// ----------------------------------------------------------------------
// Define Action Step Functions and Methods
// ----------------------------------------------------------------------

// NewExternalReference - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (s *ActionStep) NewExternalReference(r ...objects.ExternalReference) (*objects.ExternalReference, error) {
	positionThatAppendWillUse := len(s.ExternalReferences)

	if len(r) > 0 {
		for i := range r {
			// Update the value so we grab the last one entered
			positionThatAppendWillUse = len(s.ExternalReferences)
			s.ExternalReferences = append(s.ExternalReferences, r[i])
		}
		return &s.ExternalReferences[positionThatAppendWillUse], nil
	}

	// If one was not passed in, lets create one
	var er objects.ExternalReference
	s.ExternalReferences = append(s.ExternalReferences, er)
	return &s.ExternalReferences[positionThatAppendWillUse], nil
}

// ----------------------------------------------------------------------
// Define Parallel Step Functions and Methods
// ----------------------------------------------------------------------

// AddNextSteps - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a next
// steps in the workflow and adds them to the next_steps property.
func (s *ParallelStep) AddNextSteps(values interface{}) error {
	return objects.AddValuesToList(&s.NextSteps, values)
}

// ----------------------------------------------------------------------
// Define Switch Step Functions and Methods
// ----------------------------------------------------------------------

// AddCase - This method takes in two values, the first is a string value
// representing the case of the switch statement. The second is a string value
// representing the identifier to be processed by the case condition.
func (s *SwitchStep) AddCase(k string, id string) error {
	if s.Cases == nil {
		m := make(map[string]string, 0)
		s.Cases = m
	}
	s.Cases[k] = id
	return nil
}
