package pgn

import (
	"fmt"
	"testing"
)

func BenchmarkString(b *testing.B) {

	for i := 0; i < b.N; i++ {

		bb := Bitboard(uint64(i))
		_ = bb.String()
	}

}

func TestString(t *testing.T) {
	testCases := []struct {
		number uint64 //the number which we want to have a bit representation
		want   string
	}{
		{
			number: 0,
			want:   "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			number: 1,
			want:   "0000000000000000000000000000000000000000000000000000000000000001",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", tc.number), func(t *testing.T) {
			bb := Bitboard(tc.number)
			if bb.String() != tc.want {
				t.Fatalf("want: %s got: %s ", tc.want, bb.String())
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		number uint64 //the number which we want to have a bit representation
		want   string
	}{
		{
			number: 0,
			want:   "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			number: 1,
			want:   "1000000000000000000000000000000000000000000000000000000000000000",
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("[%d]", tc.number), func(t *testing.T) {
			bb := Bitboard(tc.number).Reverse()
			if bb.String() != tc.want {
				t.Fatalf("want: %s got: %s ", tc.want, bb.String())
			}
		})
	}
}

func TestOccupied(t *testing.T) {
	testCases := []struct {
		number     uint64 //the number which we want to have a bit representation
		index      Square //the position within the bit representation
		isOccupied bool   //whether the bit is 1
	}{
		{
			number:     0,
			index:      0,
			isOccupied: false,
		},
		{
			number:     1,
			index:      0,
			isOccupied: true,
		},
		{
			number:     1,
			index:      2,
			isOccupied: false,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("numbers [%d]", tc.number), func(t *testing.T) {
			bb := Bitboard(tc.number)
			if bb.Occupied(tc.index) != tc.isOccupied {
				t.Fatalf("given number: %d and index %d want: %t got: %t ", tc.number, tc.index, bb.Occupied(tc.index), tc.isOccupied)
			}
		})
	}
}
