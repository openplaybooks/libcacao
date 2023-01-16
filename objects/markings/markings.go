// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package markings implements the CACAO 1.0 data marking definition objects.
// CACAO data marking definition objects contain detailed information about a
// specific data marking. Data markings typically represent handling or sharing
// requirements and are applied via the markings property in a playbook.
//
// Data marking objects MUST NOT be versioned because it would allow for
// indirect changes to the markings on a playbook. For example, if a statement
// marking definition is changed from "Reuse Allowed" to "Reuse Prohibited", all
// playbooks marked with that statement marking definition would effectively
// have an updated marking without being updated themselves. Instead, in this
// example a new statement marking definition with the new text should be
// created and the marked objects updated to point to the new data marking
// object.
//
// Playbooks may be marked with multiple marking statements. In other words, the
// same playbook can be marked with both a statement saying "Copyright 2020" and
// a statement saying, "Terms of use are ..." and both statements apply.
package markings

import "github.com/openplaybooks/libcacao/objects"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// DataMarkingObject - This interface defines a data marking object
type DataMarkingObject interface {
	GetCommon() CommonProperties
}

// CommonProperties - Each data marking object contains some base properties
// that are common across all data markings. These common properties are defined
// in the following type.
type CommonProperties struct {
	ObjectType         string                      `json:"type,omitempty"`
	ID                 string                      `json:"id,omitempty"`
	Name               string                      `json:"name,omitempty"`
	Description        string                      `json:"description,omitempty"`
	CreatedBy          string                      `json:"created_by,omitempty"`
	Created            string                      `json:"created,omitempty"`
	Modified           string                      `json:"modified,omitempty"`
	Revoked            bool                        `json:"revoked,omitempty"`
	ValidFrom          string                      `json:"valid_from,omitempty"`
	ValidUntil         string                      `json:"valid_until,omitempty"`
	Labels             []string                    `json:"labels,omitempty"`
	ExternalReferences []objects.ExternalReference `json:"external_references,omitempty"`
}

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
