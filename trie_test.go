package bourke

import (
	"testing"
)

func Test_Trie_Flag_Empty(t *testing.T) {
	prefix := prefix[byte, byte]{value: 1, parent: nil, transitions: nil, flags: prefixEmpty}
	if prefix.isKey() {
		t.Errorf("internal prefix MUST NOT be marked as a key")
	}
	if prefix.isTombstone() {
		t.Errorf("internal prefix MUST NOT be marked as a tombstone")
	}
}

func Test_Trie_Flag_Key(t *testing.T) {
	prefix := prefix[byte, byte]{value: 1, parent: nil, transitions: nil, flags: prefixKey}
	if !prefix.isKey() {
		t.Errorf("key prefix MUST be marked as a key")
	}
	if prefix.isTombstone() {
		t.Errorf("key prefix MUST NOT be marked as a tombstone")
	}
}

func Test_Trie_Flag_Tombstone(t *testing.T) {
	prefix := prefix[byte, byte]{value: 1, parent: nil, transitions: nil, flags: prefixTombstone}
	if prefix.isKey() {
		t.Errorf("tombstone prefix MUST NOT be marked as a key")
	}
	if !prefix.isTombstone() {
		t.Errorf("tombstone prefix MUST be marked as a tombstone")
	}
}

func Test_Trie_Write_1(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 3 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 3)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}
	if result, _ := trie.Get([]byte("bear")); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if result, _ := trie.Get([]byte("beard")); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}
	if result, _ := trie.Get([]byte("beach")); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if _, err := trie.Get([]byte("does_not_exist")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}

	trie.Remove([]byte("bear"))
	if _, err := trie.Get([]byte("bear")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if size := trie.Size(); size != 2 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 2)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}
	trie.Tombstone([]byte("bear"))
	if _, err := trie.Get([]byte("bear")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if size := trie.Size(); size != 2 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 2)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}

	trie.Remove([]byte("beard"))
	if _, err := trie.Get([]byte("beard")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if size := trie.Size(); size != 1 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 1)
	}
	if size := trie.InternalSize(); size != 5 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 5)
	}

	trie.Remove([]byte("beach"))
	if _, err := trie.Get([]byte("beach")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	//if size := trie.Size(); size != 0 {
	//	t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	//}
	//if size := trie.InternalSize(); size != 0 {
	//	t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	//}
}

func Test_Trie_Write_2(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if _, err := trie.Get([]byte("does_not_exist")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
}

func Test_Trie_Write_3(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte{1}, 1)

	if size := trie.Size(); size != 1 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 1)
	}
	if size := trie.InternalSize(); size != 1 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 1)
	}
	if result, _ := trie.Get([]byte{1}); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if _, err := trie.Get([]byte("does_not_exist")); err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
}

func Test_Trie_Get_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if _, err := trie.Get([]byte("bake")); err == nil {
		t.Errorf("found")
	}
}

