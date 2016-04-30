package main

import (
	"reflect"
	"testing"
)

func TestRecord(t *testing.T) {

	lst := []uint32{1, 2, 3}
	var rs *result

	rs = &result{"a", "b", lst, lst}
	rs.Record(1, 3)
	if !reflect.DeepEqual(rs.idxlist1, []uint32{1, 2, 3, 4}) {
		t.Error("The first list not recorded.")
	}
	if !reflect.DeepEqual(rs.idxlist2, lst) {
		t.Error("The second list not remained.")
	}

	rs = &result{"a", "b", lst, lst}
	rs.Record(2, 3)
	if !reflect.DeepEqual(rs.idxlist1, lst) {
		t.Error("The first list not remained.")
	}

	if !reflect.DeepEqual(rs.idxlist2, []uint32{1, 2, 3, 4}) {
		t.Error("The second list not recorded.")
	}

}

func TestRecordNone(t *testing.T) {

	lst := []uint32{1, 2, 3}
	var rs *result

	rs = &result{"a", "b", lst, lst}
	rs.RecordNone(1)
	if !reflect.DeepEqual(rs.idxlist1, []uint32{1, 2, 3, 0}) {
		t.Error("The first list not recorded.")
	}
	if !reflect.DeepEqual(rs.idxlist2, lst) {
		t.Error("The second list not remained.")
	}

	rs = &result{"a", "b", lst, lst}
	rs.RecordNone(2)
	if !reflect.DeepEqual(rs.idxlist1, lst) {
		t.Error("The first list not recorded.")
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{1, 2, 3, 0}) {
		t.Error("The second list not remained.")
	}

}

func TestDiff(t *testing.T) {

	var rs *result

	h1 := []hashelem{
		hashelem{0, 1},
		hashelem{1, 2},
		hashelem{2, 3},
	}

	h2 := []hashelem{
		hashelem{0, 1},
		hashelem{1, 2},
		hashelem{2, 4},
	}

	rs = &result{"a", "b", []uint32{}, []uint32{}}
	rs.Diff(h1, 0, uint32(len(h1)), h2, 0, uint32(len(h2)))

	if !reflect.DeepEqual(rs.idxlist1, []uint32{3}) {
		t.Error("The first result is not expected.", rs.idxlist1)
	}
	if !reflect.DeepEqual(rs.idxlist2, []uint32{3}) {
		t.Error("The second result is not expected.", rs.idxlist2)
	}

}
