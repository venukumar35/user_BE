package response

import "time"

type ProductDetailsResponse struct {
	Id                       int       `json:"id" gorm:"column:id"`
	Title                    string    `json:"title" gorm:"column:title"`
	Price                    int16     `json:"price" gorm:"column:price"`
	SeasonalName             string    `json:"seasonalName" gorm:"column:seasonalName"`
	CommonDesId              string    `json:"commonDesId" gorm:"column:commonDesId"`
	Fit                      string    `json:"fit" gorm:"column:fit"`
	Materail                 string    `json:"materail" gorm:"column:materail"`
	Care                     string    `json:"care" gorm:"column:care"`
	BrandName                string    `json:"brandName" gorm:"column:brandName"`
	Origin                   string    `json:"origin" gorm:"column:origin"`
	ProductId                int16     `json:"productId" gorm:"column:productId"`
	Occasion                 string    `json:"occasion" gorm:"column:occasion"`
	SpecialFeature           string    `json:"specialFeature" gorm:"column:specialFeature"`
	TopId                    int16     `json:"topId" gorm:"column:topId"`
	ProductDescription       string    `json:"productDescription" gorm:"column:productDescription"`
	Weight                   int       `json:"weight" gorm:"column:weight"`
	PrintAndPattern          string    `json:"printAndPattern" gorm:"column:printAndPattern"`
	Length                   int       `json:"length" gorm:"column:length"`
	Shoulder                 int       `json:"shoulder" gorm:"column:shoulder"`
	Chest                    int       `json:"chest" gorm:"column:chest"`
	CommonDescriptionId      int       `json:"commonDescriptionId" gorm:"column:commonDescriptionId"`
	Type                     string    `json:"type" gorm:"column:type"`
	SleeveTypeId             int       `json:"sleeveTypeId" gorm:"column:sleeveTypeId"`
	NeckTypeId               int       `json:"neckTypeId" gorm:"column:neckTypeId"`
	ColorFamily              string    `json:"colorFamily" gorm:"column:colorFamily"`
	Pocket                   string    `json:"pocket" gorm:"column:pocket"`
	BtId                     int       `json:"btId" gorm:"column:btId"` //Bottom
	BottomproductDescription string    `json:"bottomproductDescription" gorm:"column:bottomproductDescription"`
	BtWeight                 int       `json:"btWeight" gorm:"column:btWeight"`
	BtPrintAndPattern        string    `json:"btPrintAndPattern" gorm:"column:btPrintAndPattern"`
	BtLength                 string    `json:"btLength" gorm:"column:btLength"`
	BtWaist                  int       `json:"btWaist" gorm:"column:btWaist"`
	BtHip                    int       `json:"btHip" gorm:"column:btHip"`
	BtcommonDescriptionId    int       `json:"btcommonDescriptionId" gorm:"column:btcommonDescriptionId"`
	BtType                   string    `json:"btType" gorm:"column:btType"`
	BtcolorFamily            string    `json:"btcolorFamily" gorm:"column:btcolorFamily"`
	BtPocket                 string    `json:"btPocket" gorm:"column:btPocket"`
	BtBeltLoop               bool      `json:"btBeltLoop" gorm:"column:btBeltLoop"`
	TypeOfPantId             int       `json:"typeOfPantId" gorm:"column:typeOfPantId"`
	TypesOfPleatsId          int       `json:"typesOfPleatsId" gorm:"column:typesOfPleatsId"`
	TypesOfLengthId          int       `json:"typesOfLengthId" gorm:"column:typesOfLengthId"`
	TopbtId                  int       `json:"topbtId" gorm:"column:topbtId"` //BtTopDescriptionId
	KurtaBt                  int       `json:"kurtaBt" gorm:"column:kurtaBt"` //BtKurtasDescriptionId
	KurtasId                 int       `json:"kurtasId" gorm:"column:kurtasId"`
	TransparencyOfTheFabric  bool      `json:"transparencyOfTheFabric" gorm:"column:transparencyOfTheFabric"`
	ShIp                     int       `json:"shIp" gorm:"column:shIp"` //shoes
	Pattern                  string    `json:"pattern" gorm:"column:pattern"`
	FootLength               string    `json:"footLength" gorm:"column:footLength"`
	SoleMaterial             string    `json:"soleMaterial" gorm:"column:soleMaterial"`
	UpperMaterial            string    `json:"upperMaterial" gorm:"column:upperMaterial"`
	Closure                  string    `json:"closure" gorm:"column:closure"`
	ToeType                  string    `json:"toeType" gorm:"column:toeType"`
	PackageContains          int       `json:"packageContains" gorm:"column:packageContains"`
	WarrentyId               int       `json:"warrentyId" gorm:"column:warrentyId"`
	WarrantyYear             int       `json:"warrantyYear" gorm:"column:warrantyYear"`
	ShoesWarrentyId          int       `json:"shoesWarrentyId" gorm:"column:shoesWarrentyId"`
	WdId                     int       `json:"wdId" gorm:"column:wdId"`
	Model                    string    `json:"model" gorm:"column:model"`
	DailShape                string    `json:"dailShape" gorm:"column:dailShape"`
	DialDiameter             string    `json:"dialDiameter" gorm:"column:dialDiameter"`
	DialColor                string    `json:"dialColor" gorm:"column:dialColor"`
	StrapColor               string    `json:"strapColor" gorm:"column:strapColor"`
	WatchesWarrentyId        int       `json:"watchesWarrentyId" gorm:"column:watchesWarrentyId"`
	OfferId                  int       `json:"offerId" gorm:"column:offerId"`
	OfferPercntage           int       `json:"offerPercntage" gorm:"column:offerPercntage"`
	OfferPrice               int16     `json:"offerPrice" gorm:"column:offerPrice"`
	OfferValidityFromdate    time.Time `json:"offerValidityFromdate" gorm:"column:offerValidityFromdate"`
	OfferValidityTodate      time.Time `json:"offerValidityTodate" gorm:"column:offerValidityTodate"`
	OfferValidityFromTime    time.Time `json:"offerValidityFromTime" gorm:"column:offerValidityFromTime"`
	OfferValidityToTime      time.Time `json:"offerValidityToTime" gorm:"column:offerValidityToTime"`
	OwnerEmail               string    `json:"ownwrEmail" gorm:"column:ownwrEmail"`
	StoreName                string    `json:"storeName" gorm:"column:storeName"`
	StoreAddress             string    `json:"storeAddress" gorm:"column:storeAddress"`
	StoreCity                string    `json:"storeCity" gorm:"column:storeCity"`
	StorePincode             string    `json:"storePincode" gorm:"column:storePincode"`
	StoreState               string    `json:"StoreState" gorm:"column:StoreState"`
	CustomerCareEmail        string    `json:"customerCareEmail" gorm:"column:customerCareEmail"`
}
type FinalProductResponse struct {
	Id                       int       `json:"id" gorm:"column:id"`
	Title                    string    `json:"title" gorm:"column:title"`
	Price                    int16     `json:"price" gorm:"column:price"`
	SeasonalName             string    `json:"seasonalName" gorm:"column:seasonalName"`
	CommonDesId              string    `json:"commonDesId" gorm:"column:commonDesId"`
	Fit                      string    `json:"fit" gorm:"column:fit"`
	Materail                 string    `json:"materail" gorm:"column:materail"`
	Care                     string    `json:"care" gorm:"column:care"`
	BrandName                string    `json:"brandName" gorm:"column:brandName"`
	Origin                   string    `json:"origin" gorm:"column:origin"`
	ProductId                int16     `json:"productId" gorm:"column:productId"`
	Occasion                 string    `json:"occasion" gorm:"column:occasion"`
	SpecialFeature           string    `json:"specialFeature" gorm:"column:specialFeature"`
	TopId                    int16     `json:"topId" gorm:"column:topId"`
	ProductDescription       string    `json:"productDescription" gorm:"column:productDescription"`
	Weight                   int       `json:"weight" gorm:"column:weight"`
	PrintAndPattern          string    `json:"printAndPattern" gorm:"column:printAndPattern"`
	Length                   int       `json:"length" gorm:"column:length"`
	Shoulder                 int       `json:"shoulder" gorm:"column:shoulder"`
	Chest                    int       `json:"chest" gorm:"column:chest"`
	CommonDescriptionId      int       `json:"commonDescriptionId" gorm:"column:commonDescriptionId"`
	Type                     string    `json:"type" gorm:"column:type"`
	SleeveTypeId             int       `json:"sleeveTypeId" gorm:"column:sleeveTypeId"`
	NeckTypeId               int       `json:"neckTypeId" gorm:"column:neckTypeId"`
	ColorFamily              string    `json:"colorFamily" gorm:"column:colorFamily"`
	Pocket                   string    `json:"pocket" gorm:"column:pocket"`
	BtId                     int       `json:"btId" gorm:"column:btId"` //Bottom
	BottomproductDescription string    `json:"bottomproductDescription" gorm:"column:bottomproductDescription"`
	BtWeight                 int       `json:"btWeight" gorm:"column:btWeight"`
	BtPrintAndPattern        string    `json:"btPrintAndPattern" gorm:"column:btPrintAndPattern"`
	BtLength                 string    `json:"btLength" gorm:"column:btLength"`
	BtWaist                  int       `json:"btWaist" gorm:"column:btWaist"`
	BtHip                    int       `json:"btHip" gorm:"column:btHip"`
	BtcommonDescriptionId    int       `json:"btcommonDescriptionId" gorm:"column:btcommonDescriptionId"`
	BtType                   string    `json:"btType" gorm:"column:btType"`
	BtcolorFamily            string    `json:"btcolorFamily" gorm:"column:btcolorFamily"`
	BtPocket                 string    `json:"btPocket" gorm:"column:btPocket"`
	BtBeltLoop               bool      `json:"btBeltLoop" gorm:"column:btBeltLoop"`
	TypeOfPantId             int       `json:"typeOfPantId" gorm:"column:typeOfPantId"`
	TypesOfPleatsId          int       `json:"typesOfPleatsId" gorm:"column:typesOfPleatsId"`
	TypesOfLengthId          int       `json:"typesOfLengthId" gorm:"column:typesOfLengthId"`
	TopbtId                  int       `json:"topbtId" gorm:"column:topbtId"` //BtTopDescriptionId
	KurtaBt                  int       `json:"kurtaBt" gorm:"column:kurtaBt"` //BtKurtasDescriptionId
	KurtasId                 int       `json:"kurtasId" gorm:"column:kurtasId"`
	TransparencyOfTheFabric  bool      `json:"transparencyOfTheFabric" gorm:"column:transparencyOfTheFabric"`
	ShIp                     int       `json:"shIp" gorm:"column:shIp"` //shoes
	Pattern                  string    `json:"pattern" gorm:"column:pattern"`
	FootLength               string    `json:"footLength" gorm:"column:footLength"`
	SoleMaterial             string    `json:"soleMaterial" gorm:"column:soleMaterial"`
	UpperMaterial            string    `json:"upperMaterial" gorm:"column:upperMaterial"`
	Closure                  string    `json:"closure" gorm:"column:closure"`
	ToeType                  string    `json:"toeType" gorm:"column:toeType"`
	PackageContains          int       `json:"packageContains" gorm:"column:packageContains"`
	WarrentyId               int       `json:"warrentyId" gorm:"column:warrentyId"`
	WarrantyYear             int       `json:"warrantyYear" gorm:"column:warrantyYear"`
	ShoesWarrentyId          int       `json:"shoesWarrentyId" gorm:"column:shoesWarrentyId"`
	WdId                     int       `json:"wdId" gorm:"column:wdId"`
	Model                    string    `json:"model" gorm:"column:model"`
	DailShape                string    `json:"dailShape" gorm:"column:dailShape"`
	DialDiameter             string    `json:"dialDiameter" gorm:"column:dialDiameter"`
	DialColor                string    `json:"dialColor" gorm:"column:dialColor"`
	StrapColor               string    `json:"strapColor" gorm:"column:strapColor"`
	WatchesWarrentyId        int       `json:"watchesWarrentyId" gorm:"column:watchesWarrentyId"`
	OfferId                  int       `json:"offerId" gorm:"column:offerId"`
	OfferPercntage           int       `json:"offerPercntage" gorm:"column:offerPercntage"`
	OfferPrice               int16     `json:"offerPrice" gorm:"column:offerPrice"`
	OfferValidityFromdate    time.Time `json:"offerValidityFromdate" gorm:"column:offerValidityFromdate"`
	OfferValidityTodate      time.Time `json:"offerValidityTodate" gorm:"column:offerValidityTodate"`
	OfferValidityFromTime    time.Time `json:"offerValidityFromTime" gorm:"column:offerValidityFromTime"`
	OfferValidityToTime      time.Time `json:"offerValidityToTime" gorm:"column:offerValidityToTime"`
	StoreName                string    `json:"storeName" gorm:"column:storeName"`
	StoreAddress             string    `json:"storeAddress" gorm:"column:storeAddress"`
	StoreCity                string    `json:"storeCity" gorm:"column:storeCity"`
	StorePincode             string    `json:"storePincode" gorm:"column:storePincode"`
	StoreState               string    `json:"StoreState" gorm:"column:StoreState"`
	CustomerCareEmail        string    `json:"customerCareEmail" gorm:"column:customerCareEmail"`
	OwnerEmail               string    `json:"ownwrEmail" gorm:"column:ownwrEmail"`
	Colors                   []Clr     `json:"colors" gorm:"column:colors"`
	Images                   []Images  `json:"images" gorm:"column:images"`
	Sizes                    []Sizes   `json:"sizes" gorm:"column:sizes"`
}
type Clr struct {
	Id        int    `json:"id" gorm:"column:id"`
	ProductId int    `json:"productId" gorm:"column:productId"`
	Colors    string `json:"colors" gorm:"column:colors"`
}
type Images struct {
	ImgId     int    `json:"imgId" gorm:"column:imgId"`
	ProductId int    `json:"productId" gorm:"column:productId"`
	ImageUrl  string `json:"imageUrl" gorm:"column:imageUrl"`
	ColorId   int    `json:"clrId" gorm:"column:clrId"`
}
type Sizes struct {
	SizeId    int    `json:"sizeId" gorm:"column:sizeId"`
	ProductId int    `json:"productId" gorm:"column:productId"`
	Size      string `json:"size" gorm:"column:size"`
	ColorId   int    `json:"clrId" gorm:"column:clrId"`
	Quantity  string `json:"quantity" gorm:"column:quantity"`
}
