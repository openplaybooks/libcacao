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
	start, _ := p.NewStartStep()
	start.Name = "Playbook Start"
	start.Description = "Example of a start step."
	step0, _ := p.NewActionStep()
	step0.Name = "Action Step 0"
	start.OnSuccess = step0.GetID()

	// After everything is done, then do this
	// Remove all of the IDs from the workflow steps since the specification only has them at the map level
	p.ClearWorkflowStepIDs()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
