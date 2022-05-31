package colors

func RGBUint16(n float64) (rgb []uint16) {
	max := float64(65535)
	return []uint16{
		uint16(trap64(n+0.25)*max + 0.25),
		uint16(trap64(n)*max + 0.25),
		uint16(trap64(n-0.25)*max + 0.25),
	}
}

func BGRUint16(n float64) (rgb []uint16) {
	max := float64(65535)
	return []uint16{
		uint16(trap64(n-0.25)*max + 0.25),
		uint16(trap64(n)*max + 0.25),
		uint16(trap64(n+0.25)*max + 0.25),
	}
}

func RGBUint8(n float32) (rgb []uint8) {
	max := float32(255)
	return []uint8{
		uint8(trapUint8(n+0.25)*max + 0.25),
		uint8(trapUint8(n)*max + 0.25),
		uint8(trapUint8(n-0.25)*max + 0.25),
	}
}

func BGRUint8(n float32) (rgb []uint8) {
	max := float32(255)
	return []uint8{
		uint8(trapUint8(n-0.25)*max + 0.25),
		uint8(trapUint8(n)*max + 0.25),
		uint8(trapUint8(n+0.25)*max + 0.25),
	}
}

func trap64(n float64) float64 {
	if n <= 0.125 || n >= 0.875 {
		return 0
	}
	switch {
	case n < 0.375:
		return (n - 0.125) * 4
	case n > 0.625:
		return (0.875 - n) * 4
	}
	return 1
}
func trapUint8(n float32) float32 {
	if n <= 0.125 || n >= 0.875 {
		return 0
	}
	switch {
	case n < 0.375:
		return (n - 0.125) * 4
	case n > 0.625:
		return (0.875 - n) * 4
	}
	return 1
}

func RevBGRUint8(rgb []uint8) float32 {
	return RevRGBUint8([]uint8{rgb[2], rgb[1], rgb[0]})
}
func RevRGBUint8(rgb []uint8) float32 {
	switch {
	case rgb[0] > rgb[1]:
		if rgb[1] == 0 {
			return (float32(rgb[0])/4 - 31.75) / 255
			//return float32(rgb[0])/4 - 0.1245
		}
		return (float32(rgb[1])/4 + 32) / 255
		//return float32(rgb[1])/4 + 0.1254
	case rgb[2] > rgb[1]:
		if rgb[1] == 0 {
			return (-float32(rgb[2])/4 + 286.75) / 255
			//return -float32(rgb[2])/4 + 1.1245
		}
		return (-float32(rgb[1])/4 + 223) / 255
		//return -float32(rgb[1])/4 + 0.8745
	}
	return (float32(rgb[2])/4 + 95.75) / 255
	//return float32(rgb[2])/4 + 0.3754
}

func RevBGRUint16(rgb []uint16) float64 {
	return RevRGBUint16([]uint16{rgb[2], rgb[1], rgb[0]})
}
func RevRGBUint16(rgb []uint16) float64 {
	switch {
	case rgb[0] > rgb[1]:
		if rgb[1] == 0 {
			return (float64(rgb[0])/4 - 8191.75) / 65535 //- 0.1245
		}
		return (float64(rgb[1])/4 + 8192.0) / 65535 //+ 0.1254
	case rgb[2] > rgb[1]:
		if rgb[1] == 0 {
			return (-float64(rgb[2])/4 + 73726.75) / 65535 //+ 1.1245
		}
		return (-float64(rgb[1])/4 + 57343) / 65535 //+ 0.8745
	}
	return (float64(rgb[2])/4 + 24575.75) / 65535 //+ 0.3754
}
