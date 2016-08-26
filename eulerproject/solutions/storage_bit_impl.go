package solutions

// StorageBitImpl ...
type StorageBitImpl struct {
	data   []byte
	length int
}

// NewStorageBoolImpl ...
func NewStorageBoolImpl(size int) *StorageBitImpl {
	byteLength := size / 8
	if size%8 > 0 {
		byteLength++
	}
	result := &StorageBitImpl{
		data:   make([]byte, byteLength, byteLength),
		length: size}
	return result
}

// Flip ...
func (storage *StorageBitImpl) Flip(index int) {
	pos := index / 8
	off := uint8(7 - (index % 8))
	mask := uint8(1) << off
	num := uint8(storage.data[pos])

	bitValue := num&mask == mask
	if bitValue {
		storage.SetData(index, false)
	} else {
		storage.SetData(index, true)
	}
}

// SetData ...
func (storage *StorageBitImpl) SetData(idx int, val bool) {
	bidx := idx / 8
	num := uint8(storage.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		//is set
		if !val {
			//so unset it
			storage.data[bidx] = num ^ mask
		}
	} else {
		//not set
		if val {
			//so set it
			storage.data[bidx] = num | mask
		}
	}

}

// GetData ...
func (storage *StorageBitImpl) GetData(idx int) bool {
	bidx := idx / 8
	num := uint8(storage.data[bidx])
	off := uint8(7 - (idx % 8))
	mask := uint8(1) << off
	if num&mask == mask {
		return true
	}
	return false
}

// GetCount ...
func (storage *StorageBitImpl) GetCount() int {
	return 0
}

// FillInPrimes ...
func (storage *StorageBitImpl) FillInPrimes() []int {
	// result := make([]int, 0, storage.GetCount())
	// for i, value := range storage.data {
	// 	if value {
	// 		result = append(result, i)
	// 	}
	// }
	// return result

	return []int{}
}
