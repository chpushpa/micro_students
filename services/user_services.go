package services

import (
	"github.com/chpushpa/micro_student/domain/httperrors"
	"github.com/chpushpa/micro_student/domain/students"
)

/*var (
	StudentService = studentsService{}

	registeredStudents = map[int64]*students.Student{}
	//currentUserId      int64 = 1
)

type studentsService struct{}*/

func Create(user students.Student) (*students.Student, *httperrors.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	//save to database
	if err := user.Save(); err != nil {
		return nil, err
	}
	/*user.Id = currentUserId
	currentUserId++

	registeredStudents[user.Id] = &user */

	return &user, nil
}

/*func (service studentsService) Get(userId int64) (*students.Student, *httperrors.HttpError) {
	if user := registeredStudents[userId]; user != nil {
		return user, nil
	}
	return nil, httperrors.NewNotFoundError(fmt.Sprintf("user %d not found", userId))
}*/
