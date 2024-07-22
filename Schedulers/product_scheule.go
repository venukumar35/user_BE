package schedulers

import (
	"TheBoys/infrastructure/config"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

func NewSchedule(db *gorm.DB, pdb *gorm.DB) *Schedule {
	return &Schedule{DB: db, productDb: pdb}
}

type Schedule struct {
	DB        *gorm.DB
	productDb *gorm.DB
}

func (s Schedule) ProductSchedulers() {
	scheduler := gocron.NewScheduler(time.Local)

	scheduler.Every(10).Minutes().Do(func() {
		if config.Config.SyncData {
			SyncProductDetails(s.DB, s.productDb)
		}
	})

	scheduler.StartAsync()

}
func SyncProductDetails(db *gorm.DB, pdb *gorm.DB) {

	var baseQuery strings.Builder

	var userData []UserDetails
	baseQuery.WriteString(`SELECT "id" AS "uId" FROM "User"`)

	err := pdb.Raw(baseQuery.String()).Scan(&userData).Error

	if err != nil {

		fmt.Errorf(`Error in user fetching : %v\n`, err)
		return
	}
	var wg sync.WaitGroup
	var lock sync.Mutex
	productChannel := make(chan ProductDetails, 100)
	productChannelForClr := make(chan ProductDetails, 100)
	productChannelForseason := make(chan ProductDetails, 100)

	clrChannel := make(chan productColorDetails, 100)

	var seasonalData []seasonalDetails

	var dataOfClrs []productColorDetails
	clrChannnelForSize := make(chan productColorDetails, 100)
	clrChannelForImage := make(chan productColorDetails, 100)
	var productDetails []ProductDetails

	var clrImage []imgDetails
	wg.Add(5)

	//Fan-out pattern gorountine
	go func() {
		defer close(productChannelForClr)
		defer close(productChannelForseason)

		for data := range productChannel {
			productChannelForClr <- data
			productChannelForseason <- data

		}
	}()

	go func() {
		defer close(clrChannnelForSize)
		defer close(clrChannelForImage)

		for v := range clrChannel {
			clrChannnelForSize <- v
			clrChannelForImage <- v
		}
	}()
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		defer close(productChannel)

		for _, data := range userData {

			var productData []ProductDetails
			var productQuery strings.Builder

			productQuery.WriteString(fmt.Sprintf(`SELECT "Product"."userId" AS "uId", "Product"."id" AS "pId" ,
			    "Product"."title" AS "title","Product"."price" AS "price",
			    "User"."username" AS "name","User"."email" AS "email","User"."mobile" AS "mobile",
				"CommonDescription"."id" AS "cmId","TopDescription"."id" AS "TpId",
				"BottomDescription"."id" AS "btId","KurtasDescription"."id" AS "kurId",
				"ShoesDescription"."id" AS "sdId","InnersDescription"."id" AS "inId",
				"WatchesDescription"."id" AS "wdId","PerfumesDescription"."id" AS "pdId",
				"ProductType"."id" AS "pdTyId","ProductCategory"."id" AS "pdCtId","ProductType"."itemsName" AS "ProductTypeName",
				"CommonDescription"."fit" AS "fit","CommonDescription"."materail" AS "material",
				"CommonDescription"."care" AS "care","CommonDescription"."brandName" AS "brandName",
				"CommonDescription"."origin" AS "origin","CommonDescription"."occasion" AS "occasion","CommonDescription"."specialFeature" AS "specialFeature",
				"TopDescription"."productDescription" AS "productDescription","TopDescription"."sleeveTypeId" AS "sleeveTypeId",
				"TopDescription"."weight" AS "weight" , "TopDescription"."chest" AS "chest" , "TopDescription"."shoulder" AS "shoulder" , "TopDescription"."neckTypeId" AS "neckTypeId",
				"TopDescription"."type" AS "type","TopDescription"."colorFamily" AS "colorFamily","TopDescription"."printAndPattern" AS "printAndPattern",
				"TopDescription"."length" AS "length","TopDescription"."pocket" AS "pocket","TopDescription"."commonDescriptionId" AS "commonDescriptionId",
			    "BottomDescription"."productDescription" AS "bottomproductDescription",
				"BottomDescription"."weight" AS "BtWeight", "BottomDescription"."printAndPattern" AS "BtPrintAndPattern",
			    "BottomDescription"."length" AS "btLength","BottomDescription"."waist" AS "btWaist","BottomDescription"."hip" AS "btHip","BottomDescription"."commonDescriptionId" AS "btcommonDescriptionId","BottomDescription"."type" AS "btType",
				"BottomDescription"."colorFamily" AS "btcolorFamily","BottomDescription"."pocket" AS "btPocket",
				"BottomDescription"."beltLoop" AS "btBeltLoop","BottomDescription"."typeOfPantId" AS "typeOfPantId","BottomDescription"."typesOfPleatsId" AS "typesOfPleatsId","BottomDescription"."typesOfLengthId" AS "typesOfLengthId",
				"KurtasDescription"."work" AS "work","KurtasDescription"."transparencyOfTheFabric" AS "transparencyOfTheFabric","KurtasDescription"."kurtasLengthTypeId" AS "length",
				"KurtasDescription"."productDescription" AS "productDescription","KurtasDescription"."chest" AS "chest","KurtasDescription"."productDescription" AS "productDescription",
				"KurtasDescription"."shoulder" AS "shoulder","KurtasDescription"."weight" AS "weight","KurtasDescription"."colorFamily" AS "colorFamily","KurtasDescription"."pocket" AS "pocket","KurtasDescription"."type" AS "type",
				"KurtasDescription"."printAndpattern" AS "printAndpattern","KurtasDescription"."kurtasNeckTypeId" AS "kurtasNeckTypeId","KurtasDescription"."kurtasSleeveTypeId" AS "kurtasSleeveTypeId","KurtasDescription"."commonDescriptionId" AS "commonDescriptionId",
			    "ShoesDescription"."printAndPattern" AS "printAndpattern", "ShoesDescription"."type" AS "type","ShoesDescription"."productDescription" AS "productDescription","ShoesDescription"."colorFamily" AS "colorFamily","ShoesDescription"."pattern" AS "pattern",
				"ShoesDescription"."footLength" AS "footLength", "ShoesDescription"."soleMaterial" AS "soleMaterial", "ShoesDescription"."upperMaterial" AS "upperMaterial", "ShoesDescription"."closure" AS "closure", "ShoesDescription"."toeType" AS "toeType",
				"ShoesDescription"."packageContains" AS "packageContains","ShoesWarranty"."id" AS "warrentyId", "ShoesWarranty"."warrantyPeriod" AS "WarrantyYear",
			    "ShoesWarranty"."shoesDescriptionId" AS "shoesWarrentyId",
				"WatchesDescription"."type" AS "type","WatchesDescription"."weight" AS "weight" ,"WatchesDescription"."printAndPattern" AS "printAndpattern","WatchesDescription"."colorFamily" AS "colorFamily","WatchesDescription"."productDescription" AS "productDescription",
			    "WatchesDescription"."dialShape" AS "dialShape","WatchesDescription"."dialDiameter" AS "dialDiameter","WatchesDescription"."dialColor" AS "dialColor","WatchesDescription"."dialShape" AS "dialShape",
				"WatchesDescription"."strapColor" AS "strapColor","WatchesWarranty"."watchsId" AS "watchesWarrentyId","WatchesWarranty"."warrantyPeriod" AS "WarrantyYear","WatchesWarranty"."id" AS "warrentyId",
				"PerfumesDescription"."id" AS "PerfumeId","PerfumesDescription"."materialDescription" AS "materialDescription",
			    "PerfumesDescription"."type" AS "type","PerfumesDescription"."weight" AS "weight","PerfumesDescription"."productDescription" AS "productDescription",
				"InnersDescription"."type" AS "type","InnersDescription"."productDescription" AS "productDescription","InnersDescription"."weight" AS "weight",
				"InnersDescription"."length" AS "length","InnersDescription"."waistRise" AS "BtWaist","InnersDescription"."printAndPattern" AS "printAndpattern",
				"InnersDescription"."packageContains" AS "packageContains","InnersDescription"."lookAndFeel" AS "lookAndFeel","InnersDescription"."colorFamily" AS "colorFamily","InnersDescription"."vestsSleeveTypeId" AS "SleeveTypeId",
				"InnersDescription"."vestsNeckTypeId" AS "neckTypeId","InnersDescription"."commonDescriptionId" AS "commonDescriptionId","InnersDescription"."multiColors" AS "multiColors",
			    "Ts"."name" AS "sleeveName","Tn"."name" AS "neckName",
				"kt"."name" AS "kurtasTypeName","ks"."name" AS "sleeveName","kn"."name" AS "neckName",
				"bt"."name" AS "btTypeName","btp"."name" AS "btmPleatName","btl"."name" AS "btmLengthName",
				"InS"."name" AS "sleeveName","InN"."name" AS "neckName","ProductCategory"."name" AS "productCategoryName",
				CASE
				WHEN "BottomDescription"."topDescriptionId" IS NOT NULL THEN "BottomDescription"."id"  ELSE NULL
				END AS "topbtId",
				CASE
				WHEN "BottomDescription"."kurtasDescriptionId" IS NOT NULL THEN "BottomDescription"."id"  ELSE NULL
				END AS "kurtaBt"
				FROM "Product"
				INNER JOIN "CommonDescription" ON "CommonDescription"."productId" = "Product"."id"
				LEFT JOIN "TopDescription" ON "TopDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "SleeveType" AS "Ts" ON "Ts"."id"="TopDescription"."sleeveTypeId"
				LEFT JOIN "NeckType" AS "Tn" ON "Tn"."id"="TopDescription"."neckTypeId"
				LEFT JOIN "BottomDescription" ON "BottomDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "TypesOfBottom" AS "bt" ON "bt"."id" = "BottomDescription"."typeOfPantId"
				LEFT JOIN "TypesOfPleats" AS "btp" ON "btp"."id" = "BottomDescription"."typesOfPleatsId"
				LEFT JOIN "TypesOfLengthBottom" AS "btl" ON "btl"."id" = "BottomDescription"."typesOfLengthId"
			    LEFT JOIN "KurtasDescription" ON "KurtasDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "SleeveType" AS "ks" ON "ks"."id"="KurtasDescription"."kurtasSleeveTypeId"
			    LEFT JOIN "NeckType" AS "kn" ON "kn"."id"="KurtasDescription"."kurtasNeckTypeId"
			    LEFT JOIN "KurtasLengthType" AS "kt" ON "kt"."id"="KurtasDescription"."kurtasLengthTypeId"
				LEFT JOIN "ShoesDescription" ON "ShoesDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "Warranty" AS "ShoesWarranty" ON "ShoesWarranty"."shoesDescriptionId"="ShoesDescription"."id"
				LEFT JOIN "InnersDescription" ON "InnersDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "SleeveType" AS "InS" ON "InS"."id"="InnersDescription"."vestsSleeveTypeId"
				LEFT JOIN "NeckType" AS "InN" ON "InN"."id"="InnersDescription"."vestsNeckTypeId"
				LEFT JOIN "WatchesDescription" ON "WatchesDescription"."commonDescriptionId"="CommonDescription"."id"
				LEFT JOIN "Warranty" AS "WatchesWarranty" ON "WatchesWarranty"."watchsId"="WatchesDescription"."id"
				LEFT JOIN "PerfumesDescription" ON "PerfumesDescription"."commonDescriptionId"="CommonDescription"."id"
				INNER JOIN "ProductType" ON "ProductType"."id"="Product"."productTypeId"
				INNER JOIN "User" ON "User"."id" = "Product"."userId"
				INNER JOIN "ProductCategory" ON "ProductCategory"."id"="ProductType"."productCategoryId" WHERE "Product"."userId"='%d' AND "Product"."isSync" ='%t' `, data.UID, false))

			err := pdb.Raw(productQuery.String()).Scan(&productData).Error

			if err != nil {
				fmt.Printf(`simulated error for user id :  %d `, data.UID)
				fmt.Printf("product sync error: %v\n", err)
				continue
			}

			mut.Lock()
			productDetails = append(productDetails, productData...)
			mut.Unlock()

			for _, pId := range productData {
				productChannel <- pId
			}

		}

	}(&wg, &lock)

	go func(wg *sync.WaitGroup, mut *sync.Mutex, channels <-chan ProductDetails) {
		defer wg.Done()

		for pid := range channels {
			var firstQuery strings.Builder
			var localSeasonal []seasonalDetails

			firstQuery.WriteString(fmt.Sprintf(`SELECT "SeasonalDresses"."id" AS "id" , "SeasonalDresses"."productId" AS "pId","Product"."userId" AS "uId",
			"Seasonal"."seasonalName" AS "seasonalName"
			FROM "SeasonalDresses"
	    	INNER JOIN "Product" ON "SeasonalDresses"."productId"="Product"."id"
	    	INNER JOIN "Seasonal" ON "Seasonal"."id"="SeasonalDresses"."seasonalId"
		    WHERE "SeasonalDresses"."productId" = %d`, pid.PID))

			err := pdb.Raw(firstQuery.String()).Scan(&localSeasonal).Error

			if err != nil {
				fmt.Printf(`simulated error for product  %d `, pid.PID)
				fmt.Printf("Product sync error: %v\n", err)
				continue
			}

			lock.Lock()
			seasonalData = append(seasonalData, localSeasonal...)
			lock.Unlock()
		}
	}(&wg, &lock, productChannelForseason)

	go func(wg *sync.WaitGroup, mut *sync.Mutex, channel <-chan ProductDetails) {
		defer wg.Done()
		defer close(clrChannel)
		for data := range channel {
			var secondBaseQuery strings.Builder
			var localDetails []productColorDetails

			secondBaseQuery.WriteString(fmt.Sprintf(`SELECT "ProductColor"."id" AS "id" , "ProductColor"."productId" AS "pId" ,"Product"."userId" AS "uId",
			"ProductColor"."colors" AS "colors"
			FROM "ProductColor"
			INNER JOIN "Product" ON "Product"."id"="ProductColor"."productId"
			WHERE "ProductColor"."productId" = %d`, data.PID))
			err := pdb.Raw(secondBaseQuery.String()).Scan(&localDetails).Error

			if err != nil {
				fmt.Printf("simulated error for product  %d ", data.PID)
				fmt.Printf("Product sync error: %v\n", err)
				continue
			}

			mut.Lock()
			dataOfClrs = append(dataOfClrs, localDetails...)
			mut.Unlock()

			for _, clId := range localDetails {
				clrChannel <- clId
			}
		}
	}(&wg, &lock, productChannelForClr)

	var sizeData []SizeDetails

	go func(wg *sync.WaitGroup, lock *sync.Mutex, channel <-chan productColorDetails) {
		defer wg.Done()
		for sData := range channel {
			var thirdBaseQuery strings.Builder
			var colorSize []SizeDetails

			thirdBaseQuery.WriteString(fmt.Sprintf(`SELECT "ProductAviableSizes"."id" AS "id" , "ProductAviableSizes"."productColorId" AS "clrId", "ProductColor"."productId" AS "pId","Product"."userId" AS "uId",
			"ProductTypeSize"."size" AS "size","ProductAviableSizes"."quantity" AS "quantity"
			FROM "ProductAviableSizes"
			INNER JOIN "ProductColor" ON "ProductColor".id = "ProductAviableSizes"."productColorId"
			INNER JOIN "ProductTypeSize" ON "ProductTypeSize"."id"= "ProductAviableSizes"."productTypeSizeId"
			INNER JOIN "Product" ON "Product"."id"="ProductColor"."productId"
			WHERE "ProductAviableSizes"."productColorId" = %d`, sData.Id))

			err := pdb.Raw(thirdBaseQuery.String()).Scan(&colorSize).Error

			if err != nil {
				fmt.Printf(`simulated error product color id %d`, sData.Id)
				fmt.Printf("Product sync error: %v\n", err)
				continue
			}
			lock.Lock()
			sizeData = append(sizeData, colorSize...)
			lock.Unlock()

		}
	}(&wg, &lock, clrChannnelForSize)

	go func(wg *sync.WaitGroup, lock *sync.Mutex, channel <-chan productColorDetails) {
		defer wg.Done()

		for imgData := range channel {

			var fourthBaseQuery strings.Builder

			var localImgDetails []imgDetails

			fourthBaseQuery.WriteString(fmt.Sprintf(`SELECT "ProductImages"."id" AS "id" , "ProductImages"."productColorId" AS "cId","Product"."userId" AS "uId", "ProductColor"."productId" AS "pId",
			"ProductImages"."imageUrl" AS "imageUrl","ProductImages"."imageBytes" AS "imageBytes"
			FROM "ProductImages"
			INNER JOIN "ProductColor" ON "ProductColor".id = "ProductImages"."productColorId"
			INNER JOIN "Product" ON "Product"."id"="ProductColor"."productId"
			WHERE "ProductImages"."productColorId" = %d`, imgData.Id))

			err := pdb.Raw(fourthBaseQuery.String()).Scan(&localImgDetails).Error

			if err != nil {
				fmt.Printf(`simulated error product color id %d`, imgData.Id)
				fmt.Printf("Product sync error: %v\n", err)
				continue
			}

			lock.Lock()
			clrImage = append(clrImage, localImgDetails...)
			lock.Unlock()
		}
	}(&wg, &lock, clrChannelForImage)

	wg.Wait()

	type commonStruct struct {
		Product           []ProductDetails
		ProductColor      []productColorDetails
		ProductClrImg     []imgDetails
		ProductClrSize    []SizeDetails
		ProductSeasonalId []seasonalDetails
	}

	mapData := make(map[int]commonStruct)

	var instedClrIds []int
	var insertedImgIds []int
	var insertedSizesIds []int
	var insertedSeasonalIds []int

	var clrCheck bool = true
	var imgCheck bool = true
	var sizeCheck bool = true
	var seasonalCheck bool = true

	for _, v := range productDetails {
		key := v.UId
		if _, exist := mapData[key]; !exist {
			mapData[key] = commonStruct{}
		}
		productData := mapData[key]
		productData.Product = append(productData.Product, ProductDetails{
			UId:                      v.UId,
			Name:                     v.Name,
			Email:                    v.Email,
			Mobile:                   v.Mobile,
			Title:                    v.Title,
			Price:                    v.Price,
			PID:                      v.PID,
			ProductTypeName:          v.ProductTypeName,
			CmId:                     v.CmId,
			Fit:                      v.Fit,
			Material:                 v.Material,
			Care:                     v.Care,
			BrandName:                v.BrandName,
			Origin:                   v.Origin,
			Occasion:                 v.Occasion,
			SpecialFeature:           v.SpecialFeature,
			TpId:                     v.TpId,
			ProductDescription:       v.ProductDescription,
			SleeveTypeId:             v.SleeveTypeId,
			Weight:                   v.Weight,
			Chest:                    v.Chest,
			Shoulder:                 v.Shoulder,
			NeckTypeId:               v.NeckTypeId,
			Type:                     v.Type,
			ColorFamily:              v.ColorFamily,
			PrintAndPattern:          v.PrintAndPattern,
			Length:                   v.Length,
			Pocket:                   v.Pocket,
			CommonDescriptionId:      v.CommonDescriptionId,
			BtId:                     v.BtId,
			BottomproductDescription: v.BottomproductDescription,
			BtWeight:                 v.BtWeight,
			BtPrintAndPattern:        v.BtPrintAndPattern,
			BtLength:                 v.BtLength,
			BtWaist:                  v.BtWaist,
			BtHip:                    v.BtHip,
			BtcommonDescriptionId:    v.CommonDescriptionId,
			BtType:                   v.BtType,
			BtcolorFamily:            v.BtcolorFamily,
			BtPocket:                 v.BtPocket,
			BtBeltLoop:               v.BtBeltLoop,
			TypeOfPantId:             v.TypeOfPantId,
			TypesOfPleatsId:          v.TypesOfPleatsId,
			TypesOfLengthId:          v.TypesOfLengthId,
			KurId:                    v.KurId,
			Work:                     v.Work,
			TransparencyOfTheFabric:  v.TransparencyOfTheFabric,
			SdIp:                     v.SdIp,
			Pattern:                  v.Pattern,
			FootLength:               v.FootLength,
			SoleMaterial:             v.SoleMaterial,
			UpperMaterial:            v.UpperMaterial,
			Closure:                  v.Closure,
			ToeType:                  v.ToeType,
			PackageContains:          v.PackageContains,
			WarrentyId:               v.WarrentyId,
			WarrantyYear:             v.WarrantyYear,
			ShoesWarrentyId:          v.ShoesWarrentyId,
			WdId:                     v.WdId,
			Model:                    v.Model,
			DailShape:                v.DailShape,
			DialDiameter:             v.DialDiameter,
			DialColor:                v.DialColor,
			StrapColor:               v.StrapColor,
			WatchesWarrentyId:        v.WatchesWarrentyId,
			PerfumeId:                v.PerfumeId,
			MaterialDescription:      v.MaterialDescription,
			InId:                     v.InId,
			PdIp:                     v.PdIp,
			PdTyId:                   v.PdTyId,
			PdCtId:                   v.PdCtId,
			TopbtId:                  v.TopbtId,
			KurtaBt:                  v.KurtaBt,
			LookAndFeel:              v.LookAndFeel,
			MultiColors:              v.MultiColors,
			SleeveName:               v.SleeveName,
			NeckName:                 v.NeckName,
			KurtasTypeName:           v.KurtasTypeName,
			BtTypeName:               v.BtTypeName,
			BtmPleatName:             v.BtmPleatName,
			BtmLengthName:            v.BtmLengthName,
			ProductCategoryName:      v.ProductCategoryName,
		})
		mapData[key] = productData

		for _, clr := range dataOfClrs {
			for i := 0; i < len(instedClrIds); i++ {
				if instedClrIds[i] == clr.Id {
					clrCheck = false
				}
			}
			if clrCheck && v.UId == clr.UId {
				instedClrIds = append(instedClrIds, clr.Id)

				clrData := mapData[key]
				clrData.ProductColor = append(mapData[key].ProductColor, productColorDetails{
					Id:     clr.Id,
					PdId:   clr.PdId,
					UId:    clr.UId,
					Colors: clr.Colors,
				})
				mapData[key] = clrData
			}
			clrCheck = true

		}

		for _, img := range clrImage {

			for i := 0; i < len(insertedImgIds); i++ {
				if insertedImgIds[i] == int(img.Id) {
					imgCheck = false
				}
			}

			if imgCheck && v.UId == img.UId {
				insertedImgIds = append(insertedImgIds, int(img.Id))
				imgData := mapData[key]
				imgData.ProductClrImg = append(imgData.ProductClrImg, imgDetails{
					Id:         img.Id,
					PdId:       img.PdId,
					UId:        img.UId,
					CId:        img.CId,
					ImageUrl:   img.ImageUrl,
					ImageBytes: img.ImageBytes,
				})
				mapData[key] = imgData
			}
			imgCheck = true
		}

		for _, size := range sizeData {

			for i := 0; i < len(insertedSizesIds); i++ {
				if insertedSizesIds[i] == int(size.Id) {
					sizeCheck = false
				}
			}
			if sizeCheck && v.UId == size.UId {
				insertedSizesIds = append(insertedSizesIds, int(size.Id))
				szData := mapData[key]
				szData.ProductClrSize = append(szData.ProductClrSize, SizeDetails{
					Id:       size.Id,
					PdId:     size.PdId,
					UId:      size.UId,
					ClrId:    size.ClrId,
					Quantity: size.Quantity,
					Size:     size.Size,
				})
				mapData[key] = szData
			}
			sizeCheck = true
		}
		for _, season := range seasonalData {

			for i := 0; i < len(insertedSeasonalIds); i++ {
				if insertedSeasonalIds[i] == int(season.Id) {
					seasonalCheck = false
				}
			}

			if seasonalCheck && v.UId == season.UId {
				insertedSeasonalIds = append(insertedSeasonalIds, int(season.Id))
				seasonData := mapData[key]
				seasonData.ProductSeasonalId = append(seasonData.ProductSeasonalId, seasonalDetails{
					Id:           season.Id,
					PdId:         season.PdId,
					UId:          season.UId,
					SeasonalName: season.SeasonalName,
				})
				mapData[key] = seasonData
			}
			seasonalCheck = true
		}
	}
	var inWg sync.WaitGroup
	var syncedProductIds []int16
	for _, insert := range mapData {
		inWg.Add(1)
		go func(insert commonStruct, inwg *sync.WaitGroup) {
			defer inWg.Done()
			db.Transaction(func(tx *gorm.DB) error {

				var (
					storeName         = "Rvk"
					address           = "5/372 anna nagar"
					city              = "coimbatore"
					pincode           = "641020"
					state             = "Tamil nadu"
					customerCareEmail = "rvk@gmail.com"
				)

				for _, pro := range insert.Product {

					type Owner struct {
						Id    int16 `json:"id" gorm:"column:id"`
						Count int   `json:"count" gorm:"column:count"`
					}
					var ownerDetails Owner
					var checkISstoreOwnerExist strings.Builder

					checkISstoreOwnerExist.WriteString(fmt.Sprintf(`SELECT "id" AS "id",COUNT(1) OVER( PARTITION BY 1) AS "count" FROM "ProductOwner" WHERE "ownerEmail" = '%s'`, pro.Email))

					err := tx.Raw(checkISstoreOwnerExist.String()).Scan(&ownerDetails).Error

					if err != nil {
						return err
					}

					var storeOwnerId int

					if ownerDetails.Count == 0 {

						var insertStoreOwner strings.Builder

						insertStoreOwner.WriteString(fmt.Sprintf(`INSERT INTO "ProductOwner" ("ownerName","ownerEmail","ownerMobile","storeName","storeAddress","storeCity","storePincode","storeState","customerCareEmail","syncUserId", "createdAt", "updatedAt")
		                VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%s','%d',NOW(),NOW()) RETURNING "id"`, pro.Name, pro.Email, pro.Mobile, storeName, address, city, pincode, state, customerCareEmail, pro.UId))
						err = tx.Raw(insertStoreOwner.String()).Scan(&storeOwnerId).Error

						if err != nil {
							return err
						}

					} else {
						storeOwnerId = int(ownerDetails.Id)
					}

					var insterProduct strings.Builder

					price, err := strconv.Atoi(pro.Price)

					if err != nil {
						return err
					}

					type CategoryId struct {
						ProductCategoryId int `json:"productCategoryId" gorm:"column:productCategoryId"`
						ProductTypeId     int `json:"productTypeId" gorm:"column:productTypeId"`
					}
					var categoryIds CategoryId

					var findProductTypeQuery strings.Builder

					findProductTypeQuery.WriteString(fmt.Sprintf(`SELECT "ProductType"."id" AS "productTypeId","ProductCategory"."id" AS "ProductCategoryId" FROM "ProductType"
					INNER JOIN "ProductCategory" ON "ProductCategory"."id"="ProductType"."productCategoryId"
					WHERE "ProductType"."itemsName"='%s' AND "ProductCategory"."name" = '%s'`, pro.ProductTypeName, pro.ProductCategoryName))

					err = db.Raw(findProductTypeQuery.String()).Scan(&categoryIds).Error

					if err != nil {
						return err
					}

					var productId int
					insterProduct.WriteString(fmt.Sprintf(`INSERT INTO "Product" ("syncId","title","price","productTypeId","productOwnerId","createdAt", "updatedAt")
				    VALUES('%d','%s','%d','%d','%d',NOW(),NOW()) RETURNING "id"`, pro.PID, pro.Title, price, categoryIds.ProductTypeId, storeOwnerId))

					err = tx.Raw(insterProduct.String()).Scan(&productId).Error

					if err != nil {
						return err
					}

					for _, season := range insert.ProductSeasonalId {
						if pro.PID == season.PdId {
							var insertSeason strings.Builder

							insertSeason.WriteString(fmt.Sprintf(`INSERT INTO "SeasonalDresses" ("seasonal","productId","createdAt", "updatedAt")
					VALUES('%s','%d',NOW(),NOW())`, season.SeasonalName, productId))

							err = tx.Exec(insertSeason.String()).Error

							if err != nil {
								return err
							}
						}
					}

					for _, clr := range insert.ProductColor {

						if clr.PdId == pro.PID {
							var insertClr strings.Builder

							var clrId int16
							insertClr.WriteString(fmt.Sprintf(`INSERT INTO "ProductColor" ("productId","colors","createdAt", "updatedAt")
					        VALUES('%d','%s',NOW(),NOW()) RETURNING "id"`, productId, clr.Colors))
							err = tx.Raw(insertClr.String()).Scan(&clrId).Error

							if err != nil {
								return err
							}

							for _, img := range insert.ProductClrImg {
								if img.PdId == pro.PID {
									cwd, err := os.Getwd()

									if err != nil {
										log.Fatalf("Error getting current directory: %s", err)
										return err
									}

									publicDir := filepath.Join(cwd, "public")
									userDir := filepath.Join(publicDir, pro.Email)

									if _, err := os.Stat(publicDir); os.IsNotExist(err) {
										err = os.Mkdir(publicDir, 0755)
										if err != nil {
											log.Fatalf("Failed to create public directory: %s", err)
											return err
										}
									}

									if _, err := os.Stat(userDir); os.IsNotExist(err) {
										err = os.Mkdir(userDir, 0755)
										if err != nil {
											log.Fatalf("Failed to create user directory: %s", err)
											return err
										}
									}

									filePath := filepath.Join(userDir, img.ImageUrl)

									file, err := os.Create(filePath)
									if err != nil {
										log.Fatalf("Failed to create file: %s", err)
										return err
									}

									_, err = file.Write([]byte(img.ImageBytes))
									if err != nil {
										log.Fatalf("Failed to write data to file: %s", err)
										return err
									}

									err = file.Close()
									if err != nil {
										log.Fatalf("Failed to close file: %s", err)
										return err
									}

									log.Println("Image successfully saved to", filePath)

									var insertClrImag strings.Builder

									insertClrImag.WriteString(fmt.Sprintf(`INSERT INTO "ProductImages" ("imageUrl","productColorId","createdAt", "updatedAt")
						        VALUES('%s','%d',NOW(),NOW())`, img.ImageUrl, clrId))

									err = tx.Exec(insertClrImag.String()).Error

									if err != nil {
										return err
									}
								}

							}

							for _, size := range insert.ProductClrSize {

								if size.PdId == pro.PID {
									var clrSize strings.Builder

									var sizeId int16
									clrSize.WriteString(fmt.Sprintf(`SELECT "ProductTypeSize"."id" FROM "ProductTypeSize"
						INNER JOIN "ProductType" ON "ProductType"."id"= "ProductTypeSize"."ietmsId"
						INNER JOIN "ProductCategory" ON "ProductCategory"."id"="ProductType"."productCategoryId"
						WHERE "ProductTypeSize"."size" = '%s' AND "ProductTypeSize"."ietmsId" = '%d' AND "ProductType"."id" = '%d' AND "ProductCategory"."id" = '%d'
						`, size.Size, categoryIds.ProductTypeId, categoryIds.ProductTypeId, categoryIds.ProductCategoryId))

									err = tx.Raw(clrSize.String()).Scan(&sizeId).Error

									if err != nil {
										return err
									}

									var insertClrSize strings.Builder

									insertClrSize.WriteString(fmt.Sprintf(`INSERT INTO "ProductAviableSizes" ("quantity","productColorId","productTypeSizeId","createdAt", "updatedAt")
						VALUES('%d','%d','%d',NOW(),NOW())`, size.Quantity, clrId, sizeId))

									err = tx.Raw(insertClrSize.String()).Scan(&sizeId).Error

									if err != nil {
										return err
									}
								}

							}
						}

					}

					var commonDesId int
					var commonDescription strings.Builder

					commonDescription.WriteString(fmt.Sprintf(`INSERT INTO "CommonDescription" ("fit","materail","care","brandName","origin","productId","occasion","specialFeature","createdAt", "updatedAt")
				VALUES('%s','%s','%s','%s','%s','%d','%s','%s',NOW(),NOW()) RETURNING "id"`, pro.Fit, pro.Material, pro.Care, pro.BrandName, pro.Origin, productId, pro.Occasion, pro.SpecialFeature))

					err = tx.Raw(commonDescription.String()).Scan(&commonDesId).Error

					if err != nil {
						return err
					}
					var sleeveId int

					if pro.SleeveName != "" {
						err = db.Raw(`SELECT "id" FROM "SleeveType" WHERE "name" = ?`, pro.SleeveName).Scan(&sleeveId).Error

						if err != nil {
							return err
						}
					}
					var neckId int

					if pro.NeckName != "" {
						err = db.Raw(`SELECT "id" FROM "NeckType" WHERE "name" = ?`, pro.NeckName).Scan(&neckId).Error

						if err != nil {
							return err
						}
					}

					var typeOfBt int

					if pro.BtTypeName != "" {
						err = db.Raw(`SELECT "id" FROM "TypesOfBottom" WHERE "name" = ?`, pro.BtTypeName).Scan(&typeOfBt).Error

						if err != nil {
							return err
						}
					}

					var typeOfBtPleats int

					if pro.BtmPleatName != "" {
						err = db.Raw(`SELECT "id" FROM "TypesOfPleats" WHERE "name" = ?`, pro.BtmPleatName).Scan(&typeOfBtPleats).Error

						if err != nil {
							return err
						}
					}
					var typeOfBtlength int

					if pro.BtmLengthName != "" {
						err = db.Raw(`SELECT "id" FROM "TypesOfLengthBottom" WHERE "name" = ?`, pro.BtmLengthName).Scan(&typeOfBtlength).Error

						if err != nil {
							return err
						}
					}

					var kurtasLen int

					if pro.KurtasTypeName != "" {
						err = db.Raw(`SELECT "id" FROM "KurtasLengthType" WHERE "name" = ?`, pro.KurtasTypeName).Scan(&kurtasLen).Error

						if err != nil {
							return err
						}
					}

					productname := pro.ProductCategoryName
					switch productname {
					case "Top":
						_, err = CreateTopDescription(tx, neckId, sleeveId, pro, commonDesId)

						if err != nil {
							return err
						}
					case "Bottom":
						err = CreateBottomDescription(tx, typeOfBt, typeOfBtPleats, typeOfBtlength, pro, commonDesId, 0, 0)

						if err != nil {
							return err
						}
					case "Ethnic":

						if pro.ProductTypeName == "Ethnic Wear Sets" {
							id, err := KurtasDescriptiondb(tx, neckId, sleeveId, kurtasLen, pro, commonDesId)

							if err != nil {
								return err
							}
							CreateBottomDescription(tx, typeOfBt, typeOfBtPleats, typeOfBtlength, pro, commonDesId, 0, int(id))

						} else if pro.ProductTypeName == "Ethnic Bottom Wear" {
							CreateBottomDescription(tx, typeOfBt, typeOfBtPleats, typeOfBtlength, pro, commonDesId, 0, 0)
						} else {
							_, err = KurtasDescriptiondb(tx, neckId, sleeveId, kurtasLen, pro, commonDesId)
							if err != nil {
								return err
							}
						}

					case "Sports":
						if pro.ProductTypeName == "T-Shirts" {

							_, err = CreateTopDescription(tx, neckId, sleeveId, pro, commonDesId)

							if err != nil {
								return err
							}
						} else if pro.ProductTypeName == "Track Suits" {

							topId, err := CreateTopDescription(tx, neckId, sleeveId, pro, commonDesId)
							if err != nil {
								return err
							}
							err = CreateBottomDescription(tx, typeOfBt, typeOfBtPleats, typeOfBtlength, pro, commonDesId, int(topId), 0)

							if err != nil {
								return err
							}

						} else {
							err = CreateBottomDescription(tx, typeOfBt, typeOfBtPleats, typeOfBtlength, pro, commonDesId, 0, 0)
							if err != nil {
								return err
							}

						}
					case "Footwear":
						err = CreateFootWear(tx, pro, commonDesId)
						if err != nil {
							return err
						}

					case "Accessories":
						break
					case "Inner":
						err = CreateInner(tx, neckId, sleeveId, pro, commonDesId)
						if err != nil {
							return err
						}

					case "Fragrances":
						err = createFragence(tx, pro, commonDesId)
						if err != nil {
							return err
						}

					case "Watches":
						err = CreateWatches(tx, pro, commonDesId)
						if err != nil {
							return err
						}

					default:
						fmt.Printf("There is no ProductCategory like this: %s\n  and product id : %d\n", pro.ProductCategoryName, pro.PID)
					}
					syncedProductIds = append(syncedProductIds, int16(pro.PID))

				}

				return nil
			})
			if err != nil {
				log.Printf("Transaction failed for insert: %v, error: %v\n", insert, err)
			}
		}(insert, &inWg)

	}
	inWg.Wait()

	for _, syncedIds := range syncedProductIds {
		var updateSyncedDaata strings.Builder

		updateSyncedDaata.WriteString(fmt.Sprintf(`UPDATE "Product" SET "isSync" = '%t' WHERE "id"= '%d'`, true, syncedIds))

		err = pdb.Exec(updateSyncedDaata.String()).Error

		if err != nil {
			fmt.Errorf(`Error occurs while updating status of synced data : %v\n`, err)
		}
	}

	fmt.Println("Data sync completed successfully")
}
func CreateTopDescription(db *gorm.DB, neck int, sleeve int, top ProductDetails, cdId int) (int16, error) {

	var insertTopDes strings.Builder

	var topId int16

	insertTopDes.WriteString(fmt.Sprintf(`INSERT INTO "TopDescription" ("productDescription","sleeveTypeId","weight","chest","shoulder","neckTypeId","type","colorFamily","printAndPattern","length","pocket","commonDescriptionId","createdAt", "updatedAt")
	VALUES('%s','%d','%d','%d','%d','%d','%s','%s','%s','%d','%s','%d',NOW(),NOW()) RETURNING "id"`, top.ProductDescription, sleeve, top.Weight, top.Chest, top.Shoulder, neck, top.Type, top.ColorFamily, top.PrintAndPattern, top.Length, top.Pocket, cdId))

	err := db.Raw(insertTopDes.String()).Scan(&topId).Error

	if err != nil {
		return 0, err
	}

	return topId, nil
}

