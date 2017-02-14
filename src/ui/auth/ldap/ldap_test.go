package ldap

import (
	"github.com/vmware/harbor/src/common/dao"
	"github.com/vmware/harbor/src/common/models"
	"testing"
)

func init() {
	dao.InitDatabase()
}
func TestAuthenticate(t *testing.T) {
	m := models.AuthModel{Principal: "harbor", Password: "harbor"}
	var auth Auth
	_, err := auth.Authenticate(m)
	if err != nil {
		t.Error("Error occured when test ldap authenticate,", err.Error())
		t.Log(err)
	}
}
