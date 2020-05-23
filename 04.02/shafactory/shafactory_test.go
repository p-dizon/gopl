package shafactory

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestInvalidShaSumAlgo(t *testing.T) {
	shaSumAlgo := 500
	_, err := GetShaSumFunc(shaSumAlgo)
	if err == nil {
		t.Errorf("Expected GetShaSumFunc to return error for invalid shaSumAlgo (%v)", shaSumAlgo)
	}
}

func TestShaSum256(t *testing.T) {
	shaSumAlgo := 256
	shaSumFunc, err := GetShaSumFunc(shaSumAlgo)
	if err != nil {
		t.Errorf("Got invalid return (%v) with valid shaSumAlgo %v", err, shaSumAlgo)
	}

	shaSum := shaSumFunc([]byte("x"))
	expectedShaSum, err := hex.DecodeString("2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881")
	if !reflect.DeepEqual(expectedShaSum, shaSum) {
		t.Errorf("Expected [%x], got [%x]", expectedShaSum, shaSum)
	}
}

func TestShaSum384(t *testing.T) {
	shaSumAlgo := 384
	shaSumFunc, err := GetShaSumFunc(shaSumAlgo)
	if err != nil {
		t.Errorf("Got invalid return (%v) with valid shaSumAlgo %v", err, shaSumAlgo)
	}

	shaSum := shaSumFunc([]byte("x"))
	expectedShaSum, err := hex.DecodeString("d752c2c51fba0e29aa190570a9d4253e44077a058d3297fa3a5630d5bd012622f97c28acaed313b5c83bb990caa7da85")
	if !reflect.DeepEqual(expectedShaSum, shaSum) {
		t.Errorf("Expected [%x], got [%x]", expectedShaSum, shaSum)
	}
}

func TestShaSum512(t *testing.T) {
	shaSumAlgo := 512
	shaSumFunc, err := GetShaSumFunc(shaSumAlgo)
	if err != nil {
		t.Errorf("Got invalid return (%v) with valid shaSumAlgo %v", err, shaSumAlgo)
	}

	shaSum := shaSumFunc([]byte("x"))
	expectedShaSum, err := hex.DecodeString("a4abd4448c49562d828115d13a1fccea927f52b4d5459297f8b43e42da89238bc13626e43dcb38ddb082488927ec904fb42057443983e88585179d50551afe62")
	if !reflect.DeepEqual(expectedShaSum, shaSum) {
		t.Errorf("Expected [%x], got [%x]", expectedShaSum, shaSum)
	}
}
