// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package tlp implements the CACAO 1.0 tlp marking object. The TLP marking
// object defines the representation of a FIRST TLP marking statement. If the
// TLP marking is externally defined, producers SHOULD use the
// external_refernces property of this object.
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
	TLPLevel string `json:"tlp_level,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new tlp marking object and return it
// as a pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *MarkingTLP {
	var m MarkingTLP
	m.Init()
	return &m
}

// Init - This method will initialize a new tlp marking object with the
// correct defaults.
func (m *MarkingTLP) Init() {
	m.ObjectType = "marking-tlp"
	m.SpecVersion = m.GetCurrentSpecVersion()
	m.SetNewID(m.ObjectType)
	m.Created = m.GetCurrentTime("milli")
	m.Modified = m.Created
	m.Revoked = false
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetWhite - This method will set the TLP Level to TLP-WHITE
func (m *MarkingTLP) SetWhite() {
	m.TLPLevel = "TLP-WHITE"
}

// SetGreen - This method will set the TLP Level to TLP-GREEN
func (m *MarkingTLP) SetGreen() {
	m.TLPLevel = "TLP-GREEN"
}

// SetAmber - This method will set the TLP Level to TLP-AMBER
func (m *MarkingTLP) SetAmber() {
	m.TLPLevel = "TLP-AMBER"
}

// SetRed - This method will set the TLP Level to TLP-RED
func (m *MarkingTLP) SetRed() {
	m.TLPLevel = "TLP-RED"
}

// GetCommon - This method returns the common data marking properties
func (m *MarkingTLP) GetCommon() markings.CommonProperties {
	return m.CommonProperties
}
