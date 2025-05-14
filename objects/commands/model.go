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
	Version            string                      `json:"version,omitempty"`
	Description        string                      `json:"description,omitempty"`
	PlaybookActivity   string                      `json:"playbook_activity,omitempty"`
	ExternalReferences []objects.ExternalReference `json:"external_references,omitempty"`
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
	Command    string         `json:"command,omitempty"`
	CommandB64 string         `json:"command_b64,omitempty"`
	Questions  []QuestionData `json:"questions,omitempty"`
}

// QuestionData - This type defines all of the properties associated with
// the response data type.
type QuestionData struct {
	ObjectType string `json:"type,omitempty"`
	ID         string `json:"id,omitempty"`
	Question   string `json:"question,omitempty"`
	DataType   string `json:"data_type,omitempty"`
}

// GetCommon - Implement the CommandObject interface and return common properties
func (c *Manual) GetCommon() CommonProperties {
	return c.CommonProperties
}

// Bash - This type implmenets the CACAO 3.0 bash command and defines all of
// its properties.
//
// The bash command represents a command that is intended to be processed via a
// shell without a login/remote connection. In addition to the inherited
// properties, this section defines the following additional properties that
// are valid for this type.
//
// * One of the following properties MUST be populated, command or command_b64.
type Bash struct {
	CommonProperties
	Command      string            `json:"command,omitempty"`
	CommandB64   string            `json:"command_b64,omitempty"`
	ReturnedData map[string]string `json:"returned_data,omitempty"`
}

// GetCommon - Implement the CommandObject interface and return common properties
func (c *Bash) GetCommon() CommonProperties {
	return c.CommonProperties
}

// HTTPAPI - This type implmenets the CACAO 3.0 http-api command and defines all
// of its properties.
//
// The HTTP API command represents a command that is intended to be processed
// via an HTTP API. In addition to the inherited properties, this section
// defines the following additional properties that are valid for this type.
type HTTPAPI struct {
	CommonProperties
	Command      string              `json:"command,omitempty"`
	Headers      map[string][]string `json:"headers,omitempty"`
	Content      string              `json:"content,omitempty"`
	ContentB64   string              `json:"content_b64,omitempty"`
	ReturnedData map[string]string   `json:"returned_data,omitempty"`
}

// GetCommon - Implement the CommandObject interface and return common properties
func (c *HTTPAPI) GetCommon() CommonProperties {
	return c.CommonProperties
}
