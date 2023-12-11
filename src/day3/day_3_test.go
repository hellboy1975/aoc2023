package day3

import (
	"testing"
)

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
	want.y = 1
	chunk := []int{4, 6}
	line := 1
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
	want.y = 139
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
	want.y = 139
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
	want.y = 140
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
