package shafactory

import (
	"crypto/sha256"
	"crypto/sha512"
	"errors"
)

func GetShaSumFunc(shaSumAlgo int) (func([]byte) []byte, error) {
	var function func([]byte) []byte
	switch shaSumAlgo {
	case 256:
		function = func(b []byte) []byte {
			hash := sha256.Sum256(b)
			return hash[:]
		}
	case 384:
		function = func(b []byte) []byte {
			hash := sha512.Sum384(b)
			return hash[:]
		}
	case 512:
		function = func(b []byte) []byte {
			hash := sha512.Sum512(b)
			return hash[:]
		}
	default:
		return nil, errors.New("Invalid sha sum algorithm")
	}
	return function, nil
}
