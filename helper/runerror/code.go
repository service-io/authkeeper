// Package runerror
// @author tabuyos
// @since 2023/8/2
// @description code
package runerror

// CODE FORMAT:
// | 服务级错误码 | 应用级错误码 | 模块级错误码 | 具体错误码 |
// 服务级错误码：1 位数进行表示，比如 1 为系统级错误；2 为普通错误，通常是由用户非法操作引起。
// 应用级错误码：2 位数进行表示，比如 01 为ERP应用；02 为车载应用。
// 模块级错误码：2 位数进行表示，比如 01 为用户模块；02 为订单模块。
// 具体的错误码：3 位数进行表示，比如 001 为手机号不合法；002 为验证码输入错误。

const (
	dft                  = 20000000
	DftError             = dft + ServerError
	DftParseError        = dft + ParseError
	DftNumberFmtError    = dft + NumberFmtError
	DftLenError          = dft + LenError
	DftNoKError          = dft + NoKError
	DftNoVError          = dft + NoVError
	DftNoPermissionError = dft + NoPermissionError
)

const (
	OK     = 0
	Failed = DftError
)

const (
	svcMin = iota
	SvcSys
	SvcUsr
	svcMax
)

const (
	AppNil = iota
	AppErp
	AppCar
	appMax
)

const (
	ModNil = iota
	ModAuth
	ModPoint
	ModOrder
	ModCache
	ModCommon
	ModPlatform
	ModTenant
	ModUser
	ModSystem
	modMax
)

const (
	minError = iota
	ServerError
	ParseError
	NumberFmtError
	LenError
	NoKError
	NoVError
	NotLoginError
	NotFoundUserError
	NotFoundTokenError
	NoPermissionError
	ConvertError
	LoginFailedError
	RegisterFailedError
	NotFoundFieldError
	LogoutFailedError
	RefreshTokenFailedError
	RefreshACLFailedError
	TooManyRequests
	ParamBindError
	TokenStoreError
	IllegalTokenError
	AccountExistError
	AccountNotExistError
	DiscardFailedError
	NotMatchError
	NotFoundAccountError
	NotFoundPlatError
	NotFoundTenantError
	TokenExpiredError
	ValidFailedError
	ValidSuccessfulError
	StatusIncorrectError
	PutObjectFailedError
	GetObjectFailedError
	StoreToDBFailedError
	maxError
)

func Get(s, a, m, c int) int {
	var code = minError
	if s < svcMin || s > minInt(9, svcMax) {
		return DftNumberFmtError
	}
	code = s
	if a < AppNil || a > minInt(99, appMax) {
		return DftNumberFmtError
	}
	code = code*100 + a
	if m < ModNil || m > minInt(99, modMax) {
		return DftNumberFmtError
	}
	code = code*100 + m
	if c < minError || c > minInt(999, maxError) {
		return DftNumberFmtError
	}
	code = code*1000 + c
	return code
}

func GetUsrErp(m, c int) int {
	return Get(SvcUsr, AppErp, m, c)
}

func GetSysErp(m, c int) int {
	return Get(SvcSys, AppErp, m, c)
}

func GetCode(sg ...int) int {
	if len(sg) == 0 {
		return DftError
	}
	if len(sg) > 4 {
		return DftLenError
	}
	if len(sg) == 1 {
		ec := sg[0]
		if ec < minError || ec > minInt(999, maxError) {
			return DftLenError
		}
		return dft + ec
	}
	if len(sg) == 4 {
		return Get(sg[0], sg[1], sg[2], sg[3])
	}
	return DftParseError
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
