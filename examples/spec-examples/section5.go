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
	variable1(prep(playbook.New()))
	marking1(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func variable1(p *playbook.Playbook) {

	// Create two variables for this playbook
	v1, _ := p.NewVariable("__ip_address__", "ipv4-addr", "1.2.3.4")
	v1.Description = "IPv4 address of malware CNC"
	header("variable 1")
	encode(p)
}

func marking1(p *playbook.Playbook) {

	// Create two data markings for this playbook
	p.NewTLPMarking("clear")
	m1, _ := p.NewStatementMarking("Copyright 2023 ACME Security Company")
	m1.CreatedBy = "identity--5abe695c-7bd5-4c31-8824-2528696cdbf1"
	header("marking 1")
	encode(p)
}

// ----------------------------------------------------------------------
// Define Supporting Functions and Methods
// ----------------------------------------------------------------------

func header(t string) {
	fmt.Println("\n// ----------------------------------------")
	fmt.Printf("// %s command example", t)
	fmt.Println("\n// ----------------------------------------")
}

func encode(p *playbook.Playbook) {
	// Encode
	data, err := p.EncodeToString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)
}