func CreateBottomDescription(db *gorm.DB, btType int, btPleats int, btLength int, bt ProductDetails, cdId int, topId int, kurtasId int) error {

	var insertBtDescription strings.Builder

	length, err := strconv.Atoi(bt.BtLength)

	if err != nil {
		return err
	}

	if topId > 0 {
		bt.TopbtId = topId
	}
	if kurtasId > 0 {
		bt.KurtaBt = kurtasId
	}
	insertBtDescription.WriteString(fmt.Sprintf(`INSERT INTO "BottomDescription" ("productDescription","weight","printAndPattern","length","waist","hip","commonDescriptionId","type","colorFamily","pocket","kurtasDescriptionId","topDescriptionId","beltLoop","typeOfPantId","typesOfPleatsId","typesOfLengthId","createdAt", "updatedAt")
	VALUES('%s','%d','%s','%d','%d','%d','%d','%s','%s','%s',NULLIF('%d',0),NULLIF('%d',0),'%t','%d','%d','%d',NOW(),NOW())`, bt.BottomproductDescription, bt.BtWeight, bt.BtPrintAndPattern, length, bt.BtWaist, bt.BtHip, cdId, bt.BtType, bt.BtcolorFamily, bt.BtPocket, bt.KurtaBt, bt.TopbtId, bt.BtBeltLoop, btType, btPleats, btLength))

	err = db.Exec(insertBtDescription.String()).Error

	if err != nil {
		return err
	}
	return nil
}

