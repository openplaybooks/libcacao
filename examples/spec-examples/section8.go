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
	manualcommand(prep(playbook.New()))
	bashcommand(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func manualcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.2.1"
	step1.Description = "Example of an action step with a manual command for an infected machine."
	cmd1, _ := step1.NewManualCommand()
	cmd1.Command = "Gather information about the infected machine."
	cmd1.AddQuestion("What OS is it running?", "string")
	cmd1.AddQuestion("What is the OS version?", "string")
	r1, _ := cmd1.AddQuestion("What is the IP address?", "string")
	cmd1.AddQuestion("What is the MAC address?", "string")

	cmd2, _ := step1.NewManualCommand()
	cmd2.Command = "Disconnect the machine from the network."
	cmd2.AddQuestion("Was the computer removed from the network?", "bool")

	fmt.Printf("\n%s.%s.%s:value\n\n", step1.ID, cmd1.ID, r1.ID)
	fmt.Printf("%s:value\n\n", r1.ID)
	encode(p)
}

func bashcommand(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 8.3.1"
	step1.Description = "Example of an action step with a bash command."
	cmd1, _ := step1.NewBashCommand()
	cmd1.Description = "View failed login attempts."
	cmd1.Command = "cat /var/log/auth.log | grep -i 'failed password'"

	fmt.Printf("\n%s.%s.stdout:value\n\n", step1.ID, cmd1.ID)
	encode(p)
}

//  cat /var/log/auth.log | grep -i 'failed password'

func encode(p *playbook.Playbook) {
	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
