package main

import (
	"fmt"
	"github.com/google/uuid"
	"encoding/binary"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}



func main() {
	
	v := String(42)
	u := uuid.NewSHA1( uuid.NameSpaceX500, []byte(v) )
	m := []byte( u.String()[0:8] )
	data := binary.BigEndian.Uint64( m )
	fmt.Printf( "%v\n%v\n%v\n%#v\n", v, data , data%256 , u.String() )

}
