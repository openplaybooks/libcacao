// Copyright 2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Workflow Step Constructors
// ----------------------------------------------------------------------

// NewStartStep - Create and initialize a new start step object and return it as
// a pointer.
func (p *Playbook) NewStartStep() (*workflow.StartStep, error) {
	var s workflow.StartStep
	s.ObjectType = "start"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewReturnStep - Create and initialize a new return step object and return it
// as a pointer.
func (p *Playbook) NewReturnStep() (*workflow.ReturnStep, error) {
	var s workflow.ReturnStep
	s.ObjectType = "return"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewEndStep - Create and initialize a new end step object and return it as a
// pointer.
func (p *Playbook) NewEndStep() (*workflow.EndStep, error) {
	var s workflow.EndStep
	s.ObjectType = "end"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewActionStep - Create and initialize a new action step object and return it
// as a pointer.
func (p *Playbook) NewActionStep() (*workflow.ActionStep, error) {
	var s workflow.ActionStep
	s.ObjectType = "action"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewPlaybookActionStep - Create and initialize a new playbook action step
// object and return it as a pointer.
func (p *Playbook) NewPlaybookActionStep() (*workflow.PlaybookActionStep, error) {
	var s workflow.PlaybookActionStep
	s.ObjectType = "playbook-action"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewParallelStep - Create and initialize a new parallel step object and return
// it as a pointer.
func (p *Playbook) NewParallelStep() (*workflow.ParallelStep, error) {
	var s workflow.ParallelStep
	s.ObjectType = "parallel"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewForeachStep - Create and initialize a new foreach step object and return
// it as a pointer.
func (p *Playbook) NewForeachStep() (*workflow.ForeachStep, error) {
	var s workflow.ForeachStep
	s.ObjectType = "foreach"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewWhileStep - Create and initialize a new while step object and return it as
// a pointer.
func (p *Playbook) NewWhileStep() (*workflow.WhileStep, error) {
	var s workflow.WhileStep
	s.ObjectType = "while"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewIfThenStep - Create and initialize a new if then step object and return it
// as a pointer.
func (p *Playbook) NewIfThenStep() (*workflow.IfThenStep, error) {
	var s workflow.IfThenStep
	s.ObjectType = "if-then"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}

// NewSwitchStep - Create and initialize a new switch step object and return it
// as a pointer.
func (p *Playbook) NewSwitchStep() (*workflow.SwitchStep, error) {
	var s workflow.SwitchStep
	s.ObjectType = "switch"
	err := s.SetNewID(s.ObjectType)
	p.AddWorkflowStep(&s)
	return &s, err
}
