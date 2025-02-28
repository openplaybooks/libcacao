// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

import (
	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Action Step Functions and Methods
// ----------------------------------------------------------------------

// NewCommand - This method creates a new empty command and returns a reference
// to it so it can be populated. However, if one or more commands are passed in
// they are all added and the reference that is returned is for the last entry
// added.
// func (s *ActionStep) NewCommand(r ...CommandData) (*CommandData, error) {
// 	positionThatAppendWillUse := len(s.Commands)

// 	if len(r) > 0 {
// 		for i := range r {
// 			// Update the value so we grab the last one entered
// 			positionThatAppendWillUse = len(s.Commands)
// 			s.Commands = append(s.Commands, r[i])
// 		}
// 		return &s.Commands[positionThatAppendWillUse], nil
// 	}

// 	// If one was not passed in, lets create one
// 	var c CommandData
// 	s.Commands = append(s.Commands, c)
// 	return &s.Commands[positionThatAppendWillUse], nil
// }

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
