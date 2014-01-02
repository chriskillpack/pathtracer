package main

import (
  "math"
  "math/rand"
  "pathtracer/vector"
)

type Material interface {
  Shade(ray Ray, intersection Intersection) vector.Vector3
}

// Generate a random direction on a hemisphere oriented around the input normal.
// The random directions have a cosine-weighted distribution.
// See formula 35 in http://people.cs.kuleuven.be/~philip.dutre/GI/TotalCompendium.pdf
func GenerateHemisphereDirection(normal vector.Vector3) vector.Vector3 {
  r1 := 2 * math.Pi * rand.Float64()
  r2 := rand.Float64()
  r2s := math.Sqrt(1-r2)

  // Compute a direction in the unit hemisphere
  x := float32(math.Cos(r1) * r2s)
  y := float32(math.Sin(r1) * r2s)
  z := float32(math.Sqrt(r2))

  // Compute a coordinate frame around the normal and then linearly combine
  // the basis vectors using the direction components as weights to transform
  // from unit hemisphere coordinate system to the normal's coordinate system.
  u, v, w := coordinateFrame(normal)
  t := vector.Add(vector.Scale(u, x), vector.Scale(v, y))
  return vector.Add(t, vector.Scale(w, z))
}

type DiffuseMaterial struct {
  diffuseColor, emissiveColor vector.Vector3
}

func (m DiffuseMaterial) Shade(ray Ray, intersection Intersection) vector.Vector3 {
  lambert := vector.Dot(ray.direction, intersection.normal)
  lambert = float32(math.Max(float64(lambert), 0))
  lambert = float32(math.Min(float64(lambert), 1))

  return vector.Scale(m.diffuseColor, lambert)
}
