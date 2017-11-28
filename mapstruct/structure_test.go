package mapstruct

import (
	"encoding/json"
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

func Test_StructToMap_2(t *testing.T) {

	type Dto2 struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	type Dto1 struct {
		Name  string `json:"name"`
		Param Dto2   `json:"dto2"`
	}

	var reqDto = Dto1{
		Name: "123",
		Param: Dto2{
			Name: "sub_123",
			Age:  12,
		},
	}
	s := structs.New(reqDto)
	s.TagName = "json"
	param := s.Map()
	fmt.Printf("\n%+v", param)

}

func Test_StructToMap_3(t *testing.T) {

	type BaseDto struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	type Dto1 struct {
		*BaseDto `json:"-"`
		Name     string `json:"name"`
	}

	var reqDto = Dto1{
		BaseDto: &BaseDto{
			Name: "sub_123",
			Age:  12,
		},
		Name: "123",
	}

	s := structs.New(reqDto)
	s.TagName = "json"
	param := s.Map()
	fmt.Printf("\n%+v", param)

}

func Test_Json(t *testing.T) {

	type BaseDto struct {
		Name string `json:"name"`
		Age  int64  `json:"age"`
	}

	type Dto1 struct {
		*BaseDto
		Name1 string `json:"name1"`
	}

	jsonStr := `{
		"name1":"123",
		"name":"234"
	}`
	var dto1 Dto1
	json.Unmarshal([]byte(jsonStr), &dto1)
	fmt.Println(dto1)
	fmt.Println(dto1.BaseDto)
}
