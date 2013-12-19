package main

import (
  "pathtracer/vector"
)

type Object interface {
  /*
    Tests if a ray intersects with the object. Returns a tuple (does_intersect,
    distance) where distance is the distance along the ray to the closest point
    of intersection.
  */
  Intersect(rayOrigin, rayDirection vector.Vector3) (doesIntersect bool, distance float32)
}
