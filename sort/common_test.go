package sort

import (
	"testing"

	"github.com/jcuan/somecode/structure"
)

// 应该每种排序针对性的测试的，就先copy吧。。。。

func TestShell(t *testing.T) {
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 86, 25}
	ShellSort(inputs)
	rights := []int{8, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	for i := range inputs {
		if inputs[i] != rights[i] {
			t.Errorf("fail expect:%v, actual:%v\n", rights, inputs)
		}
	}
}

func TestBinaryInsert(t *testing.T) {
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 86, 25}
	BinaryInsert(inputs)
	rights := []int{8, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	for i := range inputs {
		if inputs[i] != rights[i] {
			t.Errorf("fail expect:%v, actual:%v\n", rights, inputs)
			break
		}
	}
}

func TestBubble(t *testing.T) {
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 86, 25}
	Bubble(inputs)
	rights := []int{8, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	for i := range inputs {
		if inputs[i] != rights[i] {
			t.Errorf("fail expect:%v, actual:%v\n", rights, inputs)
			break
		}
	}
}

func TestQuick(t *testing.T) {
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 12, 86, 25, 12}
	Quik(inputs, 0, len(inputs)-1)
	rights := []int{8, 12, 12, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	for i := range inputs {
		if inputs[i] != rights[i] {
			t.Errorf("fail expect:%v, actual:%v\n", rights, inputs)
			break
		}
	}
}

func TestHeap(t *testing.T) {
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 12, 86, 25, 12}
	Heap(inputs)
	rights := []int{8, 12, 12, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	for i := range inputs {
		if inputs[i] != rights[i] {
			t.Fatalf("fail expect:%v, actual:%v\n", rights, inputs)
			break
		}
	}
}

func TestMerge(t *testing.T) {
	// inputs := []int{1, 4, 7, 2, 1}
	inputs := []int{23, 15, 12, 35, 46, 19, 20, 14, 8, 74, 12, 86, 25, 12}
	res := MergeSortSlice(inputs)
	rights := []int{8, 12, 12, 12, 14, 15, 19, 20, 23, 25, 35, 46, 74, 86}
	checkItems(rights, res, t)
	head, _ := structure.NewIntList(inputs)
	result := MergeList(head, len(inputs))
	result.Print()
	p := result
	for i := range rights {
		intVal := p.MustIntVal()
		if intVal != rights[i] {
			p.Print()
			t.Fatal()
		}
		p = p.Next
	}
	if p != nil {
		t.Error("list tail not nil")
	}
}

func checkItems(expect, actual []int, t *testing.T) {
	if len(expect) != len(actual) {
		t.Fatalf("fail expect:%v, actual:%v\n", expect, actual)
	}
	for i := range expect {
		if expect[i] != actual[i] {
			t.Fatalf("fail expect:%v, actual:%v\n", expect, actual)
		}
	}
}
