Vim�UnDo� uh����n�8�ŕc4�?���/���	�A��                                     V��8    _�                            ����                                                                                                                                                                                                                                                                                                                                                             V��6     �               >		goa.Error(ctx, "error getting Member", "error", err.Error())5�_�                            ����                                                                                                                                                                                                                                                                                                                                                             V��7    �                  package models       import (   	"time"       	"xor/cmd/api/app"       	"github.com/goadesign/goa"   	"github.com/jinzhu/gorm"   	"golang.org/x/net/context"   )       /// OneMember returns an array of view: default.   Vfunc (m *MemberDB) OneMemberLoaded(ctx context.Context, id int) (*app.Member, error) {   W	defer goa.MeasureSince([]string{"goa", "db", "member", "onememberloaded"}, time.Now())       	var native Member   �	err := m.Db.Scopes().Table(m.TableName()).Preload("Applications").Preload("Portfolios.Products").Preload("Purchases").Preload("Users").Where("id = ?", id).Find(&native).Error       1	if err != nil && err != gorm.ErrRecordNotFound {   A		goa.LogError(ctx, "error getting Member", "error", err.Error())   		return nil, err   	}       !	view := *native.MemberToMember()   	return &view, err   }5��