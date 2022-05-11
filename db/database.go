package db

import (
	"fmt"

	"github.com/tsukiamaoto/bookShop-go/config"
	"github.com/tsukiamaoto/bookShop-go/model"
	"github.com/tsukiamaoto/bookShop-go/utils"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect(conf *config.Config) *gorm.DB {
	conn, err := gorm.Open(postgres.Open(conf.Databases["shopCart"].Source))
	if err != nil {
		fmt.Println("使用 gorm 連線 DB 發生錯誤，原因為", err)
		log.Error(err)
		return nil
	}

	return conn
}

func AutoMigrate(db *gorm.DB) {
	cartMigration(db)
	orderMigration(db)
	productMigration(db)
	sellMigration(db)
	userMigration(db)

	migrateProductTypes2Types(db)
}

func cartMigration(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.Cart)); err != nil {
		panic("資料庫Cart migration的失敗原因是" + err.Error())
	}
	fmt.Println("cart db migration 成功！")

	if err := db.AutoMigrate(new(*model.CartItem)); err != nil {
		panic("資料庫CartItem migration的失敗原因是" + err.Error())
	}
	fmt.Println("cartItem db migration 成功！")
}

func orderMigration(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.Order)); err != nil {
		panic("資料庫Order migration的失敗原因是" + err.Error())
	}
	fmt.Println("order db migration 成功！")

	if err := db.AutoMigrate(new(*model.OrderItem)); err != nil {
		panic("資料庫OrderItem migration的失敗原因是" + err.Error())
	}
	fmt.Println("orderItem db migration 成功！")
}

func productMigration(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.Product)); err != nil {
		panic("資料庫Product migration的失敗原因是" + err.Error())
	}
	fmt.Println("product db migration 成功！")

	if err := db.AutoMigrate(new(*model.Category)); err != nil {
		panic("Category migration的失敗原因是" + err.Error())
	}
	fmt.Println("category db migration 成功！")

	if err := db.AutoMigrate(new(*model.Type)); err != nil {
		panic("type migration的失敗原因是" + err.Error())
	}
	fmt.Println("type db migration 成功！")
}

func sellMigration(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.Seller)); err != nil {
		panic("資料庫Seller migration的失敗原因是" + err.Error())
	}
	fmt.Println("seller db migration 成功！")
}

func userMigration(db *gorm.DB) {
	if err := db.AutoMigrate(new(*model.User)); err != nil {
		panic("資料庫User migration的失敗原因是" + err.Error())
	}
	fmt.Println("user db migration 成功！")
}

func migrateProductTypes2Types(db *gorm.DB) {
	products := make([]*model.Product, 0)

	if err := db.Model(&model.Product{}).Preload("Categories").Find(&products).Error; err != nil {
		log.Error("Failed to find products with Preload Categories, the reason is", err)
	}

	for _, product := range products {
		for _, category := range product.Categories {
			var keys []string
			relations := utils.RelationMap(category.Types)
			// append root key
			keys = append(keys, "root")
			keys = append(keys, category.Types...)
			types := utils.BuildTypes(keys, relations)

			var parentId *int
			var rootTypeId *int
			for index, typeValue := range types {
				var isExistedType model.Type
				if err := db.Model(&model.Type{}).Where("name = ?", typeValue.Name).Limit(1).Find(&isExistedType).Error; err != nil {
					log.Error("Failed to find type with name, the reason is ", err)
				}
				// if type not found, created a new type
				// else type has found, save id for next type value
				if (model.Type{}) == isExistedType {
					typeValue.ParentID = parentId
					if err := db.Model(&model.Type{}).Create(&typeValue).Error; err != nil {
						log.Error("Failed to create type, the reason is ", err)
					}
					// save parent id for next type
					parentId = &typeValue.ID
				} else {
					parentId = &isExistedType.ID
				}

				// save root id for type id of category
				if index == 0 {
					rootTypeId = parentId
				}
			}

			// updated type of categroies
			if rootTypeId != nil && category.TypeID == nil {
				category.TypeID = rootTypeId
				if err := db.Model(&model.Category{}).Where("id = ?", category.ID).Updates(category).Error; err != nil {
					log.Error("Faild to updated type of category, the reason is ", err)
				}
			}
		}
	}

	log.Println("Successfully migrate ProductTypes to Types!")
}