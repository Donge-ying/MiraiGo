package tlv

import "github.com/Donge-ying/MiraiGo/binary"

func T16A(arr []byte) []byte {
	return binary.NewWriterF(func(w *binary.Writer) {
		w.WriteUInt16(0x16A)
		w.WriteTlv(arr)
	})
}
