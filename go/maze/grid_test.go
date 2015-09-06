package maze

import (
	"reflect"
	"testing"
)

func TestPrepareGrid(t *testing.T) {
	want := [][]*Cell{{NewCell(0, 0)}}

	got := prepareGrid(1, 1)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}

func TestEachCell(t *testing.T) {
	want := []*Cell{NewCell(0, 0), NewCell(0, 1), NewCell(1, 0), NewCell(1, 1)}

	grid := &NormalGrid{
		rows: 2,
		cols: 2,
		grid: prepareGrid(2, 2),
	}
	got := grid.EachCell()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %#v, want: %#v", got, want)
	}
}

func TestNewNormalGrid(t *testing.T) {
	nw, ne, sw, se := NewCell(0, 0), NewCell(0, 1), NewCell(1, 0), NewCell(1, 1)
	nw.East, nw.South = ne, sw
	ne.West, ne.South = nw, se
	sw.East, sw.North = se, nw
	se.West, se.North = sw, ne
	want := &NormalGrid{
		rows: 2,
		cols: 2,
		grid: [][]*Cell{{nw, ne}, {sw, se}},
	}

	got := NewNormalGrid(2, 2)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\n got: %v\nwant: %v", got, want)
	}
}

func TestString(t *testing.T) {
	want := `
+---+---+
|   |   |
+   +   +
|       |
+---+---+
`

	grid := NewNormalGrid(2, 2)
	nw, ne, sw, se := grid.Get(0, 0), grid.Get(0, 1), grid.Get(1, 0), grid.Get(1, 1)
	nw.Link(sw)
	sw.Link(se)
	se.Link(ne)
	got := "\n" + grid.String()

	if got != want {
		t.Errorf("\n got: %s\nwant: %s\n", got, want)
	}
}
