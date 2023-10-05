package configs

import "prakerja12/models/user"

func initMigrate() {
	DB.AutoMigrate(&user.User{})
}
