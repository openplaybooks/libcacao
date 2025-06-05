// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package commands

import (
	"github.com/openplaybooks/libcacao/objects"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// CommandObject - This interface defines a command object. I needed to add
// the ClearID() function to the interface to make sure I could call it on
// an object that is defined as fullfilling this interface.
type CommandObject interface {
	GetCommon() CommonProperties
}

// CommonProperties - Each command contains some base properties that are common
// across all commands. These common properties are defined in the following
// table. The ID property here is just to help make processing easier, it will
// be removed when it is added to the playbook.
type CommonProperties struct {
	ObjectType         string                      `json:"type,omitempty"`
	ID                 string                      `json:"id,omitempty"`
	Description        string                      `json:"description,omitempty"`
	Version            string                      `json:"version,omitempty"`
	PlaybookActivity   string                      `json:"playbook_activity,omitempty"`
	ExternalReferences []objects.ExternalReference `json:"external_references,omitempty"`
}

// Assignment - This type implmenets the CACAO 3.0 assignement command and
// defines all of its properties.
//
// The assignment command enables variables to be assigned or mapped to a
// different variable.
type Assignment struct {
	CommonProperties
	VariableAssignments []VariableAssignment `json:"variable_assignments,omitempty"`
}

// VariableAssignment - This type defines all of the properties associated with
// the variable assignment data type.
type VariableAssignment struct {
	Variable         string `json:"variable,omitempty"`
	OriginalVariable string `json:"original_variable,omitempty"`
	Operator         string `json:"operator,omitempty"`
	Delimiter        string `json:"delimiter,omitempty"`
}

// GetCommon - Implement the CommandObject interface and returns common properties
func (c *Assignment) GetCommon() CommonProperties {
	return c.CommonProperties
}

// Manual - This type implmenets the CACAO 3.0 manual command and defines all of
// its properties.
//
// The manual command represents a command that is intended to be processed by a
// human or a system that acts on behalf of a human. In addition to the
// inherited properties, this section defines the following additional
// properties that are valid for this type.
//
// * Either the command property or the command_b64 property MUST be populated.
type Manual struct {
	CommonProperties
	Command    string     `json:"command,omitempty"`
	CommandB64 string     `json:"command_b64,omitempty"`
	Questions  []Question `json:"questions,omitempty"`
}

// Question - This type defines all of the properties associated with
// the response data type.
type Question struct {
	ObjectType string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	Prompt     string `json:"prompt,omitempty"`
	DataType   string `json:"data_type,omitempty"`
}

// GetCommon - Implement the CommandObject interface and returns common properties
func (c *Manual) GetCommon() CommonProperties {
	return c.CommonProperties
}

// Bash - This type implmenets the CACAO 3.0 bash command and defines all of
// its properties.
//
// The bash command represents a command that is executed via a shell without a
// login/remote connection. In addition to the inherited properties, this
// section defines the following additional properties that are valid for this
// type. Output from these commands can be referenced via the stdout and stderr
// names, see examples below.
//
// * One of the following properties MUST be populated, command or command_b64.
type Bash struct {
	CommonProperties
	Command    string `json:"command,omitempty"`
	CommandB64 string `json:"command_b64,omitempty"`
}

// GetCommon - Implement the CommandObject interface and returns common properties
func (c *Bash) GetCommon() CommonProperties {
	return c.CommonProperties
}

// SSH - This type implmenets the CACAO 3.0 ssh command and defines all of
// its properties.
//
// The ssh command represents a command that is intended to be processed via an
// SSH connection. In addition to the inherited properties, this section
// defines the following additional properties that are valid for this type.
// Output from these commands can be referenced via the stdout and stderr
// names, see examples below.
//
// * One of the following properties MUST be populated, command or command_b64.
type SSH struct {
	CommonProperties
	Command    string `json:"command,omitempty"`
	CommandB64 string `json:"command_b64,omitempty"`
}

// GetCommon - Implement the CommandObject interface and returns common properties
func (c *SSH) GetCommon() CommonProperties {
	return c.CommonProperties
}

// HTTP - This type implmenets the CACAO 3.0 http-api command and defines all
// of its properties.
//
// The HTTP API command represents a command that is intended to be processed
// via an HTTP API. In addition to the inherited properties, this section
// defines the following additional properties that are valid for this type.
type HTTP struct {
	CommonProperties
	Command    string              `json:"command,omitempty"`
	Headers    map[string][]string `json:"headers,omitempty"`
	Content    string              `json:"content,omitempty"`
	ContentB64 string              `json:"content_b64,omitempty"`
}

// GetCommon - Implement the CommandObject interface and returns common properties
func (c *HTTP) GetCommon() CommonProperties {
	return c.CommonProperties
}
