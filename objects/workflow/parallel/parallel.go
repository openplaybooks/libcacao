// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package parallel implements the CACAO 2.0 workflow parallel step object.
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
package parallel

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// ParallelStep - This type implmenets the CACAO 2.0 workflow parallel
// step and defines all of the properties associated with the parallel step.
// Some properties are inherited from the workflow.CommonProperties type.
type ParallelStep struct {
	workflow.CommonProperties
	NextSteps []string `json:"next_steps,omitempty"`
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new workflow parallel step object and
// return it as a pointer. It will also initialize the object by setting all of
// the basic properties.
func New() *ParallelStep {
	var w ParallelStep
	w.Init()
	return &w
}

// Init - This method will initialize a new workflow parallel step object with
// the correct defaults.
func (w *ParallelStep) Init() {
	w.ObjectType = "parallel"
	w.SetNewID(w.ObjectType)
}

// ----------------------------------------------------------------------
// Define Functions and Methods
// ----------------------------------------------------------------------

// GetCommon - This method returns the common step properties
func (w *ParallelStep) GetCommon() workflow.CommonProperties {
	return w.CommonProperties
}

// AddNextSteps - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// next steps in the workflow and adds them to the next_steps property.
func (w *ParallelStep) AddNextSteps(values interface{}) error {
	return objects.AddValuesToList(&w.NextSteps, values)
}
