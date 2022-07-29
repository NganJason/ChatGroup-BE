package utils

import (
	"encoding/binary"

	uuid "github.com/google/uuid"
)

func BoolPtr(b bool) *bool {
	return &b
}

func StrPtr(s string) *string {
	return &s
}

func IntPtr(i int) *int {
	return &i
}

func Uint64Ptr(i uint64) *uint64 {
	return &i
}

func GenerateUUID() *uint64 {
	byte, _ := uuid.New().MarshalBinary()
	
	id := binary.BigEndian.Uint64(byte)
	
	return Uint64Ptr(id)
}
