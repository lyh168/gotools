package mapstruct

import "github.com/mitchellh/mapstructure"

type (
	DecodeHookFunc     = mapstructure.DecodeHookFunc
	DecodeHookFuncKind = mapstructure.DecodeHookFuncKind
	DecodeHookFuncType = mapstructure.DecodeHookFuncType
	Decoder            = mapstructure.Decoder
	DecoderConfig      = mapstructure.DecoderConfig
	Error              = mapstructure.Error
	Metadata           = mapstructure.Metadata
)

var (
	ComposeDecodeHookFunc = mapstructure.ComposeDecodeHookFunc
	Decode                = mapstructure.Decode
	DecodeHookExec        = mapstructure.DecodeHookExec
	DecodeMetadata        = mapstructure.DecodeMetadata
	NewDecoder            = mapstructure.NewDecoder
	WeakDecode            = mapstructure.WeakDecode
	WeakDecodeMetadata    = mapstructure.WeakDecodeMetadata
)

var (
	StringToIPHookFunc           = mapstructure.StringToIPHookFunc
	StringToIPNetHookFunc        = mapstructure.StringToIPNetHookFunc
	StringToSliceHookFunc        = mapstructure.StringToSliceHookFunc
	StringToTimeDurationHookFunc = mapstructure.StringToTimeDurationHookFunc
	StringToTimeHookFunc         = mapstructure.StringToTimeHookFunc
	WeaklyTypedHook              = mapstructure.WeaklyTypedHook
)
