package role

//go:generate mockgen -source=role.go -destination=role_mock.go -package=role -mock_names=Interface=MockRepository

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"go-scaffold/internal/app/model"
	"gorm.io/gorm"
	"time"
)

type RepositoryInterface interface {
	// FindList 列表查询
	FindList(ctx context.Context, param FindListParam, columns []string, order string) ([]*model.Role, error)

	// FindOneById 根据 id 查询详情
	FindOneById(ctx context.Context, id uint64, columns []string) (*model.Role, error)

	// Create 创建数据
	Create(ctx context.Context, role *model.Role) (*model.Role, error)

	// Save 保存数据
	Save(ctx context.Context, role *model.Role) (*model.Role, error)

	// Delete 删除数据
	Delete(ctx context.Context, role *model.Role) error
}

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewRepository(db *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		db:  db,
		rdb: rdb,
	}
}

var (
	cacheKeyFormat = model.Role{}.TableName() + "_%d"
	cacheExpire    = 3600
)

// FindListParam 列表查询参数
type FindListParam struct {
	Keyword string
}

// FindList 列表查询
func (r *Repository) FindList(ctx context.Context, param FindListParam, columns []string, order string) ([]*model.Role, error) {
	var roles []*model.Role
	query := r.db.Select(columns)

	if param.Keyword != "" {
		query.Where(
			r.db.Where("name LIKE ?", "%"+param.Keyword+"%"),
		)
	}

	err := query.Order(order).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

// FindOneById 根据 id 查询详情
func (r *Repository) FindOneById(ctx context.Context, id uint64, columns []string) (*model.Role, error) {
	m := new(model.Role)

	cacheValue, err := r.rdb.Get(
		ctx,
		fmt.Sprintf(cacheKeyFormat, id),
	).Bytes()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return nil, err
		}
	}

	if cacheValue != nil {
		if err = jsoniter.Unmarshal(cacheValue, m); err != nil {
			return nil, err
		}

		return m, nil
	}

	err = r.db.Select(columns).Where("id = ?", id).Take(m).Error
	if err != nil {
		return nil, err
	}

	cacheValue, err = jsoniter.Marshal(m)
	if err != nil {
		return nil, err
	}

	err = r.rdb.Set(
		ctx,
		fmt.Sprintf(cacheKeyFormat, id),
		string(cacheValue),
		time.Duration(cacheExpire)*time.Second,
	).Err()
	if err != nil {
		return nil, err
	}

	return m, nil
}

// Create 创建数据
func (r *Repository) Create(ctx context.Context, role *model.Role) (*model.Role, error) {
	if err := r.db.Create(role).Error; err != nil {
		return nil, err
	}

	cacheValue, err := jsoniter.Marshal(role)
	if err != nil {
		return nil, err
	}

	err = r.rdb.Set(
		ctx,
		fmt.Sprintf(cacheKeyFormat, role.Id),
		string(cacheValue),
		time.Duration(cacheExpire)*time.Second,
	).Err()
	if err != nil {
		return nil, err
	}

	return role, nil
}

// Save 保存数据
func (r *Repository) Save(ctx context.Context, role *model.Role) (*model.Role, error) {
	if err := r.db.Save(role).Error; err != nil {
		return nil, err
	}

	cacheValue, err := jsoniter.Marshal(role)
	if err != nil {
		return nil, err
	}

	err = r.rdb.Set(
		ctx,
		fmt.Sprintf(cacheKeyFormat, role.Id),
		string(cacheValue),
		time.Duration(cacheExpire)*time.Second,
	).Err()
	if err != nil {
		return nil, err
	}

	return role, nil
}

// Delete 删除数据
func (r *Repository) Delete(ctx context.Context, role *model.Role) error {
	if err := r.db.Delete(role).Error; err != nil {
		return err
	}

	err := r.rdb.Del(
		ctx,
		fmt.Sprintf(cacheKeyFormat, role.Id),
	).Err()
	if err != nil {
		return err
	}

	return nil
}
