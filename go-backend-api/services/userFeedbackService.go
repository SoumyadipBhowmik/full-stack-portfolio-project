package services

import (
	"fmt"

	"github.com/SoumyadipBhowmik/go-backend/repositories"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

type UserFeedbackService struct {
	userRepository *repositories.UserFeedbackRepository
}

func NewUserFeedbackService(userFeedbackRepository *repositories.UserFeedbackRepository) *UserFeedbackService {
	return &UserFeedbackService{userRepository: userFeedbackRepository}
}

func (uf *UserFeedbackService) CreateNewFeedback(name, email, feedback string) (uuid.UUID, error) {

	id, err := uf.userRepository.CreateNewFeedback(name, email, feedback)
	if err != nil {
		fmt.Println(err)
	}

	var userid uuid.UUID
	copier.Copy(&userid, id)

	return userid, err

}
