package hash_map

import "math"

const HASH_MAP_SIZE = 100
var   HASH_CONSTANT = ((math.Sqrt(5) - 1) / 2)

type HashMap struct {
  hash_map []int
}

func NewHashMap() *HashMap {
  var hm *HashMap = new(HashMap)
  hm.hash_map = make([]int, HASH_MAP_SIZE)

  return hm
}

func (hm *HashMap) Insert(value int) int {
  var probe int = Hash(value)

  for hm.hash_map[probe] != 0 {
    probe = NextProbe(probe)
  }

  hm.hash_map[probe] = value

  return 0
}

func (hm *HashMap) Search(value int) int {
  var probe int = Hash(value)

  for i := 0; i < HASH_MAP_SIZE; i++ {
    if hm.hash_map[probe] == value {
      return value
    }

    probe = NextProbe(probe)
  }

  return -1;
}

func (hm *HashMap) GetHashMap() (arr []int) {
  return hm.hash_map
}


func Hash(value int) int {
  return int((HASH_MAP_SIZE * math.Mod(float64(value) * HASH_CONSTANT, 1)))
}

func NextProbe(probe int) int {
  return int(math.Mod(float64(probe + 1), HASH_MAP_SIZE))
}
