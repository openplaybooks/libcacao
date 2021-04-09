// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package parallel implements the CACAO 1.0 workflow parallel step object.
//
// This section defines how to create steps that can be processed in parallel.
// In addition to the inherited properties, this section defines one additional
// specific property that is valid for this type. A parallel step MUST execute
// all workflow steps that are part of the next_steps property before this step
// can be considered complete and the workflow logic moves on.
//
// The Parallel Step is a playbook step that allows playbook authors to define
// two or more steps that can be executed at the same time. For example, a
// playbook that responds to an incident may require both the network team and
// the desktop team to investigate and respond to a threat at the same time.
// Another example is in response to a cyber attack on an operational technology
// (OT) environment that would require releasing air / steam / water pressure
// simultaneously.
//
// The steps referenced from this object are intended to be processed in
// parallel, however, if an implementation can not support executing steps in
// parallel, then the steps MAY be executed in sequential order if the desired
// outcome stays the same.
package parallel

import (
	"github.com/jordan2175/libcacao/objects"
	"github.com/jordan2175/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// WorkflowParallelStep - This type implmenets the CACAO 1.0 workflow parallel
// step and defines all of the properties associated with the parallel step.
// Some properties are inherited from the workflow.CommonProperties type.
type WorkflowParallelStep struct {
	workflow.CommonProperties
	NextSteps []string `json:"next_steps,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow parallel step object and
// return it as a pointer. It will also initialize the object by setting all of
// the basic properties.
func New() *WorkflowParallelStep {
	var w WorkflowParallelStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow parallel step object with
// the correct defaults.
func (w *WorkflowParallelStep) Init() {
	w.ObjectType = "parallel"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *WorkflowParallelStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}

// AddNextSteps - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// next steps in the workflow and adds them to the next_steps property.
func (w *WorkflowParallelStep) AddNextSteps(values interface{}) error {
	return objects.AddValuesToList(&w.NextSteps, values)
}
