// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects"
	"github.com/openplaybooks/libcacao/objects/markings"
	"github.com/openplaybooks/libcacao/objects/playbook"
	"github.com/openplaybooks/libcacao/objects/workflow"
)

func main() {
	p := playbook.New()

	p.ID = "playbook--61a6c41e-6efc-4516-a242-dfbc5c89d562"
	p.Name = "Find Malware FuzzyPanda"
	p.Description = "This playbook will look for FuzzyPanda on the network and in a SIEM"
	p.AddPlaybookTypes("investigation")
	p.AddPlaybookActivities("analyze-collected-data,identify-indicators")
	p.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	p.Created = "2023-02-19T08:00:24.918Z"
	p.Modified = "2023-02-19T08:00:24.918Z"
	p.ValidFrom = "2023-02-19T08:00:24.918Z"
	p.ValidUntil = "2023-12-31T23:59:59.999Z"
	p.AddDerivedFrom("playbook--00ee41a2-c2ca-41da-8ea9-681344eb3926")
	p.Priority = 3
	p.Severity = 70
	p.Impact = 5
	p.AddIndustrySectors("aerospace,defense")
	p.AddLabels("malware,fuzzypanda,apt")

	r1, _ := p.NewExternalReference()
	r1.Name = "ACME Security FuzzyPanda Report"
	r1.Description = "ACME security review of FuzzyPanda 2021"
	r1.Source = "ACME Security Company, Solutions for FuzzyPanda 2021, January 2021. Available online: hxxp://www[.]example[.]com/info/fuzzypanda2021.html"
	r1.URL = "hxxp://www[.]example[.]com/info/fuzzypanda2021.html"
	r1.ExternalID = "fuzzypanda 2023.01"
	r1.ReferenceID = "malware--2008c526-508f-4ad4-a565-b84a4949b2af"

	//
	// Create a statement marking and TLP marking for this playbook
	//
	m1 := markings.NewStatementMarking()
	m1.ID = "marking-statement--6424867b-0440-4885-bd0b-604d51786d06"
	m1.Statement = "Copyright 2023 ACME Security Company"
	m1.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m1.Created = "2023-02-19T08:00:24.918Z"
	m2 := markings.NewTLPGreenMarking()
	p.AddMarkings([]string{m1.GetID(), m2.GetID()})
	p.AddMarkingDefinition(m1)
	p.AddMarkingDefinition(m2)

	v1 := objects.NewVariable()
	v1.ObjectType = "ipv4-addr"
	v1.Name = "__data_exfil_site__"
	v1.Description = "The IP address for the data exfiltration site"
	v1.Value = "1.2.3.4"
	v1.Constant = false
	v1.External = false
	p.AddVariable(*v1)

	//
	// Create workflow steps for this playbook
	//
	start, _ := workflow.NewStartStep()
	start.ID = "start--07bea005-4a36-4a77-bd1f-79a6e4682a13"
	start.Name = "Start Playbook Example 1"
	end, _ := workflow.NewEndStep()
	end.ID = "end--6b23c237-ade8-4d00-9aa1-75999738d557"
	end.Name = "End Playbook Example 1"

	step1, _ := workflow.NewActionStep()
	step1.ID = "action--7f40f9d7-de39-4027-ab97-15035beff2ff"
	step1.Name = "IP Lookup"
	step1.Description = "Lookup the IP address in the SIEM"
	cmd1, _ := step1.NewCommand()
	cmd1.SetManual()
	cmd1.Command = "Look up IP __data_exfil_site__:value in SIEM"
	cmd1.PlaybookActivity = "identify-indicators"

	// Link all of the steps together
	p.WorkflowStart = start.GetID()
	p.WorkflowException = " ... "
	start.OnCompletion = step1.GetID()
	step1.OnCompletion = end.GetID()

	// Add workflow to the playbook
	p.AddWorkflowStep(start)
	p.AddWorkflowStep(step1)
	p.AddWorkflowStep(end)

	// Remove all of the IDs from the workflow steps since the specification only has them at the map level
	p.ClearWorkflowStepIDs()

	// valid, count, details := p.Valid()

	// fmt.Println("Object valid: ", valid)
	// fmt.Println("Error Cound: ", count)

	// for _, result := range details {
	// 	fmt.Println(result)
	// }

	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
