package util

import (
	"net/http"
	"time"

	"dumpapp_server/pkg/config"
	errors2 "dumpapp_server/pkg/errors"
	"github.com/gorilla/securecookie"
)

var (
	hashKey  = []byte(config.DumpConfig.AppConfig.Cookie.HMACKey)
	blockKey = []byte(config.DumpConfig.AppConfig.Cookie.AESKey)
	s        = securecookie.New(hashKey, blockKey)
)

func SetCookie(w http.ResponseWriter, name string, value map[string]string) {
	if encoded, err := s.Encode(name, value); err == nil {
		cookie := &http.Cookie{
			Name:     name,
			Value:    encoded,
			Path:     "/",
			Secure:   false,
			HttpOnly: true,
			MaxAge:   60 * 60 * 24 * 30,
			SameSite: http.SameSiteNoneMode,
		}
		http.SetCookie(w, cookie)
	}
}

func GetCookie(r *http.Request, name string) map[string]string {
	value := make(map[string]string)
	if cookie, err := r.Cookie(name); err == nil {
		if err = s.Decode(name, cookie.Value, &value); err != nil {
			panic(errors2.ErrNotAuthorized)
		}
	}
	return value
}

func ClearCookie(w http.ResponseWriter, name string) {
	cookie := &http.Cookie{
		Name:     name,
		MaxAge:   0,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
		HttpOnly: true,
	}
	http.SetCookie(w, cookie)
}
