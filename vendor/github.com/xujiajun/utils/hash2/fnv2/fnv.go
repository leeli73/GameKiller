package fnv2

import "hash/fnv"

func Hash32(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}
