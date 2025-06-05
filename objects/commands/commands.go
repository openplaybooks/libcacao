// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package commands

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// AddVariableAssignment - This method will add a variable assignment to an assignment command
func (c *Assignment) AddVariableAssignment(newvar string, origvar string, operator ...string) (*VariableAssignment, error) {
	positionThatAppendWillUse := len(c.VariableAssignments)

	var v VariableAssignment
	v.Variable = newvar
	v.OriginalVariable = origvar

	if operator != nil {

		if !objects.IsVocabValueValid(operator[0], objects.GetOperatorsVocab()) {
			return nil, fmt.Errorf("the operator type %s is not a valid CACAO operator type", operator[0])
		}

		v.Operator = operator[0]

	}

	if v.Operator == "to_list" {

		if !objects.IsVocabValueValid(operator[1], objects.GetDelimitersVocab()) {
			return nil, fmt.Errorf("the delimiter type %s is not a valid CACAO delimiter type", operator[1])
		}
		v.Delimiter = operator[1]
	}

	c.VariableAssignments = append(c.VariableAssignments, v)
	return &c.VariableAssignments[positionThatAppendWillUse], nil
}

// AddQuestion - This method will add a question to a manual command
func (c *Manual) AddQuestion(q string, vartype string) (*Question, error) {
	positionThatAppendWillUse := len(c.Questions)

	var r Question
	r.ObjectType = "question"
	r.ID, _ = objects.CreateID(r.ObjectType)
	r.Prompt = q
	r.DataType = vartype
	c.Questions = append(c.Questions, r)
	return &c.Questions[positionThatAppendWillUse], nil
}

// AddHeader - This method takes values to be added to an HTTP header.
func (c *HTTP) AddHeader(key string, values ...string) error {
	if c.Headers == nil {
		m := make(map[string][]string, 0)
		c.Headers = m
	}

	for _, v := range values {
		c.Headers[key] = append(c.Headers[key], v)
	}

	return nil
}

// SetNewID - This method takes in a string value representing an object type
// and creates a new ID based on the specification format and updates the id
// property for the object.
func (c *CommonProperties) SetNewID(objType string) error {

	if !objects.IsVocabValueValid(objType, GetCommandTypesVocab()) {
		return fmt.Errorf("the object type %s is not a valid CACAO command type", objType)
	}

	c.ID, _ = objects.CreateID(objType)
	return nil
}

// GetID - This method returns the ID of the step object
func (c *CommonProperties) GetID() string {
	return c.ID
}

// ClearID - This method will clear the ID from the object
func (c *CommonProperties) ClearID() {
	c.ID = ""
}
