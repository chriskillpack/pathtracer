package main

import (
  "math"
  "pathtracer/vector"
)

var DefaultIntersection = Intersection{distance: math.MaxFloat32}

// Intersection holds information about the intersection between a ray and a
// scene object.
type Intersection struct {
  // Did the ray intersect the object?
  doesIntersect bool
  // Distance along the ray to the closest point of intersection. May be
  // negative indicating that the intersection is behind the ray origin.
  distance      float32
  // Normalized surface normal of the object at the point of intersection.
  normal        vector.Vector3
}

// SceneObject specifies the common interface that all scene objects must
// implement.
type SceneObject interface {
  // Returns an Intersection object that describes the intersection between a
  // ray and the object. If no intersection exists then doesIntersect will be
  // false and all other field values should be ignored.
  Intersect(ray Ray) Intersection
}
