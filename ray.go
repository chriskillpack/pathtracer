package main

import (
  "pathtracer/vector"
)

type Ray struct {
  // The origin of the ray in world space.
  origin vector.Vector3

  // The normalized direction vector of the ray in world space.
  direction vector.Vector3
}
