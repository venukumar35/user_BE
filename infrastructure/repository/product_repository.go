package repository

import (
	"TheBoys/app/model/request"
	"TheBoys/app/model/response"
	"TheBoys/domain"
	"TheBoys/utills"
	"fmt"
	"strings"
	"sync"

	"gorm.io/gorm"
)

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return &ProductRepository{db}
}

type ProductRepository struct {
	db *gorm.DB
}

func (r *ProductRepository) GetProductCategory() (interface{}, error) {
	var data []struct {
		Id   int    `json:"id" gorm:"column:id"`
		Name string `json:"name" gorm:"column:name"`
	}

	err := r.db.Raw(`SELECT "id","name" FROM"ProductCategory"`).Scan(&data).Error

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *ProductRepository) GetProducts() (*utills.PaginationResponse, error) {

	var data []response.ProductDetailsResponse

	var wg sync.WaitGroup

	wg.Add(4)

	productDetails := make(chan int16, 100)
	colorChannels := make(chan int16, 100)
	colorIdForSizes := make(chan int16, 100)
	colorIdForImages := make(chan int16, 100)

	catchErr := make(chan error, len(productDetails))
	go func() {
		defer close(colorIdForImages)
		defer close(colorIdForSizes)

		for id := range colorChannels {
			colorIdForImages <- id
			colorIdForSizes <- id
		}
	}()
	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		defer close(productDetails)
		var baseQuery strings.Builder
		baseQuery.WriteString(`SELECT 
            "Product"."id" AS "id",
            "Product"."title" AS "title",
			"Product"."price" AS "price",
			"Offer"."offerPercntage" AS "offerPercntage",
		    "Offer"."offerPrice" AS "offerPrice",
		    "OfferValidity"."fromDate" AS "OfferValidityFromdate",
		    "OfferValidity"."toDate" AS "OfferValidityTodate",
		    "OfferValidity"."fromTime" AS "OfferValidityFromTime",
		    "OfferValidity"."toTime" AS "OfferValidityToTime",
			"ProductOwner"."storeName" AS "storeName",
			"ProductOwner"."storeAddress" AS "storeAddress",
			"ProductOwner"."storeCity" AS "storeCity",
			"ProductOwner"."storePincode" AS "storePincode",
			"ProductOwner"."storeState" AS "storeState",
			"ProductOwner"."customerCareEmail" AS "customerCareEmail",
			 "ProductOwner"."ownerEmail" AS "ownwrEmail",
			"commonDes"."fit" AS "fit","commonDes"."materail" AS "materail","commonDes"."care" AS "care","commonDes"."brandName" AS "brandName",
			"commonDes"."origin" AS "origin","commonDes"."occasion" AS "occasion","commonDes"."specialFeature" AS "specialFeature",
            "season"."seasonName" AS "seasonalName","commonDes"."fit" AS "fit","commonDes"."commonId" AS "CommonDesId","commonDes"."topId" AS "topId",
            "commonDes"."topProDes" AS "productDescription","commonDes"."topWeight" AS "weight","commonDes"."topPrintPattern" AS "printAndPattern",
            "commonDes"."topLength" AS "length","commonDes"."topShoulder" AS "shoulder","commonDes"."topChest" AS "chest",
            "commonDes"."topCommonDesId" AS "CommonDescriptionId","commonDes"."topType" AS "type","commonDes"."topSleeveId" AS "sleeveTypeId",
            "commonDes"."topNeckId" AS "neckTypeId","commonDes"."topColorFamily" AS "colorFamily","commonDes"."topPocket" AS "pocket","commonDes"."btProDes" AS "bottomproductDescription",
            "commonDes"."btWeight" AS "BtWeight","commonDes"."btPrintPattern" AS "BtPrintAndPattern","commonDes"."btLength" AS "btLength",
            "commonDes"."btWaist" AS "btWaist","commonDes"."btHip" AS "btHip","commonDes"."btCommnDesId" AS "btcommonDescriptionId","commonDes"."btType" AS "btType",
            "commonDes"."btColorFamily" AS "btcolorFamily","commonDes"."btPocket" AS "btPocket","commonDes"."btBeltLoop" AS "btBeltLoop",
            "commonDes"."btPantType" AS "typeOfPantId","commonDes"."btPleatsId" AS "typesOfPleatsId","commonDes"."btLengthId" AS "typesOfLengthId",
            "commonDes"."btKurtasDesId" AS "kurtaBt","commonDes"."btTopDescId" AS "topbtId","commonDes"."kurId" AS "kurtasId","commonDes"."work" AS "work",
            "commonDes"."kurtransparencyOfTheFabric" AS "transparencyOfTheFabric","commonDes"."kurtasLengthTypeId" AS "length","commonDes"."kurChest" AS "chest",
            "commonDes"."kurProductDescription" AS "productDescription","commonDes"."kurShoulder" AS "shoulder","commonDes"."kurWeight" AS "weight",
            "commonDes"."kurColorFamily" AS "colorFamily","commonDes"."kurPocket" AS "pocket","commonDes"."kurType" AS "type","commonDes"."kurPrintAndpattern" AS "printAndpattern",
            "commonDes"."kurNeck" AS "kurtasNeckTypeId","commonDes"."kurSleeve" AS "kurtasSleeveTypeId","commonDes"."kurCommonDesId" AS "commonDescriptionId",
            "commonDes"."shId" AS "shIp","commonDes"."shPrintAndPattern" AS "printAndpattern","commonDes"."shType" AS "type",
            "commonDes"."shProductDes" AS "productDescription","commonDes"."shColorFamily" AS "colorFamily","commonDes"."shPattern" AS "pattern",
            "commonDes"."shLength" AS "footLength","commonDes"."shSolematerial" AS "soleMaterial","commonDes"."shUpperMaterial" AS "upperMaterial",
            "commonDes"."shClosure" AS "closure","commonDes"."shToeType" AS "toeType","commonDes"."shWeight" AS "weight",
            "commonDes"."shPackageContains" AS "packageContains","commonDes"."shWarrantyId" AS "warrentyId","commonDes"."shWarrantyPeriod" AS "WarrantyYear",
            "commonDes"."watchId" AS "WdId","commonDes"."watchType" AS "type","commonDes"."watchWeight" AS "weight","commonDes"."watchModel" AS "model",
	        "commonDes"."watchDialShape" AS "dailShape","commonDes"."watchPrintAndPattern" AS "printAndPattern","commonDes"."watchDialDiameter" AS "dialDiameter",
	        "commonDes"."watchDialColor" AS "dialColor","commonDes"."watchStrapColor" AS "watchStrapColor","commonDes"."watchColorFamily" AS "strapColor",
	        "commonDes"."watchProductDescription" AS "productDescription","commonDes"."watchWarrantyPeriod" AS "WarrantyYear","commonDes"."watchWarrantyId" AS "warrentyId",
		    "commonDes"."perfumeId" AS "PerfumeId","commonDes"."perfumeType" AS "type","commonDes"."perfumeMaterialDescription" AS "MaterialDescription",
	        "commonDes"."perfumeWeight" AS "weight","commonDes"."perfumeDes" AS "productDescription","commonDes"."innerId" AS "InId",
	        "commonDes"."innerType" AS "type","commonDes"."innerProductDescription" AS "productDescription","commonDes"."innerWeight" AS "weight",
	        "commonDes"."innerLength" AS "length","commonDes"."innerWaistRise" AS "btWaist","commonDes"."innerPrintAndPattern" AS "printAndPattern",
	        "commonDes"."innerPackageContains" AS "packageContains","commonDes"."innerLookAndFeel" AS "lookAndFeel","commonDes"."innerColorFamily" AS "colorFamily",
	        "commonDes"."innerVestsSleeveTypeId" AS "sleeveTypeId","commonDes"."innerVestsNeckTypeId" AS "neckTypeId","commonDes"."innerMultiColors" AS "multiColors"
            FROM "Product" INNER JOIN "ProductType" ON "ProductType"."id" = "Product"."productTypeId"
	        INNER JOIN "ProductCategory" ON "ProductCategory"."id" = "ProductType"."productCategoryId"
            INNER JOIN (SELECT "productId", STRING_AGG("seasonal", ', ') AS "seasonName"  FROM "SeasonalDresses"  GROUP BY "productId") "season" ON "season"."productId" = "Product"."id"
            INNER JOIN (SELECT 
            "CommonDescription"."id" AS "commonId", 
            "CommonDescription"."productId",
			"CommonDescription"."fit" AS "fit","CommonDescription"."materail" AS "materail",
			"CommonDescription"."care" AS "care","CommonDescription"."brandName" AS "brandName",
            "CommonDescription"."origin" AS "origin","CommonDescription"."occasion" AS "occasion",
            "CommonDescription"."specialFeature",STRING_AGG(CAST("TopDescription"."id" AS CHAR), ', ') AS "topId",
            STRING_AGG("TopDescription"."productDescription", ', ') AS "topProDes",STRING_AGG(CAST("TopDescription"."weight" AS CHAR), ', ') AS "topWeight",
            STRING_AGG("TopDescription"."printAndPattern", ', ') AS "topPrintPattern",STRING_AGG(CAST("TopDescription"."length" AS CHAR), ', ') AS "topLength",
            STRING_AGG(CAST("TopDescription"."shoulder" AS CHAR), ', ') AS "topShoulder",STRING_AGG(CAST("TopDescription"."chest" AS CHAR), ', ') AS "topChest",
            STRING_AGG(CAST("TopDescription"."commonDescriptionId" AS CHAR), ', ') AS "topCommonDesId",STRING_AGG("TopDescription"."type", ', ') AS "topType",
            STRING_AGG(CAST("TopDescription"."sleeveTypeId" AS CHAR), ', ') AS "topSleeveId",STRING_AGG(CAST("TopDescription"."neckTypeId" AS CHAR), ', ') AS "topNeckId",
            STRING_AGG("TopDescription"."colorFamily", ', ') AS "topColorFamily",STRING_AGG(CAST("TopDescription"."pocket" AS CHAR), ', ') AS "topPocket",
            STRING_AGG("BottomDescription"."productDescription", ', ') AS "btProDes",STRING_AGG(CAST("BottomDescription"."weight" AS CHAR), ', ') AS "btWeight",
            STRING_AGG("BottomDescription"."printAndPattern", ', ') AS "btPrintPattern",STRING_AGG(CAST("BottomDescription"."length" AS CHAR), ', ') AS "btLength",
            STRING_AGG(CAST("BottomDescription"."waist" AS CHAR), ', ') AS "btWaist",STRING_AGG(CAST("BottomDescription"."hip" AS CHAR), ', ') AS "btHip",
            STRING_AGG(CAST("BottomDescription"."commonDescriptionId" AS CHAR), ', ') AS "btCommnDesId",STRING_AGG("BottomDescription"."type", ', ') AS "btType",
            STRING_AGG("BottomDescription"."colorFamily", ', ') AS "btColorFamily",STRING_AGG("BottomDescription"."pocket", ', ') AS "btPocket",
            STRING_AGG(CAST("BottomDescription"."kurtasDescriptionId" AS CHAR), ', ') AS "btKurtasDesId",STRING_AGG(CAST("BottomDescription"."topDescriptionId" AS CHAR), ', ') AS "btTopDescId",
            STRING_AGG(CAST("BottomDescription"."beltLoop" AS CHAR), ', ') AS "btBeltLoop",STRING_AGG(CAST("BottomDescription"."typeOfPantId" AS CHAR), ', ') AS "btPantType",
            STRING_AGG(CAST("BottomDescription"."typesOfPleatsId" AS CHAR), ', ') AS "btPleatsId",STRING_AGG(CAST("BottomDescription"."typesOfLengthId" AS CHAR), ', ') AS "btLengthId",
            STRING_AGG(CAST("KurtasDescription"."id" AS CHAR), ', ') AS "kurId",STRING_AGG("KurtasDescription"."work", ', ') AS "work",STRING_AGG("KurtasDescription"."productDescription", ', ') AS "kurProductDescription",
            STRING_AGG(CAST("KurtasDescription"."chest" AS CHAR), ', ') AS "kurChest",STRING_AGG(CAST("KurtasDescription"."shoulder" AS CHAR), ', ') AS "kurShoulder",
            STRING_AGG(CAST("KurtasDescription"."transparencyOfTheFabric" AS CHAR), ', ') AS "kurtransparencyOfTheFabric",STRING_AGG(CAST("KurtasDescription"."kurtasLengthTypeId" AS CHAR), ', ') AS "kurtasLengthTypeId",
            STRING_AGG(CAST("KurtasDescription"."weight" AS CHAR), ', ') AS "kurWeight",STRING_AGG("KurtasDescription"."colorFamily", ', ') AS "kurColorFamily",
            STRING_AGG(CAST("KurtasDescription"."pocket" AS CHAR), ', ') AS "kurPocket",STRING_AGG("KurtasDescription"."type", ', ') AS "kurType",
            STRING_AGG("KurtasDescription"."printAndpattern", ', ') AS "kurPrintAndpattern",STRING_AGG(CAST("KurtasDescription"."kurtasNeckTypeId" AS CHAR), ', ') AS "kurNeck",
            STRING_AGG(CAST("KurtasDescription"."kurtasSleeveTypeId" AS CHAR), ', ') AS "kurSleeve",STRING_AGG(CAST("KurtasDescription"."commonDescriptionId" AS CHAR), ', ') AS "kurCommonDesId",
            STRING_AGG(CAST("ShoesDescription"."id" AS CHAR), ', ') AS "shId",STRING_AGG("ShoesDescription"."pattern", ', ') AS "shPattern",
            STRING_AGG(CAST("ShoesDescription"."footLength" AS CHAR), ', ') AS "shLength",STRING_AGG("ShoesDescription"."type", ', ') AS "shType",
            STRING_AGG("ShoesDescription"."soleMaterial", ', ') AS "shSolematerial",STRING_AGG("ShoesDescription"."printAndPattern", ', ') AS "shPrintAndPattern",
            STRING_AGG("ShoesDescription"."upperMaterial", ', ') AS "shUpperMaterial",STRING_AGG("ShoesDescription"."closure", ', ') AS "shClosure",
            STRING_AGG("ShoesDescription"."toeType", ', ') AS "shToeType",STRING_AGG(CAST("ShoesDescription"."weight" AS CHAR), ', ') AS "shWeight",
            STRING_AGG("ShoesDescription"."colorFamily", ', ') AS "shColorFamily",STRING_AGG("ShoesDescription"."productDescription", ', ') AS "shProductDes",
            STRING_AGG(CAST("ShoesDescription"."packageContains" AS CHAR), ', ') AS "shPackageContains",STRING_AGG(CAST("ShoesDescription"."commonDescriptionId" AS CHAR), ', ') AS "shCommonId",
            STRING_AGG(CAST("shWarranty"."warrantyPeriod" AS CHAR), ', ') AS "shWarrantyPeriod",STRING_AGG(CAST("shWarranty"."id" AS CHAR), ', ') AS "shWarrantyId",
		    STRING_AGG(CAST("WatchesDescription"."id" AS CHAR), ', ') AS "watchId",
            STRING_AGG("WatchesDescription"."type", ', ') AS "watchType",STRING_AGG(CAST("WatchesDescription"."weight" AS CHAR), ', ') AS "watchWeight",
            STRING_AGG("WatchesDescription"."model", ', ') AS "watchModel",STRING_AGG("WatchesDescription"."dialShape", ', ') AS "watchDialShape",
            STRING_AGG("WatchesDescription"."printAndPattern", ', ') AS "watchPrintAndPattern",STRING_AGG(CAST("WatchesDescription"."dialDiameter" AS CHAR), ', ') AS "watchDialDiameter",
            STRING_AGG("WatchesDescription"."dialColor", ', ') AS "watchDialColor",STRING_AGG("WatchesDescription"."strapColor", ', ') AS "watchStrapColor",
            STRING_AGG("WatchesDescription"."colorFamily", ', ') AS "watchColorFamily",STRING_AGG("WatchesDescription"."productDescription", ', ') AS "watchProductDescription",
            STRING_AGG(CAST("watchWarranty"."warrantyPeriod" AS CHAR), ', ') AS "watchWarrantyPeriod",STRING_AGG(CAST("watchWarranty"."id" AS CHAR), ', ') AS "watchWarrantyId",
			STRING_AGG(CAST("PerfumesDescription"."id" AS CHAR), ', ') AS "perfumeId",
            STRING_AGG("PerfumesDescription"."productDescription", ', ') AS "perfumeDes",STRING_AGG("PerfumesDescription"."type", ', ') AS "perfumeType",
            STRING_AGG("PerfumesDescription"."materialDescription", ', ') AS "perfumeMaterialDescription",STRING_AGG(CAST("PerfumesDescription"."weight" AS CHAR), ', ') AS "perfumeWeight",
		    STRING_AGG(CAST("InnersDescription"."id" AS CHAR), ', ') AS "innerId",STRING_AGG("InnersDescription"."type", ', ') AS "innerType",
	        STRING_AGG("InnersDescription"."productDescription", ', ') AS "innerProductDescription",STRING_AGG(CAST("InnersDescription"."weight" AS CHAR), ', ') AS "innerWeight",
            STRING_AGG(CAST("InnersDescription"."length" AS CHAR), ', ') AS "innerLength",STRING_AGG(CAST("InnersDescription"."waistRise" AS CHAR), ', ') AS "innerWaistRise",
	        STRING_AGG("InnersDescription"."printAndPattern", ', ') AS "innerPrintAndPattern",STRING_AGG(CAST("InnersDescription"."packageContains" AS CHAR), ', ') AS "innerPackageContains",STRING_AGG("InnersDescription"."lookAndFeel", ', ') AS "innerLookAndFeel",
	        STRING_AGG("InnersDescription"."colorFamily", ', ') AS "innerColorFamily",    STRING_AGG(CAST("InnersDescription"."vestsSleeveTypeId" AS CHAR), ', ') AS "innerVestsSleeveTypeId",
	        STRING_AGG(CAST("InnersDescription"."vestsNeckTypeId" AS CHAR), ', ') AS "innerVestsNeckTypeId",STRING_AGG(CAST("InnersDescription"."multiColors" AS CHAR), ', ') AS "innerMultiColors"
            FROM "CommonDescription" 
            LEFT JOIN "TopDescription" ON "TopDescription"."commonDescriptionId" = "CommonDescription"."id"
            LEFT JOIN "BottomDescription" ON "BottomDescription"."commonDescriptionId" = "CommonDescription"."id"
            LEFT JOIN "KurtasDescription" ON "KurtasDescription"."commonDescriptionId" = "CommonDescription"."id"
            LEFT JOIN "ShoesDescription" ON "ShoesDescription"."commonDescriptionId" = "CommonDescription"."id"
            LEFT JOIN "Warranty" AS "shWarranty" ON "shWarranty"."shoesDescriptionId" = "ShoesDescription"."id"
            LEFT JOIN "WatchesDescription" ON "WatchesDescription"."commonDescriptionId" = "CommonDescription"."id"
			LEFT JOIN "Warranty" AS "watchWarranty" ON "watchWarranty"."watchsId" = "WatchesDescription"."id"
            LEFT JOIN "PerfumesDescription" ON "PerfumesDescription"."commonDescriptionId" = "CommonDescription"."id"
            LEFT JOIN "InnersDescription" ON "InnersDescription"."commonDescriptionId" = "CommonDescription"."id"
            GROUP BY "CommonDescription"."id") "commonDes" ON "commonDes"."productId" = "Product"."id"
			INNER JOIN "ProductOwner" ON "ProductOwner"."id"="Product"."productOwnerId"
			LEFT JOIN "Offer" ON "Offer"."productId" = "Product"."id"
			LEFT JOIN "OfferValidity" ON "OfferValidity"."offerId" = "Offer"."id"
			`)

		if err := r.db.Raw(baseQuery.String()).Scan(&data).Error; err != nil {
			return err
		}
		for _, id := range data {
			productDetails <- int16(id.Id)
		}
		return nil
	}(&wg)

	type color struct {
		Id        int    `json:"id" gorm:"column:id"`
		ProductId int16  `json:"productId" gorm:"column:productId"`
		OwnerId   int16  `json:"ownerId" gorm:"column:ownerId"`
		Colors    string `json:"colors" gorm:"column:colors"`
	}

	var datas []color
	go func(wg *sync.WaitGroup) {
		wg.Done()
		defer close(colorChannels)
		for v := range productDetails {
			var colorData []color
			var colorDetailsQuery strings.Builder

			colorDetailsQuery.WriteString(fmt.Sprintf(`SELECT DISTINCT "ProductColor"."id","ProductColor"."productId",
	"Product"."productOwnerId" AS "ownerId","ProductColor"."colors" AS "colors"
	FROM "ProductColor" 
    INNER JOIN "Product" ON "Product"."id" = "ProductColor"."productId"
	INNER JOIN "ProductOwner" ON "ProductOwner"."id" = "Product"."productOwnerId"		
	INNER JOIN "ProductAviableSizes" ON "ProductAviableSizes"."productColorId" = "ProductColor"."id"
	INNER JOIN "ProductImages" ON "ProductImages"."productColorId"= "ProductColor"."id"
	WHERE "ProductColor"."productId" = '%d'`, v))

			err := r.db.Raw(colorDetailsQuery.String()).Scan(&colorData).Error

			if err != nil {
				fmt.Errorf(`Error occurs in fetching colors details`, err)
				catchErr <- err
				return
			}
			for _, ids := range colorData {
				colorChannels <- int16(ids.Id)
			}

			datas = append(datas, colorData...)

		}
	}(&wg)

	var imgDetailsData []response.Images
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for clrIds := range colorIdForImages {
			var imgData []response.Images
			var imgDetailsQuery strings.Builder

			imgDetailsQuery.WriteString(fmt.Sprintf(`SELECT DISTINCT "ProductColor"."productId" AS "productId", "ProductImages"."id" AS "imgId",
		"ProductImages"."productColorId" AS "clrId","ProductImages"."imageUrl" AS "imageUrl"
		FROM "ProductImages"
		INNER JOIN "ProductColor" ON "ProductColor"."id" = "ProductImages"."productColorId"
		WHERE "ProductImages"."productColorId" = '%d'`, clrIds))

			err := r.db.Raw(imgDetailsQuery.String()).Scan(&imgData).Error

			if err != nil {
				fmt.Errorf(`Error occurs in fetching images details`, err)
				catchErr <- err
				return
			}
			imgDetailsData = append(imgDetailsData, imgData...)
		}
	}(&wg)
	count := len(data)

	var sizeData []response.Sizes
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		for clrIds := range colorIdForSizes {

			var sizeDetails []response.Sizes

			var sizeQuery strings.Builder

			sizeQuery.WriteString(fmt.Sprintf(`SELECT "ProductAviableSizes"."id" AS "sizeId","ProductAviableSizes"."quantity" AS "quantity","ProductAviableSizes"."productColorId" AS "clrId",
			"ProductTypeSize"."size" AS "size","ProductColor"."productId" AS "productId"
		    FROM "ProductAviableSizes"
		    INNER JOIN "ProductColor" ON "ProductColor"."id" = "ProductAviableSizes"."productColorId"
		    INNER JOIN "ProductTypeSize" ON "ProductTypeSize"."id"="ProductAviableSizes"."productTypeSizeId"
		    WHERE "ProductColor"."productId" = '%d'`, clrIds))

			err := r.db.Raw(sizeQuery.String()).Scan(&sizeDetails).Error

			if err != nil {
				fmt.Errorf(`Error occurs in fetching sizes details`, err)
				catchErr <- err
				return
			}

			sizeData = append(sizeData, sizeDetails...)
		}

	}(&wg)

	wg.Wait()
	defer close(catchErr)
	// for e := range catchErr {
	// 	fmt.Println("data into error")
	// 	if e != nil {
	// 		return nil, e
	// 	}
	// }
	var results []response.FinalProductResponse
	for _, pro := range data {
		resultsItems := response.FinalProductResponse{
			Id:                       pro.Id,
			Title:                    pro.Title,
			Price:                    pro.Price,
			SeasonalName:             pro.SeasonalName,
			CommonDesId:              pro.CommonDesId,
			Fit:                      pro.Fit,
			Materail:                 pro.Materail,
			Care:                     pro.Care,
			BrandName:                pro.BrandName,
			Origin:                   pro.Origin,
			ProductId:                pro.ProductId,
			Occasion:                 pro.Occasion,
			SpecialFeature:           pro.SpecialFeature,
			TopId:                    pro.TopId,
			ProductDescription:       pro.ProductDescription,
			Weight:                   pro.Weight,
			PrintAndPattern:          pro.PrintAndPattern,
			Length:                   pro.Length,
			Shoulder:                 pro.Shoulder,
			Chest:                    pro.Chest,
			CommonDescriptionId:      pro.CommonDescriptionId,
			Type:                     pro.Type,
			SleeveTypeId:             pro.SleeveTypeId,
			NeckTypeId:               pro.NeckTypeId,
			ColorFamily:              pro.ColorFamily,
			Pocket:                   pro.Pocket,
			BtId:                     pro.BtId,
			BottomproductDescription: pro.BottomproductDescription,
			BtWeight:                 pro.BtWeight,
			BtPrintAndPattern:        pro.BtPrintAndPattern,
			BtLength:                 pro.BtLength,
			BtWaist:                  pro.BtWaist,
			BtHip:                    pro.BtHip,
			BtcommonDescriptionId:    pro.BtcommonDescriptionId,
			BtType:                   pro.BtType,
			BtcolorFamily:            pro.BtcolorFamily,
			BtPocket:                 pro.BtPocket,
			BtBeltLoop:               pro.BtBeltLoop,
			TypeOfPantId:             pro.TypeOfPantId,
			TypesOfPleatsId:          pro.TypesOfPleatsId,
			TypesOfLengthId:          pro.TypesOfLengthId,
			TopbtId:                  pro.TopbtId,
			KurtaBt:                  pro.KurtaBt,
			KurtasId:                 pro.KurtasId,
			TransparencyOfTheFabric:  pro.TransparencyOfTheFabric,
			ShIp:                     pro.ShIp,
			Pattern:                  pro.Pattern,
			FootLength:               pro.FootLength,
			SoleMaterial:             pro.SoleMaterial,
			UpperMaterial:            pro.UpperMaterial,
			Closure:                  pro.Closure,
			ToeType:                  pro.ToeType,
			PackageContains:          pro.PackageContains,
			WarrentyId:               pro.WarrentyId,
			WarrantyYear:             pro.WarrantyYear,
			ShoesWarrentyId:          pro.ShoesWarrentyId,
			WdId:                     pro.WdId,
			Model:                    pro.Model,
			DailShape:                pro.DailShape,
			DialDiameter:             pro.DialDiameter,
			DialColor:                pro.DialColor,
			StrapColor:               pro.StrapColor,
			WatchesWarrentyId:        pro.WatchesWarrentyId,
			OfferId:                  pro.OfferId,
			OfferPercntage:           pro.OfferPercntage,
			OfferPrice:               pro.OfferPrice,
			OfferValidityFromdate:    pro.OfferValidityFromdate,
			OfferValidityTodate:      pro.OfferValidityTodate,
			OfferValidityFromTime:    pro.OfferValidityFromTime,
			OfferValidityToTime:      pro.OfferValidityToTime,
			StoreName:                pro.SeasonalName,
			StoreAddress:             pro.StoreAddress,
			StoreCity:                pro.StoreCity,
			StorePincode:             pro.StorePincode,
			StoreState:               pro.StoreState,
			CustomerCareEmail:        pro.CustomerCareEmail,
			Colors:                   []response.Clr{},
			Images:                   []response.Images{},
			Sizes:                    []response.Sizes{},
		}

		for _, clr := range datas {
			if clr.ProductId == int16(pro.Id) {
				resultsItems.Colors = append(resultsItems.Colors, response.Clr{
					Colors:    clr.Colors,
					ProductId: int(clr.ProductId),
					Id:        clr.Id,
				})

				for _, img := range imgDetailsData {
					if img.ColorId == clr.Id && clr.ProductId == int16(pro.Id) {
						resultsItems.Images = append(resultsItems.Images, response.Images{
							ImgId:     img.ImgId,
							ProductId: img.ProductId,
							ImageUrl:  "/public" + "/" + pro.OwnerEmail + "/" + img.ImageUrl,
							ColorId:   img.ColorId,
						})
					}
				}

				for _, sz := range sizeData {
					if sz.ColorId == clr.Id && sz.ProductId == pro.Id {
						resultsItems.Sizes = append(resultsItems.Sizes, response.Sizes{
							SizeId:    sz.SizeId,
							ProductId: sz.ProductId,
							Size:      sz.Size,
							ColorId:   sz.ColorId,
							Quantity:  sz.Quantity,
						})
					}
				}
			}

		}

		results = append(results, resultsItems)
	}

	response := utills.PaginatedResponse(int64(count), 0, results)

	return response, nil
}

func (r *ProductRepository) GetProductById(req request.RequestProductById) (*utills.PaginationResponse, error) {
	var data []struct {
		Id        int    `json:"id" gorm:"column:id"`
		Name      string `json:"name" gorm:"column:name"`
		ItemsName string `json:"itemsName" gorm:"column:itemsName"`
		ItemsId   int    `json:"itemsId" gorm:"column:itemsId"`
	}

	var baseQuery strings.Builder

	baseQuery.WriteString(fmt.Sprintf(`SELECT "pc"."id","pc"."name","ProductType"."itemsName" as "itemsName","ProductType"."id" as "itemsId" FROM "ProductCategory" as "pc" INNER JOIN "ProductType" ON "ProductType"."productCategoryId" = "pc"."id" WHERE "pc"."id"= '%d'`, req.Id))

	if err := r.db.Raw(baseQuery.String()).Scan(&data).Error; err != nil {
		return nil, err
	}

	count := len(data)

	response := utills.PaginatedResponse(int64(count), 0, data)

	return response, nil
}
