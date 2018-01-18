package random

import (
	"bytes"
	"encoding/binary"
	"strconv"
	"time"

	"github.com/relax-space/go-kit/base"

	uuid "github.com/satori/go.uuid"
)

func UuIdInt64() (uuId uint64) {
	u1, _ := uuid.NewV4()
	buf := bytes.NewBuffer(u1.Bytes())
	binary.Read(buf, binary.BigEndian, &uuId)
	return
}

func Uuid(randomType string) string {
	return randomType + strconv.FormatUint(UuIdInt64(), 10)
}
func NewUuid(randomType string) string {
	t := time.Now()
	return randomType + base.NewDateFormat(t, "yyMMdd") + strconv.FormatUint(UuIdInt64(), 10)
}
