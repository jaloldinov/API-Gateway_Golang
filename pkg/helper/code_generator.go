package helper

import (
	"crypto/rand"
	"encoding/json"
	"io"
	"math/big"
)

var (
	//table for code generator
	table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
)

// GenerateCode is function generating n-digit random code
func GenerateCode(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func RandomInt(max int) int {
	n, err := rand.Int(rand.Reader, big.NewInt(100000000))
	if err != nil {
		return 0
	}

	return int(n.Int64()) % max
}

func MarshalToStruct(data interface{}, resp interface{}) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = json.Unmarshal(js, resp)
	if err != nil {
		return err
	}

	return nil
}
