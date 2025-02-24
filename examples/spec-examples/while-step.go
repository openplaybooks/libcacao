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
	while, _ := p.NewWhileStep()
	while.Name = "While Step 1"
	while.Description = "Example of a while step."
	while.Condition = "__variable__:value = '10.0.0.0/8'"
	ontrue1, _ := p.NewActionStep()
	ontrue1.Name = "Action Step 1"
	onfalse1, _ := p.NewActionStep()
	onfalse1.Name = "Action Step 2"
	while.OnTrue = ontrue1.GetID()
	while.OnFalse = onfalse1.GetID()

	ret1, _ := p.NewReturnStep()
	ret1.Name = "Return To While Step 1"
	ret1.Description = "Example of a return step from branch 1."
	ontrue1.OnSuccess = ret1.GetID()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
