package auth

import (
	"fmt"
	"testing"

	"github.com/relax-space/go-kitt/testt"
)

func Test_NewToken(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	fmt.Println(token)
	testt.Ok(t, err)
}

func Test_Extract(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	testt.Ok(t, err)
	mapClaims, err := Extract(token)
	testt.Ok(t, err)
	testt.Equals(t, "wx1234", mapClaims["app_id"])
}

func Test_Renew(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	testt.Ok(t, err)
	token, err = Renew(token)
	testt.Ok(t, err)
}

func Test_EditPayload(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	testt.Ok(t, err)
	addParam := map[string]interface{}{"name": "limon"}
	_, err = EditPayload(token, addParam)
	testt.Ok(t, err)
}

func Test_EditPayloadAdd(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	testt.Ok(t, err)
	addParam := map[string]interface{}{"name": "limon"}
	token, err = EditPayload(token, addParam)
	testt.Ok(t, err)
	claim, err := Extract(token)
	testt.Assert(t, "limon" == claim["name"], "add token failure")
	testt.Assert(t, "wx1234" == claim["app_id"], "add token failure")
}

func Test_EditPayloadUpdate(t *testing.T) {
	token, err := NewToken(map[string]interface{}{"app_id": "wx1234", "role_id": "1"})
	testt.Ok(t, err)
	addParam := map[string]interface{}{"app_id": "wx1111"}
	token, err = EditPayload(token, addParam)
	testt.Ok(t, err)
	claim, err := Extract(token)
	testt.Assert(t, "wx1111" == claim["app_id"], "edit token failure")
}
