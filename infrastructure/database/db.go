package database

import (
	"TheBoys/country"
	"TheBoys/infrastructure/config"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func IntDB() (*gorm.DB, *gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.Config.DbDsn), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   true,
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	productDb, err := gorm.Open(postgres.Open(config.Config.ProdutDb), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
			NoLowerCase:   true,
			NameReplacer:  strings.NewReplacer("CID", "Cid"),
		},
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	err = intiData(db, productDb)

	if err != nil {
		return nil, nil, err
	}
	return db, productDb, nil

}

func intiData(db *gorm.DB, pdb *gorm.DB) error {
	err := initCountry(db)
	if err != nil {
		return err
	}

	err = initState(db)

	if err != nil {
		return err
	}

	err = initRole(db)
	if err != nil {
		return err
	}

	err = initGender(db)
	if err != nil {
		return err
	}

	err = initUser(db)

	if err != nil {
		return err
	}

	err = productTypes(db, pdb)

	if err != nil {
		return err
	}

	err = initProductCategory(db)
	if err != nil {
		return err
	}
	return nil
}

func productTypes(db *gorm.DB, pdb *gorm.DB) error {
	var sleeveCount int

	err := db.Raw(`SELECT  COUNT(*) FROM "SleeveType"`).Scan(&sleeveCount).Error

	if err != nil {
		return err
	}

	type Common struct {
		Id   int64  `json:"id" gorm:"column:id"`
		Name string `json:"name" gorm:"column:name"`
	}
	var sleeveData []Common

	var sleeveBaseQuery strings.Builder

	sleeveBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "SleeveType"`)

	err = pdb.Raw(sleeveBaseQuery.String()).Scan(&sleeveData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding sleeve type`)
		fmt.Printf("Product sync error: %v\n", err)
		return err
	}
	if sleeveCount == 0 {
		for _, slee := range sleeveData {
			err = db.Exec(`INSERT INTO "SleeveType" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, slee.Name).Error
			if err != nil {
				return err
			}
		}
	}

	/**NECK TYPE****/

	var neckCount int

	err = db.Raw(`SELECT  COUNT(*) FROM "NeckType"`).Scan(&neckCount).Error

	if err != nil {
		return err
	}

	var neckData []Common

	var neckBaseQuery strings.Builder

	neckBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "NeckType"`)

	err = pdb.Raw(neckBaseQuery.String()).Scan(&neckData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding neck type`)
		fmt.Printf("Product sync error: %v\n", err)
		return err
	}
	if neckCount == 0 {
		for _, neck := range neckData {
			err = db.Exec(`INSERT INTO "NeckType" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, neck.Name).Error
			if err != nil {
				return err
			}
		}
	}

	/**KURTAS LENGTH*****/

	var kurtasCount int

	err = db.Raw(`SELECT  COUNT(*) FROM "KurtasLengthType"`).Scan(&kurtasCount).Error

	if err != nil {
		return err
	}

	var kurtasLengthData []Common

	var kurtasLengthBaseQuery strings.Builder

	kurtasLengthBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "KurtasLengthType"`)

	err = pdb.Raw(kurtasLengthBaseQuery.String()).Scan(&kurtasLengthData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding kurtas length type`)
		fmt.Printf("Product sync error: %v\n", err)
		return err
	}
	if kurtasCount == 0 {
		for _, kurtasData := range kurtasLengthData {
			err = db.Exec(`INSERT INTO "KurtasLengthType" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, kurtasData.Name).Error
			if err != nil {
				return err
			}
		}
	}
	/**BOTTOM LENGTH******/

	var btmTypeCount int

	err = db.Raw(`SELECT  COUNT(*) FROM "TypesOfBottom"`).Scan(&btmTypeCount).Error

	if err != nil {
		return err
	}

	var bottomTypeData []Common

	var bottomTypeDataBaseQuery strings.Builder

	bottomTypeDataBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "TypesOfBottom"`)

	err = pdb.Raw(bottomTypeDataBaseQuery.String()).Scan(&bottomTypeData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding bottom type`)
		fmt.Printf("Product error: %v\n", err)
		return err
	}
	if btmTypeCount == 0 {
		for _, btm := range bottomTypeData {
			err = db.Exec(`INSERT INTO "TypesOfBottom" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, btm.Name).Error
			if err != nil {
				return err
			}
		}
	}
	/**BOTTOM PLEATS*******/
	var pleatsTypeCount int

	err = db.Raw(`SELECT  COUNT(*) FROM "TypesOfPleats"`).Scan(&pleatsTypeCount).Error

	if err != nil {
		return err
	}

	var pleatsTypeData []Common

	var pleatsTypeDataBaseQuery strings.Builder

	pleatsTypeDataBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "TypesOfPleats"`)

	err = pdb.Raw(pleatsTypeDataBaseQuery.String()).Scan(&pleatsTypeData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding pleats type`)
		fmt.Printf("Product error: %v\n", err)
		return err
	}
	if pleatsTypeCount == 0 {
		for _, btmPleats := range pleatsTypeData {
			err = db.Exec(`INSERT INTO "TypesOfPleats" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, btmPleats.Name).Error
			if err != nil {
				return err
			}
		}
	}
	/**BOTTOM LENGTH TYPE****/
	var btmLengthTypeCount int

	err = db.Raw(`SELECT  COUNT(*) FROM "TypesOfLengthBottom"`).Scan(&btmLengthTypeCount).Error

	if err != nil {
		return err
	}
	var bottomLengthTypeData []Common

	var bottomLengthBaseQuery strings.Builder

	bottomLengthBaseQuery.WriteString(`SELECT "id" AS "id","name" AS "name" FROM "TypesOfLengthBottom"`)

	err = pdb.Raw(bottomLengthBaseQuery.String()).Scan(&bottomLengthTypeData).Error

	if err != nil {
		fmt.Printf(`simulated error while adding bottom length type`)
		fmt.Printf("Product error: %v\n", err)
		return err
	}
	if btmLengthTypeCount == 0 {
		for _, btmlen := range bottomLengthTypeData {
			err = db.Exec(`INSERT INTO "TypesOfLengthBottom" ("name","updatedAt","createdAt") VALUES (?,NOW(),NOW())`, btmlen.Name).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func initCountry(db *gorm.DB) error {
	var count int

	var dbName string
	db.Raw("SELECT current_database()").Row().Scan(&dbName)

	fmt.Println("Current database:", dbName)

	err := db.Raw(`SELECT COUNT(*) FROM "Country";`).Scan(&count).Error
	if err != nil {
		return err
	}

	if count == 0 {
		for _, countryData := range country.AllCountryData {
			err = db.Exec(`INSERT INTO "Country" ("name","timezoneOffset","dialCode","updatedAt","createdAt") VALUES(?,?,?, NOW() , NOW() )`, countryData.Name, countryData.TimeOffset, countryData.DailCode).Error
			if err != nil {
				return err
			}
		}
	}

	return nil

}
func initState(db *gorm.DB) error {
	var count int
	err := db.Raw(`SELECT COUNT(*) FROM "State" `).Scan(&count).Error
	if err != nil {
		return nil
	}
	if count == 0 {
		for _, stateData := range country.IndianStates {
			err = db.Exec(`INSERT INTO "State" ("name","countryId","updatedAt","createdAt") VALUES(?,77, NOW(), NOW())`, stateData).Error

			if err != nil {
				return err
			}
		}
	}
	return nil
}

func initUser(db *gorm.DB) error {
	var count int

	err := db.Raw(`SELECT COUNT(*) FROM "User" `).Scan(&count).Error
	if err != nil {
		return nil
	}
	if count == 0 {
		var id int
		err = db.Raw(`INSERT INTO "User" ("username","email","mobile","roleId","genderId","updatedAt","createdAt") VALUES('venukumar','venukumar.rvk@gmail.com','6380666892',1,1, NOW(), NOW()) RETURNING "id"`).Scan((&id)).Error
		if err != nil {
			return err
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("Venu@123"), bcrypt.DefaultCost)
		if err != nil {
			return err
		}

		fmt.Println("hash password", hashedPassword, []byte("Venu@123"))

		err = db.Exec(`INSERT INTO "UserPassword" ("password","userId","updatedAt","createdAt")VALUES(?,?,NOW(),NOW())`, string(hashedPassword), id).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func initRole(db *gorm.DB) error {
	var count int

	err := db.Raw(`SELECT COUNT(*) FROM "Roles" `).Scan(&count).Error
	if err != nil {
		return nil
	}

	roleName := []string{"Admin", "User", "StoreOwners"}

	if count == 0 {
		for _, roles := range roleName {
			err = db.Exec(`INSERT INTO "Roles" ("name","updatedAt","createdAt") VALUES(?, NOW(), NOW())`, roles).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func initGender(db *gorm.DB) error {
	var count int

	err := db.Raw(`SELECT COUNT(*) FROM "Gender" `).Scan(&count).Error
	if err != nil {
		return nil
	}

	var gen = []string{"Male", "Female"}
	if count == 0 {

		for _, data := range gen {
			err := db.Exec(`INSERT INTO "Gender"("name","updatedAt","createdAt") VALUES(?, NOW(),NOW())`, data).Error

			if err != nil {
				return nil
			}

		}

	}
	return nil
}

func initProductCategory(db *gorm.DB) error {
	var count int

	err := db.Raw(`SELECT COUNT(*) FROM "ProductCategory"`).Scan(&count).Error
	if err != nil {
		return nil
	}
	var (
		top = []string{
			"T-Shirts", "Polo T Shirts", "Casual Shirts", "Formal Shirts",
			"Suits & Blazers", "Jackets", "Sweaters & Sweatshirts",
		}
		bottom = []string{
			"Jeans", "Casual Trousers", "Formal Trousers", "Joggers",
			"Shorts", "Three Fourth",
		}
		ethnic = []string{
			"Kurtas", "Ethnic Wear Sets", "Nehru Jackets", "Ethnic Bottom Wear",
		}
		sports = []string{
			"T-Shirts", "Shorts", "Track Pants", "Track Suits",
		}
		footwear = []string{
			"Casual Shoes", "Formal Shoes", "Sports Shoes",
			"Slippers", "Sandals", "Socks",
		}
		accessories = []string{
			"Caps And Hats", "Ties", "Handkerchiefs", "Belts",
			"Bags", "Wallets", "Watches", "Sun glass",
		}
		innerWear      = []string{"Briefs", "Boxers", "Vests"}
		sizesForShirts = []string{"S", "M", "L", "XL", "XXL", "XXXL"}
		sizesForPants  = []string{"28", "30", "32", "36", "38", "40", "42"}
		footWears      = []string{"6", "7", "8", "9", "10"}
		empty          = []string{"Quanity"}
	)

	if count == 0 {
		categories := []string{
			"Top", "Bottom", "Ethnic", "Sports", "Fragrances",
			"Footwear", "Accessories", "Inner", "Watches",
		}

		for _, ctg := range categories {

			var ctData struct {
				Name string `json:"name" gorm:"column:name"`
				Id   int    `json:"id" gorm:"column:id"`
			}

			var ctQuery strings.Builder

			ctQuery.WriteString(fmt.Sprintf(`INSERT INTO "ProductCategory" ("name", "genderId", "createdAt", "updatedAt") VALUES('%s', '%d', NOW(), NOW()) RETURNING "name", "id"`, ctg, 1))
			err = db.Raw(ctQuery.String()).Scan(&ctData).Error

			if err != nil {
				fmt.Printf("Error inserting category %s: %s\n", ctg, err.Error())
				return err
			}
			ctgName := ctData.Name

			fmt.Println("data from the name", ctgName)
			switch ctgName {
			case "Top":
				AddProductTypes(uint(ctData.Id), top, sizesForShirts, db)
			case "Bottom":
				AddProductTypes(uint(ctData.Id), bottom, sizesForPants, db)
			case "Ethnic":
				AddProductTypes(uint(ctData.Id), ethnic, sizesForShirts, db)
			case "Sports":
				AddProductTypes(uint(ctData.Id), sports, sizesForShirts, db)
			case "Footwear":
				AddProductTypes(uint(ctData.Id), footwear, footWears, db)
			case "Accessories":
				AddProductTypes(uint(ctData.Id), accessories, empty, db)
			case "Inner":
				AddProductTypes(uint(ctData.Id), innerWear, sizesForPants, db)
			case "Fragrances":
				AddProductTypes(uint(ctData.Id), []string{"Fragrances"}, empty, db)
			case "Watches":
				AddProductTypes(uint(ctData.Id), []string{"Watches"}, empty, db)
			default:
				fmt.Printf("There is no ProductCategory like this: %s\n", ctData.Name)
			}
		}
	}

	return nil
}
func AddProductTypes(categoryID uint, productTypes []string, sizes []string, db *gorm.DB) error {

	for _, productType := range productTypes {

		var pt struct {
			ItemsName string `json:"itemsName" gorm:"column:itemsName"`
			Id        int    `json:"id" gorm:"column:id"`
		}

		var ptQuery strings.Builder
		ptQuery.WriteString(fmt.Sprintf(`INSERT INTO "ProductType" ("itemsName", "productCategoryId", "createdAt", "updatedAt") VALUES('%s', %d, NOW(), NOW()) RETURNING "itemsName", "id"`, productType, categoryID))
		err := db.Raw(ptQuery.String()).Scan(&pt).Error

		if err != nil {
			return err
		}

		if pt.ItemsName == "Socks" {
			sizes = []string{"6-8", "8-10", "10-12", "12-14"}
		} else if pt.ItemsName == "Track Pants" || pt.ItemsName == "Shorts" {
			sizes = []string{"28", "30", "32", "36", "38", "40", "42"}
		}

		for _, size := range sizes {
			var sizeQuery strings.Builder
			sizeQuery.WriteString(fmt.Sprintf(`INSERT INTO "ProductTypeSize" ("size", "ietmsId", "createdAt", "updatedAt") VALUES('%s', %d, NOW(), NOW())`, size, pt.Id))
			err := db.Exec(sizeQuery.String()).Error

			if err != nil {
				return err
			}
		}
	}

	return nil
}
