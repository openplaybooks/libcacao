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
	parallel, _ := p.NewParallelStep()
	parallel.Name = "Parallel Step 1"
	parallel.Description = "Example of a parallel step."
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 1"
	step2, _ := p.NewActionStep()
	step2.Name = "Action Step 2"
	step3, _ := p.NewActionStep()
	step3.Name = "Action Step 3"
	parallel.AddNextSteps(step1.GetID())
	parallel.AddNextSteps(step2.GetID())
	parallel.OnSuccess = step3.GetID()

	ret1, _ := p.NewReturnStep()
	ret1.Name = "Return To Parallel Step 1"
	ret1.Description = "Example of a return step from branch 1."
	step1.OnSuccess = ret1.GetID()

	ret2, _ := p.NewReturnStep()
	ret2.Name = "Return To Parallel Step 1"
	ret2.Description = "Example of a return step from branch 2."
	step2.OnSuccess = ret2.GetID()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
