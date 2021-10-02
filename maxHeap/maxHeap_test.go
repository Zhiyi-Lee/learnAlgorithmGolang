package maxHeap

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestMaxHeap(t *testing.T) {
	maxHeap := NewMaxHeap()
	fmt.Println(maxHeap.Size())

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 50; i++ {
		maxHeap.Insert(rand.Int() % 100)
	}

	fmt.Println(maxHeap.Size())

	for !maxHeap.IsEmpty() {
		item, has := maxHeap.ExtractMax()
		if has {
			fmt.Print(item, " ")
		}
	}

}
