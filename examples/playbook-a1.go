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

	p.Name = "Prevent FuzzyPanda Malware"
	p.Description = "This playbook will block traffic to the FuzzyPanda data exfil site"
	p.AddPlaybookTypes("prevention")
	p.CreatedBy = "identity--uuid2"
	p.ValidFrom = p.GetCurrentTime("micro")
	p.ValidUntil = "2021-12-31T23:59:59.999999Z"
	p.AddDerivedFrom("playbook--uuid50")
	p.Priority = 3
	p.Severity = 70
	p.Impact = 5
	p.AddLabels("malware,fuzzypanda,apt")

	r1, _ := p.NewExternalReference()
	r1.Name = "ACME Security FuzzyPanda Report"
	r1.Description = "ACME security review of FuzzyPanda 2021"
	r1.Source = "ACME Security Company, Solutions for FuzzyPanda 2021, January 2021. Available online: http://www.example.com/info/fuzzypanda2021.html"
	r1.URL = "http://www.example.com/info/fuzzypanda2020.html"
	r1.ExternalID = "fuzzypanda 2021.01"
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
	// Create workflow steps for this playbook
	//
	start := start.New()
	end := end.New()

	step1 := single.New()
	step1.Name = "Receive IOC"
	step1.Description = "Get FuzzyPanda Data Exfil Site IP Address of 1.2.3.4"
	cmd1, _ := step1.NewCommand()
	cmd1.SetManual()
	cmd1.Command = "Get IOC from threat feed"

	step2 := parallel.New()
	step2.Name = "Update Protection Tools"
	step2.Description = "This step will update the firewall and client EDR in parallel"

	step3 := single.New()
	step3.Name = "Add IP to Firewall Blocklist"
	step3.Description = "This step will add the IP address of the FuzzyPanda data exfil site to the firewall"
	cmd3, _ := step3.NewCommand()
	cmd3.SetManual()
	cmd3.Command = "Open firewall console and add 1.2.3.4 to the firewall blocking policy"

	step4 := single.New()
	step4.Name = "Add IP to Client EDR Blocklist"
	step4.Description = "This step will add the IP address of the FuzzyPanda data exfil site to the client EDR solution"
	cmd4, _ := step4.NewCommand()
	cmd4.SetManual()
	cmd4.Command = "Open EDR console and add 1.2.3.4 to the blocking policy"

	step5 := single.New()
	step5.Name = "Create Ticket"
	step5.Description = "This step will create a ticket for this issue"
	cmd5, _ := step5.NewCommand()
	cmd5.SetManual()
	cmd5.Command = "Open case management tool and create a ticket with the details of what was done"

	step6 := single.New()
	step6.Name = "Update SIEM"
	step6.Description = "This step will update the SIEM to look for traffic attempts to the FuzzyPanda data exfil site"
	cmd6, _ := step6.NewCommand()
	cmd6.SetManual()
	cmd6.Command = "Open SIEM solution and add rule to look for 1.2.3.4"

	// Link all of the steps together
	p.WorkflowStart = start.GetID()
	start.OnCompletion = step1.GetID()
	step1.OnCompletion = step2.GetID()
	step2.AddNextSteps(step3.GetID())
	step2.AddNextSteps(step4.GetID())
	step3.OnCompletion = step5.GetID()
	step4.OnCompletion = step5.GetID()
	step5.OnCompletion = step6.GetID()
	step6.OnCompletion = end.GetID()

	// Add workflow to the playbook
	p.AddWorkflowStep(start)
	p.AddWorkflowStep(step1)
	p.AddWorkflowStep(step2)
	p.AddWorkflowStep(step3)
	p.AddWorkflowStep(step4)
	p.AddWorkflowStep(step5)
	p.AddWorkflowStep(step6)
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
