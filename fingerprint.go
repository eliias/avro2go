package avro2go

const empty uint64 = 0xc15d213aa4d7a795

var table [256]uint64

func init() {
	for i := uint64(0); i < 256; i++ {
		fp := i
		for j := 0; j < 8; j++ {
			fp = (fp >> uint64(1)) ^ (empty & -(fp & uint64(1)))
			table[i] = fp
		}
	}
}

func Fingerprint64(buf []byte) uint64 {
	fp := empty
	for i := uint64(0); i < uint64(len(buf)); i++ {
		fp = (fp >> 8) ^ table[(int)(fp^uint64(buf[i]))&0xff]
	}
	return fp
}
