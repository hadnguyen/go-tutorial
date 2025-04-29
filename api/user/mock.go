package user

import (
	"go-tutorial/api/user/dto"
	"go-tutorial/api/user/model"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetUserPrivateProfile(user *model.User) (*dto.InfoPrivateUser, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.InfoPrivateUser), args.Error(1)
}

func (m *MockService) GetUserPublicProfile(userId primitive.ObjectID) (*dto.InfoPublicUser, error) {
	args := m.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*dto.InfoPublicUser), args.Error(1)
}

func (m *MockService) FindRoleByCode(code model.RoleCode) (*model.Role, error) {
	args := m.Called(code)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Role), args.Error(1)
}

func (m *MockService) FindRoles(roleIds []primitive.ObjectID) ([]*model.Role, error) {
	args := m.Called(roleIds)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Role), args.Error(1)
}

func (m *MockService) FindUserById(id primitive.ObjectID) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockService) FindUserByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockService) CreateUser(user *model.User) (*model.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockService) FindUserPrivateProfile(user *model.User) (*model.User, error) {
	args := m.Called(user)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockService) FindUserPublicProfile(userId primitive.ObjectID) (*model.User, error) {
	args := m.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockService) DeleteUserByEmail(email string) (bool, error) {
	args := m.Called(email)
	return args.Bool(0), args.Error(1)
}
