// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"github.com/openplaybooks/libcacao/objects/steps"
)

// ----------------------------------------------------------------------
// Workflow Step Constructors
// ----------------------------------------------------------------------

// AddStep - This method takes in an interface represening a workflow step
// object that satisfies the workflow.StepObject interface and adds it to the
// map.
func (p *Playbook) AddStep(s steps.StepObject) error {
	k := s.GetCommon().ID
	if p.Workflow == nil {
		m := make(map[string]steps.StepObject, 0)
		p.Workflow = m
	}
	p.Workflow[k] = s
	// After we add it to the playbook lets clear out the embedded ID
	//v.ClearID()

	// Make sure we call you the logic features as needed
	if p.PlaybookProcessingSummary == nil {
		var ps ProcessingSummary
		p.PlaybookProcessingSummary = &ps
	}

	switch s.GetCommon().ObjectType {
	case "playbook-action":
		p.PlaybookProcessingSummary.ExternalPlaybooks = true
	case "parallel":
		p.PlaybookProcessingSummary.ParallelProcessing = true
	case "foreach":
		p.PlaybookProcessingSummary.ForeachLogic = true
	case "while":
		p.PlaybookProcessingSummary.WhileLogic = true
	case "if-then":
		p.PlaybookProcessingSummary.IfLogic = true
	case "switch":
		p.PlaybookProcessingSummary.SwitchLogic = true
	}
	return nil
}

// NewStartStep - Create and initialize a new start step object and return it as
// a pointer.
func (p *Playbook) NewStartStep() (*steps.StartStep, error) {
	var s steps.StartStep
	s.ObjectType = "start"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewReturnStep - Create and initialize a new return step object and return it
// as a pointer.
func (p *Playbook) NewReturnStep() (*steps.ReturnStep, error) {
	var s steps.ReturnStep
	s.ObjectType = "return"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewEndStep - Create and initialize a new end step object and return it as a
// pointer.
func (p *Playbook) NewEndStep() (*steps.EndStep, error) {
	var s steps.EndStep
	s.ObjectType = "end"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewActionStep - Create and initialize a new action step object and return it
// as a pointer.
func (p *Playbook) NewActionStep() (*steps.ActionStep, error) {
	var s steps.ActionStep
	s.ObjectType = "action"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewPlaybookActionStep - Create and initialize a new playbook action step
// object and return it as a pointer.
func (p *Playbook) NewPlaybookActionStep() (*steps.PlaybookActionStep, error) {
	var s steps.PlaybookActionStep
	s.ObjectType = "playbook-action"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewParallelStep - Create and initialize a new parallel step object and return
// it as a pointer.
func (p *Playbook) NewParallelStep() (*steps.ParallelStep, error) {
	var s steps.ParallelStep
	s.ObjectType = "parallel"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewForeachStep - Create and initialize a new foreach step object and return
// it as a pointer.
func (p *Playbook) NewForeachStep() (*steps.ForeachStep, error) {
	var s steps.ForeachStep
	s.ObjectType = "foreach"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewWhileStep - Create and initialize a new while step object and return it as
// a pointer.
func (p *Playbook) NewWhileStep() (*steps.WhileStep, error) {
	var s steps.WhileStep
	s.ObjectType = "while"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewIfThenStep - Create and initialize a new if then step object and return it
// as a pointer.
func (p *Playbook) NewIfThenStep() (*steps.IfThenStep, error) {
	var s steps.IfThenStep
	s.ObjectType = "if-then"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}

// NewSwitchStep - Create and initialize a new switch step object and return it
// as a pointer.
func (p *Playbook) NewSwitchStep() (*steps.SwitchStep, error) {
	var s steps.SwitchStep
	s.ObjectType = "switch"
	err := s.SetNewID(s.ObjectType)
	p.AddStep(&s)
	return &s, err
}
