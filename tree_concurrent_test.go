package bourke

import (
	"sync"
	"testing"
)

//const size uint32 = 100

//func printDurationsInNanos(durations []int64) {
//	sort.Slice(durations, func(i, j int) bool {
//		return durations[i] < durations[j]
//	})
//	percentile := durations[int(float64(size)*0.99)]
//	var sum uint32
//	for _, duration := range durations {
//		sum = sum + uint32(duration)
//	}
//	average := sum / size
//	maximum := durations[size-1]
//	minimum := durations[0]
//
//	fmt.Printf("sample size: %d\nin nanoseconds\n", size)
//	fmt.Printf("99th p.: %d\n", percentile)
//	fmt.Printf("average: %d\n", average)
//	fmt.Printf("maximum: %d\n", maximum)
//	fmt.Printf("minimum: %d\n", minimum)
//}

//func Test_TreeConcurrent_Perf_1_Shard(t *testing.T) {
//	tree := NewTreeConcurrent[uint32, uint64](uint32(1), FixedLengthKeyHasher)
//
//	putDurations := make([]int64, size)
//	var wg sync.WaitGroup
//	wg.Add(int(size))
//
//	var i uint32 = 0
//	for ; i < size; i++ {
//		go func(iteration uint32) {
//			defer wg.Done()
//			beginPut := time.Now().UnixNano()
//			err := tree.Put(iteration, uint64(iteration))
//			endPut := time.Now().UnixNano()
//			if err != nil {
//				panic(err.Error())
//			}
//			putDurations[iteration] = endPut - beginPut
//		}(i)
//	}
//	wg.Wait()
//
//	printDurationsInNanos(putDurations)
//}

//func Test_TreeConcurrent_Perf_8_Shards(t *testing.T) {
//	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)
//
//	putDurations := make([]int64, size)
//	var wg sync.WaitGroup
//	wg.Add(int(size))
//
//	var i uint32 = 0
//	for ; i < size; i++ {
//		go func(iteration uint32) {
//			defer wg.Done()
//			beginPut := time.Now().UnixNano()
//			err := tree.Put(iteration, uint64(iteration))
//			endPut := time.Now().UnixNano()
//			if err != nil {
//				panic(err.Error())
//			}
//			putDurations[iteration] = endPut - beginPut
//		}(i)
//	}
//	wg.Wait()
//
//	printDurationsInNanos(putDurations)
//}

//func Test_TreeConcurrent_Perf_8_Shards_Seq(t *testing.T) {
//	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)
//
//	putDurations := make([]int64, size)
//	var wg sync.WaitGroup
//	wg.Add(int(size))
//
//	var i uint32 = 0
//	for ; i < size; i++ {
//		func(iteration uint32) {
//			defer wg.Done()
//			beginPut := time.Now().UnixNano()
//			err := tree.Put(iteration, uint64(iteration))
//			endPut := time.Now().UnixNano()
//			if err != nil {
//				panic(err.Error())
//			}
//			putDurations[iteration] = endPut - beginPut
//		}(i)
//	}
//	wg.Wait()
//
//	printDurationsInNanos(putDurations)
//}

//func Test_TreeConcurrent_Perf_32_Shards(t *testing.T) {
//	tree := NewTreeConcurrent[uint32, uint64](uint32(32), FixedLengthKeyHasher)
//
//	putDurations := make([]int64, size)
//	var wg sync.WaitGroup
//	wg.Add(int(size))
//
//	var i uint32 = 0
//	for ; i < size; i++ {
//		go func(iteration uint32) {
//			defer wg.Done()
//			beginPut := time.Now().UnixNano()
//			tree.Put(iteration, uint64(iteration))
//			endPut := time.Now().UnixNano()
//			putDurations[iteration] = endPut - beginPut
//		}(i)
//	}
//	wg.Wait()
//
//	printDurationsInNanos(putDurations)
//}

func Test_TreeConcurrent_Get(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)
	var size uint32 = 8
	var wg sync.WaitGroup
	wg.Add(int(size))
	var i uint32 = 0
	for ; i < size; i++ {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(i)
	}
	wg.Wait()

	if result, _ := tree.Get(0); result != 0 {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
	if result, _ := tree.Get(1); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if result, _ := tree.Get(2); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}
	if result, _ := tree.Get(3); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if result, _ := tree.Get(4); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if result, _ := tree.Get(5); result != 5 {
		t.Errorf("actual: %d, expected: %d", result, 5)
	}
	if result, _ := tree.Get(6); result != 6 {
		t.Errorf("actual: %d, expected: %d", result, 6)
	}
	if result, _ := tree.Get(7); result != 7 {
		t.Errorf("actual: %d, expected: %d", result, 7)
	}
	if result, err := tree.Get(8); err == nil {
		t.Errorf("should return not_found %d", result)
	}
}

