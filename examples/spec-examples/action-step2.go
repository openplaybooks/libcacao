// Copyright 2021 Bret Jordan, All rights reserved.
//
// Use of this source code is governed by an Apache 2.0 license that can be
// found in the LICENSE file in the root of the source tree.

package main

import (
	"fmt"

	"github.com/openplaybooks/libcacao/objects/playbook"
)

func main() {
	p := playbook.New()
	p.Created = ""
	p.Modified = ""

	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 2"
	step1.Description = "Example of an http-api action step."

	// Create workflow steps for this playbook
	end, _ := p.NewEndStep()
	end.Name = "Playbook End"
	end.Description = "Example of an end step."

	step1.OnSuccess = end.GetID()
	step1.Agent = "individual--328a89ab-3b8f-40c4-a491-24a40bcd3cd4"
	cmd1, _ := step1.NewCommand()
	cmd1.ObjectType = "http-api"
	cmd1.Command = "/v1/blockSystem?id=192.168.0.100"

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
