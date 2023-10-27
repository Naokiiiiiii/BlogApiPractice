package apperrors

type ErrCode string

const (
	Unknown ErrCode = "U0000"

	InsertDataFailed ErrCode = "S001"
	GetDataFailed    ErrCode = "S002"
	NAData           ErrCode = "S003"
	NoTargetData     ErrCode = "S004"
	UpdateDataFailed ErrCode = "S005"
	DeleteDataFailed ErrCode = "S006"

	ReqBodyDecodeFailed ErrCode = "R001"
	BadParam            ErrCode = "R002"

	RequiredAuthorizationHeader ErrCode = "A001"
	CannotMakeValidatior        ErrCode = "A001"
	Unauthorizated              ErrCode = "A003"
	NotMatchUser                ErrCode = "A004"
	GetUserInfoFailed           ErrCode = "A005"
	ExchangeTokenFailed         ErrCode = "A006"
	DecodeUserInfoFailed        ErrCode = "A007"
	ExchangeRefreshTokenFailed  ErrCode = "A008"
	GetIDTokenFailed            ErrCode = "A009"
	ParsePayloadFailed          ErrCode = "A010"
)
