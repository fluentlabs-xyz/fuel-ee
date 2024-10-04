package common

func IsBitSet(n int8, pos uint) bool {
	return (n & (1 << pos)) != 0
}
