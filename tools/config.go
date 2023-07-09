package tools

import (
	"encoding/json"
	"errors"
	"github.com/mymmrac/telego"
	"strings"
)

func checkProtocol(url string) (string, error) {
	if checkStringStartWith(url, VLESS+"://") {
		return VLESS, nil
	} else if checkStringStartWith(url, VMESS+"://") {
		return VMESS, nil
	} else {
		return "", errors.New(ProtocolNotFoundErr)
	}
}
func convertVmessToVmessBody(url string) (*VmessBody, error) {
	//remove vmess://
	newURL := url[8:]

	//decode base64
	decodedURL, err := decodeBase64(newURL)
	if err != nil {
		return nil, err
	}

	//unmarshal to vmessBody
	vmessBody := &VmessBody{}
	err = json.Unmarshal(decodedURL, vmessBody)
	if err != nil {
		return nil, err
	}

	return vmessBody, nil
}
func getUserEmailFromVlessConfigURL(url string) (string, string) {
	//get email
	splitString := strings.Split(url, "#")
	configUser := splitString[len(splitString)-1]
	splitConfigUser := strings.Split(configUser, "-")
	userEmail := splitConfigUser[len(splitConfigUser)-1]

	serverAddress := GetStringInBetween(url, "@", ":")

	return userEmail, serverAddress
}
func getUserEmailFromVmessConfigURL(url string) (string, string, error) {
	vmessBody, err := convertVmessToVmessBody(url)
	if err != nil {
		return "", "", err
	}
	splitUserInfo := strings.Split(vmessBody.UserInfo, "-")
	userEmail := splitUserInfo[len(splitUserInfo)-1]

	return userEmail, vmessBody.ServerAddress, nil
}
func GetUserEmailFromConfigURL(url string) (*ServerInfo, string, error) {
	protocol, err := checkProtocol(url)
	if err != nil {
		return nil, "", err
	}

	userEmail := ""
	serverAddress := ""
	if protocol == VLESS {
		userEmail, serverAddress = getUserEmailFromVlessConfigURL(url)
	} else if protocol == VMESS {
		userEmail, serverAddress, err = getUserEmailFromVmessConfigURL(url)
		if err != nil {
			return nil, "", err
		}
	}

	//find server info from requirements
	serverInfo := findServerInfo(serverAddress)

	return serverInfo, userEmail, nil
}
func readClientTraffic(clientTraffic *UserInboundResponse) string {
	text := ""

	text += "*" + ClientEmail + "*" + clientTraffic.Email + "\n"
	text += "*" + DownloadClientUsage + "*" + clearUsage(clientTraffic.Down) + "\n"
	text += "*" + UploadClientUsage + "*" + clearUsage(clientTraffic.Up) + "\n"
	text += "*" + TotalClientUsage + "*" + clearUsage(clientTraffic.Down+clientTraffic.Up) + "\n"
	text += "*" + AllowedClientUsage + "*" + clearUsage(clientTraffic.Total) + "\n"
	text += "*" + ConfigExpireTime + "*" + convertTimestampToPersian(clientTraffic.ExpiredAt) + "\n"

	return text
}
func getClientTraffic(url string) string {
	serverInfo, mail, err := GetUserEmailFromConfigURL(url)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	clientTraffic, err := GetClientTraffic(serverInfo, mail)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	return readClientTraffic(clientTraffic)
}
func getClientTrafficByEmail(ip, mail string) string {
	//find server info
	serverInfo := findServerInfo(ip)

	clientTraffic, err := GetClientTraffic(serverInfo, mail)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	return readClientTraffic(clientTraffic)
}
func getConfigUrlFromMessage(message telego.Message) string {
	photo, err := getImageFromBot(message)
	configUrl := ""

	if err != nil {
		if err.Error() == MessageIsNotImageTypeErr {
			//message is text
			configUrl = message.Text
		} else {
			sendMessage(message.Chat.ID, convertErrorMessage(err.Error()))
		}
	} else {
		//message is photo
		configUrl, err = scanQRCode(photo)
		if err != nil {
			sendMessage(message.Chat.ID, convertErrorMessage(err.Error()))
		}
	}

	return configUrl
}
