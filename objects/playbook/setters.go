// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"github.com/jordan2175/libcacao/objects"
	"github.com/jordan2175/libcacao/objects/markings"
	"github.com/jordan2175/libcacao/objects/workflow"
)

// SetNewID - This method takes in a string value representing an object type and
// creates a new ID based on the specification format and update the id property
// for the object.
func (p *Playbook) SetNewID(s string) error {
	// TODO Add check to validate input value
	p.ID, _ = objects.CreateID(s)
	return nil
}

// GetCurrentSpecVersion - This method returns the current specification version
// that this library is using.
func (p *Playbook) GetCurrentSpecVersion() string {
	return objects.GetCurrentSpecVersion()
}

// GetCurrentTime - This method takes in a value of either milli or micro and
// returns the current time in RFC 3339 format
func (p *Playbook) GetCurrentTime(precision string) string {
	return objects.GetCurrentTime(precision)
}

// SetModified - This method takes in a timestamp in either time.Time or string
// format and updates the modified property with it. The value is stored as a
// string, so if the value is in time.Time format, it will be converted to the
// correct timestamp format.
func (p *Playbook) SetModified(t interface{}) error {
	ts, _ := objects.TimeToString(t, "milli")
	p.Modified = ts
	return nil
}

// AddPlaybookTypes - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// playbook type and adds it to the playbook_types property.
func (p *Playbook) AddPlaybookTypes(values interface{}) error {
	return objects.AddValuesToList(&p.PlaybookTypes, values)
}

// AddDerivedFrom - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// playbook uuid and adds it to the derived_from property.
func (p *Playbook) AddDerivedFrom(values interface{}) error {
	return objects.AddValuesToList(&p.DerivedFrom, values)
}

// AddLabels - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// label and adds it to the labels property.
func (p *Playbook) AddLabels(values interface{}) error {
	return objects.AddValuesToList(&p.Labels, values)
}

// AddMarkings - This method takes in a string value, a comma separated list of
// string values, or a slice of string values that all representing a
// data marking and adds it to the markings property.
func (p *Playbook) AddMarkings(values interface{}) error {
	return objects.AddValuesToList(&p.Markings, values)
}

// NewExternalReference - This method creates a new empty external reference and
// returns a reference to it so it can be populated. However, if one or more
// external references are passed in they are all added and the reference that
// is returned is for the last entry added.
func (p *Playbook) NewExternalReference(r ...objects.ExternalReference) (*objects.ExternalReference, error) {
	positionThatAppendWillUse := len(p.ExternalReferences)

	if len(r) > 0 {
		for i := range r {
			// Update the value so we grab the last one entered
			positionThatAppendWillUse = len(p.ExternalReferences)
			p.ExternalReferences = append(p.ExternalReferences, r[i])
		}
		return &p.ExternalReferences[positionThatAppendWillUse], nil
	}

	// If one was not passed in, lets create one
	var er objects.ExternalReference
	p.ExternalReferences = append(p.ExternalReferences, er)
	return &p.ExternalReferences[positionThatAppendWillUse], nil
}

// AddMarkingDefinition - This method takes in an interface represening a
// marking definition object that satisfies the markings.DataMarkingObject
// interface and adds it to the map.
func (p *Playbook) AddMarkingDefinition(v markings.DataMarkingObject) error {
	k := v.GetCommon().ID
	if p.DataMarkingDefinitions == nil {
		m := make(map[string]markings.DataMarkingObject, 0)
		p.DataMarkingDefinitions = m
	}
	p.DataMarkingDefinitions[k] = v

	// Since we are using datamarkings, make sure the features property captures that
	if p.Features == nil {
		var f Features
		p.Features = &f
	}
	p.Features.DataMarkings = true
	return nil
}

// AddWorkflowStep - This method takes in an interface represening a workflow
// step object that satisfies the workflow.StepObject interface and adds it to
// the map.
func (p *Playbook) AddWorkflowStep(v workflow.StepObject) error {
	k := v.GetCommon().ID
	if p.Workflow == nil {
		m := make(map[string]workflow.StepObject, 0)
		p.Workflow = m
	}
	p.Workflow[k] = v

	// Make sure we call you the logic features as needed
	if p.Features == nil {
		var f Features
		p.Features = &f
	}

	switch v.GetCommon().ObjectType {
	case "parallel":
		p.Features.ParallelProcessing = true
	case "if-condition":
		p.Features.IfLogic = true
	}

	return nil
}
