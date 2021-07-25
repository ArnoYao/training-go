package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/arnoyao/training-go/02week/dao"
)

func main() {

	cfg := dao.DefaultDBConfig()

	dao.Connect(cfg)

	userDB := dao.UserDB{DB: dao.GetDB()}

	user, err := userDB.Get(10000)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("未查到数据")
			return
		}
		fmt.Printf("未处理的异常:\n%+v\n", err)
		return
	}

	fmt.Printf("%v\n", *user)
}
