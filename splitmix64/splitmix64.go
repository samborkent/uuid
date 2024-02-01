// splitmix64 (https://rosettacode.org/wiki/Pseudo-random_numbers/Splitmix64)
package splitmix64

const maxUInt64 = ^uint64(0)

var state uint64

func NextInt() uint64 {
	state += 0x9e3779b97f4a7c15
	z := state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}
