package models

import (
    "github.com/astaxie/beego/orm"
)

type AuthUser struct {
    Id             int
    Role_id        int
    Account        string
    Realname       string
    Phone          string
    Position       string
    Password       string
    Password_token string
    Status         int
    Login_at       int
    Created_at     int
    Updated_at     int
}

func (this *AuthUser) TableName() string {
    return TableName("admin_auth_user")
}

func (this *AuthUser) One(filters ...interface{}) (*AuthUser, error) {
    model := new(AuthUser)
    query := orm.NewOrm().QueryTable(this.TableName())

    if len(filters) > 0 {
        l := len(filters)
        for k := 0; k < l; k += 2 {
            query = query.Filter(filters[k].(string), filters[k+1])
        }
    }

    errs := query.One(model)
    if errs != nil {
        return nil, errs
    }
    return model, errs
}
