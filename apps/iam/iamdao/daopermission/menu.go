package daopermission

import (
	"context"
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"

	"github.com/morehao/golib/gerror"
	"github.com/morehao/golib/gutil"
	"gorm.io/gorm"
)

type MenuCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type MenuDao struct {
	iamdao.Base
}

func NewMenuDao() *MenuDao {
	return &MenuDao{}
}

func (d *MenuDao) TableName() string {
	return iammodel.TableNameMenu
}

func (d *MenuDao) WithTx(db *gorm.DB) *MenuDao {
	return &MenuDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *MenuDao) Insert(ctx context.Context, entity *iammodel.MenuEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[MenuDao] Insert fail, entity:%s", gutil.ToJsonString(entity))
	}
	return nil
}

func (d *MenuDao) BatchInsert(ctx context.Context, entityList iammodel.MenuEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gerror.DBInsertErr).Wrapf(nil, "[MenuDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gerror.DBInsertErr).Wrapf(err, "[MenuDao] BatchInsert fail, entityList:%s", gutil.ToJsonString(entityList))
	}
	return nil
}

func (d *MenuDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.MenuEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[MenuDao] UpdateByID fail, id:%d entity:%s", id, gutil.ToJsonString(entity))
	}
	return nil
}

func (d *MenuDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gerror.DBUpdateErr).Wrapf(err, "[MenuDao] UpdateMap fail, id:%d, updateMap:%s", id, gutil.ToJsonString(updateMap))
	}
	return nil
}

func (d *MenuDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gerror.DBDeleteErr).Wrapf(err, "[MenuDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *MenuDao) GetById(ctx context.Context, id uint) (*iammodel.MenuEntity, error) {
	var entity iammodel.MenuEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *MenuDao) GetByCond(ctx context.Context, cond *MenuCond) (*iammodel.MenuEntity, error) {
	var entity iammodel.MenuEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] GetById fail, cond:%s", gutil.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *MenuDao) GetListByCond(ctx context.Context, cond *MenuCond) (iammodel.MenuEntityList, error) {
	var entityList iammodel.MenuEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] GetListByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *MenuDao) GetPageListByCond(ctx context.Context, cond *MenuCond) (iammodel.MenuEntityList, int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] GetPageListByCond count fail, cond:%s", gutil.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.MenuEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] GetPageListByCond find fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *MenuDao) CountByCond(ctx context.Context, cond *MenuCond) (int64, error) {
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gerror.DBFindErr).Wrapf(err, "[MenuDao] CountByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return count, nil
}

func (d *MenuDao) BuildCondition(db *gorm.DB, cond *MenuCond) {
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
