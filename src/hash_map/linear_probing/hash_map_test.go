package hash_map

import "testing"

func TestHashMap(t *testing.T) {
  var hm *HashMap = NewHashMap()

  for i := 0; i < 42; i++ {
    hm.Insert(i)
  }

  for i := 0; i < 42; i++ {
    var result int = hm.Search(i)

    if result == -1 {
      t.Errorf("expected %v, got %v", i, result)
    }
  }

  for i := 42; i < 100; i++ {
    var result int = hm.Search(i)

    if result != -1 {
      t.Errorf("expected -1, got %v", result)
    }
  }
}
