package schedulers

type ProductDetails struct {
	UId                      int    `json:"uId" gorm:"column:uId"`
	Name                     string `json:"name" gorm:"column:name"`
	Email                    string `json:"email" gorm:"column:email"`
	Mobile                   string `json:"mobile" gorm:"column:mobile"`
	PID                      int    `json:"pId" gorm:"column:pId"`
	Title                    string `json:"title" gorm:"column:title"`
	Price                    string `json:"price" gorm:"column:price"`
	PdTyId                   int    `josn:"pdTyId" gorm:"column:pdTyId"` //product type id
	ProductTypeName          string `json:"productTypeName" gorm:"column:productTypeName"`
	PdCtId                   int    `json:"pdCtId" gorm:"column:pdCtId"` //Product category id
	ProductCategoryName      string `json:"productCategoryName" gorm:"column:productCategoryName"`
	CmId                     int    `json:"cmId" gorm:"column:cmId"`
	Fit                      string `json:"fit" gorm:"column:fit"`
	Material                 string `json:"material" gorm:"column:material"`
	Care                     string `json:"care" gorm:"column:care"`
	BrandName                string `json:"brandName" gorm:"column:brandName"`
	Origin                   string `json:"origin" gorm:"column:origin"`
	Occasion                 string `json:"occasion" gorm:"column:occasion"`
	SpecialFeature           string `json:"specialFeature" gorm:"column:specialFeature"`
	TpId                     int    `json:"tpId" gorm:"column:tpId"` //Top
	ProductDescription       string `json:"productDescription" gorm:"column:productDescription"`
	SleeveTypeId             int    `json:"sleeveTypeId" gorm:"column:sleeveTypeId"`
	Weight                   int    `json:"weight" gorm:"column:weight"`
	Chest                    int    `json:"chest" gorm:"column:chest"`
	Shoulder                 int    `json:"shoulder" gorm:"column:shoulder"`
	NeckTypeId               int    `json:"neckTypeId" gorm:"column:neckTypeId"`
	Type                     string `json:"type" gorm:"column:type"`
	ColorFamily              string `json:"colorFamily" gorm:"column:colorFamily"`
	PrintAndPattern          string `json:"printAndPattern" gorm:"column:printAndPattern"`
	Length                   int    `json:"length" gorm:"column:length"`
	Pocket                   string `json:"pocket" gorm:"column:pocket"`
	CommonDescriptionId      int    `json:"commonDescriptionId" gorm:"column:commonDescriptionId"`
	BtId                     int    `json:"btId" gorm:"column:btId"` //Bottom
	BottomproductDescription string `json:"bottomproductDescription" gorm:"column:bottomproductDescription"`
	BtWeight                 int    `json:"btWeight" gorm:"column:btWeight"`
	BtPrintAndPattern        string `json:"btPrintAndPattern" gorm:"column:btPrintAndPattern"`
	BtLength                 string `json:"btLength" gorm:"column:btLength"`
	BtWaist                  int    `json:"btWaist" gorm:"column:btWaist"`
	BtHip                    int    `json:"btHip" gorm:"column:btHip"`
	BtcommonDescriptionId    int    `json:"btcommonDescriptionId" gorm:"column:btcommonDescriptionId"`
	BtType                   string `json:"btType" gorm:"column:btType"`
	BtcolorFamily            string `json:"btcolorFamily" gorm:"column:btcolorFamily"`
	BtPocket                 string `json:"btPocket" gorm:"column:btPocket"`
	BtBeltLoop               bool   `json:"btBeltLoop" gorm:"column:btBeltLoop"`
	TypeOfPantId             int    `json:"typeOfPantId" gorm:"column:typeOfPantId"`
	TypesOfPleatsId          int    `json:"typesOfPleatsId" gorm:"column:typesOfPleatsId"`
	TypesOfLengthId          int    `json:"typesOfLengthId" gorm:"column:typesOfLengthId"`
	TopbtId                  int    `json:"topbtId" gorm:"column:topbtId"` //BttopDescriptionId
	KurtaBt                  int    `json:"kurtaBt" gorm:"column:kurtaBt"` //BtkurtasDescriptionId
	KurId                    int    `json:"kurId" gorm:"column:kurId"`
	Work                     string `json:"work" gorm:"column:work"` //kurtas
	TransparencyOfTheFabric  bool   `json:"transparencyOfTheFabric" gorm:"column:transparencyOfTheFabric"`
	SdIp                     int    `json:"sdIp" gorm:"column:sdId"` //shoes
	Pattern                  string `json:"pattern" gorm:"column:pattern"`
	FootLength               string `json:"footLength" gorm:"column:footLength"`
	SoleMaterial             string `json:"soleMaterial" gorm:"column:soleMaterial"`
	UpperMaterial            string `json:"upperMaterial" gorm:"column:upperMaterial"`
	Closure                  string `json:"closure" gorm:"column:closure"`
	ToeType                  string `json:"toeType" gorm:"column:toeType"`
	PackageContains          int    `json:"packageContains" gorm:"column:packageContains"`
	WarrentyId               int    `json:"warrentyId" gorm:"column:warrentyId"`
	WarrantyYear             int    `json:"warrantyYear" gorm:"column:warrantyYear"`
	ShoesWarrentyId          int    `json:"shoesWarrentyId" gorm:"column:shoesWarrentyId"`
	WdId                     int    `json:"wdId" gorm:"column:wdId"`
	Model                    string `json:"model" gorm:"column:model"`
	DailShape                string `json:"dailShape" gorm:"column:dailShape"`
	DialDiameter             string `json:"dialDiameter" gorm:"column:dialDiameter"`
	DialColor                string `json:"dialColor" gorm:"column:dialColor"`
	StrapColor               string `json:"strapColor" gorm:"column:strapColor"`
	WatchesWarrentyId        int    `json:"watchesWarrentyId" gorm:"column:watchesWarrentyId"`
	PerfumeId                int    `json:"perfumeId" gorm:"column:perfumeId"`
	MaterialDescription      string `json:"materialDescription" gorm:"column:materialDescription"`
	InId                     int    `json:"inId" gorm:"column:inId"`
	LookAndFeel              string `json:"lookAndFeel" gorm:"column:lookAndFeel"`
	MultiColors              bool   `json:"multiColors" gorm:"column:multiColors"`
	PdIp                     int    `jon:"pdIp" gorm:"column:pdId"`
	SleeveName               string `json:"sleeveName" gorm:"column:sleeveName"`
	NeckName                 string `json:"neckName" gorm:"column:neckName"`
	KurtasTypeName           string `json:"kurtasTypeName" gorm:"column:kurtasTypeName"`
	BtTypeName               string `json:"btTypeName" gorm:"column:btTypeName"`
	BtmPleatName             string `json:"btmPleatName" gorm:"column:btmPleatName"`
	BtmLengthName            string `json:"btmLengthName" gorm:"column:btmLengthName"`
}
type UserDetails struct {
	UID int `json:"uId" gorm:"column:uId"`
}

