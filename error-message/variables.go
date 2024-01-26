package errormessage

var ERR_INVALID_BODY_PARAMETERS ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_BODY_PARAMETERS",
	Default: "Invalid body parameters",
	Language: Language{
		TR: "Geçersiz veri parametreleri",
	},
}

var ERR_INVALID_QUERY_PARAMETERS ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_QUERY_PARAMETERS",
	Default: "Invalid query parameters",
	Language: Language{
		TR: "Geçersiz sorgu parametreleri",
	},
}

var ERR_INVALID_FORM_PARAMETERS ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_FORM_PARAMETERS",
	Default: "Invalid form parameters",
	Language: Language{
		TR: "Geçersiz form parametreleri",
	},
}

var ERR_INVALID_BEARER_AUTH_HEADER ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_BEARER_AUTH_HEADER",
	Default: "Invalid bearer token authorization header",
	Language: Language{
		TR: "Geçersiz bearer token formatı",
	},
}

var ERR_INVALID_BEARER_AUTH_HEADER_FORMAT ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_BEARER_AUTH_HEADER_FORMAT",
	Default: "Invalid bearer token format",
	Language: Language{
		TR: "Geçersiz bearer token formatı",
	},
}

var ERR_INVALID_BASIC_AUTH_HEADER ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_BASIC_AUTH_HEADER",
	Default: "Invalid basic auth authorization header",
	Language: Language{
		TR: "Geçersiz basic auth parametresi",
	},
}

var ERR_INVALID_BASIC_AUTH_HEADER_FORMAT ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_BASIC_AUTH_HEADER_FORMAT",
	Default: "Invalid basic auth format",
	Language: Language{
		TR: "Geçersiz basic auth formatı",
	},
}

var ERR_INVALID_CUSTOM_HEADER ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_CUSTOM_HEADER",
	Default: "Invalid custom header",
	Language: Language{
		TR: "Geçersiz header parametresi",
	},
}

var ERR_INVALID_PASSWORD ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_PASSWORD",
	Default: "Invalid password",
	Language: Language{
		TR: "Geçersiz şifre",
	},
}

