package xsr256pp

import "github.com/samborkent/uuid/splitmix64"

var s = [4]uint64{splitmix64.NextInt(), splitmix64.NextInt(), splitmix64.NextInt(), splitmix64.NextInt()}

func Next() uint64 {
	result := rotl(s[0]+s[3], 23) + s[0]

	t := s[1] << 17

	s[2] ^= s[0]
	s[3] ^= s[1]
	s[1] ^= s[2]
	s[0] ^= s[3]

	s[2] ^= t

	s[3] = rotl(s[3], 45)

	return result
}

func rotl(x uint64, k int) uint64 {
	return (x << k) | (x >> (64 - k))
}
