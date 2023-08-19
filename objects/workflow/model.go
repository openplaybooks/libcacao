// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package workflow

import "github.com/openplaybooks/libcacao/objects"

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// CommonProperties - Each workflow step contains some base properties that are
// common across all steps. These common properties are defined in the following
// table. The ID property here is just to help make processing easier, it will
// be removed when it is added to the playbook.
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
	OnCompletion       string                       `json:"on_completion,omitempty"`
	OnSuccess          string                       `json:"on_success,omitempty"`
	OnFailure          string                       `json:"on_failure,omitempty"`
	//StepExtensions
}

// StartStep - This type implmenets the CACAO 2.0 workflow start step and
// defines all of the properties associated with the start step.
//
// The Start Step workflow step is the starting point of a playbook and
// represents an explicit entry in the workflow to signify the start of a
// playbook. This workflow step MUST NOT use the on_success or on_failure
// properties.
type StartStep struct {
	CommonProperties
}

// EndStep - This type implmenets the CACAO 2.0 workflow end step and defines
// all of the properties associated with the end step.
//
// The End Step workflow step is the ending point of a playbook or branch of
// step (e.g., a list of steps that are part of a parallel processing branch)
// and represents an explicit point in the workflow to signify the end of a
// playbook or branch of steps. When a playbook or branch of a playbook
// terminates it MUST call an End Step. This workflow step MUST NOT use the
// on_completion, on_success, or on_failure properties.
type EndStep struct {
	CommonProperties
}

// CommandData - This type implement the CACAO 2.0 command data type.
//
// The CACAO command object (command-data) contains detailed information about
// the commands that are to be executed or processed automatically or manually
// as part of an action step (see section 4.5). Each command listed in an
// action step may be of a different command type, however, all commands listed
// in a single step MUST be processed or executed by all of the agents defined
// in that step.
//
// Commands can use and refer to variables just like other parts of the
// playbook. For each command either the command property or the command_b64
// property MUST be present.
//
// The individual commands MAY be defined in other specifications, and when
// possible will be mapped to the JSON structure of this specification. When
// that is not possible, they will be base64 encoded.
type CommandData struct {
	ObjectType       string `json:"type,omitempty"`
	Description      string `json:"description,omitempty"`
	Command          string `json:"command,omitempty"`
	CommandB64       string `json:"command_b64,omitempty"`
	Version          string `json:"version,omitempty"`
	PlaybookActivity string `json:"playbook_activity,omitempty"`
}

// ActionStep - This type implmenets the CACAO 2.0 workflow action step and
// defines all of the properties associated with the action step.
//
// The Action Step workflow step contains the actual commands to be executed on
// one or more agents. These commands are intended to be processed
// sequentially.
type ActionStep struct {
	CommonProperties
	Commands []CommandData `json:"commands,omitempty"`
	Agent    string        `json:"agent,omitempty"`
	Targets  []string      `json:"targets,omitempty"`
	InArgs   []string      `json:"in_args,omitempty"`
	OutArgs  []string      `json:"out_args,omitempty"`
}

// PlaybookActionStep - This type implmenets the CACAO 2.0 workflow playbook
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

// ParallelStep - This type implmenets the CACAO 2.0 workflow parallel
// step and defines all of the properties associated with the parallel step.
//
// The Parallel Step workflow step defines how to create steps that are
// processed in parallel. This workflow step allows playbook authors to define
// two or more steps that can be executed at the same time. For example, a
// playbook that responds to an incident may require both the network team and
// the desktop team to investigate and respond to a threat at the same time.
// Another example is a response to a cyber attack on an operational
// technology (OT) environment that requires releasing air / steam / water
// pressure simultaneously. In addition to the inherited properties, this
// section defines one additional specific property that is valid for this
// type. Implementations MUST wait for all steps referenced in the next_steps
// property to complete before moving on.
//
// The steps referenced from this object are intended to be processed in
// parallel, however, if an implementation cannot support executing steps in
// parallel, then the steps MAY be executed in sequential order if the desired
// outcome is the same.
type ParallelStep struct {
	CommonProperties
	NextSteps []string `json:"next_steps,omitempty"`
}

// IfStep - This type implmenets the CACAO 2.0 workflow if condition
// step and defines all of the properties associated with the if condition step.
type IfStep struct {
	CommonProperties
	Condition string   `json:"condition,omitempty"`
	OnTrue    []string `json:"on_true,omitempty"`
	OnFalse   []string `json:"on_false,omitempty"`
}

// WhileStep - This type implmenets the CACAO 2.0 workflow while condition step
// and defines all of the properties associated with the if condition step.
type WhileStep struct {
	CommonProperties
	Condition string   `json:"condition,omitempty"`
	OnTrue    []string `json:"on_true,omitempty"`
}

// SwitchStep - This type implmenets the CACAO 2.0 workflow switch condition
// step and defines all of the properties associated with the switch condition
// step. Some properties are inherited from the workflow.CommonProperties
// type.
type SwitchStep struct {
	CommonProperties
	Switch string              `json:"switch,omitempty"`
	Cases  map[string][]string `json:"cases,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// NewStartStep - This will create and initialize a new workflow start step object
// and return it as a pointer.
func NewStartStep() (*StartStep, error) {
	var w StartStep
	w.ObjectType = "start"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewEndStep - This will create and initialize a new workflow end step object
// and return it as a pointer.
func NewEndStep() (*EndStep, error) {
	var w EndStep
	w.ObjectType = "end"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewActionStep - This will create and initialize a new workflow action step
// object and return it as a pointer.
func NewActionStep() (*ActionStep, error) {
	var w ActionStep
	w.ObjectType = "action"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewPlaybookActionStep - This will create and initialize a new workflow
// playbook action step object and return it as a pointer.
func NewPlaybookActionStep() (*PlaybookActionStep, error) {
	var w PlaybookActionStep
	w.ObjectType = "playbook-action"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewParallelStep - This will create and initialize a new workflow parallel
// step object and return it as a pointer.
func NewParallelStep() (*ParallelStep, error) {
	var w ParallelStep
	w.ObjectType = "parallel"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewIfStep - This will create and initialize a new workflow if condition step
// object and return it as a pointer.
func NewIfStep() (*IfStep, error) {
	var w IfStep
	w.ObjectType = "if-condition"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewWhileStep - This will create and initialize a new workflow if condition
// step object and return it as a pointer.
func NewWhileStep() (*WhileStep, error) {
	var w WhileStep
	w.ObjectType = "while-condition"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}

// NewSwitchStep - This will create and initialize a new workflow switch condition
// step object and return it as a pointer.
func NewSwitchStep() (*SwitchStep, error) {
	var w SwitchStep
	w.ObjectType = "switch-condition"
	err := w.SetNewID(w.ObjectType)
	return &w, err
}
