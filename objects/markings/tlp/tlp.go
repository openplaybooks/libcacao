// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package tlp implements the CACAO 1.0 tlp marking object. The TLP marking
// object defines the representation of a FIRST TLP marking statement. If the
// TLP marking is externally defined, producers SHOULD use the
// external_references property of this object.
package tlp

import "github.com/openplaybooks/libcacao/objects/markings"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// MarkingTLP - This type implements the CACAO 1.0 TLP data marking object and
// defines all of the properties associated with the TLP marking type. Some
// properties are inherited from the markings.CommonProperties type.
type MarkingTLP struct {
	markings.CommonProperties
	TLPv2Level string `json:"tlpv2_level,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new tlp marking object and return it
// as a pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *MarkingTLP {
	var m MarkingTLP
	// m.Init()
	return &m
}

// Init - This method will initialize a new tlp marking object with the
// correct defaults.
// func (m *MarkingTLP) Init() {
// 	m.ObjectType = "marking-tlp"
// 	m.SetNewID(m.ObjectType)
// 	m.Created = m.GetCurrentTime("milli")
// 	m.Modified = m.Created
// 	m.Revoked = false
// }

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetClear - This method will set the TLP Level to TLP-CLEAR
func (m *MarkingTLP) SetClear() {
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--94868c89-83c2-464b-929b-a1a8aa3c8487"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.Modified = m.Created
	m.TLPv2Level = "TLP-CLEAR"
}

// SetGreen - This method will set the TLP Level to TLP-GREEN
func (m *MarkingTLP) SetGreen() {
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--bab4a63c-aed9-4cf5-a766-dfca5abac2bb"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.Modified = m.Created
	m.TLPv2Level = "TLP-GREEN"
}

// SetAmber - This method will set the TLP Level to TLP-AMBER
func (m *MarkingTLP) SetAmber() {
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--55d920b0-5e8b-4f79-9ee9-91f868d9b421"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.Modified = m.Created
	m.TLPv2Level = "TLP-AMBER"
}

// SetAmberStrict - This method will set the TLP Level to TLP-AMBER+STRICT
func (m *MarkingTLP) SetAmberStrict() {
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--939a9414-2ddd-4d32-a0cd-375ea402b003"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.Modified = m.Created
	m.TLPv2Level = "TLP-AMBER+STRICT"
}

// SetRed - This method will set the TLP Level to TLP-RED
func (m *MarkingTLP) SetRed() {
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--e828b379-4e03-4974-9ac4-e53a884c97c1"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.Modified = m.Created
	m.TLPv2Level = "TLP-RED"
}

// GetCommon - This method returns the common data marking properties
func (m *MarkingTLP) GetCommon() markings.CommonProperties {
	return m.CommonProperties
}
