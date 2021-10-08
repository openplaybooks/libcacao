// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package playbook

import (
	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/markings"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

// TODO
// finish adding targets and extensionDefinitions for the playbook object

// ----------------------------------------------------------------------
// Define Object Model
// ----------------------------------------------------------------------

// Playbook - This type implements the CACAO Playbook object and defines all of
// the properties and methods needed to create and work with this object.
type Playbook struct {
	ObjectType         string                         `json:"type,omitempty"`
	SpecVersion        string                         `json:"spec_version,omitempty"`
	ID                 string                         `json:"id,omitempty"`
	Name               string                         `json:"name,omitempty"`
	Description        string                         `json:"description,omitempty"`
	PlaybookTypes      []string                       `json:"playbook_types,omitempty"`
	CreatedBy          string                         `json:"created_by,omitempty"`
	Created            string                         `json:"created,omitempty"`
	Modified           string                         `json:"modified,omitempty"`
	Revoked            bool                           `json:"revoked,omitempty"`
	ValidFrom          string                         `json:"valid_from,omitempty"`
	ValidUntil         string                         `json:"valid_until,omitempty"`
	DerivedFrom        []string                       `json:"derived_from,omitempty"`
	Priority           int                            `json:"priority,omitempty"`
	Severity           int                            `json:"severity,omitempty"`
	Impact             int                            `json:"impact,omitempty"`
	IndustrySectors    []string                       `json:"industry_sectors,omitempty"`
	Labels             []string                       `json:"labels,omitempty"`
	ExternalReferences []objects.ExternalReference    `json:"external_references,omitempty"`
	Features           *Features                      `json:"features,omitempty"`
	Markings           []string                       `json:"markings,omitempty"`
	PlaybookVariables  map[string]objects.Variables   `json:"playbook_variables,omitempty"`
	WorkflowStart      string                         `json:"workflow_start,omitempty"`
	WorkflowException  string                         `json:"workflow_exception,omitempty"`
	Workflow           map[string]workflow.StepObject `json:"workflow,omitempty"`
	//Targets              map[string]XXXXX `json:"targets,omitempty"`
	//ExtensionDefinitions map[string]XXXXX `json:extension_definitions,omitempty"`
	DataMarkingDefinitions map[string]markings.DataMarkingObject `json:"data_marking_definitions,omitempty"`
	Signatures             []Signature                           `json:"signatures,omitempty"`
}

// Features - This type defines a list of playbook features
type Features struct {
	ParallelProcessing bool `json:"parallel_processing,omitempty"`
	IfLogic            bool `json:"if_logic,omitempty"`
	WhileLogic         bool `json:"while_logic,omitempty"`
	SwitchLogic        bool `json:"switch_logic,omitempty"`
	TemporalLogic      bool `json:"temporal_logic,omitempty"`
	DataMarkings       bool `json:"data_markings,omitempty"`
	Extensions         bool `json:"extensions,omitempty"`
}

// Signature - This type captures the actual digital signature of an object
type Signature struct {
	//Header    SignatureHeader `json:"header,omitempty"`
	//Excluded  []string `json:"excluded,omitempty"`
	ObjectType     string   `json:"type,omitempty"`
	SpecVersion    string   `json:"spec_version,omitempty"`
	ID             string   `json:"id,omitempty"`
	CreatedBy      string   `json:"created_by,omitempty"`
	Created        string   `json:"created,omitempty"`
	Modified       string   `json:"modified,omitempty"`
	Revoked        bool     `json:"revoked,omitempty"`
	Signee         string   `json:"signee,omitempty"`
	ValidFrom      string   `json:"valid_from,omitempty"`
	ValidUntil     string   `json:"valid_until,omitempty"`
	RelatedTo      string   `json:"related_to,omitempty"`
	RelatedVersion string   `json:"related_version,omitempty"`
	SHA256         string   `json:"sha256,omitempty"`
	Algorithm      string   `json:"algorithm,omitempty"`
	PublicKeys     []string `json:"public_keys,omitempty"`
	Value          string   `json:"value,omitempty"`
}

// This type is used to capture results from the Valid() and Compare() functions
type results struct {
	debug         bool
	problemsFound int
	resultDetails []string
}

// Valid - Implement the Claims class needed for JWT
func (s Signature) Valid() error {
	return nil
}

// ----------------------------------------------------------------------
// Initialization Functions
// ----------------------------------------------------------------------

// New - This function will create a new Playbook object and return it as a
// pointer. It will also initialize the object by setting all of the basic
// properties.
func New() *Playbook {
	p := new(Playbook)
	p.Init()
	return p
}

// Init - This method will initialize a new playbook object with the correct
// defaults.
func (p *Playbook) Init() {
	p.ObjectType = "playbook"
	p.SpecVersion = p.GetCurrentSpecVersion()
	p.SetNewID(p.ObjectType)
	p.Created = p.GetCurrentTime("milli")
	p.Modified = p.Created
	p.Revoked = false
}
