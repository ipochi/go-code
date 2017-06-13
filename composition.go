package main

import "fmt"

// we wil try to recreate the example in the below link to solidify the understanding of compoisiton in go.
// https://www.goinggo.net/2015/09/composition-with-go.html

/*
The idea behind this program is we have a contractor that we hired to renovate our house.
In particular, there are some boards in the house that have rotted and need to be yanked out,
as well as new boards that need to be nailed in. The contractor will be given a supply of nails,
boards to work with and tools to perform the work.

*/

type Board struct {
	NailsNeeded int
	NailsDriven int
}

// behavior to drive the nails in a board can be incorporated in an interface.

type NailDriver interface {
	DriveNail(nailSupply int, b *Board)
}

// behavior to pull nails can be incorporated in an interface.

type NailPuller interface {
	PullNail(nailSupply int, b *Board)
}

// both the behaviours are now part of another interface.

type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// Lets create a tool called Mallet that has the behavior of driving the nails.
// its an empty struct
type Mallet struct{}

// let this Mallet type implement the interface which has the behavior of driving the nails.
// so we create a method  [with a reciever of Mallet type].
func (Mallet) DriveNail(nailSupply int, b *Board) {

	// Take a nail from the supply
	fmt.Println("Taking a nail from the supply")
	nailSupply = nailSupply - 1
	// drive it in the board.
	fmt.Println("Driving a nail into the Board")
	b.NailsDriven = b.NailsDriven + 1
}

// lets create a tool called CrowBar that has the behaviour of pulling the nails.

type Crowbar struct{}

// Let this Crowbar implement the interface which has the behaviour of pulling the nails.

func (Crowbar) PullNail(nailSupply int, b *Board) {

	// Remove a nial from the board
	b.NailsDriven = b.NailsDriven - 1
	// Add that nail to the suuply.
	nailSupply = nailSupply + 1
}

// Lets create a type of contractor  to do our work.

type Contractor struct{}

// Contractor has two behaviours driving a nail and pulling a nail.
// Hence the Contractor must implement the interfaces NailDriver and NailPuller
// He does it by using the appropriate tool.

// Lets create a method for Constractor that drives a nail

// Tool to be used in this method is a Mallet , which implements the behaviour
// of driving a nail ie inteface - NailDriver
func (Contractor) Fasten(nd NailDriver, nailSupply int, b *Board) {
	for b.NailsNeeded > b.NailsDriven {
		nd.DriveNail(nailSupply, b)
	}
}

// Tool to be used in this method is a CrowBar , which implements the behaviour
// of pulling a nail ie inteface - NailPuller
func (Contractor) Unfasten(np NailPuller, nailSupply int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		np.PullNail(nailSupply, b)
	}
}

// Another behaviour of a contractor is processing the Boards.

func (c Contractor) ProcessBoards(ndp NailDrivePuller, nailSupply int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		switch {
		case b.NailsNeeded > b.NailsDriven:
			c.Fasten(ndp, nailSupply, b)
		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(ndp, nailSupply, b)
		}
	}
}

// Toolbox type can have tools that can drive the nails or pull the nails
// Hence the Toolbox struct implements both the behaviours and the no of nails
type ToolBox struct {
	NailDriver
	NailPuller
	nails int
}

func main() {
	boards := []Board{
		// Boards to tbe removed
		{NailsDriven: 4},
		{NailsDriven: 5},
		{NailsDriven: 6},

		// Boards to be fastened

		{NailsNeeded: 6},
		{NailsNeeded: 8},
		{NailsNeeded: 9},
	}

	//lets get the toolbox ready

	tb := ToolBox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	// Display the current state of our toolbox and boards.
	displayState(&tb, boards)

	var c Contractor
	c.ProcessBoards(&tb, tb.nails, boards)

	displayState(&tb, boards)
}

//  Gives the  information about all the boards.
func displayState(tb *ToolBox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}
