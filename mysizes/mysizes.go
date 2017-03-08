package mysizes

import "fmt"

type ByteSize float64

const (
	_            = iota // ignorer første verdien siden den er 0
	KiB ByteSize = 1 << (10 * iota)
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

func (b ByteSize) String() string {
	switch {
	case b >= YiB:
		return fmt.Sprintf("%.2fYiB (yobibyte)", b/YiB)
	case b >= ZiB:
		return fmt.Sprintf("%.2fZiB (zebibyte)", b/ZiB)
	case b >= EiB:
		return fmt.Sprintf("%.2fEiB (exbibyte)", b/EiB)
	case b >= PiB:
		return fmt.Sprintf("%.2fPiB (pebibyte)", b/PiB)
	case b >= TiB:
		return fmt.Sprintf("%.2fTiB (tebibyte)", b/TiB)
	case b >= GiB:
		return fmt.Sprintf("%.2fGiB (gibibyte)", b/GiB)
	case b >= MiB:
		return fmt.Sprintf("%.2fMiB (mebibyte) = %.2f bytes", b/MiB, b)
	case b >= KiB:
		return fmt.Sprintf("%.2fKiB (kibibyte)", b/KiB)
	}
	return fmt.Sprintf("%.2fB (byte = 8 bits)", b)
}
