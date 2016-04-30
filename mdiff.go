package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"sort"
)

type hashelem struct {
	idx  uint32
	hash uint64
}

type hashelemlist []hashelem

func (a hashelemlist) Len() int           { return len(a) }
func (a hashelemlist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a hashelemlist) Less(i, j int) bool { return a[i].hash < a[j].hash }

type result struct {
	fname1   string
	fname2   string
	idxlist1 []uint32
	idxlist2 []uint32
}

func (rs *result) record(n uint32, idx uint32) {
	if n == 1 {
		rs.idxlist1 = append(rs.idxlist1, idx+1)
	} else {
		rs.idxlist2 = append(rs.idxlist2, idx+1)
	}
}

func (rs *result) recordnone(n uint32) {
	if n == 1 {
		rs.idxlist1 = append(rs.idxlist1, 0)
	} else {
		rs.idxlist2 = append(rs.idxlist2, 0)
	}
}

func (rs *result) diff(hashes1 hashelemlist, idx1 uint32, size1 uint32, hashes2 hashelemlist, idx2 uint32, size2 uint32) {

	if idx1 == size1 && idx2 < size2 {
		rs.recordnone(1)
		rs.record(2, hashes2[idx2].idx)
		rs.diff(hashes1, idx1, size1, hashes2, idx2+1, size2)
	}

	if idx1 < size1 && idx2 == size2 {
		rs.record(1, hashes1[idx1].idx)
		rs.recordnone(2)
		rs.diff(hashes1, idx1+1, size1, hashes2, idx2, size2)
	}

	if idx1 == size1 || idx2 == size2 {
		return
	}

	if hashes1[idx1].hash == hashes2[idx2].hash {
		rs.diff(hashes1, idx1+1, size1, hashes2, idx2+1, size2)
	} else {
		if (size1 - idx1) < (size2 - idx2) {
			rs.recordnone(1)
			rs.record(2, hashes2[idx2].idx)
			rs.diff(hashes1, idx1, size1, hashes2, idx2+1, size2)
		} else if (size1 - idx1) > (size2 - idx2) {
			rs.record(1, hashes1[idx1].idx)
			rs.recordnone(2)
			rs.diff(hashes1, idx1+1, size1, hashes2, idx2, size2)
		} else {
			rs.record(1, hashes1[idx1].idx)
			rs.record(2, hashes2[idx2].idx)
			rs.diff(hashes1, idx1+1, size1, hashes2, idx2+1, size2)
		}
	}

}

func (rs *result) stat(n uint32, fd *os.File) {

	var fname string
	var idxlist []uint32
	if n == 1 {
		fname = rs.fname1
		idxlist = rs.idxlist1
	} else {
		fname = rs.fname2
		idxlist = rs.idxlist2
	}
	fmt.Println(fname)

	scanner := bufio.NewScanner(fd)
	var idx = 0
	var line uint32 = 1 // 0 is reserved for none
	for scanner.Scan() {
		str := scanner.Text()
		if line == idxlist[idx] {
			fmt.Printf("%d: %s\n", line, str)
			idx += 1
			if len(idxlist) == idx {
				break
			}
		}
		line += 1
	}

}

func hashID(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	sum := h.Sum64()
	return sum
}

func hashLines(fd *os.File, hashes *hashelemlist, ch chan int) {
	scanner := bufio.NewScanner(fd)
	var idx uint32 = 0
	for scanner.Scan() {
		str := scanner.Text()
		*hashes = append(*hashes, hashelem{idx, hashID(str)})
		idx += 1
	}
	ch <- 0
}

func runDiff(rs *result) {

	fd1, err := os.Open(rs.fname1)
	if err != nil {
		os.Exit(-1)
	}
	defer fd1.Close()

	fd2, err := os.Open(rs.fname2)
	if err != nil {
		os.Exit(-1)
	}
	defer fd2.Close()

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	hashes1 := make(hashelemlist, 0, 10000)
	go hashLines(fd1, &hashes1, ch1)

	hashes2 := make(hashelemlist, 0, 10000)
	go hashLines(fd2, &hashes2, ch2)

	<-ch1
	<-ch2

	sort.Sort(hashes1)
	sort.Sort(hashes2)

	size1 := uint32(len(hashes1))
	size2 := uint32(len(hashes2))

	rs.diff(hashes1, 0, size1, hashes2, 0, size2)

}

func runStat(rs *result) {

	fd1, err := os.Open(rs.fname1)
	if err != nil {
		os.Exit(-1)
	}
	defer fd1.Close()

	fd2, err := os.Open(rs.fname2)
	if err != nil {
		os.Exit(-1)
	}
	defer fd2.Close()

	rs.stat(1, fd1)
	rs.stat(2, fd2)

}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Not enough arguments.")
		os.Exit(-1)
	}

	f1 := os.Args[1]
	f2 := os.Args[2]

	rs := &result{
		f1,
		f2,
		make([]uint32, 0, 1000),
		make([]uint32, 0, 1000),
	}

	runDiff(rs)
	runStat(rs)
}