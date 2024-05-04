package repository

import (
	"gorm.io/gorm"
	"winton/app/domain/dao"
)

type RoleRepository interface {
	FindAllRole() ([]dao.Role, error)
	FindRoleById(id int) (dao.Role, error)
	Save(role *dao.Role) (dao.Role, error)
	DeleteRoleById(id int) error
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func (r RoleRepositoryImpl) FindAllRole() ([]dao.Role, error) {
	var roles []dao.Role

	var err = r.db.Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r RoleRepositoryImpl) FindRoleById(id int) (dao.Role, error) {
	role := dao.Role{
		ID: id,
	}
	err := r.db.First(&role).Error
	if err != nil {
		return dao.Role{}, err
	}
	return role, nil
}

func (r RoleRepositoryImpl) Save(role *dao.Role) (dao.Role, error) {
	var err = r.db.Save(role).Error
	if err != nil {
		return dao.Role{}, err
	}
	return *role, nil
}

func (r RoleRepositoryImpl) DeleteRoleById(id int) error {
	err := r.db.Delete(&dao.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func RoleRepositoryInit(db *gorm.DB) *RoleRepositoryImpl {
	db.AutoMigrate(&dao.Role{})
	return &RoleRepositoryImpl{
		db: db,
	}
}
