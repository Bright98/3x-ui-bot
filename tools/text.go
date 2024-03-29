package tools

// texts
const (
	HelloMessage         = "خوش اومدین 👋\n برای دریافت باقیمانده حجم و مدت زمان استفاده از vpn، کانفیگ خود را وارد کنین:\n\n"
	RetryCalcClientUsage = "دریافت مجدد"
	ClientEmail          = "👤 نام: "
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
	CantGetImage        = "مشکل در دریافت تصویر"
	CantDecodeImage     = "مشکل در خواندن بارکد"
)

// variables
const (
	VLESS = "vless"
	VMESS = "vmess"
)

// .env
const (
	BotToken = "BotToken"
)

// errors
const (
	AuthErr                  = "AUTH_ERR"
	LoginFailedErr           = "LOGIN_FAILED"
	CantConnectErr           = "CANT_CONNECT_ERR"
	InvalidationErr          = "INVALIDATION_ERR"
	ProtocolNotFoundErr      = "PROTOCOL_NOT_FOUND"
	UserNotFoundErr          = "USER_NOT_FOUND"
	MessageIsNotImageTypeErr = "MESSAGE_IS_NOT_IMAGE_TYPE"
	CantGetImageErr          = "CANT_GET_IMAGE"
	CantDecodeImageErr       = "CANT_DECODE_IMAGE"
)
