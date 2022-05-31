// Copyright 2022 pschou (https://github.com/pschou/go-colormap)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package colormap_test

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	//colormap "pkg/github.com/pschou/go-colormap"
	"testing"

	colormap "github.com/pschou/go-colormap"
)

func ExampleBGRUint8() {
	for i := 0; i < 16; i++ {
		rgb := colormap.BGRUint8(float32(i) / 255)
		fmt.Println(rgb)
	}
	// Output:
	// [0 0 127]
	// [0 0 131]
	// [0 0 135]
	// [0 0 139]
	// [0 0 143]
	// [0 0 147]
	// [0 0 151]
	// [0 0 155]
	// [0 0 159]
	// [0 0 163]
	// [0 0 167]
	// [0 0 171]
	// [0 0 175]
	// [0 0 179]
	// [0 0 183]
	// [0 0 187]
}
func TestBGRUint8(t *testing.T) {
	for i := 0; i < 256; i++ {
		rgb := colormap.BGRUint8(float32(i) / 255)
		rev := colormap.RevBGRUint8(rgb) * 255
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0]), float32(rgb[1]), float32(rgb[2]), rev)
			t.Errorf("FAIL on %d", i)
			///fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
		}
	}
}

func ExampleRGBUint8() {
	for i := 0; i < 16; i++ {
		rgb := colormap.RGBUint8(float32(i) / 255)
		fmt.Println(rgb)
	}
	// Output:
	// [127 0 0]
	// [131 0 0]
	// [135 0 0]
	// [139 0 0]
	// [143 0 0]
	// [147 0 0]
	// [151 0 0]
	// [155 0 0]
	// [159 0 0]
	// [163 0 0]
	// [167 0 0]
	// [171 0 0]
	// [175 0 0]
	// [179 0 0]
	// [183 0 0]
	// [187 0 0]
}

func TestRGBUint8(t *testing.T) {
	for i := 0; i < 256; i++ {
		rgb := colormap.RGBUint8(float32(i) / 255)
		rev := colormap.RevRGBUint8(rgb) * 255
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}
func TestBGRUint16(t *testing.T) {
	for i := 0; i < 65536; i++ {
		rgb := colormap.BGRUint16(float64(i) / 65535)
		rev := colormap.RevBGRUint16(rgb) * 65535
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}

func TestRGBUint16(t *testing.T) {
	for i := 0; i < 65536; i++ {
		rgb := colormap.RGBUint16(float64(i) / 65535)
		rev := colormap.RevRGBUint16(rgb) * 65535
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}

func Example_makePNG() {
	width := 512
	height := 100

	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Create RGB Palette image
	for x := 0; x < width; x++ {
		rgb := colormap.RGBUint8(float32(x) / float32(width))
		c := color.RGBA{rgb[0], rgb[1], rgb[2], 0xff}
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}

	// Write out RGB colorbar
	f2, _ := os.Create("RGB.png")
	png.Encode(f2, img)

	// Create RGB Palette image
	for x := 0; x < width; x++ {
		rgb := colormap.BGRUint8(float32(x) / float32(width))
		c := color.RGBA{rgb[0], rgb[1], rgb[2], 0xff}
		for y := 0; y < height; y++ {
			img.Set(x, y, c)
		}
	}

	// Write out BGR colorbar
	f1, _ := os.Create("BGR.png")
	png.Encode(f1, img)
	// Output:
}
