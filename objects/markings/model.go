// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

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
	Revoked            bool                        `json:"revoked,omitempty"`
	ValidFrom          string                      `json:"valid_from,omitempty"`
	ValidUntil         string                      `json:"valid_until,omitempty"`
	Labels             []string                    `json:"labels,omitempty"`
	ExternalReferences []objects.ExternalReference `json:"external_references,omitempty"`
	//marking_extensions
}

// MarkingTLP - This type implements the FIRST TLPv2 object as a CACAO 2.0 TLP
// data marking object and defines all of the properties associated with the
// TLP marking type.
type MarkingTLP struct {
	CommonProperties
	TLPv2Level string `json:"tlpv2_level,omitempty"`
}

// MarkingStatement - This type implements the CACAO 2.0 statement marking
// object and defines all of the properties associated with the statement
// marking type.
//
// The statement marking object defines the representation of a textual marking
// statement (e.g., copyright, terms of use, etc.). Statement markings are
// generally not machine-readable.
type MarkingStatement struct {
	CommonProperties
	Statement string `json:"statement,omitempty"`
}

// MarkingIEP - This type implements the CACAO 2.0 IEP data marking object and
// defines all of the properties associated with the IEP marking type.
//
// The IEP marking object defines the representation of a FIRST IEP marking
// statement. For more information about the properties from the IEP
// specification, please refer to that document [IEP].
type MarkingIEP struct {
	CommonProperties
	TLP                        string   `json:"tlp,omitempty"`
	IEPVersion                 string   `json:"iep_version,omitempty"`
	StartDate                  string   `json:"start_date,omitempty"`
	EndDate                    string   `json:"end_date,omitempty"`
	EncryptInTransit           string   `json:"encrypt_in_transit,omitempty"`
	PermittedActions           string   `json:"permitted_actions,omitempty"`
	AffectedPartyNotifications string   `json:"affected_party_notifications,omitempty"`
	Attribution                string   `json:"attribution,omitempty"`
	UnmodifiedResale           string   `json:"unmodified_resale,omitempty"`
	ExternalReferences         []string `json:"external_references,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// NewTLPClearMarking - This will create and initialize a new tlp clear marking
// object and return it as a pointer.
func NewTLPClearMarking() *MarkingTLP {
	var m MarkingTLP
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--94868c89-83c2-464b-929b-a1a8aa3c8487"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.TLPv2Level = "TLP-CLEAR"
	return &m
}

// NewTLPGreenMarking - This will create and initialize a new tlp green marking
// object and return it as a pointer.
func NewTLPGreenMarking() *MarkingTLP {
	var m MarkingTLP
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--bab4a63c-aed9-4cf5-a766-dfca5abac2bb"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.TLPv2Level = "TLP-GREEN"
	return &m
}

// NewTLPAmberMarking - This will create and initialize a new tlp amber marking
// object and return it as a pointer.
func NewTLPAmberMarking() *MarkingTLP {
	var m MarkingTLP
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--55d920b0-5e8b-4f79-9ee9-91f868d9b421"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.TLPv2Level = "TLP-AMBER"
	return &m
}

// NewTLPAmberStrictMarking - This will create and initialize a new tlp amber
// strict marking object and return it as a pointer.
func NewTLPAmberStrictMarking() *MarkingTLP {
	var m MarkingTLP
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--939a9414-2ddd-4d32-a0cd-375ea402b003"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.TLPv2Level = "TLP-AMBER+STRICT"
	return &m
}

// NewTLPRedMarking - This will create and initialize a new tlp red marking
// object and return it as a pointer.
func NewTLPRedMarking() *MarkingTLP {
	var m MarkingTLP
	m.ObjectType = "marking-tlp"
	m.ID = "marking-tlp--e828b379-4e03-4974-9ac4-e53a884c97c1"
	m.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m.Created = "2022-10-01T00:00:00.000Z"
	m.TLPv2Level = "TLP-RED"
	return &m
}

// NewStatementMarking - This will create and initialize a new statement marking
// object and return it as a pointer.
func NewStatementMarking() *MarkingStatement {
	var m MarkingStatement
	m.ObjectType = "marking-statement"
	m.SetNewID(m.ObjectType)
	m.Created = m.GetCurrentTime("milli")
	m.Revoked = false
	return &m
}

// NewIEPMarking - This function will create a new iep marking object and return it
// as a pointer.
func NewIEPMarking() *MarkingIEP {
	var m MarkingIEP
	return &m
}
