package bitmap

import (
	"unsafe"
)

var size = int(unsafe.Sizeof(uint(1)) * 8)

type BitMap struct {
	bits []uint
}

func New(length int) *BitMap {
	return &BitMap{
		make([]uint, length/size+1),
	}
}

func (b *BitMap) Set(idx uint) {
	index := idx / 32
	bitIndex := idx % 32

	b.bits[index] = b.bits[index] | (1 << bitIndex)
}

func (b *BitMap) Exists(idx uint) bool {
	index := idx / 32
	bitIndex := idx % 32
	return b.bits[index]&(1<<bitIndex) != 0
}
