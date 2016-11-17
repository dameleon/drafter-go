package drafter

type DrafterErrorCode int

const (
	// https://github.com/apiaryio/drafter/blob/491d2644b63e516c2e82719c570980b34384ee96/src/drafter.cc#L30-L39
	// https://github.com/apiaryio/drafter/blob/491d2644b63e516c2e82719c570980b34384ee96/src/drafter_private.cc#L29
	DRAFTER_ERROR DrafterErrorCode = -1
	// https://github.com/apiaryio/snowcrash/blob/fc0c005c5ba3e0098e54571217a40cc790034a0b/src/SourceAnnotation.h#L97-L106
	DRAFTER_NO_ERROR
	DRAFTER_APPLICATION_ERROR
	DRAFTER_BUSINESS_ERROR
	DRAFTER_MODEL_ERROR
	DRAFTER_MSON_ERROR
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
	switch code {
	case DRAFTER_ERROR:
		return DRAFTER_ERROR
	case DRAFTER_NO_ERROR:
		return DRAFTER_NO_ERROR
	case DRAFTER_APPLICATION_ERROR:
		return DRAFTER_APPLICATION_ERROR
	case DRAFTER_BUSINESS_ERROR:
		return DRAFTER_BUSINESS_ERROR
	case DRAFTER_MODEL_ERROR:
		return DRAFTER_MODEL_ERROR
	case DRAFTER_MSON_ERROR:
		return DRAFTER_MSON_ERROR
	default:
		panic("undefined error")
	}
}