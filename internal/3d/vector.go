package three_d

type Vector2 struct {
	X, Y float64
}

type Vector3 struct {
	X, Y, Z float64
}

// See /home/per/Documents/3d projection (latex)/3d projection.tex
func projectTo2d(v Vector3) Vector2 {
	return Vector2{v.X / v.Z, v.Y / v.Z}
}

func projectToScreen(v Vector2) Vector2 {
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

	return Vector2{(v.X + 1) / 2 * width, (1 - (v.Y+1)/2) * height}
}
