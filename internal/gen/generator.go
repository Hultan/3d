package gen

import (
	"github.com/hultan/3d/internal/vec"
)

func GenerateCube() []vec.Vector3 {
	var points3d []vec.Vector3
	z := 1.5

	points3d = append(points3d, vec.Vector3{X: -0.5, Y: -0.5, Z: z})
	points3d = append(points3d, vec.Vector3{X: 0.5, Y: -0.5, Z: z})
	points3d = append(points3d, vec.Vector3{X: 0.5, Y: 0.5, Z: z})
	points3d = append(points3d, vec.Vector3{X: -0.5, Y: 0.5, Z: z})
	points3d = append(points3d, vec.Vector3{X: -0.5, Y: -0.5, Z: z + 1})
	points3d = append(points3d, vec.Vector3{X: 0.5, Y: -0.5, Z: z + 1})
	points3d = append(points3d, vec.Vector3{X: 0.5, Y: 0.5, Z: z + 1})
	points3d = append(points3d, vec.Vector3{X: -0.5, Y: 0.5, Z: z + 1})

	return points3d
}