func Test_Trie_Get(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bake"), 4)
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 4 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 4)
	}
	if size := trie.InternalSize(); size != 10 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 10)
	}
	if result, _ := trie.Get([]byte("bake")); result != 4 {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if result, _ := trie.Get([]byte("bear")); result != 1 {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if result, _ := trie.Get([]byte("beard")); result != 2 {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}
	if result, _ := trie.Get([]byte("beach")); result != 3 {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if _, err := trie.Get([]byte("notfound")); err == nil {
		t.Errorf("found")
	}
}

func Test_Trie_Predecessoor_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, result, _ := trie.Predecessor([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_Predeccessor(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bake"), 4)
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 4 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 4)
	}
	if size := trie.InternalSize(); size != 10 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 10)
	}
	if _, _, err := trie.Predecessor([]byte("a")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Predecessor([]byte("ba")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Predecessor([]byte("bak")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Predecessor([]byte("bage")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Predecessor([]byte("bake")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}

	if key, _, _ := trie.Predecessor([]byte("bea")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Predecessor([]byte("bax")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Predecessor([]byte("bare")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Predecessor([]byte("beach")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}

	if key, _, _ := trie.Predecessor([]byte("bead")); string(key) != "beach" {
		t.Errorf("actual: %s, expected: %s", key, "beach")
	}
	if key, _, _ := trie.Predecessor([]byte("bear")); string(key) != "beach" {
		t.Errorf("actual: %s, expected: %s", key, "beach")
	}

	if key, _, _ := trie.Predecessor([]byte("beara")); string(key) != "bear" {
		t.Errorf("actual: %s, expected: %s", key, "bear")
	}
	if key, _, _ := trie.Predecessor([]byte("beard")); string(key) != "bear" {
		t.Errorf("actual: %s, expected: %s", key, "bear")
	}

	if key, _, _ := trie.Predecessor([]byte("bearx")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
	if key, _, _ := trie.Predecessor([]byte("zeard")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
	if key, _, _ := trie.Predecessor([]byte("x")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
}

func Test_Trie_Successor_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, result, _ := trie.Successor([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_Successor(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bake"), 4)
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 4 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 4)
	}
	if size := trie.InternalSize(); size != 10 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 10)
	}
	if key, result, _ := trie.Successor([]byte("a")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Successor([]byte("b")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Successor([]byte("ba")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Successor([]byte("bak")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}

	if key, result, _ := trie.Successor([]byte("bea")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Successor([]byte("bake")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Successor([]byte("bax")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Successor([]byte("bare")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}

	if key, result, _ := trie.Successor([]byte("beach")); result != 1 && string(key) != "bear" {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if key, result, _ := trie.Successor([]byte("bead")); result != 1 && string(key) != "bear" {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}

	if key, result, _ := trie.Successor([]byte("bear")); result != 2 && string(key) != "beard" {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}

	if key, result, _ := trie.Successor([]byte("beard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
	if key, result, _ := trie.Successor([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_Floor_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, result, _ := trie.Floor([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_Floor(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bake"), 4)
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 4 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 4)
	}
	if size := trie.InternalSize(); size != 10 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 10)
	}
	if _, _, err := trie.Floor([]byte("a")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Floor([]byte("ba")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Floor([]byte("bak")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}
	if _, _, err := trie.Floor([]byte("bage")); err == nil || err.Error() != notFoundMessage {
		t.Errorf("should not be found")
	}

	if key, _, _ := trie.Floor([]byte("bake")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Floor([]byte("bea")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Floor([]byte("bax")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}
	if key, _, _ := trie.Floor([]byte("bare")); string(key) != "bake" {
		t.Errorf("actual: %s, expected: %s", key, "bake")
	}

	if key, _, _ := trie.Floor([]byte("beach")); string(key) != "beach" {
		t.Errorf("actual: %s, expected: %s", key, "beach")
	}
	if key, _, _ := trie.Floor([]byte("bead")); string(key) != "beach" {
		t.Errorf("actual: %s, expected: %s", key, "beach")
	}

	if key, _, _ := trie.Floor([]byte("bear")); string(key) != "bear" {
		t.Errorf("actual: %s, expected: %s", key, "bear")
	}
	if key, _, _ := trie.Floor([]byte("beara")); string(key) != "bear" {
		t.Errorf("actual: %s, expected: %s", key, "bear")
	}

	if key, _, _ := trie.Floor([]byte("beard")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
	if key, _, _ := trie.Floor([]byte("bearx")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
	if key, _, _ := trie.Floor([]byte("zeard")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
	if key, _, _ := trie.Floor([]byte("x")); string(key) != "beard" {
		t.Errorf("actual: %s, expected: %s", key, "beard")
	}
}

func Test_Trie_Ceiling_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, result, _ := trie.Ceiling([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_Ceiling(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bake"), 4)
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 4 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 4)
	}
	if size := trie.InternalSize(); size != 10 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 10)
	}
	if key, result, _ := trie.Ceiling([]byte("a")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Ceiling([]byte("ba")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Ceiling([]byte("bak")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}
	if key, result, _ := trie.Ceiling([]byte("bake")); result != 4 && string(key) != "bake" {
		t.Errorf("actual: %d, expected: %d", result, 4)
	}

	if key, result, _ := trie.Ceiling([]byte("bea")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Ceiling([]byte("bax")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Ceiling([]byte("bare")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}
	if key, result, _ := trie.Ceiling([]byte("beach")); result != 3 && string(key) != "beach" {
		t.Errorf("actual: %d, expected: %d", result, 3)
	}

	if key, result, _ := trie.Ceiling([]byte("bead")); result != 1 && string(key) != "bear" {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}
	if key, result, _ := trie.Ceiling([]byte("bear")); result != 1 && string(key) != "bear" {
		t.Errorf("actual: %d, expected: %d", result, 1)
	}

	if key, result, _ := trie.Ceiling([]byte("beard")); result != 2 && string(key) != "beard" {
		t.Errorf("actual: %d, expected: %d", result, 2)
	}

	if key, result, _ := trie.Ceiling([]byte("zeard")); result != 0 && key != nil {
		t.Errorf("actual: %d, expected: %d", result, 0)
	}
}

func Test_Trie_First(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 3 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 3)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}
	if key, value, _ := trie.First(); string(key) != "beach" || value != 3 {
		t.Errorf("actual: %d, expected: %d", value, 3)
	}
}

func Test_Trie_First_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, value, err := trie.First(); key != nil || value != 0 || err.Error() != notFoundMessage {
		t.Errorf("actual: %d, expected: %d", value, 3)
	}
}

func Test_Trie_Last_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	if size := trie.Size(); size != 0 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 0)
	}
	if size := trie.InternalSize(); size != 0 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 0)
	}
	if key, val, err := trie.Last(); err != nil && err.Error() != notFoundMessage {
		t.Errorf("key: %s, val: %v", key, val)
	}
}

func Test_Trie_Last_1(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	trie.Tombstone([]byte("beard"))
	trie.Tombstone([]byte("bear"))

	if size := trie.Size(); size != 1 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 1)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}
	if key, val, _ := trie.Last(); string(key) != "beach" {
		t.Errorf("key: %s, val: %v", key, val)
	}
}

func Test_Trie_Last_2(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("bear"), 1)
	trie.Put([]byte("beard"), 2)
	trie.Put([]byte("beach"), 3)

	if size := trie.Size(); size != 3 {
		t.Errorf("size: [actual: %d, expected: %d]", size, 3)
	}
	if size := trie.InternalSize(); size != 7 {
		t.Errorf("internal size: [actual: %d, expected: %d]", size, 7)
	}
	if key, val, _ := trie.Last(); string(key) != "beard" {
		t.Errorf("key: %s, val: %v", key, val)
	}
}

func Test_Trie_from_Empty(t *testing.T) {
	trie := newTrie[byte, byte]()

	next := trie.forward(trie.root, 0, trie.root)
	for current, _ := next(); current != nil; current, _ = next() {
		t.Errorf("should not iterate when empty")
	}
}

func Test_Trie_from_Single(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("a"), 1)

	result := make([]struct {
		depth  uint
		prefix *prefix[byte, byte]
	}, 9)
	index := 0
	next := trie.forward(trie.root, 0, trie.root)
	for current, depth := next(); current != nil; current, depth = next() {
		result[index] = struct {
			depth  uint
			prefix *prefix[byte, byte]
		}{depth, current}
		index++
	}

	if index != 1 {
		t.Errorf("incorrect length")
	}
	if string(result[0].prefix.parentEdge) != "a" || result[0].depth != 0 || result[0].prefix.isKey() == false || result[0].prefix.value != 1 {
		t.Errorf("edge: %c, depth: %d, key: %t, value: %d\n", result[0].prefix.parentEdge, result[0].depth, result[0].prefix.isKey(), result[0].prefix.value)
	}
}

func Test_Trie_from_Single_Chain(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("abc"), 1)

	result := make([]struct {
		depth  uint
		prefix *prefix[byte, byte]
	}, 3)
	index := 0
	next := trie.forward(trie.root, 0, trie.root)
	for current, depth := next(); current != nil; current, depth = next() {
		result[index] = struct {
			depth  uint
			prefix *prefix[byte, byte]
		}{depth, current}
		index++
	}

	if index != 3 {
		t.Errorf("incorrect length")
	}
	if string(result[0].prefix.parentEdge) != "a" || result[0].depth != 0 || result[0].prefix.isKey() == true || result[0].prefix.value != 0 {
		t.Errorf("edge: %c, depth: %d, key: %t, value: %d\n", result[0].prefix.parentEdge, result[0].depth, result[0].prefix.isKey(), result[0].prefix.value)
	}
	if string(result[1].prefix.parentEdge) != "b" || result[1].depth != 1 || result[1].prefix.isKey() == true || result[1].prefix.value != 0 {
		t.Errorf("edge: %c, depth: %d, key: %t, value: %d\n", result[1].prefix.parentEdge, result[1].depth, result[1].prefix.isKey(), result[1].prefix.value)
	}
	if string(result[2].prefix.parentEdge) != "c" || result[2].depth != 2 || result[2].prefix.isKey() == false || result[2].prefix.value != 1 {
		t.Errorf("edge: %c, depth: %d, key: %t, value: %d\n", result[2].prefix.parentEdge, result[2].depth, result[2].prefix.isKey(), result[2].prefix.value)
	}
}

func Test_Trie_from(t *testing.T) {
	trie := newTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)
	trie.Put([]byte("beach"), 2)

	result := make([]struct {
		depth  uint
		prefix *prefix[byte, byte]
	}, 9)
	index := 0
	next := trie.forward(trie.root, 0, trie.root)
	for current, depth := next(); current != nil; current, depth = next() {
		result[index] = struct {
			depth  uint
			prefix *prefix[byte, byte]
		}{depth, current}
		index++
	}

	if index != 9 {
		t.Errorf("incorrect length")
	}
	if string(result[0].prefix.parentEdge) != "a" || result[0].depth != 0 || result[0].prefix.isKey() == true || result[0].prefix.value != 0 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[0].prefix.parentEdge, result[0].depth, result[0].prefix.isKey(), result[0].prefix.value)
	}
	if string(result[1].prefix.parentEdge) != "b" || result[1].depth != 1 || result[1].prefix.isKey() == false || result[1].prefix.value != 1 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[1].prefix.parentEdge, result[1].depth, result[1].prefix.isKey(), result[1].prefix.value)
	}
	if string(result[2].prefix.parentEdge) != "b" || result[2].depth != 0 || result[2].prefix.isKey() == true || result[2].prefix.value != 0 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[2].prefix.parentEdge, result[2].depth, result[2].prefix.isKey(), result[2].prefix.value)
	}
	if string(result[3].prefix.parentEdge) != "e" || result[3].depth != 1 || result[3].prefix.isKey() == true || result[3].prefix.value != 0 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[3].prefix.parentEdge, result[3].depth, result[3].prefix.isKey(), result[3].prefix.value)
	}
	if string(result[4].prefix.parentEdge) != "a" || result[4].depth != 2 || result[4].prefix.isKey() == true || result[4].prefix.value != 0 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[4].prefix.parentEdge, result[4].depth, result[4].prefix.isKey(), result[4].prefix.value)
	}
	if string(result[5].prefix.parentEdge) != "c" || result[5].depth != 3 || result[5].prefix.isKey() == true || result[5].prefix.value != 0 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[5].prefix.parentEdge, result[5].depth, result[5].prefix.isKey(), result[5].prefix.value)
	}
	if string(result[6].prefix.parentEdge) != "h" || result[6].depth != 4 || result[6].prefix.isKey() == false || result[6].prefix.value != 2 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[6].prefix.parentEdge, result[6].depth, result[6].prefix.isKey(), result[6].prefix.value)
	}
	if string(result[7].prefix.parentEdge) != "r" || result[7].depth != 3 || result[7].prefix.isKey() == false || result[7].prefix.value != 3 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[7].prefix.parentEdge, result[7].depth, result[7].prefix.isKey(), result[7].prefix.value)
	}
	if string(result[8].prefix.parentEdge) != "d" || result[8].depth != 4 || result[8].prefix.isKey() == false || result[8].prefix.value != 4 {
		t.Errorf("char: %c, depth: %d, cursor: %t, value: %d\n", result[8].prefix.parentEdge, result[8].depth, result[8].prefix.isKey(), result[8].prefix.value)
	}
}

