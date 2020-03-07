package position

import (
	utils "../utils"
	"math"
)

// AxialPosition is a position interface that treats just one axis
// and is useful for composing relative positions with axial positions
// of different types
type AxialPosition interface {
	Get(start, end int) int
}

// FixedAxialPosition is a wrapper on a fixed value
type FixedAxialPosition struct {
	value int
}

func (p *FixedAxialPosition) Get(_, _ int) int {
	return p.value
}

func NewFixedAxialPosition(val int) *FixedAxialPosition {
	fap := new(FixedAxialPosition)
	fap.value = val

	return fap
}

// RelativeAxialPosition is an axial position representation that uses
// a fixed value added to a chosen relative key point of the image
type RelativeAxialPosition struct {
	value int
	mode  int
}

func (p *RelativeAxialPosition) Get(start, end int) int {
	return getRelative(start, end, p.mode) + p.value
}

func NewRelativeAxialPosition(val, mode int) *RelativeAxialPosition {
	rap := new(RelativeAxialPosition)
	rap.value = val
	rap.mode = mode

	return rap
}

// RelativeAxialPosition is an axial position representation that uses
// a ratio that is used to interpolate between the key points of the image
type PercentAxialPosition struct {
	percent float64
}

func (p *PercentAxialPosition) Get(start, end int) int {
	return int(math.Round(utils.LERP(float64(start), float64(end), p.percent)))
}

func NewPercentAxialPosition(val float64) *PercentAxialPosition {
	par := new(PercentAxialPosition)
	par.percent = val

	return par
}
