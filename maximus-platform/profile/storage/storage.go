package storage

type Storage struct {
	UserStorage
	UserSettingsStorage
	UserContactStorage
}

type UUID string

func PtrUID(uuid UUID) *UUID {
	return &uuid
}

func PtrS(s string) *string {
	return &s
}

func PtrB(b bool) *bool {
	return &b
}
