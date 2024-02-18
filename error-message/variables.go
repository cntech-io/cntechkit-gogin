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

var ERR_INVALID_EMAIL ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_EMAIL",
	Default: "Invalid e-mail address",
	Language: Language{
		TR: "Geçersiz e-posta adresi",
	},
}

var ERR_INVALID_CHARACTERS ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_CHARACTERS",
	Default: "Invalid characters",
	Language: Language{
		TR: "Geçersiz karakterler",
	},
}

var ERR_INVALID_DATA_TYPE ErrorMessage = ErrorMessage{
	Code:    "ERR_INVALID_DATA_TYPE",
	Default: "Invalid data type",
	Language: Language{
		TR: "Geçersiz veri tipi",
	},
}

var ERR_USER_EXISTS ErrorMessage = ErrorMessage{
	Code:    "ERR_USER_EXISTS",
	Default: "User already exists",
	Language: Language{
		TR: "Kullanıcı zaten var",
	},
}

var ERR_USERNAME_EXISTS ErrorMessage = ErrorMessage{
	Code:    "ERR_USERNAME_EXISTS",
	Default: "Username already exists",
	Language: Language{
		TR: "Kullanıcı adı zaten var",
	},
}

var ERR_USER_NOT_FOUND ErrorMessage = ErrorMessage{
	Code:    "ERR_USER_NOT_FOUND",
	Default: "User not found",
	Language: Language{
		TR: "Kullanıcı bulunamadı",
	},
}

var ERR_USER_NOT_VERIFIED ErrorMessage = ErrorMessage{
	Code:    "ERR_USER_NOT_VERIFIED",
	Default: "User not verified",
	Language: Language{
		TR: "Kullanıcı onay bekleniyor",
	},
}

var ERR_USER_NOT_AUTHENTICATED ErrorMessage = ErrorMessage{
	Code:    "ERR_USER_NOT_AUTHENTICATED",
	Default: "User not authenticated",
	Language: Language{
		TR: "Kullanıcı kimliği doğrulanamadı",
	},
}

var ERR_MISSING_FIELDS = ErrorMessage{
	Code:    "ERR_MISSING_FIELDS",
	Default: "Missing fields",
	Language: Language{
		TR: "Gerekli alanlar eksik",
	},
}

var ERR_TOO_MANY_REQUEST = ErrorMessage{
	Code:    "ERR_TOO_MANY_REQUEST",
	Default: "Too many requests",
	Language: Language{
		TR: "İstek sayısı limitine ulaşıldı",
	},
}

var ERR_UNKNOWN = ErrorMessage{
	Code:    "ERR_UNKNOWN",
	Default: "Unknown error",
	Language: Language{
		TR: "Bilinmeyen hata",
	},
}
