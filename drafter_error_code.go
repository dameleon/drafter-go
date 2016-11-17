package main

type DrafterErrorCode int

const (
	// https://github.com/apiaryio/drafter/blob/491d2644b63e516c2e82719c570980b34384ee96/src/drafter.cc#L30-L39
	// https://github.com/apiaryio/drafter/blob/491d2644b63e516c2e82719c570980b34384ee96/src/drafter_private.cc#L29
	DRAFTER_ERROR DrafterErrorCode = -1
	// https://github.com/apiaryio/snowcrash/blob/fc0c005c5ba3e0098e54571217a40cc790034a0b/src/SourceAnnotation.h#L97-L106
	DRAFTER_NO_ERROR = 0
	DRAFTER_APPLICATION_ERROR = 1
	DRAFTER_BUSINESS_ERROR = 2
	DRAFTER_MODEL_ERROR = 3
	DRAFTER_MSON_ERROR = 4
)

func (ec DrafterErrorCode) String() string {
	switch ec {
	case DRAFTER_ERROR:
		return "drafter_error"
	case DRAFTER_NO_ERROR:
		return "drafter_no_error"
	case DRAFTER_APPLICATION_ERROR:
		return "drafter_application_error"
	case DRAFTER_BUSINESS_ERROR:
		return "drafter_business_error"
	case DRAFTER_MODEL_ERROR:
		return "drafter_model_error"
	case DRAFTER_MSON_ERROR:
		return "drafter_mson_error"
	default:
		return "unknown"
	}
}

func newDrafterErrorCode(code int) DrafterErrorCode {
	return DrafterErrorCode(code)
}