func KurtasDescriptiondb(db *gorm.DB, neck int, sleeve int, kurtasLen int, kur ProductDetails, cdId int) (int16, error) {

	var insertKurtas strings.Builder

	var kurId int16
	insertKurtas.WriteString(fmt.Sprintf(`INSERT INTO "KurtasDescription" ("work","productDescription","chest","shoulder","transparencyOfTheFabric","kurtasLengthTypeId","weight","colorFamily","pocket","type","printAndpattern","kurtasNeckTypeId","kurtasSleeveTypeId","commonDescriptionId","createdAt", "updatedAt")
	VALUES('%s','%s','%d','%d','%t','%d','%d','%s','%s','%s','%s','%d','%d','%d',NOW(),NOW()) RETURNING "id"`, kur.Work, kur.ProductDescription, kur.Chest, kur.Shoulder, kur.TransparencyOfTheFabric, kurtasLen, kur.Weight, kur.ColorFamily, kur.Pocket, kur.Type, kur.PrintAndPattern, neck, sleeve, cdId))

	err := db.Raw(insertKurtas.String()).Scan(&kurId).Error

	if err != nil {
		return 0, err
	}

	return kurId, nil
}
func CreateFootWear(db *gorm.DB, fw ProductDetails, cdId int) error {
	var insertFootwear strings.Builder
	var id int16

	insertFootwear.WriteString(fmt.Sprintf(`INSERT INTO "ShoesDescription" ("pattern","footLength","type","soleMaterial","printAndPattern","upperMaterial","closure","toeType","weight","colorFamily","productDescription","packageContains","commonDescriptionId","createdAt", "updatedAt")
	VALUES('%s','%s','%s','%s','%s','%s','%s','%s','%d','%s','%s','%d','%d',NOW(),NOW()) RETURNING "id"`, fw.Pattern, fw.FootLength, fw.Type, fw.SoleMaterial, fw.PrintAndPattern, fw.UpperMaterial, fw.Closure, fw.ToeType, fw.Weight, fw.ColorFamily, fw.ProductDescription, fw.PackageContains, cdId))

	err := db.Raw(insertFootwear.String()).Scan(&id).Error

	if err != nil {
		return err
	}

	var warrentyInsert strings.Builder

	warrentyInsert.WriteString(fmt.Sprintf(`INSERT INTO "Warranty" ("shoesDescriptionId","watchsId","warrantyPeriod","createdAt", "updatedAt")
	VALUES ('%d',NULL,'%d',NOW(),NOW())`, id, fw.WarrantyYear))

	err = db.Raw(warrentyInsert.String()).Error

	if err != nil {
		return err
	}
	return nil
}
func CreateWatches(db *gorm.DB, watch ProductDetails, cdId int) error {
	var insertWatches strings.Builder

	insertWatches.WriteString(fmt.Sprintf(`INSERT INTO "WatchesDescription" ("type","weight","model","dialShape","printAndPattern","dialDiameter","dialColor","strapColor","colorFamily","productDescription","commonDescriptionId","createdAt","updatedAt")
	VALUES('%s','%d','%s','%s','%s','%s','%s','%s','%s','%s','%d',NOW(),NOW())`, watch.Type, watch.Weight, watch.Model, watch.DailShape, watch.PrintAndPattern, watch.DialDiameter, watch.DialColor, watch.StrapColor, watch.ColorFamily, watch.ProductDescription, cdId))

	err := db.Raw(insertWatches.String()).Error

	if err != nil {
		return err
	}
	return nil
}
func createFragence(db *gorm.DB, fg ProductDetails, cdId int) error {
	var insertFragence strings.Builder

	insertFragence.WriteString(fmt.Sprintf(`INSERT INTO "PerfumesDescription" ("productDescription","type","materialDescription","weight","commonDescriptionId","createdAt","updatedAt")
	VALUES('%s','%s','%s','%d','%d',NOW(),NOW())`, fg.ProductDescription, fg.Type, fg.MaterialDescription, fg.Weight, cdId))

	err := db.Raw(insertFragence.String()).Error

	if err != nil {
		return err
	}
	return nil
}
func CreateInner(db *gorm.DB, neck int, sleeve int, in ProductDetails, cdId int) error {
	var insertInner strings.Builder

	insertInner.WriteString(fmt.Sprintf(`INSERT INTO "InnersDescription" ("type","productDescription","weight","length","waistRise","printAndPattern","packageContains","lookAndFeel","colorFamily","vestsSleeveTypeId","vestsNeckTypeId","commonDescriptionId","multiColors","createdAt","updatedAt")
	VALUES('%s','%s','%d','%d','%d','%s','%d','%s','%s',NULLIF('%d',0),NULLIF('%d',0),'%d','%t',NOW(),NOW())`, in.Type, in.ProductDescription, in.Weight, in.Length, in.BtWaist, in.PrintAndPattern, in.PackageContains, in.LookAndFeel, in.ColorFamily, sleeve, neck, cdId, in.MultiColors))

	err := db.Raw(insertInner.String()).Error

	if err != nil {
		return err
	}
	return nil
}
