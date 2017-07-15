package service

import (
	"fmt"
	"golangidev/domain"
	"golangidev/respository"
)

type UserService interface {
	QueryForList(sql string) (objs []*domain.User)
	Save(domain.User)
	Update(domain.User) (result int64)
	Remove(domain.User) (result int64)
	GetTotalUsers(sql string) (result int64)
	GetTotalUsersByPaging(current_page int, per_page_num int) (objs []domain.User)
	SearchUsersByPaging(current_page int, per_page_num int) (objs []domain.User)
	GetUserById(id int) (user domain.User)
	BaseService
}
type UserServiceImpl struct {
}

var userDao *respository.UserRepoImpl = new(respository.UserRepoImpl)

func (us *UserServiceImpl) codingEveryday() string {
	return fmt.Sprintf("just for test")
}

func (us *UserServiceImpl) QueryForList(sql string) (objs []domain.User) {

	//result := []domain.User{}

	//result = userDao.QueryForList("select * from userinfo")
	return userDao.QueryForList(sql)
}
func (us *UserServiceImpl) Save(user domain.User) (result int64) {
	return userDao.Save(user)
}
func (us *UserServiceImpl) Update(user domain.User) (result int64) {
	return userDao.Update(user)
}
func (us *UserServiceImpl) GetTotalUsers(sql string) (result int) {

	return userDao.GetTotalUsers(sql)
}
func (us *UserServiceImpl) GetUserById(id int) (user domain.User) {
	return userDao.GetUserById(id)
}
func (us *UserServiceImpl) GetTotalUsersByPaging(current_page int, per_page_num int) (objs []domain.User) {

	return userDao.GetTotalUsersByPaging(current_page, per_page_num)
}
func (us *UserServiceImpl) SearchUsersByPaging(current_page int, per_page_num int) (objs []domain.User) {

	return userDao.SearchUsersByPaging(current_page, per_page_num)
}
