package solutions

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCoordinate_Sum(t *testing.T) {
	type args struct {
		a Coordinate
		b Coordinate
	}
	tests := []struct {
		name string
		c    Coordinate
		args args
		want Coordinate
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Sum(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Coordinate.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangeleData_onBoundary(t *testing.T) {
	type args struct {
		s Coordinate
	}
	tests := []struct {
		name string
		data *TriangeleData
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.onBoundary(tt.args.s); got != tt.want {
				t.Errorf("TriangeleData.onBoundary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangeleData_inRange(t *testing.T) {
	type args struct {
		s Coordinate
	}
	tests := []struct {
		name string
		data *TriangeleData
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.inRange(tt.args.s); got != tt.want {
				t.Errorf("TriangeleData.inRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangeleData_iterate(t *testing.T) {
	type args struct {
		input map[Coordinate]*big.Int
	}
	tests := []struct {
		name string
		data *TriangeleData
		args args
		want map[Coordinate]*big.Int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.iterate(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TriangeleData.iterate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTriangeleData_loadData(t *testing.T) {
	tests := []struct {
		name    string
		data    *TriangeleData
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.data.loadData(); (err != nil) != tt.wantErr {
				t.Errorf("TriangeleData.loadData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTriangeleData_do(t *testing.T) {
	tests := []struct {
		name string
		data *TriangeleData
		want *big.Int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.data.do(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TriangeleData.do() = %v, want %v", got, tt.want)
			}
		})
	}
}
