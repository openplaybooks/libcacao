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
	sw1, _ := p.NewSwitchStep()
	sw1.Name = "Switch Step 1"
	sw1.Description = "Example of a switch step."
	sw1.Switch = "__variable__:value"
	case1, _ := p.NewActionStep()
	case1.Name = "Action Step 1"
	case2, _ := p.NewActionStep()
	case2.Name = "Action Step 2"

	sw1.AddCase("192.168.0.11", case1.GetID())
	sw1.AddCase("192.168.0.12", case2.GetID())

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
