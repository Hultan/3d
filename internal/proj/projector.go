package proj

import (
	"github.com/hultan/3d/internal/vec"
)

// See /home/per/Documents/3d projection (latex)/3d projection.tex

func ProjectTo2d(v vec.Vector3) vec.Vector2 {
	return vec.Vector2{
		X: v.X / v.Z,
		Y: v.Y / v.Z,
	}
}

func ProjectToScreen(v vec.Vector2, screenWidth, screenHeight float64) vec.Vector2 {
	// We want :
	// v.X => -1 .. 1 => 0 .. width
	// So we do:
	// v.X + 1 => 0 .. 2
	// (v.X + 1)/2 => 0 .. 1
	// (v.X + 1)/2*width => 0 .. width

	// We also want :
	// v.Y => -1 .. 1 => height .. 0
	// So we do:
	// v.Y + 1 => 0 .. 2
	// (v.Y + 1)/2 => 0 .. 1
	// (1-(v.Y + 1)/2) => 1 .. 0				// Invert it since Y axis is inverted
	// (1-(v.Y + 1)/2)*height => height .. 0

	return vec.Vector2{
		X: (v.X + 1) / 2 * screenWidth,
		Y: (1 - (v.Y+1)/2) * screenHeight,
	}
}
