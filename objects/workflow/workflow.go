// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package workflow implements the CACAO 1.0 workflow step objects.
//
// Workflows contain a series of steps that are stored in a dictionary, where
// the key is the step ID and the value is a workflow step. These workflow steps
// along with the associated commands form the building blocks for playbooks and
// are used to control the commands that need to be executed. Workflows process
// steps either sequentially, in parallel, or both depending on the type of
// steps required by the playbook. In addition to simple processing, workflow
// steps MAY also contain conditional and/or temporal operations to control the
// execution of the playbook.
//
// Conditional processing means executing steps or commands after some sort of
// condition is met. Temporal processing means executing steps or commands
// either during a certain time window or after some period of time has passed.
//
// This section defines the various workflow steps and how they may be used to
// define a playbook.
package workflow

import "github.com/openplaybooks/libcacao/objects"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// StepObject - This interface defines a workflow step object
type StepObject interface {
	GetCommon() CommonProperties
}

// CommonProperties - Each workflow step contains some base properties that are
// common across all steps. These common properties are defined in the following
// table.
type CommonProperties struct {
	ObjectType         string                       `json:"type,omitempty"`
	ID                 string                       `json:"id,omitempty"`
	Name               string                       `json:"name,omitempty"`
	Description        string                       `json:"description,omitempty"`
	ExternalReferences []objects.ExternalReference  `json:"external_references,omitempty"`
	Delay              int                          `json:"delay,omitempty"`
	Timeout            int                          `json:"timeout,omitempty"`
	StepVariables      map[string]objects.Variables `json:"playbook_variables,omitempty"`
	Owner              string                       `json:"owner,omitempty"`
	OnSuccess          string                       `json:"on_success,omitempty"`
	OnFailure          string                       `json:"on_failure,omitempty"`
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// SetNewID - This method takes in a string value representing an object type and
// creates a new ID based on the specification format and update the id property
// for the object.
func (w *CommonProperties) SetNewID(s string) error {
	// TODO Add check to validate input value
	w.ID, _ = objects.CreateID(s)
	return nil
}

// GetID - This method returns the ID of the step object
func (w *CommonProperties) GetID() string {
	return w.ID
}