func Test_Trie_All(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)

	trie.Tombstone([]byte("bake"))
	trie.Tombstone([]byte("invalid"))

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.All() {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 4 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "ab" || result[0].value != 1 {
		t.Errorf("%s", result[0].key)
	}
	if string(result[1].key) != "beach" || result[1].value != 2 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[2].key) != "bear" || result[2].value != 3 {
		t.Errorf("%s", result[2].key)
	}
	if string(result[3].key) != "beard" || result[3].value != 4 {
		t.Errorf("%s", result[4].key)
	}
}

func Test_Trie_LessThan_Inclusive_Empty(t *testing.T) {
	trie := NewTrie[byte, byte]()

	for range trie.LessThan([]byte("bake"), true) {
		t.Error("should not iterate")
	}
}

func Test_Trie_LessThan_Inclusive(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.LessThan([]byte("bear"), true) {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 4 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "ab" || result[0].value != 1 {
		t.Errorf("%s", result[0].key)
	}
	if string(result[1].key) != "bake" || result[1].value != 5 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[2].key) != "beach" || result[2].value != 2 {
		t.Errorf("%s", result[2].key)
	}
	if string(result[3].key) != "bear" || result[3].value != 3 {
		t.Errorf("%s", result[3].key)
	}
}

func Test_Trie_LessThan_Exclusive_Empty(t *testing.T) {
	trie := NewTrie[byte, byte]()

	for range trie.LessThan([]byte("a"), false) {
		t.Error("Empty. Should not iterate.")
	}
}

