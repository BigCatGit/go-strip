package core

import (
	"math/rand"
	"sync"
	"time"
)

func GenerateSameWords(n int, s byte) string {
	bytes := make([]byte, n)
	for i := 0; i < n; i++ {
		bytes[i] = s
	}
	return string(bytes)
}
func GetRandomString(n int) string {
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandomSymbol() string {
	str := "!@#$%^&*()-+~."
	bytes := []byte(str)
	return string(bytes[rand.Intn(len(bytes))])
}

type TypeData struct {
	Offset uint64
	Length uint64
}
type TypeStringOffset struct {
	Base      uint64
	PCLnTab   uint64
	Datas     []TypeData
	GoVersion TypeData
	Func      []TypeData
	FileName  []TypeData
	GoMod     []TypeData
}

var TypeStringOffsets *TypeStringOffset

func init() {
	a := sync.Once{}
	a.Do(func() {
		rand.Seed(time.Now().Unix())
		TypeStringOffsets = new(TypeStringOffset)
	})
}
