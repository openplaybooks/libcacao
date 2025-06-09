// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package steps

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/commands"
	"github.com/openplaybooks/libcacao/objects/variables"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// StepObject - This interface defines a workflow step object. I needed to add
// the ClearID() function to the interface to make sure I could call it on
// an object that is defined as fullfilling this interface.
type StepObject interface {
	GetCommon() CommonProperties
	ClearID()
}

// CommonProperties - Each workflow step contains some base properties that are
// common across all steps. These common properties are defined in the following
// table. The ID property here is just to help make processing easier, it will
// be removed when it is added to the playbook.
type CommonProperties struct {
	ObjectType    string                              `json:"type,omitempty"`
	ID            string                              `json:"id,omitempty"`
	Name          string                              `json:"name,omitempty"`
	Description   string                              `json:"description,omitempty"`
	Owner         string                              `json:"owner,omitempty"`
	Delay         int                                 `json:"delay,omitempty"`
	StepVariables map[string]variables.VariableObject `json:"step_variables,omitempty"`
	//Coordinates   CanvasLayout                        `json:"coordinates,omitempty"`
	// StepExtensions
}

// ClearID - This method will clear the ID from the object
func (s *CommonProperties) ClearID() {
	s.ID = ""
}

// CanvasLayout - captures the relative position (in pixels) of a visualized
// CACAO object from its midpoint to the upper left corner (origin) of a
// canvas. In addition, it also captures any connections (paths) from it to
// other objects.
type CanvasLayout struct {
	X                   int          `json:"x,omitempty"`
	Y                   int          `json:"y,omitempty"`
	OutgoingConnections []Connection `json:"connections,omitempty"`
}

type Connection struct {
	ConnectionType string `json:"connection_type,omitempty"`
	X              []int  `json:"x,omitempty"`
	Y              []int  `json:"y,omitempty"`
}

// StartStep - This type implmenets the CACAO 3.0 start step and defines all of
// its properties.
//
// The start step object is used to define an explicit starting point of a
// playbook.
type StartStep struct {
	CommonProperties
	OnSuccess string `json:"on_success,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *StartStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// ReturnStep - This type implmenets the CACAO 3.0 return step and defines all
// of its properties.
//
// The return step object is used to define when processing MUST return to the
// step that started the branch.
type ReturnStep struct {
	CommonProperties
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *ReturnStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// EndStep - This type implmenets the CACAO 3.0 end step and defines all of its
// properties.
//
// The end step object is used to define an explicit ending point of a playbook.
// When a playbook terminates it MUST call an end step.
type EndStep struct {
	CommonProperties
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *EndStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// ActionStep - This type implmenets the CACAO 3.0 workflow action step and
// defines all of the properties associated with the action step.
//
// The Action Step workflow step contains the actual commands to be executed on
// one or more agents. These commands are intended to be processed
// sequentially.
type ActionStep struct {
	CommonProperties
	Commands           []commands.CommandObject    `json:"commands,omitempty"`
	Timeout            int                         `json:"timeout,omitempty"`
	OnTimeout          string                      `json:"on_timeout,omitempty"`
	OnSuccess          string                      `json:"on_success,omitempty"`
	OnFailure          string                      `json:"on_failure,omitempty"`
	Agent              string                      `json:"agent,omitempty"`
	Targets            []string                    `json:"targets,omitempty"`
	ExternalReferences []objects.ExternalReference `json:"external_references,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *ActionStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// PlaybookActionStep - This type implmenets the CACAO 3.0 workflow playbook
// action step and defines all of the properties associated with the playbook
// action step.
//
// The Playbook Action Step workflow step executes a referenced playbook. In
// addition to the inherited properties, this section defines four more
// specific properties that are valid for this type.
type PlaybookActionStep struct {
	CommonProperties
	PlaybookID      string   `json:"playbook_id,omitempty"`
	PlaybookVersion string   `json:"playbook_version,omitempty"`
	InArgs          []string `json:"in_args,omitempty"`
	OutArgs         []string `json:"out_args,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *PlaybookActionStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// ParallelStep - This type implmenets the CACAO 3.0 parallel step and defines
// all of its properties.
//
// The parallel step object can be used to enable two or more steps to be
// processed at the same time, or in other words processed in parallel, see
// Fig. 7-2. The definition of parallel execution and how many parallel steps
// can be processed at once is implementation dependent and is not part of this
// specification.
type ParallelStep struct {
	CommonProperties
	NextSteps []string `json:"next_steps,omitempty"`
	OnSuccess string   `json:"on_success,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *ParallelStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// ForeachStep - This type implmenets the CACAO 3.0 foreach step and defines all
// of its properties.
//
// The foreach step object can be used to perform one or more operations on each
// element in a list data.
type ForeachStep struct {
	CommonProperties
	Collection string `json:"collection,omitempty"`
	Element    string `json:"element,omitempty"`
	Do         string `json:"do,omitempty"`
	OnSuccess  string `json:"on_success,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *ForeachStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// WhileStep - This type implmenets the CACAO 3.0 while step and defines all of
// its properties.
//
// The while step object can be used to enable a 'while' loop within the
// workflow of a playbook.
type WhileStep struct {
	CommonProperties
	Condition string `json:"condition,omitempty"`
	OnTrue    string `json:"on_true,omitempty"`
	OnFalse   string `json:"on_false,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *WhileStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// IfThenStep - This type implmenets the CACAO 3.0 if then step and defines all
// of its properties.
//
// The if-then step object can be used to enable 'if-then' logic within the
// workflow of a playbook. This object is processed in-line with the general
// workflow of the playbook and therefore does NOT have the on_success or the
// on_failure properties.
type IfThenStep struct {
	CommonProperties
	Condition string `json:"condition,omitempty"`
	OnTrue    string `json:"on_true,omitempty"`
	OnFalse   string `json:"on_false,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *IfThenStep) GetCommon() CommonProperties {
	return s.CommonProperties
}

// SwitchStep - This type implmenets the CACAO 3.0 switch step and defines all
// of its properties.
//
// The switch step object can be used to enable the 'switch' logic within the
// workflow of a playbook. This object is processed in-line with the general
// workflow of the playbook and therefore does NOT have the on_success or the
// on_failure properties.
type SwitchStep struct {
	CommonProperties
	Switch string            `json:"switch,omitempty"`
	Cases  map[string]string `json:"cases,omitempty"`
}

// GetCommon - Implement the StepObject interface and return common properties
func (s *SwitchStep) GetCommon() CommonProperties {
	return s.CommonProperties
}
