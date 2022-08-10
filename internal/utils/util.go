package utils

import (
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"net/http"
	"strings"

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

func Uint32Ptr(i uint32) *uint32 {
	return &i
}

func GenerateUUID() *uint64 {
	byte, _ := uuid.New().MarshalBinary()

	id := binary.BigEndian.Uint64(byte)

	return Uint64Ptr(id)
}

func GetServerAddress(ctx context.Context) string {
	srvAddr := ctx.Value(http.LocalAddrContextKey).(net.Addr)

	s := strings.Replace(srvAddr.String(), LocalHostIP, LocalHostAddr, 1)

	return fmt.Sprintf("http://%s/", s)
}

func GetImgUrl(ctx context.Context, path string) string {
	serverAddr := GetServerAddress(ctx)

	stringList := strings.Split(path, "/")
	imgPath := stringList[len(stringList)-1]

	return fmt.Sprintf("%sapi/user/get_image?path=%s", serverAddr, imgPath)
}
