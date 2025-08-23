package auth

import (
	"net/http"

	"github.com/ahmaruff/eav-platform/internal/shared"
	"github.com/alexedwards/scs/sqlite3store"
	"github.com/alexedwards/scs/v2"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	sessionManager *scs.SessionManager
}

func NewService(db *sqlx.DB) *Service {
	config := shared.LoadConfig()

	sessionManager := scs.New()
	sessionManager.Store = sqlite3store.New(db.DB)
	// configure session settings
	sessionManager.Lifetime = config.GetSessionLifetime()
	sessionManager.Cookie.Name = config.Session.Name
	sessionManager.Cookie.HttpOnly = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = config.Session.Secure // true for HTTPS

	return &Service{
		sessionManager: sessionManager,
	}
}

func (s *Service) CreateSession(w http.ResponseWriter, r *http.Request, userID string) {
	s.sessionManager.Put(r.Context(), "user_id", userID)
}

func (s *Service) DestroySession(w http.ResponseWriter, r *http.Request) {
	s.sessionManager.Destroy(r.Context())
}

func (s *Service) GetUserID(r *http.Request) string {
	return s.sessionManager.GetString(r.Context(), "user_id")
}
