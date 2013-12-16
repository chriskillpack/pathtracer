package main

type Object interface {
  /*
    Tests the ray against the object. Returns the distance along the ray to the
    closest point of intersection or MaxFloat32 in case of no intersection.
  */
  Intersect(ray_origin, ray_direction Vector3) float32
}