func Test_TreeConcurrent_Successor_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	if _, _, err := tree.Successor(1); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Successor_Single(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	tree.Put(8, 8)

	if _, result, _ := tree.Successor(1); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Successor(7); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, _, err := tree.Successor(8); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Successor(16); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Successor(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{8, 9, 11, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	if _, result, _ := tree.Successor(1); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Successor(2); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Successor(7); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Successor(8); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Successor(9); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Successor(10); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Successor(11); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, _, err := tree.Successor(12); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Successor(13); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Ceiling_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	if _, _, err := tree.Ceiling(1); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Ceiling_Single(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	tree.Put(8, 8)

	if _, result, _ := tree.Ceiling(1); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(7); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, _, err := tree.Ceiling(16); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Ceiling(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{8, 9, 11, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	if _, result, _ := tree.Ceiling(1); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(2); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(7); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Ceiling(10); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Ceiling(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Ceiling(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, _, err := tree.Successor(13); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Floor_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	if _, _, err := tree.Floor(1); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Floor_Single(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	_ = tree.Put(8, 8)

	if _, result, _ := tree.Floor(20); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Floor(9); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Floor(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, _, err := tree.Floor(4); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Floor(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{8, 9, 11, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			_ = tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	if _, _, err := tree.Floor(1); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Floor(7); err == nil {
		t.Errorf("should return not_found")
	}
	if _, result, _ := tree.Floor(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Floor(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Floor(10); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Floor(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Floor(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, result, _ := tree.Floor(20); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
}

func Test_TreeConcurrent_Predecessor_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	if _, _, err := tree.Predecessor(20); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_TreeConcurrent_Predecessor_Single(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	tree.Put(8, 8)

	if _, _, err := tree.Predecessor(1); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Predecessor(8); err == nil {
		t.Errorf("should return not_found")
	}
	if _, result, _ := tree.Predecessor(9); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Predecessor(20); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
}

func Test_TreeConcurrent_Predecessor(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{8, 9, 11, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			err := tree.Put(iteration, uint64(iteration))
			if err != nil {
				panic(err.Error())
			}
		}(key)
	}
	wg.Wait()

	if _, _, err := tree.Predecessor(1); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Predecessor(2); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Predecessor(8); err == nil {
		t.Errorf("should return not_found")
	}
	if _, result, _ := tree.Predecessor(9); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Predecessor(10); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Predecessor(11); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Predecessor(12); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Predecessor(13); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, result, _ := tree.Predecessor(20); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
}

func Test_TreeConcurrent_All_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	for range tree.All() {
		t.Errorf("should not iterate")
	}
}

func Test_TreeConcurrent_All(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.All() {
		results[key] = value
	}

	if len(results) != 9 {
		t.Errorf("wrong size")
	}
	if results[1] != 1 {
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 2 {
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[4] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[5] != 5 {
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 9 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 11 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 12 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 13 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_GreaterThan_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	for range tree.GreaterThan(0, true) {
		t.Errorf("should not iterate")
	}
}

func Test_TreeConcurrent_GreaterThan_Inclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.GreaterThan(6, true) {
		results[key] = value
	}

	if len(results) != 6 {
		t.Errorf("wrong size")
	}
	if results[1] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 9 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 11 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 12 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 13 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_GreaterThan_Exclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.GreaterThan(6, false) {
		results[key] = value
	}

	if len(results) != 5 {
		t.Errorf("wrong size")
	}
	if results[1] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 9 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 11 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 12 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 13 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_LessThan_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	for range tree.LessThan(8, true) {
		t.Errorf("should not iterate")
	}
}

func Test_TreeConcurrent_LessThan_Inclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.LessThan(8, true) {
		results[key] = value
	}

	if len(results) != 5 {
		t.Errorf("wrong size")
	}
	if results[1] != 1 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 2 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 5 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 0 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 { // zero initial value 10 is not contained
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 0 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 0 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 0 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_LessThan_Exclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.LessThan(8, false) {
		results[key] = value
	}

	if len(results) != 4 {
		t.Errorf("wrong size")
	}
	if results[1] != 1 {
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 2 {
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 {
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 {
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 5 {
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 {
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 0 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 0 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 {
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 0 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 0 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 0 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_Between_Empty(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	for range tree.Between(8, true, 100, true) {
		t.Errorf("should not iterate")
	}
}

func Test_TreeConcurrent_Between_Inclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.Between(5, true, 11, true) {
		results[key] = value
	}

	if len(results) != 5 {
		t.Errorf("wrong size")
	}
	if results[1] != 0 {
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 0 {
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 {
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 {
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 5 {
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 {
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 9 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 {
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 11 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 0 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 0 {
		t.Errorf("actual: %d", results[13])
	}
}

func Test_TreeConcurrent_Between_Exclusive(t *testing.T) {
	tree := NewTreeConcurrent[uint32, uint64](uint32(8), FixedLengthKeyHasher)

	var wg sync.WaitGroup
	keys := []uint32{13, 2, 5, 11, 1, 6, 8, 9, 12}
	wg.Add(len(keys))
	for _, key := range keys {
		go func(iteration uint32) {
			defer wg.Done()
			tree.Put(iteration, uint64(iteration))
		}(key)
	}
	wg.Wait()

	results := make(map[uint32]uint64)
	for key, value := range tree.Between(2, false, 12, false) {
		results[key] = value
	}

	if len(results) != 5 {
		t.Errorf("wrong size")
	}
	if results[1] != 0 {
		t.Errorf("actual: %d", results[1])
	}
	if results[2] != 0 {
		t.Errorf("actual: %d", results[2])
	}
	if results[3] != 0 {
		t.Errorf("actual: %d", results[3])
	}
	if results[4] != 0 {
		t.Errorf("actual: %d", results[4])
	}
	if results[5] != 5 {
		t.Errorf("actual: %d", results[5])
	}
	if results[6] != 6 {
		t.Errorf("actual: %d", results[6])
	}
	if results[7] != 0 {
		t.Errorf("actual: %d", results[7])
	}
	if results[8] != 8 {
		t.Errorf("actual: %d", results[8])
	}
	if results[9] != 9 {
		t.Errorf("actual: %d", results[9])
	}
	if results[10] != 0 {
		t.Errorf("actual: %d", results[10])
	}
	if results[11] != 11 {
		t.Errorf("actual: %d", results[11])
	}
	if results[12] != 0 {
		t.Errorf("actual: %d", results[12])
	}
	if results[13] != 0 {
		t.Errorf("actual: %d", results[13])
	}
}
