package bourke

import (
	"testing"
)

func Test_Tree_Get_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	if _, err := tree.Get(1); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
}

func Test_Tree_Get_Single(t *testing.T) {
	tree := NewTree[int, int]()

	tree.Put(8, 8)

	if result, _ := tree.Get(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, err := tree.Get(10); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
}

func Test_Tree_Ceiling_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	if _, _, err := tree.Ceiling(9); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Ceiling_Single(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)

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
	if _, _, err := tree.Ceiling(9); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Ceiling(20); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Ceiling(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)
	_ = tree.Put(9, 9)
	_ = tree.Put(11, 11)
	_ = tree.Put(12, 12)

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
	if _, _, err := tree.Ceiling(13); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Ceiling(20); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Floor_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	if _, _, err := tree.Floor(8); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Floor_Single(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)

	if _, _, err := tree.Floor(1); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Floor(4); err == nil {
		t.Errorf("should return not_found")
	}
	if _, result, _ := tree.Floor(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Floor(9); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Floor(100); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
}

func Test_Tree_Floor(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)
	_ = tree.Put(9, 9)
	_ = tree.Put(11, 11)
	_ = tree.Put(12, 12)

	if _, _, err := tree.Floor(1); err == nil {
		t.Errorf("should return not_found")
	}
	if _, _, err := tree.Floor(4); err == nil {
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
	if _, result, _ := tree.Floor(13); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, result, _ := tree.Floor(20); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
}

func Test_Tree_Successor_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	if _, _, err := tree.Successor(1); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Successor_Single(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)

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

func Test_Tree_Successor(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)
	_ = tree.Put(9, 9)
	_ = tree.Put(11, 11)
	_ = tree.Put(12, 12)

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

func Test_Tree_Predecessor_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	if _, _, err := tree.Predecessor(20); err == nil {
		t.Errorf("should return not_found")
	}
}

func Test_Tree_Predecessor_Single(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)

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

func Test_Tree_Predecessor(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)
	_ = tree.Put(9, 9)
	_ = tree.Put(11, 11)
	_ = tree.Put(12, 12)

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

func Test_Tree_All_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	for range tree.All() {
		t.Errorf("should not iterate")
	}
}

func Test_Tree_All(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(10, 10)
	_ = tree.Put(9, 9)
	_ = tree.Put(12, 12)
	_ = tree.Put(8, 8)
	_ = tree.Put(11, 11)

	result := make(map[int]int)
	for key, value := range tree.All() {
		result[key] = value
	}

	result8 := result[8]
	result9 := result[9]
	result10 := result[10]
	result11 := result[11]
	result12 := result[12]

	if len(result) != 5 {
		t.Errorf("wrong size")
	}
	if result8 != 8 {
		t.Errorf("actual: %d", result8)
	}
	if result9 != 9 {
		t.Errorf("actual: %d", result9)
	}
	if result10 != 10 {
		t.Errorf("actual: %d", result10)
	}
	if result11 != 11 {
		t.Errorf("actual: %d", result11)
	}
	if result12 != 12 {
		t.Errorf("actual: %d", result12)
	}
}

func Test_Tree_GreaterThan(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(11, 11)
	_ = tree.Put(13, 13)
	_ = tree.Put(9, 9)
	_ = tree.Put(12, 12)
	_ = tree.Put(8, 8)

	results10 := make(map[int]int)
	for key, value := range tree.GreaterThan(10, false) {
		results10[key] = value
	}
	results11 := make(map[int]int)
	for key, value := range tree.GreaterThan(11, true) {
		results11[key] = value
	}
	results11Strict := make(map[int]int)
	for key, value := range tree.GreaterThan(11, false) {
		results11Strict[key] = value
	}
	results13 := make(map[int]int)
	for key, value := range tree.GreaterThan(13, true) {
		results13[key] = value
	}
	for range tree.GreaterThan(13, false) {
		t.Errorf("should not iterate")
	}
	for range tree.GreaterThan(20, true) {
		t.Errorf("should not iterate")
	}

	if len(results10) != 3 {
		t.Errorf("wrong size")
	}
	if len(results11) != 3 {
		t.Errorf("wrong size")
	}
	if len(results11Strict) != 2 {
		t.Errorf("wrong size")
	}
	if len(results13) != 1 {
		t.Errorf("wrong size")
	}

	if results10[11] != 11 {
		t.Errorf("actual: %d, expected: %d", results10[11], 11)
	}
	if results10[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results10[12], 12)
	}
	if results10[13] != 13 {
		t.Errorf("actual: %d, expected: %d", results10[11], 13)
	}

	if results11[11] != 11 {
		t.Errorf("actual: %d, expected: %d", results11[11], 11)
	}
	if results11[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results11[12], 12)
	}
	if results11[13] != 13 {
		t.Errorf("actual: %d, expected: %d", results11[11], 13)
	}

	if results11Strict[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results11Strict[12], 12)
	}
	if results11Strict[13] != 13 {
		t.Errorf("actual: %d, expected: %d", results11Strict[11], 13)
	}

	if results13[13] != 13 {
		t.Errorf("actual: %d, expected: %d", results13[13], 13)
	}
}

