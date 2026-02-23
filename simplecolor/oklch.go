package simplecolor

import "math"

// OKLCH represents a color in the OKLCH color space.
type OKLCH struct {
	L float64 // Lightness: 0.0 to 1.0
	C float64 // Chroma: 0.0 to ~0.4
	H float64 // Hue: 0 to 360 degrees
}

// sRGBToLinear converts a single sRGB channel value to linear RGB.
func sRGBToLinear(c float64) float64 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

// linearToSRGB converts a single linear RGB channel value to sRGB.
func linearToSRGB(c float64) float64 {
	if c <= 0.0031308 {
		return c * 12.92
	}
	return 1.055*math.Pow(c, 1.0/2.4) - 0.055
}

// ToOKLCH converts a SimpleColor to the OKLCH color space.
func (c SimpleColor) ToOKLCH() OKLCH {
	r, g, b, _ := c.RGBA()
	// Normalize to 0-1
	rf := sRGBToLinear(float64(r) / 255.0)
	gf := sRGBToLinear(float64(g) / 255.0)
	bf := sRGBToLinear(float64(b) / 255.0)

	// Linear RGB to LMS (Ottosson)
	l_ := 0.4122214708*rf + 0.5363325363*gf + 0.0514459929*bf
	m_ := 0.2119034982*rf + 0.6806995451*gf + 0.1073969566*bf
	s_ := 0.0883024619*rf + 0.2220049174*gf + 0.6896926208*bf

	l := math.Cbrt(l_)
	m := math.Cbrt(m_)
	s := math.Cbrt(s_)

	// LMS to OKLab
	L := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	a := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	bLab := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s

	// OKLab to OKLCH
	C := math.Sqrt(a*a + bLab*bLab)
	H := math.Atan2(bLab, a) * 180.0 / math.Pi
	if H < 0 {
		H += 360.0
	}

	return OKLCH{L: L, C: C, H: H}
}

// FromOKLCH creates a SimpleColor from OKLCH values.
func FromOKLCH(l, c, h float64) SimpleColor {
	// OKLCH to OKLab
	hRad := h * math.Pi / 180.0
	a := c * math.Cos(hRad)
	b := c * math.Sin(hRad)

	// OKLab to LMS (inverse of the second matrix)
	l_ := l + 0.3963377774*a + 0.2158037573*b
	m_ := l - 0.1055613458*a - 0.0638541728*b
	s_ := l - 0.0894841775*a - 1.2914855480*b

	// Cube the LMS values
	lc := l_ * l_ * l_
	mc := m_ * m_ * m_
	sc := s_ * s_ * s_

	// LMS to linear RGB (inverse of the first matrix)
	rf := 4.0767416621*lc - 3.3077115913*mc + 0.2309699292*sc
	gf := -1.2684380046*lc + 2.6097574011*mc - 0.3413193965*sc
	bf := -0.0041960863*lc - 0.7034186147*mc + 1.7076147010*sc

	// Clamp to [0, 1]
	rf = math.Max(0, math.Min(1, rf))
	gf = math.Max(0, math.Min(1, gf))
	bf = math.Max(0, math.Min(1, bf))

	// Linear RGB to sRGB
	ri := uint32(math.Round(linearToSRGB(rf) * 255.0))
	gi := uint32(math.Round(linearToSRGB(gf) * 255.0))
	bi := uint32(math.Round(linearToSRGB(bf) * 255.0))

	return SimpleColor(ri<<16 | gi<<8 | bi)
}

// ClampOKLCH clamps a color's lightness and chroma within the given bounds,
// preserving hue.
func (c SimpleColor) ClampOKLCH(minL, maxL, minC, maxC float64) SimpleColor {
	oklch := c.ToOKLCH()
	oklch.L = math.Max(minL, math.Min(maxL, oklch.L))
	oklch.C = math.Max(minC, math.Min(maxC, oklch.C))
	return FromOKLCH(oklch.L, oklch.C, oklch.H)
}

// ClampOKLCH clamps all colors in the palette within the given OKLCH bounds.
func (s SimplePalette) ClampOKLCH(minL, maxL, minC, maxC float64) SimplePalette {
	result := make(SimplePalette, len(s))
	for i, c := range s {
		result[i] = c.ClampOKLCH(minL, maxL, minC, maxC)
	}
	return result
}

// NormalizeLightness sets all colors in the palette to the same OKLCH lightness.
func (s SimplePalette) NormalizeLightness(targetL float64) SimplePalette {
	result := make(SimplePalette, len(s))
	for i, c := range s {
		oklch := c.ToOKLCH()
		result[i] = FromOKLCH(targetL, oklch.C, oklch.H)
	}
	return result
}

// NormalizeChroma sets all colors in the palette to the same OKLCH chroma.
func (s SimplePalette) NormalizeChroma(targetC float64) SimplePalette {
	result := make(SimplePalette, len(s))
	for i, c := range s {
		oklch := c.ToOKLCH()
		result[i] = FromOKLCH(oklch.L, targetC, oklch.H)
	}
	return result
}
