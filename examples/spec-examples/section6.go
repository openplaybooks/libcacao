// Copyright 2019-2025 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects/playbook"
)

func main() {
	playbook1(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func playbook1(p *playbook.Playbook) {

	p.ID = "playbook--61a6c41e-6efc-4516-a242-dfbc5c89d562"
	p.Name = "Find Malware FuzzyPanda"
	p.Description = "This playbook will look for FuzzyPanda on the network and in a SIEM"
	p.AddPlaybookTypes("investigation")
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

	p.CanvasHeight = 500
	p.CanvasWidth = 1000

	r1, _ := p.NewExternalReference()
	r1.Name = "ACME Security FuzzyPanda Report"
	r1.Description = "ACME security review of FuzzyPanda 2021"
	r1.Source = "ACME Security Company, Solutions for FuzzyPanda 2021, January 2021. Available online: hxxp://www[.]example[.]com/info/fuzzypanda2021.html"
	r1.URL = "hxxp://www[.]example[.]com/info/fuzzypanda2021.html"
	r1.ExternalID = "fuzzypanda 2023.01"
	r1.ReferenceID = "malware--2008c526-508f-4ad4-a565-b84a4949b2af"

	p.NewTLPMarking("clear")
	m1, _ := p.NewStatementMarking("Copyright 2023 ACME Security Company")
	m1.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	m1.Created = "2023-02-19T08:00:24.918Z"

	v1, _ := p.NewVariable("__data_exfil_site__", "ipv4-addr", "1.2.3.4")
	v1.Description = "The IP address for the data exfiltration site"
	v1.Constant = false
	v1.External = false

	start, _ := p.NewStartStep()
	step1, _ := p.NewActionStep()
	end, _ := p.NewEndStep()

	start.Name = "Start Playbook Example 1"
	start.Description = "Example of a start step."
	start.OnSuccess = step1.GetID()

	step1.Name = "IP Lookup"
	step1.Description = "Lookup the IP address in the SIEM"
	cmd1, _ := step1.NewManualCommand()
	cmd1.Command = "Look up IP __data_exfil_site__:value in SIEM"
	step1.OnSuccess = end.GetID()

	end.Name = "Playbook End"
	end.Description = "Example of an end step."

	header("playbook1")
	encode(p)

}

// ----------------------------------------------------------------------
// Define Supporting Functions and Methods
// ----------------------------------------------------------------------

func header(t string) {
	fmt.Println("\n// ----------------------------------------")
	fmt.Printf("// %s command example", t)
	fmt.Println("\n// ----------------------------------------")
}

func encode(p *playbook.Playbook) {
	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
