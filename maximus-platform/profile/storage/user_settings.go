package storage

type UserSettingsStorage interface {
	StoreUserSettings(in StoreUserSettings) (UserSettings, error)
	UpdateUserSettings(userID UUID, in UpdateUserSettings) (UserSettings, error)
	GetUserSettings(in GetUserSettings) (UserSettings, error)
}

type UserSettings struct {
	UserID UUID `json:"userID" bson:"userID"`

	TwoFAChannelType string `json:"2FAChannelType" bson:"2FAChannelType"`
	Locale           string `json:"locale" bson:"locale"`
}

type StoreUserSettings struct {
	UserID UUID `json:"userID" bson:"userID"`

	TwoFAChannelType string `json:"2FAChannelType" bson:"2FAChannelType"`
	Locale           string `json:"locale" bson:"locale"`
}

type UpdateUserSettings struct {
	TwoFAChannelType *string
	Locale           *string
}

type GetUserSettings struct {
	UserID UUID
}
