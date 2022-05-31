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

package colors

import (
	"fmt"
	"testing"
)

func TestBGRUint8(t *testing.T) {
	for i := 0; i < 256; i++ {
		rgb := BGRUint8(float32(i) / 255)
		rev := RevBGRUint8(rgb) * 255
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0]), float32(rgb[1]), float32(rgb[2]), rev)
			t.Errorf("FAIL on %d", i)
			///fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
		}
	}
}
func TestRGBUint8(t *testing.T) {
	for i := 0; i < 256; i++ {
		rgb := RGBUint8(float32(i) / 255)
		rev := RevRGBUint8(rgb) * 255
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}
func TestBGRUint16(t *testing.T) {
	for i := 0; i < 65535; i++ {
		rgb := BGRUint16(float64(i) / 65535)
		rev := RevBGRUint16(rgb) * 65535
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}
func TestRGBUint16(t *testing.T) {
	for i := 0; i < 65535; i++ {
		rgb := RGBUint16(float64(i) / 65535)
		rev := RevRGBUint16(rgb) * 65535
		if int64(i)*100 != int64(rev*100+0.5) {
			fmt.Printf("% 3d  %0.4f  %0.4f  %0.4f  %f\n", i, float32(rgb[0])/255, float32(rgb[1])/255, float32(rgb[2])/255, rev)
			t.Errorf("FAIL on %d", i)
		}
	}
}
