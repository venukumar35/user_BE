package utills

import (
	"TheBoys/infrastructure/config"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func GenerateOTP() string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := rand.Intn(900000) + 100000
	return strconv.Itoa(otp)
}

func SendSms(timeZoneOffset int, mobile string, username string, location string, otp string) error {
	currentDate := time.Now().UTC()

	localTime := currentDate.Add(time.Duration(timeZoneOffset) * time.Minute)

	client := &http.Client{}

	var baseUrl string = "https://otp2.maccesssmspush.com/OTP_ACL_Web/OtpRequestListener"

	queryParams := url.Values{}

	queryParams.Add("enterpriseid", "sternaotp")
	queryParams.Add("subEnterpriseid", "sternaotp")
	queryParams.Add("pusheid", "sternaotp")
	queryParams.Add("pushepwd", "sternaotp")
	queryParams.Add("msisdn", mobile)
	queryParams.Add("sender", "STERNA")
	queryParams.Add("templateId", "2")
	queryParams.Add("msgtext", fmt.Sprintf(`Hi %s, Request Time : %v, Location : %s, Code : %s. Regards STERNA`, username, localTime.Format("02 January 2006, 03:04 PM"), location, otp))

	requestURL := fmt.Sprintf("%s?%s", baseUrl, queryParams.Encode())

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		return err
	}

	response, err := client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	return nil
}

func SendMail(to, otp string, offset int) error {
	currentDate := time.Now().UTC()

	displayName := config.Config.SmtpDisplayName
	from := config.Config.SmtpUserName
	password := config.Config.SmtpPassword
	smtpHost := config.Config.SmtpHost
	smtpPort := config.Config.SmtpPort

	localTime := currentDate.Add(time.Duration(offset) * time.Minute)

	subject := "The boys login otp"

	body := fmt.Sprintf(`<p>Your  OTP is <b>%s</b> requested at: <b>%s</b></p>`, otp, localTime.Format("02 January 2006, 03:04 PM"))

	message := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s", from, to, subject, body)

	auth := smtp.PlainAuth(displayName, from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}

type PaginationResponse struct {
	From       int         `json:"from"`
	To         int         `json:"to"`
	TotalCount int         `json:"totalCount"`
	TotalPages int         `json:"totalPages"`
	Data       interface{} `json:"data"`
}

func PaginatedResponse(totalCount int64, page int, data interface{}) *PaginationResponse {
	limit := 10

	var offset int
	from := 0
	to := 0
	totalPages := 0

	if page > 0 {
		offset = (page - 1) * limit
	}

	if totalCount > 0 && page > 0 {
		from = offset + 1
		if (offset + limit) > int(totalCount) {
			to = int(totalCount)
		} else {
			to = offset + 10
		}
		totalPages = int(math.Ceil(float64(totalCount) / float64(limit)))
	}

	return &PaginationResponse{
		From:       from,
		To:         to,
		TotalCount: int(totalCount),
		TotalPages: totalPages,
		Data:       data,
	}
}

type DateFilter string

const (
	All           DateFilter = "All"
	Today         DateFilter = "Today"
	Yesterday     DateFilter = "Yesterday"
	MonthTillDate DateFilter = "MonthTillDate"
	DateRange     DateFilter = "DateRange"
)

func DateFilterResponse(dateFilter DateFilter, startDate string, endDate string) (map[string]interface{}, error) {
	dateFilterRes := make(map[string]interface{})

	switch dateFilter {
	case Today:
		today := time.Now().UTC().Format("2006-01-02")
		dateFilterRes["gte"] = today + "T00:00:00.000Z"
		dateFilterRes["lte"] = today + "T23:59:59.000Z"
	case Yesterday:
		yesterday := time.Now().UTC().AddDate(0, 0, -1).Format("2006-01-02")
		dateFilterRes["gte"] = yesterday + "T00:00:00.000Z"
		dateFilterRes["lte"] = yesterday + "T23:59:59.000Z"
	case DateRange:
		_, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			return nil, fmt.Errorf("error parsing start date: %v", err)
		}
		_, err = time.Parse("2006-01-02", endDate)
		if err != nil {
			return nil, fmt.Errorf("error parsing end date: %v", err)
		}
		dateFilterRes["gte"] = startDate + "T00:00:00.000Z"
		dateFilterRes["lte"] = endDate + "T23:59:59.000Z"
	case MonthTillDate:
		date := time.Now().UTC()
		date = time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
		dateFilterRes["gte"] = date.Format("2006-01-02") + "T00:00:00.000Z"
		dateFilterRes["lte"] = time.Now().UTC().Format("2006-01-02") + "T23:59:59.000Z"
	case All:
		dateFilterRes["lte"] = time.Now().UTC().Format("2006-01-02") + "T23:59:59.000Z"
		dateFilterRes["gte"] = "2000-01-01T00:00:00.000Z"
	default:
		return nil, fmt.Errorf("invalid date filter")
	}
	return dateFilterRes, nil
}

func SqlParamValidator(input string) string {
	unsafeChars := []string{"'", "\"", ";", "--"}
	sanitizedInput := input
	for _, char := range unsafeChars {
		if char == "'" {
			sanitizedInput = strings.ReplaceAll(sanitizedInput, char, "''")
		} else {
			sanitizedInput = strings.ReplaceAll(sanitizedInput, char, "")
		}
	}
	return sanitizedInput
}

func Duration(fromDate, toDate, fromTime, toTime string) (float64, error) {
	const dateFormat = "2006-01-02 15:04"

	startStr := fromDate + " " + fromTime
	endStr := toDate + " " + toTime

	startTime, err := time.Parse(dateFormat, startStr)
	if err != nil {
		return 0, err
	}
	endTime, err := time.Parse(dateFormat, endStr)
	if err != nil {
		return 0, err
	}

	duration := endTime.Sub(startTime)
	return duration.Hours(), nil
}
func PadLeft(str string) string {
	value, err := strconv.Atoi(str)

	if err != nil {
		return ""
	}

	if value < 10 {
		return fmt.Sprintf("0%s", str)
	}
	return str
}

type ProductCategory string

const (
	TopWear     ProductCategory = "Top"
	BottomWear  ProductCategory = "Bottom"
	EthnicWear  ProductCategory = "Ethnic"
	SportsWear  ProductCategory = "Sports"
	Fragrances  ProductCategory = "Fragrances"
	Footwear    ProductCategory = "Footwear"
	Accessories ProductCategory = "Accessories"
	Innerwear   ProductCategory = "Inner"
)
