package tlv

import "github.com/Donge-ying/MiraiGo/binary"

func T17A(value int32) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x17a)
		w.WriteTlv(binary.NewWriterF(func(w *binary.Writer) {
			w.WriteUInt32(uint32(value))
		}))
	})
}
