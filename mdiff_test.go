package main

import (
	"reflect"
	"testing"
)

func TestRecord(t *testing.T) {

	lst := []uint32{1, 2, 3}
	var rs *result

	rs = &result{"a", "b", lst, lst}
	rs.Record(1, 4)
	if !reflect.DeepEqual(rs.idxlist1, []uint32{1, 2, 3, 4}) {
		t.Error("The first list not recorded.")
	}
	if !reflect.DeepEqual(rs.idxlist2, lst) {
		t.Error("The second list not remained.")
	}

	rs = &result{"a", "b", lst, lst}
	rs.Record(2, 4)
	if !reflect.DeepEqual(rs.idxlist1, lst) {
		t.Error("The first list not remained.")
	}

	if !reflect.DeepEqual(rs.idxlist2, []uint32{1, 2, 3, 4}) {
		t.Error("The second list not recorded.")
	}

}

func TestDiff1(t *testing.T) {

	var rs *result

	h1 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 30},
	}

	rs = &result{"a", "b", []uint32{}, []uint32{}}
	rs.Diff(h1, 0, uint32(len(h1)), h1, 0, uint32(len(h1)))

	if !reflect.DeepEqual(rs.idxlist1, []uint32{}) {
		t.Error("The first result is not expected.", rs.idxlist1)
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{}) {
		t.Error("The second result is not expected.", rs.idxlist2)
	}

}

func TestDiff2(t *testing.T) {

	var rs *result

	h1 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 30},
	}

	h2 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 40},
	}

	rs = &result{"a", "b", []uint32{}, []uint32{}}
	rs.Diff(h1, 0, uint32(len(h1)), h2, 0, uint32(len(h2)))

	if !reflect.DeepEqual(rs.idxlist1, []uint32{2}) {
		t.Error("The first result is not expected.", rs.idxlist1)
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{2}) {
		t.Error("The second result is not expected.", rs.idxlist2)
	}

}

func TestDiff3(t *testing.T) {

	var rs *result

	h1 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 30},
	}

	h2 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 50},
		hashelem{2, 30},
	}

	rs = &result{"a", "b", []uint32{}, []uint32{}}
	rs.Diff(h1, 0, uint32(len(h1)), h2, 0, uint32(len(h2)))

	if !reflect.DeepEqual(rs.idxlist1, []uint32{1}) {
		t.Error("The first result is not expected.", rs.idxlist1)
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{1}) {
		t.Error("The second result is not expected.", rs.idxlist2)
	}

}

func TestDiff4(t *testing.T) {

	var rs *result

	h1 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 30},
	}

	h2 := []hashelem{
		hashelem{0, 10},
		hashelem{1, 20},
		hashelem{2, 30},
		hashelem{3, 40},
		hashelem{4, 50},
	}

	rs = &result{"a", "b", []uint32{}, []uint32{}}
	rs.Diff(h1, 0, uint32(len(h1)), h2, 0, uint32(len(h2)))

	if !reflect.DeepEqual(rs.idxlist1, []uint32{}) {
		t.Error("The first result is not expected.", rs.idxlist1)
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{3, 4}) {
		t.Error("The second result is not expected.", rs.idxlist2)
	}

}

func TestSortByHash(t *testing.T) {

	h1 := hashelemlist{
		hashelem{0, 20},
		hashelem{1, 10},
		hashelem{2, 40},
		hashelem{3, 30},
	}

	sortByHash(&h1)

	a1 := hashelemlist{
		hashelem{1, 10},
		hashelem{0, 20},
		hashelem{3, 30},
		hashelem{2, 40},
	}

	if !reflect.DeepEqual(h1, a1) {
		t.Error("Not sorted by hash properly.")
	}

}
