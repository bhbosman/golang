package solutions

// StorageBoolImpl ...
type StorageBoolImpl struct {
	data  []bool
	count int
}

// Flip ...
func (storage *StorageBoolImpl) Flip(index int) {
	storage.data[index] = !storage.data[index]
	if storage.data[index] {
		storage.count++
	} else {
		storage.count--
	}
}

// SetData ...
func (storage *StorageBoolImpl) SetData(index int, value bool) {
	before := storage.data[index]
	storage.data[index] = value
	after := storage.data[index]
	if before != after {
		if after {
			storage.count++
		} else {
			storage.count--
		}
	}
}

// GetData ...
func (storage *StorageBoolImpl) GetData(index int) bool {
	return storage.data[index]
}

// GetCount ...
func (storage *StorageBoolImpl) GetCount() int {
	return storage.count
}

// FillInPrimes ...
func (storage *StorageBoolImpl) FillInPrimes() []int {
	result := make([]int, 0, storage.GetCount())
	for i, value := range storage.data {
		if value {
			result = append(result, i)
		}
	}
	return result
}
