package render

import (
	"bnb-booking/internal/config"
	"bnb-booking/internal/models"
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	// what to store in the session
	gob.Register(models.Reservation{})

	// change to true when in production
	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = testApp.InProduction

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}

func (w *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (w *myWriter) Write(b []byte) (int, error) {
	length := len(b)
	return length, nil
}

func (w *myWriter) WriteHeader(statusCode int) {

}