func Test_Tree_LessThan_Empty(t *testing.T) {
	tree := NewTree[int, int]()

	for range tree.LessThan(10, true) {
		t.Errorf("should not iterate")
	}
}

func Test_Tree_LessThan_Single(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(8, 8)

	results := make(map[int]int)
	for key, value := range tree.LessThan(10, true) {
		results[key] = value
	}

	if results[8] != 8 {
		t.Errorf("actual: %d, expected: %d", results[8], 8)
	}
}

func Test_Tree_LessThan(t *testing.T) {
	tree := NewTree[int, int]()

	_ = tree.Put(11, 11)
	_ = tree.Put(9, 9)
	_ = tree.Put(12, 12)
	_ = tree.Put(8, 8)

	resultInclusive := make(map[int]int)
	for key, value := range tree.LessThan(10, true) {
		resultInclusive[key] = value
	}
	resultExclusive := make(map[int]int)
	for key, value := range tree.LessThan(10, false) {
		resultExclusive[key] = value
	}
	resultGreaterThanUpperBound := make(map[int]int)
	for key, value := range tree.LessThan(12, true) {
		resultGreaterThanUpperBound[key] = value
	}

	for range tree.LessThan(7, true) {
		t.Errorf("should not iterate")
	}
	for range tree.LessThan(8, false) {
		t.Errorf("should not iterate")
	}

	if len(resultInclusive) != 2 {
		t.Errorf("wrong size")
	}
	if len(resultExclusive) != 2 {
		t.Errorf("wrong size")
	}
	if len(resultGreaterThanUpperBound) != 4 {
		t.Errorf("wrong size")
	}

	if resultInclusive[8] != 8 {
		t.Errorf("actual: %d", resultInclusive[8])
	}
	if resultInclusive[9] != 9 {
		t.Errorf("actual: %d", resultInclusive[9])
	}

	if resultExclusive[8] != 8 {
		t.Errorf("actual: %d", resultExclusive[8])
	}
	if resultExclusive[9] != 9 {
		t.Errorf("actual: %d", resultExclusive[9])
	}

	if resultGreaterThanUpperBound[8] != 8 {
		t.Errorf("actual: %d", resultGreaterThanUpperBound[8])
	}
	if resultGreaterThanUpperBound[9] != 9 {
		t.Errorf("actual: %d", resultGreaterThanUpperBound[9])
	}
	if resultGreaterThanUpperBound[11] != 11 {
		t.Errorf("actual: %d", resultGreaterThanUpperBound[11])
	}
	if resultGreaterThanUpperBound[12] != 12 {
		t.Errorf("actual: %d", resultGreaterThanUpperBound[12])
	}
}