func Test_Trie_LessThan_Exclusive(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.LessThan([]byte("bear"), false) {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 3 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "ab" || result[0].value != 1 {
		t.Errorf("%s", result[0].key)
	}
	if string(result[1].key) != "bake" || result[1].value != 5 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[2].key) != "beach" || result[2].value != 2 {
		t.Errorf("%s", result[2].key)
	}
}

func Test_Trie_GreaterThan_Inclusive_Empty(t *testing.T) {
	trie := NewTrie[byte, byte]()

	for range trie.GreaterThan([]byte("bake"), true) {
		t.Error("should not iterate")
	}
}

func Test_Trie_GreaterThan_Inclusive(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.GreaterThan([]byte("bake"), true) {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 4 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "bake" || result[0].value != 5 {
		t.Errorf("%s", result[0].key)
	}
	if string(result[1].key) != "beach" || result[1].value != 2 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[2].key) != "bear" || result[2].value != 3 {
		t.Errorf("%s", result[2].key)
	}
	if string(result[3].key) != "beard" || result[3].value != 4 {
		t.Errorf("%s", result[4].key)
	}
}

func Test_Trie_GreaterThan_Exclusive_Empty(t *testing.T) {
	trie := NewTrie[byte, byte]()

	for range trie.GreaterThan([]byte("a"), false) {
		t.Error("Empty. Should not iterate.")
	}
}

