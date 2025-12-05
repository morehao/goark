package daotenant

import (
	"context"
	"fmt"
	"time"

	"github.com/morehao/goark/apps/iam/iamdao"
	"github.com/morehao/goark/apps/iam/iammodel"
	"github.com/morehao/goark/pkg/code"
	"github.com/morehao/golib/biz/gconstant"
	"github.com/morehao/golib/gutil"
	"gorm.io/gorm"
)

type TenantConfigCond struct {
	ID             uint
	IDs            []uint
	IsDelete       bool
	Page           int
	PageSize       int
	CreatedAtStart int64
	CreatedAtEnd   int64
	OrderField     string
}

type TenantConfigDao struct {
	iamdao.Base
}

func NewTenantConfigDao() *TenantConfigDao {
	return &TenantConfigDao{}
}

func (d *TenantConfigDao) TableName() string {
	return iammodel.TableNameTenantConfig
}

func (d *TenantConfigDao) WithTx(db *gorm.DB) *TenantConfigDao {
	return &TenantConfigDao{
		Base: iamdao.Base{Tx: db},
	}
}

func (d *TenantConfigDao) Insert(ctx context.Context, entity *iammodel.TenantConfigEntity) error {
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entity).Error; err != nil {
		return code.GetError(gconstant.DBInsertErr).Wrapf(err, "[TenantConfigDao] Insert fail, entity:%s", gutil.ToJsonString(entity))
	}
	return nil
}

func (d *TenantConfigDao) BatchInsert(ctx context.Context, entityList iammodel.TenantConfigEntityList) error {
	if len(entityList) == 0 {
		return code.GetError(gconstant.DBInsertErr).Wrapf(nil, "[TenantConfigDao] BatchInsert fail, entityList is empty")
	}

	db := d.DB(ctx).Table(d.TableName())
	if err := db.Create(entityList).Error; err != nil {
		return code.GetError(gconstant.DBInsertErr).Wrapf(err, "[TenantConfigDao] BatchInsert fail, entityList:%s", gutil.ToJsonString(entityList))
	}
	return nil
}

func (d *TenantConfigDao) UpdateByID(ctx context.Context, id uint, entity *iammodel.TenantConfigEntity) error {
	db := d.DB(ctx).Model(&iammodel.TenantConfigEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(entity).Error; err != nil {
		return code.GetError(gconstant.DBUpdateErr).Wrapf(err, "[TenantConfigDao] UpdateByID fail, id:%d entity:%s", id, gutil.ToJsonString(entity))
	}
	return nil
}

func (d *TenantConfigDao) UpdateMap(ctx context.Context, id uint, updateMap map[string]interface{}) error {
	db := d.DB(ctx).Model(&iammodel.TenantConfigEntity{}).Table(d.TableName())
	if err := db.Where("id = ?", id).Updates(updateMap).Error; err != nil {
		return code.GetError(gconstant.DBUpdateErr).Wrapf(err, "[TenantConfigDao] UpdateMap fail, id:%d, updateMap:%s", id, gutil.ToJsonString(updateMap))
	}
	return nil
}

func (d *TenantConfigDao) Delete(ctx context.Context, id, deletedBy uint) error {
	db := d.DB(ctx).Model(&iammodel.TenantConfigEntity{}).Table(d.TableName())
	updatedField := map[string]interface{}{
		"deleted_time": time.Now(),
		"deleted_by":   deletedBy,
	}
	if err := db.Where("id = ?", id).Updates(updatedField).Error; err != nil {
		return code.GetError(gconstant.DBDeleteErr).Wrapf(err, "[TenantConfigDao] Delete fail, id:%d, deletedBy:%d", id, deletedBy)
	}
	return nil
}

func (d *TenantConfigDao) GetById(ctx context.Context, id uint) (*iammodel.TenantConfigEntity, error) {
	var entity iammodel.TenantConfigEntity
	db := d.DB(ctx).Table(d.TableName())
	if err := db.Where("id = ?", id).Find(&entity).Error; err != nil {
		return nil, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] GetById fail, id:%d", id)
	}
	return &entity, nil
}

func (d *TenantConfigDao) GetByCond(ctx context.Context, cond *TenantConfigCond) (*iammodel.TenantConfigEntity, error) {
	var entity iammodel.TenantConfigEntity
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entity).Error; err != nil {
		return nil, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] GetById fail, cond:%s", gutil.ToJsonString(cond))
	}
	return &entity, nil
}

func (d *TenantConfigDao) GetListByCond(ctx context.Context, cond *TenantConfigCond) (iammodel.TenantConfigEntityList, error) {
	var entityList iammodel.TenantConfigEntityList
	db := d.DB(ctx).Table(d.TableName())

	d.BuildCondition(db, cond)

	if err := db.Find(&entityList).Error; err != nil {
		return nil, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] GetListByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, nil
}

func (d *TenantConfigDao) GetPageListByCond(ctx context.Context, cond *TenantConfigCond) (iammodel.TenantConfigEntityList, int64, error) {
	db := d.DB(ctx).Model(&iammodel.TenantConfigEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)

	var count int64
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] GetPageListByCond count fail, cond:%s", gutil.ToJsonString(cond))
	}
	if cond.PageSize > 0 && cond.Page > 0 {
		db.Offset((cond.Page - 1) * cond.PageSize).Limit(cond.PageSize)
	}
	var entityList iammodel.TenantConfigEntityList
	if err := db.Find(&entityList).Error; err != nil {
		return nil, 0, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] GetPageListByCond find fail, cond:%s", gutil.ToJsonString(cond))
	}
	return entityList, count, nil
}

func (d *TenantConfigDao) CountByCond(ctx context.Context, cond *TenantConfigCond) (int64, error) {
	db := d.DB(ctx).Model(&iammodel.TenantConfigEntity{}).Table(d.TableName())

	d.BuildCondition(db, cond)
	var count int64
	if err := db.Count(&count).Error; err != nil {
		return 0, code.GetError(gconstant.DBFindErr).Wrapf(err, "[TenantConfigDao] CountByCond fail, cond:%s", gutil.ToJsonString(cond))
	}
	return count, nil
}

func (d *TenantConfigDao) BuildCondition(db *gorm.DB, cond *TenantConfigCond) {
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
