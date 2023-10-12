package service

import (
	"bic-gin/internal/schema"
	"bic-gin/pkg/db"
	"bic-gin/pkg/gen/api"
	"github.com/gin-gonic/gin"
)

type UserService struct {
}

func (UserService) GetUsers(_ *gin.Context, in *api.RequestUsers) (out *api.ResponseUsers, code api.ErrCode) {
	code = api.ECSuccess
	var dataList []*schema.User
	find := db.SqlDB().Model(&schema.User{})
	if in.Username != "" {
		find.Where("username like ?", "%"+in.Username+"%")
	}
	if in.Email != "" {
		find.Where("email like ?", "%"+in.Email+"%")
	}
	var count int64
	if err := find.Count(&count).
		Limit(int(in.PageSize)).Offset(int((in.Page - 1) * in.PageSize)).
		Find(&dataList).Error; err != nil {
		code = api.ECDbFind.Wrap(err)
		return
	}
	out = &api.ResponseUsers{
		Page:     in.Page,
		PageSize: in.PageSize,
		Total:    uint32(count),
	}
	for _, v := range dataList {
		out.DataList = append(out.DataList, &api.User{
			Id:         uint32(v.ID),
			UpdateTime: v.UpdatedAt.UnixMilli(),
			Username:   v.Username,
			Name:       v.Name,
			Email:      v.Email,
			Ban:        v.Ban,
		})
	}
	return
}
