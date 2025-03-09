package port

import (
	"context"
	"mime/multipart"
	"time"
)

type (
	IHelper interface {
		GenUUID() string
		GenPromptPayQrCodeString(phoneno string, amount float64) string
		ConvertStructToStrJson(v interface{}) (string, error)
		ConvertJsonToStruct(jsonStr string, result interface{}) error
		ConvertStrToFloat64(value string) (float64, error)
	}

	IStorageService interface {
		UploadImage(image *multipart.FileHeader, path string) error
	}

	ICache interface {
		Get(ctx context.Context, key string) (string, error)
		Cache(ctx context.Context, key string, data string, expireTime time.Time) (bool, error)
		Delete(ctx context.Context, key string) error
		// ExcuteTransaction(ctx context.Context , f func() error) error
		CacheIgnoreDuplcateKey(ctx context.Context, key string, data string, expireTime time.Time) error
	}

	IMessageQueue interface {
		Produce(topic, message string) error
		Consumer(topic string) error
	}
)
