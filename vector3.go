package main

import (
  "math"
)

type Vector3 struct {
  X, Y, Z float32
}

func (v Vector3) Len() float32 {
  len := math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z))
  return float32(len)
}

func Dot(a, b Vector3) float32 {
  return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

func Cross(a, b Vector3) Vector3 {
  return Vector3{a.Y*b.Z - a.Z*b.Y,
                 a.Z*b.X - a.X*b.Z,
                 a.X*b.Y - a.Y*b.X}
}

func Add(a, b Vector3) Vector3 {
  return Vector3{a.X+b.X, a.Y+b.Y, a.Z+b.Z}
}

func Sub(a, b Vector3) Vector3 {
  return Vector3{a.X-b.X, a.Y-b.Y, a.Z-b.Z}
}
