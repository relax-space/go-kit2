package mapstruct

import (
	"fmt"
	"testing"

	"github.com/fatih/structs"
)

func Test_StructToMap(t *testing.T) {

	var reqDto = struct {
		AppId string `json:"app_id"`
	}{"xiao.xinmiao"}
	s := structs.New(reqDto)
	s.TagName = "json"
	param := s.Map()
	fmt.Printf("\n%+v", param)

}
