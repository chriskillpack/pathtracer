package main

import (
  "fmt"
  "image"
  "image/png"
  "math"
  "os"
  "pathtracer/vector"
)

var _ = fmt.Println

const (
  ImageWidth int = 256
  ImageHeight = 256
)

var (
  WorldUp = vector.Vector3{0,1,0}

  EyePosition = vector.Vector3{0,15,-10}
  EyeLookAtTarget = vector.Vector3{0,15,0}
  EyePlaneDist = float32(1.0)

  Scene []SceneObject
)

// Populate the scene with objects.
func populateScene(scene []SceneObject) []SceneObject {
  whiteDiffuse := DiffuseMaterial{vector.Vector3{1,1,1}, vector.Vector3{}}
  redDiffuse := DiffuseMaterial{vector.Vector3{1,0,0}, vector.Vector3{}}
  blueDiffuse := DiffuseMaterial{vector.Vector3{0,0,1}, vector.Vector3{}}
  whiteEmissive := DiffuseMaterial{vector.Vector3{1,1,1}, vector.Vector3{10,10,10}}

  scene = append(scene,
    Sphere{vector.Vector3{2,10,2}, 2.5, whiteDiffuse},
    Plane{vector.Vector3{-1,0,0}, 10, redDiffuse}, // Left
    Plane{vector.Vector3{0,0,1}, 20, whiteDiffuse},  // Back
    Plane{vector.Vector3{1,0,0}, 10, blueDiffuse}, // Right
    Plane{vector.Vector3{0,1,0}, 0, whiteDiffuse},  // Bottom
    Plane{vector.Vector3{0,1,0}, 30, whiteEmissive})  // Top

  return scene
}

func main() {
  m := image.NewRGBA(image.Rect(0, 0, ImageWidth, ImageHeight))
  img,_ := os.Create("foo.png")
  defer img.Close()

  Scene = populateScene(Scene)

  // Compute the eye-space to world-space transformation basis vectors. We use
  // these basis vectors to generate target points for ray generation. The
  // whole process is a little long-winded but working up from first principles.
  lookVector := vector.Sub(EyeLookAtTarget, EyePosition).Normalize()
  xAxis, yAxis, _ := coordinateFrame(lookVector)

  for i := 0; i < ImageHeight; i++ {
    for j := 0; j < ImageWidth; j++ {
      // Compute the 2D position of this pixel on the view plane. The view plane
      // is centered around the view vector and extends [-1,1] in both axis.
      npx := float32(j - ImageWidth/2) / float32(ImageWidth/2)
      npy := float32(ImageHeight/2 - i) / float32(ImageHeight/2)

      // Scale the world-space basis vectors of the view plane to generate
      // offset vectors from the center point of the view plane.
      xAxisOffset := vector.Scale(xAxis, npx)
      yAxisOffset := vector.Scale(yAxis, npy)
      // Move down the view vector from the eye to generate the center point
      // on the view plane in world-space.
      midPoint := vector.Add(EyePosition, vector.Scale(lookVector, EyePlaneDist))

      // Add the offset vectors on to the center point to generate the ray
      // target in world-space.
      target := vector.Add(vector.Add(midPoint, xAxisOffset), yAxisOffset)

      // Finally compute a ray direction from the eye to the target.
      rayDirection := vector.Sub(target, EyePosition).Normalize()

      // Sample the scene using the ray.
      ray := Ray{EyePosition, rayDirection}
      var pixelColor vector.Vector3
      for s := 0 ; s < 20 ; s++ {
        pixelColor = vector.Add(pixelColor, sampleScene(ray, 10))
      }

      r, g, b := colorToColor8(pixelColor)

      index := i * m.Stride + j * 4
      m.Pix[index+0] = r
      m.Pix[index+1] = g
      m.Pix[index+2] = b
      m.Pix[index+3] = 255;
    }
  }
  png.Encode(img, m)
}

// Samples the scene with the ray and returns the color along that path.
func sampleScene(ray Ray, numLevels int32) vector.Vector3 {
  if numLevels == 0 {
    return vector.Vector3{} // No more recursing, return black
  }

  var closestIntersection Intersection = DefaultIntersection
  var closestObject SceneObject = nil
  for _, object := range Scene {
    intersection := object.Intersect(ray)
    if intersection.doesIntersect {
      if intersection.distance < closestIntersection.distance {
        closestIntersection = intersection
        closestObject = object
      }
    }
  }

  var color vector.Vector3
  if closestIntersection.doesIntersect {
    var diffuseMaterial DiffuseMaterial
    diffuseMaterial = closestObject.Material().(DiffuseMaterial)

    // Compute new sampling direction
    newRayOrigin := vector.Add(ray.origin, vector.Scale(ray.direction, closestIntersection.distance))
    newRayDirection := GenerateHemisphereDirection(closestIntersection.normal)

    newRay := Ray{newRayOrigin, newRayDirection}
    sampledColor := sampleScene(newRay, numLevels-1)
    modulatedColor := vector.Mul(diffuseMaterial.diffuseColor, sampledColor)
    color = vector.Add(diffuseMaterial.emissiveColor, modulatedColor)
    // color = diffuseMaterial.Shade(ray, closestIntersection)
  } else {
    // Return black (for now)
  }

  return color
}

func colorToColor8(color vector.Vector3) (r, g, b uint8) {
  r = colorToUint8(color.X * 255)
  g = colorToUint8(color.Y * 255)
  b = colorToUint8(color.Z * 255)
  return
}

// Convert a normal into a color. Normal components are in the range [-1,1] and
// are converted to colors in the range [0,254].
func normalToColor8(normal vector.Vector3) (r, g, b uint8) {
  r = colorToUint8((normal.X + 1) * 127)
  g = colorToUint8((normal.Y + 1) * 127)
  b = colorToUint8((normal.Z + 1) * 127)
  return
}

// Convert a floating point value to a clamped value that fits in the [0,255]
// range of a uint8.
func colorToUint8(x float32) uint8 {
  if (x > 255) {
    return 255
  } else if (x < 0) {
    return 0
  } else {
    return uint8(x)
  }
}

// Given an input vector n this function generates a coordinate frame with the
// input vector as the Z axis.
func coordinateFrame(n vector.Vector3) (u, v, w vector.Vector3) {
  if (math.Abs(float64(vector.Dot(n, WorldUp))) > 0.9) {
    // Input too close to world up, use another world up.
    u = vector.Cross(vector.Vector3{0,0,-1}, n)
  } else {
    u = vector.Cross(WorldUp, n).Normalize()
  }

  v = vector.Cross(n, u).Normalize()
  w = vector.Normalize(n)

  return
}
