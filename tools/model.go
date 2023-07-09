package tools

type VmessBody struct {
	ID            string `json:"id"`
	Version       string `json:"v"`
	ServerAddress string `json:"add"`
	ServerPort    any    `json:"port"`
	UserInfo      string `json:"ps"`
	UserEmail     string `json:"email"`
}

type ApiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"msg"`
	Object  any    `json:"obj"`
}

type UserInboundResponse struct {
	ID        int64  `json:"id"`
	InboundID int64  `json:"inbound_Id"`
	Enable    bool   `json:"enable"`
	Email     string `json:"email"`
	Up        int64  `json:"up"`
	Down      int64  `json:"down"`
	ExpiredAt int64  `json:"expiryTime"`
	Total     int64  `json:"total"`
}
type Requirements struct {
	Servers []ServerInfo `json:"servers" yaml:"servers"`
}
type ServerInfo struct {
	IP       string `json:"ip" yaml:"ip"`
	Port     string `json:"port" yaml:"port"`
	Username string `json:"username" yaml:"username"`
	Password string `json:"password" yaml:"password"`
}
