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
	assignmentcommand(prep(playbook.New()))
	manualcommand(prep(playbook.New()))
	bashcommand(prep(playbook.New()))
	sshcommand(prep(playbook.New()))
	httpcommand1(prep(playbook.New()))
	httpcommand2(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func assignmentcommand(p *playbook.Playbook) {

	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.2.1"
	step1.Description = "Example of an action step with an assignment command."

	// Create command
	cmd1, _ := step1.NewAssignmentCommand()
	cmd1.Description = "Copy various variables to new variables."
	cmd1.AddVariableAssignment("__new_variable_1__", "__existing_variable__")
	cmd1.AddVariableAssignment("__new_variable_2__", "__http--481b0ab0-df88-409e-ba58-7d77d7ffa4af.body__", "to_list", ",")

	header("assignment")
	encode(p)
}

func manualcommand(p *playbook.Playbook) {

	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.3.1"
	step1.Description = "Example of an action step with a manual command for an infected machine."

	// Create command
	cmd1, _ := step1.NewManualCommand()
	cmd1.Command = "Gather information about the infected machine."

	// Create four questions
	cmd1.AddQuestion("What OS is it running?", "string")
	cmd1.AddQuestion("What is the OS version?", "string")
	r3, _ := cmd1.AddQuestion("What is the IP address in ddd.ddd.ddd.ddd format?", "ipv4-addr")
	cmd1.AddQuestion("What is the MAC address in xx:xx:xx:xx:xx:xx format?", "mac-addr")

	// Create command
	cmd2, _ := step1.NewManualCommand()
	cmd2.Command = "Quarantine the machine to vlan __vlan_quarantine_number__:value."

	// Create one question
	cmd2.AddQuestion("What is the new IP address in ddd.ddd.ddd.ddd format?", "ipv4-addr")

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	resid := r3.ID

	header("manual")
	encode(p)
	fmt.Printf("__%s.response__:value\n\n", resid)
}

func bashcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.4.1"
	step1.Description = "Example of an action step with a bash command."
	cmd1, _ := step1.NewBashCommand()
	cmd1.Description = "View failed login attempts."
	cmd1.Command = "cat /var/log/auth.log | grep -i 'failed password'"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("bash")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func sshcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.5.1"
	step1.Description = "Example of an action step with an ssh command."
	cmd1, _ := step1.NewSSHCommand()
	cmd1.Description = "View failed login attempts."
	cmd1.Command = "cat /var/log/auth.log | grep -i 'failed password'"

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("ssh")
	encode(p)
	fmt.Printf("__%s.stdout__:value\n\n", cmdid)
}

func httpcommand1(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.6.1"
	step1.Description = "Example of an action step with an http GET command."
	cmd1, _ := step1.NewHTTPCommand()
	cmd1.Description = "Get current data for an ID"
	cmd1.Command = "GET /v1/getData?id=__some_data_id__:value HTTP/1.1"
	cmd1.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) Chrome/109.0.0.0 Safari/537.36")

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("http 1")
	encode(p)
	fmt.Printf("__%s.body__:value\n\n", cmdid)
}

func httpcommand2(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.6.2"
	step1.Description = "Example of an action step with an http POST command."
	cmd1, _ := step1.NewHTTPCommand()
	cmd1.Description = "Post data to endpoint"
	cmd1.Command = "POST /api1/newObjects/ HTTP/1.1"
	cmd1.AddHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) Chrome/109.0.0.0 Safari/537.36")
	cmd1.ContentB64 = "ewogICJ0eXBlIjogImhvdXNlIiwKICAidmFsdWUiOiAic29tZSBkYXRhIiwKICAuLi4KfQ=="

	// Make a copy of these IDs for the summary view below since the IDs will
	// get cleared out by the encode function
	cmdid := cmd1.ID

	header("http 2")
	encode(p)
	fmt.Printf("__%s.status__:value\n\n", cmdid)
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
