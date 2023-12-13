package day3

import (
	"reflect"
	"testing"
)

func init() {
	// need to setup the symbols variable
	symbols = append(symbols, []bool{false, false, false, false, false, false, false, false})
	symbols = append(symbols, []bool{false, false, false, true, false, false, false, false})
	symbols = append(symbols, []bool{false, false, false, false, false, false, false, false})
}

func TestXcoord(t *testing.T) {
	var want coord
	want.x = 3
	want.y = 2
	chunk := []int{4, 6}
	line := 3
	got := Xcoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestXcoordLineOne(t *testing.T) {
	var want coord
	want.x = 3
	want.y = 0
	chunk := []int{4, 6}
	line := 0
	got := Xcoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestXcoordFinalLine(t *testing.T) {
	var want coord
	want.x = 3
	want.y = 138
	chunk := []int{4, 6}
	line := DataExtent
	got := Xcoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestXcoordFirstCol(t *testing.T) {
	var want coord
	want.x = 0
	want.y = 138
	chunk := []int{0, 2}
	line := DataExtent
	got := Xcoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestYcoord(t *testing.T) {
	var want coord
	want.x = 7
	want.y = 4
	chunk := []int{4, 6}
	line := 3
	got := Ycoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestYcoordFirst(t *testing.T) {
	var want coord
	want.x = 3
	want.y = 4
	chunk := []int{0, 2}
	line := 3
	got := Ycoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestYcoordLast(t *testing.T) {
	var want coord
	want.x = 139
	want.y = 4
	chunk := []int{137, 139}
	line := 3
	got := Ycoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestYcoordLastCol(t *testing.T) {
	var want coord
	want.x = 139
	want.y = 139
	chunk := []int{137, 139}
	line := DataExtent
	got := Ycoord(chunk, line)

	if got.x != want.x {
		t.Errorf("X Coordinate got %d, want %d", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("Y Coordinate got %d, want %d", got.y, want.y)
	}
}

func TestCharType(t *testing.T) {
	types := []rune(".3*")
	want := []int{CharBlank, CharNum, CharSymbol}

	for i, r := range types {
		got := CharType(r)
		if got != want[i] {
			t.Errorf("CharType %c got %d, want %d", r, got, want[i])
		}
	}
}

func TestIsChunkNextToSymbolTrue(t *testing.T) {
	var x, y coord

	x.x = 3
	x.y = 0
	y.x = 5
	y.y = 2

	got := IsChunkNextToSymbol(x, y)

	want := true
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestIsChunkNextToSymbolFalse(t *testing.T) {
	var x, y coord

	x.x = 5
	x.y = 0
	y.x = 7
	y.y = 2

	got := IsChunkNextToSymbol(x, y)

	want := false
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func TestLineNumberChunks(t *testing.T) {
	line := "......456..789..."
	want := [][]int{{6, 9}, {11, 14}}
	got := LineNumberChunks(line)

	if reflect.DeepEqual(got, want) != true {
		t.Errorf("got %v, want %v", got, want)
	}
}