func Test_Tree_Between(t *testing.T) {
	tree := newTree[int, int]()

	_ = tree.Put(11, 11)
	_ = tree.Put(8, 8)
	_ = tree.Put(9, 9)
	_ = tree.Put(13, 13)
	_ = tree.Put(12, 12)
	verifyInvariants(tree)

	results10_12 := make(map[int]int)
	for key, value := range tree.Between(10, true, 12, true) {
		results10_12[key] = value
	}
	results7_13 := make(map[int]int)
	for key, value := range tree.Between(7, false, 13, true) {
		results7_13[key] = value
	}
	results8strict_13strict := make(map[int]int)
	for key, value := range tree.Between(8, false, 13, false) {
		results8strict_13strict[key] = value
	}
	results1_10 := make(map[int]int)
	for key, value := range tree.Between(1, true, 10, false) {
		results1_10[key] = value
	}
	for range tree.Between(11, true, 8, true) {
		t.Errorf("should not iterate")
	}
	for range tree.Between(1, true, 7, true) {
		t.Errorf("should not iterate")
	}
	for range tree.Between(1, true, 8, false) {
		t.Errorf("should not iterate")
	}
	for range tree.Between(14, true, 20, false) {
		t.Errorf("should not iterate")
	}

	if len(results10_12) != 2 {
		t.Errorf("wrong size")
	}
	if len(results7_13) != 5 {
		t.Errorf("wrong size")
	}
	if len(results8strict_13strict) != 3 {
		t.Errorf("wrong size")
	}
	if len(results1_10) != 2 {
		t.Errorf("wrong size")
	}

	if results10_12[11] != 11 {
		t.Errorf("actual: %d, expected: %d", results10_12[11], 11)
	}
	if results10_12[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results10_12[12], 12)
	}

	if results7_13[8] != 8 {
		t.Errorf("actual: %d, expected: %d", results7_13[8], 8)
	}
	if results7_13[9] != 9 {
		t.Errorf("actual: %d, expected: %d", results7_13[9], 9)
	}
	if results7_13[11] != 11 {
		t.Errorf("actual: %d, expected: %d", results7_13[11], 11)
	}
	if results7_13[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results7_13[12], 12)
	}
	if results7_13[13] != 13 {
		t.Errorf("actual: %d, expected: %d", results7_13[13], 13)
	}

	if results8strict_13strict[9] != 9 {
		t.Errorf("actual: %d, expected: %d", results8strict_13strict[9], 9)
	}
	if results8strict_13strict[11] != 11 {
		t.Errorf("actual: %d, expected: %d", results8strict_13strict[11], 11)
	}
	if results8strict_13strict[12] != 12 {
		t.Errorf("actual: %d, expected: %d", results8strict_13strict[12], 12)
	}

	if results1_10[8] != 8 {
		t.Errorf("actual: %d, expected: %d", results1_10[8], 8)
	}
	if results1_10[9] != 9 {
		t.Errorf("actual: %d, expected: %d", results1_10[9], 9)
	}
}

