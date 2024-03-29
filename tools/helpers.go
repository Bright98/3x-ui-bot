package tools

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	ptime "github.com/yaa110/go-persian-calendar"
	"image/jpeg"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func checkStringStartWith(str string, startWith string) bool {
	return strings.HasPrefix(str, startWith)
}
func GetStringInBetween(str string, start string, end string) string {
	s := strings.Index(str, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return ""
	}
	e = s + e
	return str[s:e]
}
func decodeBase64(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}
func convertTimestampToPersian(timestamp int64) string {
	if timestamp == 0 {
		return Unlimited
	}

	pt := ptime.Unix(0, timestamp*int64(time.Millisecond))
	return pt.Format("yyyy/MM/dd")
}
func clearUsage(usage int64) string {
	if usage == 0 {
		return Unlimited
	}

	unit := ""
	unitCount := 0
	_usage := float64(0)

	for getDigitOfNumber(usage) > 3 {
		_usage = float64(usage) / float64(1024)
		usage = usage / 1024
		unitCount++
	}

	switch unitCount {
	case 1:
		unit = KiloByte
	case 2:
		unit = MegaByte
	case 3:
		unit = GigaByte
	case 4:
		unit = TeraByte
	case 5:
		unit = PetaByte
	default:
	}

	return fmt.Sprintf("%0.2f ", _usage) + unit
}
func getDigitOfNumber(i int64) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
func scanQRCode(photo []byte) (string, error) {
	// >>>> old package with errors
	// >>>> "github.com/tuotoo/qrcode"
	//qr, err := qrcode.Decode(bytes.NewReader(photo))
	//if err != nil {
	//	return "", errors.New(CantDecodeImageErr)
	//}
	//return qr.Content, nil

	// >>>> use new package
	img, err := jpeg.Decode(bytes.NewBuffer(photo))
	if err != nil {
		return "", errors.New(CantDecodeImageErr)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", errors.New(CantDecodeImageErr)
	}

	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", errors.New(CantDecodeImageErr)
	}

	return result.String(), nil
}
func getImageFromBot(message telego.Message) ([]byte, error) {
	photo := message.Photo
	if photo == nil {
		return nil, errors.New(MessageIsNotImageTypeErr)
	}
	lastPhoto := len(photo) - 1

	file, err := bot.GetFile(&telego.GetFileParams{FileID: photo[lastPhoto].FileID})
	if err != nil {
		return nil, errors.New(CantGetImageErr)
	}

	url := "https://api.telegram.org/file/bot" + os.Getenv(BotToken) + "/" + file.FilePath
	bytes, err := tu.DownloadFile(url)
	if err != nil {
		return nil, errors.New(CantGetImageErr)
	}

	return bytes, err
}

func postApi(body map[string]any, ip, port, route string) ([]byte, string, error) {
	bodyByte, _ := json.Marshal(body)
	url := "http://" + ip + ":" + port + route
	req, err := http.NewRequest(
		"POST", url, bytes.NewBuffer(bodyByte),
	)
	if err != nil {
		return nil, "", errors.New(CantConnectErr)
	}
	req.Header.Set("Accept", "/")
	req.Header.Set("Content-Type", "application/json")

	req.Close = true
	client := &http.Client{Transport: &http.Transport{}}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil, "", errors.New(CantConnectErr)
	}

	defer res.Body.Close()
	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, "", errors.New(InvalidationErr)
	}
	return responseBody, findSessionInCookies(res.Cookies()), nil
}
func getApi(ip, port, route string, loginSession string) ([]byte, error) {
	url := "http://" + ip + ":" + port + route
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, errors.New(CantConnectErr)
	}
	req.Close = true
	req.Header.Set("accept", "application/json")
	cookie := http.Cookie{Name: "session", Value: loginSession}
	req.AddCookie(&cookie)

	client := &http.Client{Transport: &http.Transport{}}
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.New(CantConnectErr)
	}
	if res.Header.Get("Content-Type") == "text/html; charset=utf-8" {
		return nil, errors.New(AuthErr)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New(InvalidationErr)
	}
	return body, nil
}
func UnmarshalBody(bodyBytes []byte) (*ApiResponse, error) {
	response := &ApiResponse{}
	err := json.Unmarshal(bodyBytes, response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
func findSessionInCookies(cookies []*http.Cookie) string {
	for _, cookie := range cookies {
		if cookie.Name == "session" {
			return cookie.Value
		}
	}
	return ""
}
func convertErrorMessage(err string) string {
	switch err {
	case AuthErr:
		return CantConnectToServer
	case LoginFailedErr:
		return CantConnectToServer
	case CantConnectErr:
		return CantConnectToServer
	case InvalidationErr:
		return CantConnectToServer
	case ProtocolNotFoundErr:
		return InvalidConfig
	case UserNotFoundErr:
		return UserNotExist
	case CantGetImageErr:
		return CantGetImage
	case CantDecodeImageErr:
		return CantDecodeImage
	default:
		return SomethingGetWrong
	}
}
func findServerInfo(ip string) *ServerInfo {
	serverInfo := &ServerInfo{}
	for _, server := range RequirementsValue.Servers {
		if server.IP == ip {
			serverInfo = &server
			return serverInfo
		}
	}

	return serverInfo
}
func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
