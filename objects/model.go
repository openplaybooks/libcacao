// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package objects

// ExternalReference - This type defines all of the properties associated with
// the External Reference type. The external-reference data type captures the
// location of information represented outside of a CACAO playbook and uses the
// JSON object type [RFC8259] for serialization. For example, a playbook could
// reference external documentation about a specific piece of malware that the
// playbook is trying to address. In addition to the name properties at least
// one of the following properties MUST be present: description, source, url, or
// external_id.
type ExternalReference struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Source      string `json:"source,omitempty"`
	URL         string `json:"url,omitempty"`
	ExternalID  string `json:"external_id,omitempty"`
	ReferenceID string `json:"reference_id,omitempty"`
}

// Variables - This type defines all of the properties assocaited with the
// Variables data type.  The variable data type captures variable information
// and uses the JSON object type [RFC8259] for serialization. Variables can be
// defined and used as the playbook is executed and are stored in a dictionary
// where the key is the name of the variable and the value is a variable data
// type. Variables can represent stateful elements that may need to be captured
// to allow for the successful execution of the playbook. All playbook variables
// are mutable unless identified as a constant.
type Variables struct {
	ObjectType  string `json:"type,omitempty"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value,omitempty"`
	Constant    bool   `json:"constant,omitempty"`
	External    bool   `json:"external,omitempty"`
}