func Test_Trie_GreaterThan_Exclusive(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.GreaterThan([]byte("bake"), false) {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 3 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "beach" || result[0].value != 2 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[1].key) != "bear" || result[1].value != 3 {
		t.Errorf("%s", result[2].key)
	}
	if string(result[2].key) != "beard" || result[2].value != 4 {
		t.Errorf("%s", result[4].key)
	}
}

func Test_Trie_Prefix_Empty(t *testing.T) {
	trie := NewTrie[byte, byte]()

	for range trie.Prefix([]byte("x")) {
		t.Errorf("should not iterate")
	}
}

func Test_Trie_Prefix(t *testing.T) {
	trie := NewTrie[byte, byte]()
	trie.Put([]byte("ab"), 1)
	trie.Put([]byte("bake"), 5)
	trie.Put([]byte("be"), 7)
	trie.Put([]byte("beach"), 2)
	trie.Put([]byte("bear"), 3)
	trie.Put([]byte("beard"), 4)
	trie.Put([]byte("zeard"), 6)

	for range trie.Prefix([]byte("x")) {
		t.Errorf("should not iterate")
	}

	result := make([]struct {
		key   []byte
		value byte
	}, 4)
	index := 0
	for key, value := range trie.Prefix([]byte("be")) {
		result[index] = struct {
			key   []byte
			value byte
		}{key, value}
		index++
	}

	if index != 4 {
		t.Errorf("incorrect length")
	}
	if string(result[0].key) != "be" || result[0].value != 7 {
		t.Errorf("%s", result[0].key)
	}
	if string(result[1].key) != "beach" || result[1].value != 2 {
		t.Errorf("%s", result[1].key)
	}
	if string(result[2].key) != "bear" || result[2].value != 3 {
		t.Errorf("%s", result[2].key)
	}
	if string(result[3].key) != "beard" || result[3].value != 4 {
		t.Errorf("%s", result[3].key)
	}
}
