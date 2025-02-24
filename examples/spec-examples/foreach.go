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
	foreach, _ := p.NewForeachStep()
	foreach.Name = "Foreach Step 1"
	foreach.Description = "Example of a foreach step."
	foreach.Collection = "__list_of_ip_addresses__"
	foreach.Element = "__ip_address__"

	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 1"
	step2, _ := p.NewActionStep()
	step2.Name = "Action Step 2"

	ret1, _ := p.NewReturnStep()
	ret1.Name = "Return To Foreach Step 1"
	ret1.Description = "Example of a return step from branch 1."
	step1.OnSuccess = ret1.GetID()

	foreach.Do = step1.GetID()
	foreach.OnSuccess = step2.GetID()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
