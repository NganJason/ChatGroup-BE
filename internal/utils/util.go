package utils

import (
	"strconv"

	uuid "github.com/nu7hatch/gouuid"
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
	u, _ := uuid.NewV4()

	id, _ := strconv.ParseUint(u.String(), 10, 64)

	return Uint64Ptr(id)
}