type productColorDetails struct {
	Id     int    `json:"id" gorm:"column:id"`
	PdId   int    `json:"pId" gorm:"column:pId"`
	UId    int    `json:"uId" gorm:"column:uId"`
	Colors string `json:"colors" gorm:"column:colors"`
}
type seasonalDetails struct {
	Id           int64  `json:"id" gorm:"column:id"`
	PdId         int    `json:"pId" gorm:"column:pId"`
	UId          int    `json:"uId" gorm:"column:uId"`
	SeasonalName string `json:"seasonalName" gorm:"column:seasonalName"`
}
type SizeDetails struct {
	Id       int64  `json:"id" gorm:"column:id"`
	PdId     int    `json:"pId" gorm:"column:pId"`
	UId      int    `json:"uId" gorm:"column:uId"`
	ClrId    int    `json:"clrId" gorm:"column:clrId"`
	Quantity int    `json:"quantity" gorm:"column:quantity"`
	Size     string `json:"size" gorm:"column:size"`
}
type imgDetails struct {
	Id         int64  `json:"id" gorm:"column:id"`
	PdId       int    `json:"pId" gorm:"column:pId"`
	UId        int    `json:"uId" gorm:"column:uId"`
	CId        int    `json:"cId" gorm:"column:cId"`
	ImageUrl   string `json:"imageUrl" gorm:"column:imageUrl"`
	ImageBytes string `json:"imageBytes" gorm:"column:imageBytes"`
}
