// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package statement implements the CACAO 1.0 statement marking object. The
// statement marking object defines the representation of a textual marking
// statement (e.g., copyright, terms of use, etc.). Statement markings are
// generally not machine-readable, and this specification does not define any
// behavior or actions based on their values.
package statement

import "github.com/openplaybooks/libcacao/objects/markings"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// MarkingStatement - This type implements the CACAO 1.0 statement marking
// object and defines all of the properties associated with the statement
// marking type. Some properties are inherited from the
// markings.CommonProperties type.
type MarkingStatement struct {
	markings.CommonProperties
	Statement string `json:"statement,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new statement marking object and return it
// as a pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *MarkingStatement {
	var m MarkingStatement
	m.Init()
	return &m
}

// Init - This method will initialize a new statement marking object with the
// correct defaults.
func (m *MarkingStatement) Init() {
	m.ObjectType = "marking-statement"
	m.SetNewID(m.ObjectType)
	m.Created = m.GetCurrentTime("milli")
	m.Modified = m.Created
	m.Revoked = false
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common data marking properties
func (m *MarkingStatement) GetCommon() markings.CommonProperties {
	return m.CommonProperties
}
