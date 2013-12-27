package main

import (
  "math"
  "pathtracer/vector"
)

type Material interface {
  Shade(ray Ray, intersection Intersection) vector.Vector3
}

type DiffuseMaterial struct {
  color vector.Vector3
}

func (m DiffuseMaterial) Shade(ray Ray, intersection Intersection) vector.Vector3 {
  lambert := vector.Dot(ray.direction, intersection.normal)
  lambert = float32(math.Max(float64(lambert), 0))
  lambert = float32(math.Min(float64(lambert), 1))

  return vector.Scale(m.color, lambert)
}
