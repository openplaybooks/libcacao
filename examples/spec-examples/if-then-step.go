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
	ifthen, _ := p.NewIfThenStep()
	ifthen.Name = "If Then Step 1"
	ifthen.Description = "Example of an if then step."
	ifthen.Condition = "__variable__:value = '10.0.0.0/8'"
	ontrue1, _ := p.NewActionStep()
	ontrue1.Name = "Action Step 1"
	onfalse1, _ := p.NewActionStep()
	onfalse1.Name = "Action Step 2"
	ifthen.OnTrue = ontrue1.GetID()
	ifthen.OnFalse = onfalse1.GetID()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
