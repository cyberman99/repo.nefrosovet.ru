package utils

import (
    _default "repo.nefrosovet.ru/maximus-platform/mailer/storage/default"
)

func GetStorage() *_default.Storage {
    return _default.GetStorage()
}
