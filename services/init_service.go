package services

import (
	"github.com/Rolas444/apigo_base/domain/models"
	"github.com/Rolas444/apigo_base/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type InitService struct {
	RoleRepo repository.RoleRepository
	UserRepo repository.UserRepository
}

func (s *InitService) InitRolesAndAdmin() error {
	existingRoles, err := s.RoleRepo.FindAll()
	if err != nil {
		return err
	}

	if len(existingRoles) == 0 {
		roles := []string{"admin", "user"}
		for _, roleName := range roles {
			// role, _ := s.RoleRepo.FindByName(roleName)
			// if err != nil && err != gorm.ErrRecordNotFound {
			// 	return err
			// }
			// println("Role: ", role)
			// if role == nil {
			println("Creating role: ", roleName)
			_, err := s.RoleRepo.Create(&models.Role{Name: roleName})
			if err != nil {
				return err
			}
			// }
		}
	}

	adminRole, err := s.RoleRepo.FindByName("admin")
	if err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("sadmin"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminUser := &models.User{
		Name:     "super admin",
		Email:    "7200ws@gmail.com",
		Password: string(hashedPassword),
		RoleID:   adminRole.ID,
	}

	_, err = s.UserRepo.FindByEmail(adminUser.Email)
	if err == gorm.ErrRecordNotFound {
		_, err = s.UserRepo.Create(adminUser)
		if err != nil {
			return err
		}
	}

	return nil

}
