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

	p.Name = "Prevent FuzzyPanda Malware"
	p.Description = "This playbook will block traffic to the FuzzyPanda data exfil site"
	p.AddPlaybookTypes("prevention,test,test2")
	p.CreatedBy = "identity--uuid2"
	p.ValidFrom = p.GetCurrentTime("micro")
	p.ValidUntil = "2021-12-31T23:59:59.999999Z"
	p.AddDerivedFrom("playbook--uuid50")
	p.Priority = 3
	p.Severity = 70
	p.Impact = 5
	p.AddLabels("malware,fuzzypanda,apt")

	valid, count, details := p.Valid()

	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	fmt.Println("Object valid: ", valid)
	fmt.Println("Error Count: ", count)

	for _, result := range details {
		fmt.Println(result)
	}

}
