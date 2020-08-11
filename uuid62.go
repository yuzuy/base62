package uuid62

import (
	"github.com/eknkc/basex"
	"github.com/google/uuid"
)

var encoding *basex.Encoding

func init() {
	encoding, _ = basex.NewEncoding("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
}

func New() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return Encode(uid), nil
}

func MustNew() string {
	return Encode(uuid.New())
}

func NewV1() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return Encode(uid), nil
}

func MustNewV1() string {
	uid := uuid.Must(uuid.NewUUID())
	return Encode(uid)
}

func Encode(v uuid.UUID) string {
	return encoding.Encode([]byte(v.String()))
}

func Decode(v string) (uuid.UUID, error) {
	uid, err := encoding.Decode(v)
	if err != nil {
		return uuid.UUID{}, err
	}
	return uuid.ParseBytes(uid)
}
