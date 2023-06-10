package tools

import (
	"encoding/json"
	"errors"
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
func getUserEmailFromVlessConfigURL(url string) string {
	splitString := strings.Split(url, "#")
	configUser := splitString[len(splitString)-1]
	splitConfigUser := strings.Split(configUser, "-")
	userEmail := splitConfigUser[len(splitConfigUser)-1]

	return userEmail
}
func getUserEmailFromVmessConfigURL(url string) (string, error) {
	vmessBody, err := convertVmessToVmessBody(url)
	if err != nil {
		return "", err
	}
	splitUserInfo := strings.Split(vmessBody.UserInfo, "-")
	userEmail := splitUserInfo[len(splitUserInfo)-1]

	return userEmail, nil
}
func GetUserEmailFromConfigURL(url string) (string, error) {
	protocol, err := checkProtocol(url)
	if err != nil {
		return "", err
	}

	userEmail := ""
	if protocol == VLESS {
		userEmail = getUserEmailFromVlessConfigURL(url)
	} else if protocol == VMESS {
		userEmail, err = getUserEmailFromVmessConfigURL(url)
		if err != nil {
			return "", err
		}
	}
	return userEmail, nil
}
func readClientTraffic(clientTraffic *UserInboundResponse) string {
	text := ""

	text += "*" + DownloadClientUsage + "*" + clearUsage(clientTraffic.Down) + "\n"
	text += "*" + UploadClientUsage + "*" + clearUsage(clientTraffic.Up) + "\n"
	text += "*" + TotalClientUsage + "*" + clearUsage(clientTraffic.Down+clientTraffic.Up) + "\n"
	text += "*" + AllowedClientUsage + "*" + clearUsage(clientTraffic.Total) + "\n"
	text += "*" + ConfigExpireTime + "*" + convertTimestampToPersian(clientTraffic.ExpiredAt) + "\n"

	return text
}
func getClientTraffic(url string) string {
	mail, err := GetUserEmailFromConfigURL(url)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	clientTraffic, err := GetClientTraffic(mail)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	return readClientTraffic(clientTraffic)
}
func getClientTrafficByEmail(mail string) string {
	clientTraffic, err := GetClientTraffic(mail)
	if err != nil {
		return convertErrorMessage(err.Error())
	}

	return readClientTraffic(clientTraffic)
}
