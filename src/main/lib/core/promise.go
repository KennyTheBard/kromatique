package core

import "image"

// Promise encapsulates the image being worked on and a
// reference to the order contract of the appliance instance
type Promise struct {
	img      image.Image
	contract Contract
}

// Result blocks the current goroutine until the image effect
// is applied completely or returns it immediately if it was
// already completed when is called
func (p Promise) Result() image.Image {
	p.contract.Deadline()
	return p.img
}
