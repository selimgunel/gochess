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
		number uint64
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
		number uint64
		want   string
	}{
		{
			number: 0,
			want:   "0000000000000000000000000000000000000000000000000000000000000000",
		},
		{
			number: 128,
			want:   "0000000000000000000000000000000000000000000000000000000010000000",
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
