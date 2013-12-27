package main

import (
  "math"
  "pathtracer/vector"
)

type Plane struct {
  // The orientation of the plane.
  normal vector.Vector3
  // The distance of the plane from the origin along it's normal.
  offset float32

  material Material
}

func (plane Plane) Intersect(ray Ray) Intersection {
  Vd := vector.Dot(plane.normal, ray.direction)
  if math.Abs(float64(Vd)) < 1e-2 {
    // Ray is parallel to plane, no intersection.
    return Intersection{}
  }
  V0 := -(vector.Dot(plane.normal, ray.origin) - plane.offset)
  t := V0 / Vd
  if t < 0 {
    // Intersection is behind the ray origin, ignore.
    return Intersection{}
  }

  return Intersection{true, t, plane.normal}
}

func (plane Plane) Material() Material {
  return plane.material
}
