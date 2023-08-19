// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package markings

import "github.com/openplaybooks/libcacao/objects"

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type and
// creates a new ID based on the specification format and update the id property
// for the object.
func (m *CommonProperties) SetNewID(s string) error {
	// TODO Add check to validate input value
	m.ID, _ = objects.CreateID(s)
	return nil
}

// GetID - This method returns the ID of the marking object
func (m *CommonProperties) GetID() string {
	return m.ID
}

// GetCurrentTime - This method takes in a value of either milli or micro and
// returns the current time in RFC 3339 format
func (m *CommonProperties) GetCurrentTime(precision string) string {
	return objects.GetCurrentTime(precision)
}

// NewExternalReference - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (m *CommonProperties) NewExternalReference(r ...objects.ExternalReference) (*objects.ExternalReference, error) {
	positionThatAppendWillUse := len(m.ExternalReferences)

	if len(r) > 0 {
		for i := range r {
			// Update the value so we grab the last one entered
			positionThatAppendWillUse = len(m.ExternalReferences)
			m.ExternalReferences = append(m.ExternalReferences, r[i])
		}
		return &m.ExternalReferences[positionThatAppendWillUse], nil
	}

	// If one was not passed in, lets create one
	var er objects.ExternalReference
	m.ExternalReferences = append(m.ExternalReferences, er)
	return &m.ExternalReferences[positionThatAppendWillUse], nil
}
