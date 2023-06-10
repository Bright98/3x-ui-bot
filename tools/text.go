package tools

// texts
const (
	HelloMessage         = "خوش اومدین 👋\n برای دریافت باقیمانده حجم و مدت زمان استفاده از vpn، کانفیگ خود را وارد کنین:\n\n"
	RetryCalcClientUsage = "دریافت مجدد"
	ClientEmail          = "📬 ایمیل: "
	AllowedClientUsage   = "✅ حجم مجاز مصرف: "
	UploadClientUsage    = "⬆️ مقدار آپلود شده: "
	DownloadClientUsage  = "⬇️ مقدار دانلود شده: "
	TotalClientUsage     = "⚠️ مجموع مصرف: "
	ConfigExpireTime     = "⏳ زمان انقضا: "

	KiloByte  = "کیلوبایت"
	MegaByte  = "مگابایت"
	GigaByte  = "گیگابایت"
	TeraByte  = "ترابایت"
	PetaByte  = "پتابایت"
	Unlimited = "♾ بی نهایت"

	SomethingGetWrong   = "مشکلی پیش آمده"
	InvalidConfig       = "کانفیگ معتبر نمی باشد"
	CantConnectToServer = "مشکل در اتصال به سرور"
	UserNotExist        = "کاربر در این سرور نمی باشد"
)

// variables
const (
	VLESS = "vless"
	VMESS = "vmess"
)

// .env variables
const (
	PanelPassword = "PANEL_PASSWORD"
	PanelUsername = "PANEL_USERNAME"
	PanelPort     = "PANEL_PORT"
	ServerIP      = "SERVER_IP"
	BotToken      = "BOT_TOKEN"
)

// errors
const (
	AuthErr             = "AUTH_ERR"
	LoginFailedErr      = "LOGIN_FAILED"
	CantConnectErr      = "CANT_CONNECT_ERR"
	InvalidationErr     = "INVALIDATION_ERR"
	ProtocolNotFoundErr = "PROTOCOL_NOT_FOUND"
	UserNotFoundErr     = "USER_NOT_FOUND"
)
