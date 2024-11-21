package noise

import (
	"math"
	"math/rand"
)

type Noise struct {
	value []float64
	size  int
	freq  float64
}

func New(seed int, freq, amp float64, size int) *Noise {
	n := &Noise{
		value: make([]float64, size*size),
		size:  size,
		freq:  freq,
	}

	rng := rand.New(rand.NewSource(int64(seed)))
	for i := range n.value {
		n.value[i] = rng.Float64() * amp
	}
	return n
}

func (n *Noise) Sample(x, y float64) float64 {
	x, y = x*n.freq, y*n.freq
	x0, y0 := math.Floor(x), math.Floor(y)
	fx, fy := x-x0, y-y0

	fx = fx * fx * (3 - 2*fx)
	fy = fy * fy * (3 - 2*fy)

	ix, iy := int(x0)&(n.size-1), int(y0)&(n.size-1)
	ix1, iy1 := (ix+1)&(n.size-1), (iy+1)&(n.size-1)

	v00 := n.value[ix+iy*n.size]
	v10 := n.value[ix1+iy*n.size]
	v01 := n.value[ix+iy1*n.size]
	v11 := n.value[ix1+iy1*n.size]

	return v00*(1-fx)*(1-fy) + v10*fx*(1-fy) + v01*(1-fx)*fy + v11*fx*fy
}
