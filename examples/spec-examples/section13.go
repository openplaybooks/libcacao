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
	steplayout1(prep(playbook.New()))
}

func prep(p *playbook.Playbook) *playbook.Playbook {
	p.Created = ""
	p.Modified = ""
	return p
}

func steplayout1(p *playbook.Playbook) {
	p.CanvasWidth = 1000
	p.CanvasHeight = 500

	start, _ := p.NewStartStep()
	end, _ := p.NewEndStep()

	start.Name = "Playbook Start"
	start.Description = "Example of a start step."
	start.OnSuccess = end.GetID()
	start.XCoordinate = 100
	start.YCoordinate = 50
	c1, _ := start.NewConnection("on-success")
	c1.AddPoint(100, 50)
	c1.AddPoint(200, 50)

	// A few other ways of doing this.
	// points1 := steps.ConnectionPoints{5, 10}
	// points2 := steps.ConnectionPoints{10, 20}
	// start.NewConnection("on-success", points1, points2)
	// start.NewConnection("on-success", steps.ConnectionPoints{5, 10}, steps.ConnectionPoints{10, 20})

	end.Name = "Playbook End"
	end.Description = "Example of an end step."
	end.XCoordinate = 200
	end.YCoordinate = 50

	header("canvaslayout1")
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