func Test_Tree_Put_RandomOrder(t *testing.T) {
	tree := newTree[int, int]()

	_ = tree.Put(9, 9)
	verifyInvariants(tree)
	_ = tree.Put(1, 1)
	verifyInvariants(tree)
	_ = tree.Put(3, 3)
	verifyInvariants(tree)
	_ = tree.Put(7, 7)
	verifyInvariants(tree)
	_ = tree.Put(5, 5)
	verifyInvariants(tree)
	_ = tree.Put(4, 4)
	verifyInvariants(tree)
	_ = tree.Put(2, 2)
	verifyInvariants(tree)
	_ = tree.Put(12, 12)
	verifyInvariants(tree)
	_ = tree.Put(6, 6)
	verifyInvariants(tree)
	_ = tree.Put(10, 10)
	verifyInvariants(tree)
	_ = tree.Put(11, 11)
	verifyInvariants(tree)
	_ = tree.Put(8, 8)
	verifyInvariants(tree)

	if result, _ := tree.Get(1); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if _, result, _ := tree.Ceiling(1); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if _, result, _ := tree.Successor(1); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}

	if result, _ := tree.Get(2); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}
	if _, result, _ := tree.Ceiling(2); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}
	if _, result, _ := tree.Successor(2); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}

	if result, _ := tree.Get(3); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if _, result, _ := tree.Ceiling(3); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if _, result, _ := tree.Successor(3); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}

	if result, _ := tree.Get(4); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if _, result, _ := tree.Ceiling(4); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if _, result, _ := tree.Successor(4); result != 5 {
		t.Errorf("actual: %d, expected: %d", result, 5)
	}

	if result, _ := tree.Get(5); result != 5 {
		t.Errorf("actual: %d, expected: %d", result, 5)
	}
	if _, result, _ := tree.Ceiling(5); result != 5 {
		t.Errorf("actual: %d, expected: %d", result, 5)
	}
	if _, result, _ := tree.Successor(5); result != 6 {
		t.Errorf("actual: %d, expected: %d", result, 6)
	}

	if result, _ := tree.Get(6); result != 6 {
		t.Errorf("actual: %d, expected: %d", result, 6)
	}
	if _, result, _ := tree.Ceiling(6); result != 6 {
		t.Errorf("actual: %d, expected: %d", result, 6)
	}
	if _, result, _ := tree.Successor(6); result != 7 {
		t.Errorf("actual: %d, expected: %d", result, 7)
	}

	if result, _ := tree.Get(7); result != 7 {
		t.Errorf("actual: %d, expected: %d", result, 7)
	}
	if _, result, _ := tree.Ceiling(7); result != 7 {
		t.Errorf("actual: %d, expected: %d", result, 7)
	}
	if _, result, _ := tree.Successor(7); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}

	if result, _ := tree.Get(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Ceiling(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if _, result, _ := tree.Successor(8); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}

	if result, _ := tree.Get(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Ceiling(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if _, result, _ := tree.Successor(9); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}

	if result, _ := tree.Get(10); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}
	if _, result, _ := tree.Ceiling(10); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}
	if _, result, _ := tree.Successor(10); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}

	if result, _ := tree.Get(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Ceiling(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if _, result, _ := tree.Successor(11); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}

	if result, _ := tree.Get(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, result, _ := tree.Ceiling(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, _, err := tree.Successor(12); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}

	if _, err := tree.Get(13); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
}

func Test_Tree_Put_AscendingOrder(t *testing.T) {
	tree := newTree[int, int]()

	_ = tree.Put(1, 1)
	verifyInvariants(tree)
	_ = tree.Put(2, 2)
	verifyInvariants(tree)
	_ = tree.Put(3, 3)
	verifyInvariants(tree)
	_ = tree.Put(4, 4)
	verifyInvariants(tree)
	_ = tree.Put(5, 5)
	verifyInvariants(tree)
	_ = tree.Put(6, 6)
	verifyInvariants(tree)
	_ = tree.Put(7, 7)
	verifyInvariants(tree)
	_ = tree.Put(8, 8)
	verifyInvariants(tree)
	_ = tree.Put(9, 9)
	verifyInvariants(tree)
	_ = tree.Put(10, 10)
	verifyInvariants(tree)
	_ = tree.Put(11, 11)
	verifyInvariants(tree)
	_ = tree.Put(12, 12)
	verifyInvariants(tree)

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
	if result, _ := tree.Get(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if result, _ := tree.Get(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if result, _ := tree.Get(10); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}
	if result, _ := tree.Get(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if result, _ := tree.Get(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, err := tree.Get(13); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
}

func Test_Tree_Remove_AscendingOrder(t *testing.T) {
	tree := newTree[int, int]()

	_ = tree.Put(1, 1)
	verifyInvariants(tree)
	_ = tree.Put(2, 2)
	verifyInvariants(tree)
	_ = tree.Put(3, 3)
	verifyInvariants(tree)
	_ = tree.Put(4, 4)
	verifyInvariants(tree)
	_ = tree.Put(5, 5)
	verifyInvariants(tree)
	_ = tree.Put(6, 6)
	verifyInvariants(tree)
	_ = tree.Put(7, 7)
	verifyInvariants(tree)
	_ = tree.Put(8, 8)
	verifyInvariants(tree)
	_ = tree.Put(9, 9)
	verifyInvariants(tree)
	_ = tree.Put(10, 10)
	verifyInvariants(tree)
	_ = tree.Put(11, 11)
	verifyInvariants(tree)
	_ = tree.Put(12, 12)
	verifyInvariants(tree)

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
	if result, _ := tree.Get(8); result != 8 {
		t.Errorf("actual: %d, expected: %d", result, 8)
	}
	if result, _ := tree.Get(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if result, _ := tree.Get(10); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}
	if result, _ := tree.Get(11); result != 11 {
		t.Errorf("actual: %d, expected: %d", result, 11)
	}
	if result, _ := tree.Get(12); result != 12 {
		t.Errorf("actual: %d, expected: %d", result, 12)
	}
	if _, err := tree.Get(13); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}

	_ = tree.Remove(1)
	verifyInvariants(tree)
	_ = tree.Remove(2)
	verifyInvariants(tree)
	_ = tree.Remove(3)
	verifyInvariants(tree)
	_ = tree.Remove(5)
	verifyInvariants(tree)
	_ = tree.Remove(8)
	verifyInvariants(tree)
	_ = tree.Remove(11)
	verifyInvariants(tree)

	if _, err := tree.Get(1); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
	if _, err := tree.Get(2); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
	if _, err := tree.Get(3); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
	if result, _ := tree.Get(4); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if _, err := tree.Get(5); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
	if result, _ := tree.Get(6); result != 6 {
		t.Errorf("actual: %d, expected: %d", result, 6)
	}
	if result, _ := tree.Get(7); result != 7 {
		t.Errorf("actual: %d, expected: %d", result, 7)
	}
	if _, err := tree.Get(8); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
	if result, _ := tree.Get(9); result != 9 {
		t.Errorf("actual: %d, expected: %d", result, 9)
	}
	if result, _ := tree.Get(10); result != 10 {
		t.Errorf("actual: %d, expected: %d", result, 10)
	}
	if _, err := tree.Get(11); err.Error() != "not_found" {
		t.Errorf("expected 'not_found' error")
	}
}
