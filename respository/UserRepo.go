package respository

import (
	"fmt"
	"golangidev/domain"
	"log"
)

type UserInfo struct {
	Id         int
	Name       string
	Password   string
	Createtime string
	Address    string
	Age        int
}

type UserRepo interface {
	//QueryForList(sql string) (objs []*domain.User)
	//Save(domain.User)
	//Update(domain.User)
	//Delete(domain.User)

	BaseRepo
}

type UserRepoImpl struct {
}

func (ud *UserRepoImpl) GetTotalUsers(sql string) (totalRecords int) {
	db := dbConn()
	defer db.Close()

	row := db.Raw(sql).Row()
	var count int
	err := row.Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	return count

}
func (ud *UserRepoImpl) GetUserById(u_id int) (result domain.User) {
	db := dbConn()
	defer db.Close()
	userinfo := domain.User{}
	db.Table("userinfo").Find(&userinfo, "id=?", u_id)

	fmt.Println(userinfo)

	return userinfo

}
func (ud *UserRepoImpl) QueryForOne(sql string) (objs domain.User) {
	db := dbConn()
	defer db.Close()
	userinfo := domain.User{}
	var id, age int
	var name, password, address string
	var createtime string
	row := db.Raw("select id,name,password,create_time, address,age from userinfo").Row()

	row.Scan(&id, &name, &password, &createtime, &address, &age)

	userinfo.Id = id
	userinfo.Name = name
	userinfo.Password = password
	userinfo.CreateTime = createtime
	userinfo.Address = address
	userinfo.Age = age

	return

}
func (ud *UserRepoImpl) QueryForList(sql string) (objs []domain.User) {

	db := dbConn()
	defer db.Close()

	// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()

	rows, err := db.Raw("select id,name,password,create_time, address,age,enabled from userinfo where 1").Rows()
	userinfo := domain.User{}
	fmt.Println(rows.Columns())

	//	objs := []domain.User{}

	res := []domain.User{}
	/*
		cols, _ := rows.Columns()

		var length int = len(cols)
		//log.Println("rows length:", length)

		objs = make([]*domain.User, length)
		i := 0
	*/
	for rows.Next() {
		var id, age, enabled int
		var name, password, address string
		var createtime string

		err = rows.Scan(&id, &name, &password, &createtime, &address, &age, &enabled)
		if err != nil {
			panic(err.Error())
		}

		userinfo.Id = id
		userinfo.Name = name
		userinfo.Password = password
		userinfo.CreateTime = createtime
		userinfo.Address = address
		userinfo.Age = age
		userinfo.Enabled = enabled

		// Join each row on struct inside the Slice
		res = append(res, userinfo)

	}

	return res
}

func (ud *UserRepoImpl) SearchUsersByPaging(current_page int, per_page_num int) (objs []domain.User) {

	db := dbConn()
	defer db.Close()
	var offset = current_page * per_page_num
	// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()

	rows, err := db.Raw("select id,name,password,create_time, address,age,enabled from userinfo where 1 and (name like ? or address like ?) and (ENABLEd=1) and  (age BETWEEN 1 and 46)  limit ?,?", "%pan%", "%demo%", offset, per_page_num).Rows()
	userinfo := domain.User{}
	fmt.Println(rows.Columns())

	//	objs := []domain.User{}

	res := []domain.User{}
	/*
		cols, _ := rows.Columns()

		var length int = len(cols)
		//log.Println("rows length:", length)

		objs = make([]*domain.User, length)
		i := 0
	*/
	for rows.Next() {
		var id, age, enabled int
		var name, password, address string
		var createtime string

		err = rows.Scan(&id, &name, &password, &createtime, &address, &age, &enabled)
		if err != nil {
			panic(err.Error())
		}

		userinfo.Id = id
		userinfo.Name = name
		userinfo.Password = password
		userinfo.CreateTime = createtime
		userinfo.Address = address
		userinfo.Age = age
		userinfo.Enabled = enabled

		// Join each row on struct inside the Slice
		res = append(res, userinfo)

	}

	return res
}

func (ud *UserRepoImpl) GetTotalUsersByPaging(current_page int, per_page_num int) (objs []domain.User) {

	db := dbConn()
	defer db.Close()
	var offset = current_page * per_page_num
	// rows, err := db.Raw("select name, age, email from users where name = ?", "jinzhu").Rows()

	rows, err := db.Raw("select id,name,password,create_time, address,age,enabled from userinfo where 1  limit ?,?", offset, per_page_num).Rows()
	userinfo := domain.User{}
	fmt.Println(rows.Columns())

	//	objs := []domain.User{}

	res := []domain.User{}
	/*
		cols, _ := rows.Columns()

		var length int = len(cols)
		//log.Println("rows length:", length)

		objs = make([]*domain.User, length)
		i := 0
	*/
	for rows.Next() {
		var id, age, enabled int
		var name, password, address string
		var createtime string

		err = rows.Scan(&id, &name, &password, &createtime, &address, &age, &enabled)
		if err != nil {
			panic(err.Error())
		}

		userinfo.Id = id
		userinfo.Name = name
		userinfo.Password = password
		userinfo.CreateTime = createtime
		userinfo.Address = address
		userinfo.Age = age
		userinfo.Enabled = enabled

		// Join each row on struct inside the Slice
		res = append(res, userinfo)

	}

	return res
}

func (ud *UserRepoImpl) Save(user domain.User) (result int64) {
	db := dbConn()
	defer db.Close()

	log.Printf("new record is created")
	return db.Table("userinfo").Save(&user).RowsAffected

}

func (ud *UserRepoImpl) Hello(sql string) (result string) {

	return sql

}

func (ud *UserRepoImpl) Update(user domain.User) (result int64) {
	db := dbConn()
	defer db.Close()

	//return db.Model(domain.User{}).Update(&userinfo).RowsAffected
	return db.Table("userinfo").Where("id = ?", user.Id).Update(&user).RowsAffected

}

func (ud *UserRepoImpl) Delete(user domain.User) (result int64) {
	db := dbConn()
	defer db.Close()
	return db.Table("userinfo").Where("id = ?", user.Id).Delete(&user).RowsAffected

}
func (ud *UserRepoImpl) Remove(user domain.User) (result int64) {
	db := dbConn()
	defer db.Close()
	user.Enabled = 0
	return db.Table("userinfo").Where("id = ?", user.Id).Update(&user).RowsAffected
}
