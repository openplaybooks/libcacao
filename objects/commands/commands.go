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

// AddQuestion -
func (c *Manual) AddQuestion(q string, vartype string) (*QuestionData, error) {
	var r QuestionData
	r.ObjectType = "question"
	r.ID, _ = objects.CreateID(r.ObjectType)
	r.Question = q
	r.DataType = vartype
	c.Questions = append(c.Questions, r)
	return &r, nil
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
