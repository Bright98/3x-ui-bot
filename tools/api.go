package tools

import (
	"encoding/json"
	"errors"
	"fmt"
)

var session = ""

func Login(serverInfo *ServerInfo) (err error) {
	var response []byte

	loginBody := make(map[string]any)
	loginBody["username"] = serverInfo.Username
	loginBody["password"] = serverInfo.Password

	response, session, err = postApi(loginBody, serverInfo.IP, serverInfo.Port, "/login")
	if err != nil {
		return err
	}

	res, err := UnmarshalBody(response)
	if err != nil {
		return err
	}

	if !res.Success {
		return errors.New(LoginFailedErr)
	}
	return nil
}
func GetClientTraffic(serverInfo *ServerInfo, userEmail string) (*UserInboundResponse, error) {
	response, err := getApi(serverInfo.IP, serverInfo.Port, "/panel/api/inbounds/getClientTraffics/"+userEmail, session)
	if err != nil {
		if err.Error() == AuthErr {
			err = Login(serverInfo)
			if err != nil {
				fmt.Println("login error")
				panic(err)
			}
			response, err = getApi(serverInfo.IP, serverInfo.Port, "/panel/api/inbounds/getClientTraffics/"+userEmail, session)
		} else {
			return nil, err
		}
	}

	res, err := UnmarshalBody(response)
	if err != nil {
		return nil, err
	}

	if res.Object != nil {
		clientTraffic := &UserInboundResponse{}
		responseObj, err := json.Marshal(res.Object)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(responseObj, clientTraffic)
		if err != nil {
			return nil, err
		}

		return clientTraffic, nil
	} else {
		return nil, errors.New(UserNotFoundErr)
	}
}
