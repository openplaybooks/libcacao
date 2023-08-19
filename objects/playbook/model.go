// Copyright 2023 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

// Package playbook implements the CACAO 2.0 playbook object.
//
// CACAO playbooks are made up of five parts; playbook metadata, the workflow
// logic, a list of agents, a list of extensions, and a list of data markings.
// Playbooks MAY refer to other playbooks in the workflow, similar to how
// programs refer to function calls or modules that comprise the program. The
// definition and normative requirements for all data types listed in the
// property table below and other property tables in this document can be found
// in Section 9.
package playbook

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/agents"
	"github.com/openplaybooks/libcacao/objects/markings"
	"github.com/openplaybooks/libcacao/objects/signature"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// Playbook - This type implements the CACAO 2.0 Playbook object and defines all
// of the properties and methods needed to create and work with this object.
type Playbook struct {
	ObjectType                string                         `json:"type,omitempty"`
	SpecVersion               string                         `json:"spec_version,omitempty"`
	ID                        string                         `json:"id,omitempty"`
	Name                      string                         `json:"name,omitempty"`
	Description               string                         `json:"description,omitempty"`
	PlaybookTypes             []string                       `json:"playbook_types,omitempty"`
	PlaybookActivities        []string                       `json:"playbook_activities,omitempty"`
	PlaybookProcessingSummary *ProcessingSummary             `json:"playbook_processing_summary,omitempty"`
	CreatedBy                 string                         `json:"created_by,omitempty"`
	Created                   string                         `json:"created,omitempty"`
	Modified                  string                         `json:"modified,omitempty"`
	Revoked                   bool                           `json:"revoked,omitempty"`
	ValidFrom                 string                         `json:"valid_from,omitempty"`
	ValidUntil                string                         `json:"valid_until,omitempty"`
	DerivedFrom               []string                       `json:"derived_from,omitempty"`
	RelatedTo                 []string                       `json:"related_to,omitempty"`
	Priority                  int                            `json:"priority,omitempty"`
	Severity                  int                            `json:"severity,omitempty"`
	Impact                    int                            `json:"impact,omitempty"`
	IndustrySectors           []string                       `json:"industry_sectors,omitempty"`
	Labels                    []string                       `json:"labels,omitempty"`
	ExternalReferences        []objects.ExternalReference    `json:"external_references,omitempty"`
	Markings                  []string                       `json:"markings,omitempty"`
	PlaybookVariables         map[string]objects.Variables   `json:"playbook_variables,omitempty"`
	WorkflowStart             string                         `json:"workflow_start,omitempty"`
	WorkflowException         string                         `json:"workflow_exception,omitempty"`
	Workflow                  map[string]workflow.StepObject `json:"workflow,omitempty"`
	//PlaybookExtensions            map[string]????                     `json:"playbook_extensions,omitempty"`
	//AuthenticationInfoDefinitions map[string]????                     `json:"authentication_info_definitions,omitempty"`
	AgentDefinitions  map[string]agents.AgentObject `json:"agent_definitions,omitempty"`
	TargetDefinitions map[string]agents.AgentObject `json:"target_definitions,omitempty"`
	//ExtensionDefinitions          map[string]????                     `json:extension_definitions,omitempty"`
	DataMarkingDefinitions map[string]markings.DataMarkingObject `json:"data_marking_definitions,omitempty"`
	Signatures             []signature.Signature                 `json:"signatures,omitempty"`
}

// ProcessingSummary - This type defines a list of playbook processing features
type ProcessingSummary struct {
	ManualPlaybook     bool `json:"manual_playbook,omitempty"`
	ExternalPlaybooks  bool `json:"external_playbooks,omitempty"`
	ParallelProcessing bool `json:"parallel_processing,omitempty"`
	IfLogic            bool `json:"if_logic,omitempty"`
	WhileLogic         bool `json:"while_logic,omitempty"`
	SwitchLogic        bool `json:"switch_logic,omitempty"`
	TemporalLogic      bool `json:"temporal_logic,omitempty"`
	DataMarkings       bool `json:"data_markings,omitempty"`
	Extensions         bool `json:"extensions,omitempty"`
}

// This type is used to capture results from the Valid() and Compare() functions
type results struct {
	debug         bool
	problemsFound int
	resultDetails []string
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new Playbook object and return it as a
// pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *Playbook {
	p := new(Playbook)
	p.ObjectType = "playbook"
	p.SpecVersion = p.GetCurrentSpecVersion()
	p.SetNewID(p.ObjectType)
	p.Created = p.GetCurrentTime("milli")
	p.Modified = p.Created
	p.Revoked = false
	return p
}
