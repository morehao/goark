package daopermission

import (
	"context"
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"

	"github.com/morehao/golib/gerror"
	"github.com/morehao/golib/gutils"
	"gorm.io/gorm"
)

type RoleMenuCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type RoleMenuDao struct {
	iamdao.Base
}

func NewRoleMenuDao() *RoleMenuDao {
	return &RoleMenuDao{}
}

func (d *RoleMenuDao) TableName() string {
	return iammodel.TableNameRoleMenu
}

func (d *RoleMenuDao) WithTx(db *gorm.DB) *RoleMenuDao {
	return &RoleMenuDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *RoleMenuDao) Insert(ctx context.Context, entity *iammodel.RoleMenuEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[RoleMenuDao] Insert fail, entity:%s", gutils.ToJsonString(entity))
	}
	return nil
}

func (d *RoleMenuDao) BatchInsert(ctx context.Context, entityList iammodel.RoleMenuEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gerror.DBInsertErr).Wrapf(nil, "[RoleMenuDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[RoleMenuDao] BatchInsert fail, entityList:%s", gutils.ToJsonString(entityList))
	}
	return nil
}

func (d *RoleMenuDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.RoleMenuEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[RoleMenuDao] UpdateByID fail, id:%d entity:%s", id, gutils.ToJsonString(entity))
	}
	return nil
}

func (d *RoleMenuDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[RoleMenuDao] UpdateMap fail, id:%d, updateMap:%s", id, gutils.ToJsonString(updateMap))
	}
	return nil
}

func (d *RoleMenuDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gerror.DBDeleteErr).Wrapf(err, "[RoleMenuDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *RoleMenuDao) GetById(ctx context.Context, id uint) (*iammodel.RoleMenuEntity, error) {
	var entity iammodel.RoleMenuEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *RoleMenuDao) GetByCond(ctx context.Context, cond *RoleMenuCond) (*iammodel.RoleMenuEntity, error) {
	var entity iammodel.RoleMenuEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] GetById fail, cond:%s", gutils.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *RoleMenuDao) GetListByCond(ctx context.Context, cond *RoleMenuCond) (iammodel.RoleMenuEntityList, error) {
	var entityList iammodel.RoleMenuEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] GetListByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *RoleMenuDao) GetPageListByCond(ctx context.Context, cond *RoleMenuCond) (iammodel.RoleMenuEntityList, int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] GetPageListByCond count fail, cond:%s", gutils.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.RoleMenuEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] GetPageListByCond find fail, cond:%s", gutils.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *RoleMenuDao) CountByCond(ctx context.Context, cond *RoleMenuCond) (int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[RoleMenuDao] CountByCond fail, cond:%s", gutils.ToJsonString(cond))
	}
	return count, nil
}

func (d *RoleMenuDao) BuildCondition(db *gorm.DB, cond *RoleMenuCond) {
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
