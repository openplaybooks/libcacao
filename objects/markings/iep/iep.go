// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package iep implements the CACAO 2.0 iep marking object.
//
// The IEP marking object defines the representation of a FIRST IEP marking
// statement. For more information about the properties from the IEP
// specification, please refer to that document [IEP].
package iep

import "github.com/openplaybooks/libcacao/objects/markings"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// MarkingIEP - This type implements the CACAO 2.0 IEP data marking object and
// defines all of the properties associated with the IEP marking type. Some
// properties are inherited from the markings.CommonProperties type.
type MarkingIEP struct {
	markings.CommonProperties
	Name             string `json:"name,omitempty"`
	TLPLevel         string `json:"tlp_level,omitempty"`
	Description      string `json:"description,omitempty"`
	IEPVersion       string `json:"iep_version,omitempty"`
	StartDate        string `json:"start_date,omitempty"`
	EndDate          string `json:"end_date,omitempty"`
	EncryptInTransit string `json:"encrypt_in_transit,omitempty"`
	PermittedActions string `json:"permitted_actions,omitempty"`
	Attribution      string `json:"attribution,omitempty"`
	UnmodifiedResale string `json:"unmodified_resale,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new tlp marking object and return it
// as a pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *MarkingIEP {
	var m MarkingIEP
	// m.Init()
	return &m
}

// Init - This method will initialize a new tlp marking object with the
// correct defaults.
// func (m *MarkingIEP) Init() {
// 	m.ObjectType = "marking-iep"
// 	m.SetNewID(m.ObjectType)
// 	m.Created = m.GetCurrentTime("milli")
// 	m.Modified = m.Created
// 	m.Revoked = false
// }

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common data marking properties
func (m *MarkingIEP) GetCommon() markings.CommonProperties {
	return m.CommonProperties
}
