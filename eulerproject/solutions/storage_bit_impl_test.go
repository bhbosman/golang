package solutions

import "testing"

func TestStorage_bit_impl(t *testing.T) {
	data := NewStorageBoolImpl(8)
	if len(data.data) != 1 {
		t.Fatalf("length failed. Length should be 1, but is is %d", data.length)
	}
	data.SetData(0, true)
	if data.data[0] != 0x01 {
		t.Fatalf("index 0. %d", data.data[0])
	}
	data.SetData(1, true)
	if data.data[1] != 0x03 {
		t.Fatal("index 1")
	}
	data.SetData(2, true)
	if data.data[0] != 0x07 {
		t.Fatal("index 2")
	}
	data.SetData(3, true)
	if data.data[0] != 0x0F {
		t.Fatal("index 3")
	}
	data.SetData(4, true)
	if data.data[0] != 0x1F {
		t.Fatal("index 4")
	}
	data.SetData(5, true)
	if data.data[0] != 0x3F {
		t.Fatal("index 5")
	}
	data.SetData(6, true)
	if data.data[0] != 0x7F {
		t.Fatal("index 6")
	}
	data.SetData(7, true)
	if data.data[0] != 0xFF {
		t.Fatal("index 7")
	}
}
