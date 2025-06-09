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
	startstep(prep(playbook.New()))
	returnstep(prep(playbook.New()))
	endstep(prep(playbook.New()))
	parallelstep(prep(playbook.New()))
	foreachstep(prep(playbook.New()))
	whilestep(prep(playbook.New()))
	ifthemstep(prep(playbook.New()))
	switchstep(prep(playbook.New()))
	actionstep1(prep(playbook.New()))
	actionstep2(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func startstep(p *playbook.Playbook) {
	start, _ := p.NewStartStep()
	start.Name = "Playbook Start"
	start.Description = "Example of a start step."
	step0, _ := p.NewActionStep()
	step0.Name = "Action Step 0"
	start.OnSuccess = step0.GetID()

	header("start step")
	encode(p)
}

func returnstep(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	ret, _ := p.NewReturnStep()
	ret.Name = "Return To Parallel Step 1"
	ret.Description = "Example of a return step."

	header("return step")
	encode(p)
}

func endstep(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	end, _ := p.NewEndStep()
	end.Name = "Playbook End"
	end.Description = "Example of an end step."

	header("end step")
	encode(p)
}

func parallelstep(p *playbook.Playbook) {
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

	header("parallel step")
	encode(p)
}

func foreachstep(p *playbook.Playbook) {
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

	header("foreach step")
	encode(p)
}

func whilestep(p *playbook.Playbook) {
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

	header("while step")
	encode(p)
}

func ifthemstep(p *playbook.Playbook) {
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

	header("if-then step")
	encode(p)
}

func switchstep(p *playbook.Playbook) {
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

	header("switch step")
	encode(p)
}

func actionstep1(p *playbook.Playbook) {
	// Create workflow steps for this playbook
	step1, _ := p.NewActionStep()
	step1.Name = "Action Step 1"
	step1.Description = "Example of a manual action step."

	// Create workflow steps for this playbook
	end, _ := p.NewEndStep()
	end.Name = "Playbook End"
	end.Description = "Example of an end step."

	step1.OnSuccess = end.GetID()
	step1.Agent = "individual--328a89ab-3b8f-40c4-a491-24a40bcd3cd4"
	cmd1, _ := step1.NewManualCommand()
	cmd1.Command = "Disconnect the infected machines from the network"

	header("action step 1")
	encode(p)
}

func actionstep2(p *playbook.Playbook) {
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
	cmd1, _ := step1.NewHTTPCommand()
	cmd1.Command = "/v1/blockSystem?id=192.168.0.100"

	header("action step 2")
	encode(p)
}

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
