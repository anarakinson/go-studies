package words

import (
    "crypto/rand"
    "math/big"
)

var Hello = "This is package words"
var Prefix = "Random word is: "
var Words = []string {"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven"}

func Random() (string) {
    max := len(Words)
    r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
    return get(r.Int64())
}

func get(index int64) (string) {
    return Prefix + Words[index]
}
