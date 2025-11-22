package daouser

import (
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"

	"github.com/gin-gonic/gin"
	"github.com/morehao/golib/gutils"
	"gorm.io/gorm"
)

type UserRoleCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type UserRoleDao struct {
	iamdao.Base
}

func NewUserRoleDao() *UserRoleDao {
	return &UserRoleDao{}
}

func (d *UserRoleDao) TableName() string {
	return iammodel.TableNameUserRole
}

func (d *UserRoleDao) WithTx(db *gorm.DB) *UserRoleDao {
	return &UserRoleDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *UserRoleDao) Insert(ctx *gin.Context, entity *iammodel.UserRoleEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[UserRoleDao] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (d *UserRoleDao) BatchInsert(ctx *gin.Context, entityList iammodel.UserRoleEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(code.DBInsertErr).Wrapf(nil, "[UserRoleDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(code.DBInsertErr).Wrapf(err, "[UserRoleDao] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (d *UserRoleDao) UpdateByID(ctx *gin.Context, id uint, entity *iammodel.UserRoleEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserRoleDao] UpdateByID fail, id:%d entity:%s", id, gutils.ToJsonString(entity))
	}
	return nil
}

func (d *UserRoleDao) UpdateMap(ctx *gin.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserRoleDao] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (d *UserRoleDao) Delete(ctx *gin.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(code.DBUpdateErr).Wrapf(err, "[UserRoleDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *UserRoleDao) GetById(ctx *gin.Context, id uint) (*iammodel.UserRoleEntity, error) {
	var entity iammodel.UserRoleEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *UserRoleDao) GetByCond(ctx *gin.Context, cond *UserRoleCond) (*iammodel.UserRoleEntity, error) {
	var entity iammodel.UserRoleEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *UserRoleDao) GetListByCond(ctx *gin.Context, cond *UserRoleCond) (iammodel.UserRoleEntityList, error) {
	var entityList iammodel.UserRoleEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *UserRoleDao) GetPageListByCond(ctx *gin.Context, cond *UserRoleCond) (iammodel.UserRoleEntityList, int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.UserRoleEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *UserRoleDao) CountByCond(ctx *gin.Context, cond *UserRoleCond) (int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(code.DBFindErr).Wrapf(err, "[UserRoleDao] CountByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return count, nil
}

func (d *UserRoleDao) BuildCondition(db *gorm.DB, cond *UserRoleCond) {
	if cond.ID > 0 {
		query := fmt.Sprintf("%s.id = ?", d.TableName())
		db.Where(query, cond.ID)
	}
	if len(cond.IDs) > 0 {
		query := fmt.Sprintf("%s.id in (?)", d.TableName())
		db.Where(query, cond.IDs)
	}
	if cond.CreatedAtStart > 0 {
		query := fmt.Sprintf("%s.created_at >= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtStart, 0))
	}
	if cond.CreatedAtEnd > 0 {
		query := fmt.Sprintf("%s.created_at <= ?", d.TableName())
		db.Where(query, time.Unix(cond.CreatedAtEnd, 0))
	}
	if cond.IsDelete {
		db.Unscoped()
	}

	if cond.OrderField != "" {
		db.Order(cond.OrderField)
	}
}
