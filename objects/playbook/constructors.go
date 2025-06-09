// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"errors"
	"fmt"
	"strings"

	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/markings"
	"github.com/openplaybooks/libcacao/objects/steps"
	"github.com/openplaybooks/libcacao/objects/variables"
)

// ----------------------------------------------------------------------
// Variable Constructor
// ----------------------------------------------------------------------

// AddVariable - This method takes in an interface represening a variable object
// that satisfies the variables.VariableObject interface and adds it to the
// map.
func (p *Playbook) AddVariable(v variables.VariableObject) error {
	if !objects.IsVocabValueValid(v.GetCommon().ObjectType, variables.GetVariableTypesVocab()) {
		return fmt.Errorf("the variable type %s is not valid", v.GetCommon().ObjectType)
	}

	if p.PlaybookVariables == nil {
		m := make(map[string]variables.VariableObject, 0)
		p.PlaybookVariables = m
	}
	k := v.GetCommon().Name
	p.PlaybookVariables[k] = v
	return nil
}

// NewStringVariable - Create and initialize a new string variable with a name
// (n) and a type (t) and return it as a pointer.
func (p *Playbook) NewStringVariable(n string, t string) (*variables.StringVariable, error) {
	var v variables.StringVariable
	v.ObjectType = t
	v.Name = n
	err := p.AddVariable(&v)
	return &v, err
}

// ----------------------------------------------------------------------
// Workflow Step Constructors
// ----------------------------------------------------------------------

// AddStep - This method takes in an interface represening a workflow step
// object that satisfies the workflow.StepObject interface and adds it to the
// map.
func (p *Playbook) AddStep(s steps.StepObject) error {
	if p.Workflow == nil {
		m := make(map[string]steps.StepObject, 0)
		p.Workflow = m
	}
	k := s.GetCommon().ID
	p.Workflow[k] = s

	// Identify any logic features that are in use and flag them
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
	var err error
	s.ObjectType = "action"
	err = s.SetNewID(s.ObjectType)
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

// ----------------------------------------------------------------------
// Playbook Marking Constructors
// ----------------------------------------------------------------------

// NewTLPMarking - This method will take in (s) as one of the following marking
// types (i.e., clear, green, amber, amber+strict, red) and assign and assign
// that marking to the playbook.
func (p *Playbook) NewTLPMarking(s string) error {

	s = strings.ToLower(s)
	switch s {
	case "clear":
		p.AddMarkingDefinition(markings.NewTLPClearMarking())
		return nil
	case "green":
		p.AddMarkingDefinition(markings.NewTLPGreenMarking())
		return nil
	case "amber":
		p.AddMarkingDefinition(markings.NewTLPAmberMarking())
		return nil
	case "amber+strict":
		p.AddMarkingDefinition(markings.NewTLPAmberStrictMarking())
		return nil
	case "red":
		p.AddMarkingDefinition(markings.NewTLPRedMarking())
		return nil
	}
	return errors.New("not a valid TLP type")
}

// NewStatementMarking - This method will take in a statement (s) and assign
// that as a marking to the playbook. It will also return a pointer to the
// marking for further changes.
func (p *Playbook) NewStatementMarking(s string) (*markings.MarkingStatement, error) {
	m := markings.NewStatementMarking()
	m.Statement = s
	p.AddMarkingDefinition(m)
	return m, nil
}

// AddMarkingDefinition - This method takes in an interface (m) represening a
// marking definition object that satisfies the markings.DataMarkingObject
// interface and adds it to the map.
func (p *Playbook) AddMarkingDefinition(m markings.DataMarkingObject) error {
	k := m.GetCommon().ID
	if p.DataMarkingDefinitions == nil {
		dm := make(map[string]markings.DataMarkingObject, 0)
		p.DataMarkingDefinitions = dm
	}
	p.DataMarkingDefinitions[k] = m
	p.AddMarkings(k)
	return nil
}

// AddMarkings - This method takes in a string value, a comma separated list of
// string values, or a slice of string values (s) that all representing a
// data marking and adds it to the markings property.
func (p *Playbook) AddMarkings(s interface{}) error {
	// Since we are applying a data marking to this playbook, we need to capture
	// that in the features property
	if p.PlaybookProcessingSummary == nil {
		var ps ProcessingSummary
		p.PlaybookProcessingSummary = &ps
	}
	p.PlaybookProcessingSummary.DataMarkings = true

	return objects.AddValuesToList(&p.Markings, s)
}
