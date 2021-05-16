package structure

import (
	"bytes"
	"log"
	"testing"
)

func setLogBuffer(b *bytes.Buffer) {
	log.SetOutput(b)
	log.SetPrefix("")
	log.SetFlags(0)
}
func TestNewLinkListAndPrint(t *testing.T) {
	intValuesList := [][]int{
		{1, 2, 3, 4, 5},
		{1},
	}
	trueStr := "1->2->3->4->5\n1\n"
	var b bytes.Buffer
	setLogBuffer(&b)
	for _, values := range intValuesList {
		l, _ := NewIntList(values)
		l.Print()
	}
	out := b.String()
	if trueStr != out {
		t.Errorf("error\nexpect:%s\nactual:%s\n", trueStr, out)
	}
}

func TestMergeTwoList(t *testing.T) {
	intValuesList := [][][]int{
		{
			{1, 2, 3}, {1, 2, 3},
		},
		{
			{1, 3, 7}, {2, 4, 6},
		},
	}
	trueStr := "1->1->2->2->3->3\n1->2->3->4->6->7\n"
	var b bytes.Buffer
	setLogBuffer(&b)
	var l *ListNode
	for _, values := range intValuesList {
		l1, _ := NewIntList(values[0])
		l2, _ := NewIntList(values[1])
		l = MergeTwoIntList2(l1, l2)
		l.Print()
	}
	out := b.String()
	if trueStr != out {
		t.Errorf("error\nexpect:%s\nactual:%s\n", trueStr, out)
	}
}
