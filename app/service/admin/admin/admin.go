package admin

import (
	"bbs/app/model/admins"
	"bbs/app/request/Admin"
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/frame/g"
)

func Add(data *Admin.AddReqEntity) error {
	admin, _ := g.DB().Table(admins.Table).Where("email = ?", data.Email).One()
	if admin != nil {
		return errors.New("邮箱已存在")
	}
	password, _ := gmd5.Encrypt(data.Password)
	res, err := g.DB().Table(admins.Table).Insert(g.Map{
		"name":       data.Name,
		"email":       data.Email,
		"password":		password,
		"status":        data.Status,
	})
	if err != nil {
		g.Log().Error("入库失败：", err)
		return errors.New("添加失败")
	}
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		g.Log().Error("入库失败：", err)
		return errors.New("添加失败")
	}

	return nil
}

func Edit(data *Admin.UpdateReqEntity, id int) error {
	admin, _ := g.DB().Table(admins.Table).Where("email = ? and id != ?", data.Email, id).One()
	if admin != nil {
		return errors.New("邮箱已存在")
	}
	gMap := g.Map{
		"name": data.Name,
		"email": data.Email,
		"status": data.Status,
	}
	if data.Password != "" {
		password, _ := gmd5.Encrypt(data.Password)
		gMap["password"] = password
	}
	res, err := g.DB().Table(admins.Table).WherePri(id).Update(gMap)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}

func Delete(id int) error {
	res, err := g.DB().Table(admins.Table).Delete("id", id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil || rows <= 0 {
		return err
	}

	return nil
}
