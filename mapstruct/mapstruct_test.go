package mapstruct

import (
	"testing"

	"github.com/relax-space/go-kitt/testt"
)

func Test_Decode(t *testing.T) {

	var apple struct {
		Color string `json:"color"`
		Price int    `json:"price"`
	}

	appleMap := map[string]interface{}{
		"color": "red",
		"price": 14,
	}

	testt.Ok(t, Decode(appleMap, &apple))

	testt.Equals(t, "red", apple.Color)

}
