package configs

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() *gorm.DB {

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	_ = viper.ReadInConfig()

	viperDBURL := viper.Get("DB_URL").(string)

	fmt.Println("dbUrl is\t\t", viperDBURL)

	db, err := gorm.Open(postgres.Open(viperDBURL), &gorm.Config{})
	//
	if err != nil {
		fmt.Println("err", err)
	}

	return db
}
