// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects/markings/statement"
	"github.com/openplaybooks/libcacao/objects/markings/tlp"
	"github.com/openplaybooks/libcacao/objects/playbook"
	"github.com/openplaybooks/libcacao/objects/workflow/end"
	"github.com/openplaybooks/libcacao/objects/workflow/parallel"
	"github.com/openplaybooks/libcacao/objects/workflow/single"
	"github.com/openplaybooks/libcacao/objects/workflow/start"
)

func main() {
	p := playbook.New()

	p.Name = "Find Malware FuzzyPanda"
	p.Description = "This playbook will look for FuzzyPanda on the network and in a SIEM"
	p.AddPlaybookTypes("investigation")
	p.CreatedBy = "identity--uuid2"
	p.ValidFrom = p.GetCurrentTime("micro")
	p.ValidUntil = "2021-12-31T23:59:59.999999Z"
	p.AddDerivedFrom("playbook--uuid99")
	p.Priority = 3
	p.Severity = 70
	p.Impact = 5
	p.AddLabels("malware,fuzzypanda,apt")

	r1, _ := p.NewExternalReference()
	r1.Name = "ACME Security FuzzyPanda Report"
	r1.Description = "ACME security review of FuzzyPanda 2020"
	r1.Source = "ACME Security Company, Solutions for FuzzyPanda 2020, January 2020. [Online]. Available: http://www.example.com/info/fuzzypanda2020.html"
	r1.URL = "http://www.example.com/info/fuzzypanda2020.html"
	r1.ExternalID = "fuzzypanda 2020.01"
	r1.ReferenceID = "malware--2008c526-508f-4ad4-a565-b84a4949b2af"

	//
	// Create a statement marking and TLP marking for this playbook
	//
	m1 := statement.New()
	m1.Statement = "Copyright 2021 ACME Security Company"
	m2 := tlp.New()
	m2.SetGreen()
	p.AddMarkings([]string{m1.GetID(), m2.GetID()})
	p.AddMarkingDefinition(m1)
	p.AddMarkingDefinition(m2)

	//
	// Create a workflow for this playbook
	//
	start := start.New()
	p.WorkflowStart = start.GetID()
	p.AddWorkflowStep(start)

	step1 := single.New()
	start.OnSuccess = step1.GetID()
	step1.Name = "Receive IOC"
	step1.Description = "Get Fuzzy Panda Data Exfil Site of 1.2.3.4"
	cmd1, _ := step1.NewCommand()
	cmd1.SetManual()
	cmd1.Command = "Get IOC from threat feed"
	p.AddWorkflowStep(step1)

	step2 := single.New()
	step1.OnSuccess = step2.GetID()
	step2.Name = "Start Investigation"
	step2.Description = "Open ticket for level 1 SOC admin to start investigation"
	cmd2, _ := step2.NewCommand()
	cmd2.SetManual()
	cmd2.Command = "Open ticket in ticketing system"
	p.AddWorkflowStep(step2)

	step3 := parallel.New()
	step2.OnSuccess = step3.GetID()
	step3.Name = "Start Parallel Steps"
	step3.Description = "We can run two steps in parallel"
	p.AddWorkflowStep(step3)

	step4 := single.New()
	step3.AddNextSteps(step4.GetID())
	step4.Name = "Query Network Tools"
	step4.Description = "Look at firewall logs for traffic to IOC"
	cmd4, _ := step4.NewCommand()
	cmd4.SetManual()
	cmd4.Command = "Open firewall console and filter for IOC value"
	p.AddWorkflowStep(step4)

	step5 := single.New()
	step3.AddNextSteps(step5.GetID())
	step5.Name = "Query SIEM for IP Address"
	step5.Description = "Look at SIEM for traffic to IOC"
	cmd5, _ := step5.NewCommand()
	cmd5.SetManual()
	cmd5.Command = "Open SIEM and filter for IOC value"
	p.AddWorkflowStep(step5)

	end := end.New()
	p.AddWorkflowStep(end)

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
