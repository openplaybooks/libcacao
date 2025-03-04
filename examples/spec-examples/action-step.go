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
	p := playbook.New()

	v1, _ := p.NewStringVariable("__list_of_zeus_ip_addresses__", "list")
	v1.Description = "A list of zeus botnet IP addresses"
	v2, _ := p.NewStringVariable("__bad_ip_address__", "string")
	v2.Description = "A single zeus botnet IP addresses"
	v3, _ := p.NewStringVariable("__block_status__", "string")
	v3.Description = "Block IP address status code"

	start, _ := p.NewStartStep()
	start.Name = "Start of Playbook"
	start.Description = "A playbook example of using variables in two steps"

	// Create action step 1
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 1"
	step1.Description = "Get list of Botnet IP addresses."
	cmd1, _ := step1.NewHTTPAPICommand()
	cmd1.Command = "https://zeustracker.abuse.ch/blocklist.php?download=badips"
	cmd1.ReturnedData["response_body"] = "__list_of_zeus_ip_addresses__"
	cmd1.ReturnedData["response_headers"] = ""
	cmd1.ReturnedData["status_code"] = ""

	// Create foreach step
	step2, _ := p.NewForeachStep()
	step2.Name = "Foreach Step 1"
	step2.Description = "Loop through the list of bad IP addresses"
	step2.Collection = "__list_of_zeus_ip_addresses__:value"
	step2.Element = "__bad_ip_address__"

	// Create action step 2
	step3, _ := p.NewActionStep()
	step3.Name = "Action Step 3"
	step3.Description = "Block IP address."
	cmd2, _ := step2.NewHTTPAPICommand()
	cmd2.Command = "https://my.firewall.com/block.php?ip=__bad_ip_address__"
	cmd2.ReturnedData["status_code"] = "__block_status__"

	ret, _ := p.NewReturnStep()
	ret.Name = "Return to Foreach Step 1"
	ret.Description = "This will return to the foreach step"

	// Create workflow steps for this playbook
	end, _ := p.NewEndStep()
	end.Name = "Playbook End"
	end.Description = "Example of an end step."

	// --------------------------------------------------
	// Stitch the playbook together
	// --------------------------------------------------
	start.OnSuccess = step1.GetID()
	step1.OnSuccess = step2.GetID()
	step2.Do = step3.GetID()
	step3.OnSuccess = ret.GetID()
	step2.OnSuccess = end.GetID()

	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

}
