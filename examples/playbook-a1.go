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
	"github.com/openplaybooks/libcacao/objects/workflow/ifcondition"
	"github.com/openplaybooks/libcacao/objects/workflow/parallel"
	"github.com/openplaybooks/libcacao/objects/workflow/single"
	"github.com/openplaybooks/libcacao/objects/workflow/start"
	"github.com/openplaybooks/libcacao/objects/workflow/switchcondition"
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

	end := end.New()
	p.AddWorkflowStep(end)

	step1 := single.New()
	start.OnCompletion = step1.GetID()
	step1.Name = "Receive IOC"
	step1.Description = "Get Fuzzy Panda Data Exfil Site of 1.2.3.4"
	cmd1, _ := step1.NewCommand()
	cmd1.SetManual()
	cmd1.Command = "Get IOC from threat feed"
	p.AddWorkflowStep(step1)

	step2 := single.New()
	step1.OnCompletion = step2.GetID()
	step2.Name = "Start Investigation"
	step2.Description = "Open ticket for level 1 SOC admin to start investigation"
	cmd2, _ := step2.NewCommand()
	cmd2.SetManual()
	cmd2.Command = "Open ticket in ticketing system"
	p.AddWorkflowStep(step2)

	step3 := parallel.New()
	step2.OnCompletion = step3.GetID()
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

	step41 := ifcondition.New()
	step4.OnCompletion = step41.GetID()
	step41.Name = "Found 1"
	step41.Description = "Was traffic found in the firewall matching the IP from the IOC"
	step41.Condition = "Was traffic found in the firewall matching the IP from IOC?"
	step41.AddOnFalse(end)
	p.AddWorkflowStep(step41)

	step5 := single.New()
	step3.AddNextSteps(step5.GetID())
	step5.Name = "Query SIEM for IP Address"
	step5.Description = "Look at SIEM for traffic to IOC"
	cmd5, _ := step5.NewCommand()
	cmd5.SetManual()
	cmd5.Command = "Open SIEM and filter for IOC value"
	p.AddWorkflowStep(step5)

	step51 := ifcondition.New()
	step5.OnCompletion = step51.GetID()
	step51.Name = "Found 2"
	step51.Description = "Was traffic found in the SIEM matching the IP from the IOC"
	step51.Condition = "Was traffic found in the SIEM matching the IP from IOC?"
	step51.AddOnFalse(end)
	p.AddWorkflowStep(step51)

	step6 := switchcondition.New()
	step41.AddOnTrue(step6)
	step51.AddOnTrue(step6)
	step6.Name = "Number Found"
	step6.Description = "Depending on the number of entries found, do something different"
	step6.Switch = "Number of sessions found"
	p.AddWorkflowStep(step6)

	step71 := single.New()
	step6.AddCase("less than 3", step71.GetID())
	step71.Name = "Create Case Priority 3"
	step71.Description = "If the number of entries found is less than 3, create a priority 3 case"
	cmd71, _ := step71.NewCommand()
	cmd71.SetManual()
	cmd71.Command = "Open support ticket case with a priority level of 3"
	p.AddWorkflowStep(step71)

	step72 := single.New()
	step6.AddCase("greater than or equal to 3 and less than 10", step72.GetID())
	step72.Name = "Create Case Priority 2"
	step72.Description = "If the number of entries found is greater than or equal to 3 and less than 10, create a priority 2 case"
	cmd72, _ := step72.NewCommand()
	cmd72.SetManual()
	cmd72.Command = "Open support ticket case with a priority level of 2"
	p.AddWorkflowStep(step72)

	step73 := single.New()
	step6.AddCase("greater than or equal to 10", step73.GetID())
	step73.Name = "Create Case Priority 1"
	step73.Description = "If the number of entries found is greater than or equal to 10, create a priority 1 case"
	cmd73, _ := step73.NewCommand()
	cmd73.SetManual()
	cmd73.Command = "Open support ticket case with a priority level of 3"
	p.AddWorkflowStep(step73)

	step8 := single.New()
	step71.OnCompletion = step8.GetID()
	step72.OnCompletion = step8.GetID()
	step73.OnCompletion = step8.GetID()
	step8.Name = "Create Recommendations"
	step8.Description = "Create a series of recommendations based on the findings of this investigation"
	cmd8, _ := step8.NewCommand()
	cmd8.SetManual()
	cmd8.Command = "Create knowledge base article on recommendations based on findings from this investigation"
	step8.OnCompletion = end.GetID()
	p.AddWorkflowStep(step8)

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
