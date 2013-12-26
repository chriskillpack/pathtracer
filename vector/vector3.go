package vector

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

func (v Vector3) Scale(factor float32) Vector3 {
  return Vector3{v.X * factor, v.Y * factor, v.Z * factor}
}

func (v Vector3) Normalize() Vector3 {
  return Normalize(v)
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

func Scale(v Vector3, x float32) Vector3 {
  return Vector3{v.X * x, v.Y * x, v.Z * x}
}

func Normalize(x Vector3) Vector3 {
  length := x.Len()
  if length > 1e-2 {
    recip_length := 1 / length
    return Vector3{x.X * recip_length, x.Y * recip_length, x.Z * recip_length}
  }
  return x
}
