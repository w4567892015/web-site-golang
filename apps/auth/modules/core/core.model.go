package core

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/storage"
)

var (
	OAuth2 fosite.OAuth2Provider
)

func init() {
	store := storage.NewMemoryStore()

	store.Clients["my-client"] = &fosite.DefaultClient{
		ID:            "my-client",
		Secret:        []byte("my-client-secret"),
		RedirectURIs:  []string{"http://localhost:8080/callback"}, // The redirect URI
		GrantTypes:    []string{"authorization_code", "refresh_token"},
		ResponseTypes: []string{"code"},
		Scopes:        []string{"openid", "profile"},
	}

	secret := []byte("my super secret signing password")

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("could not generate rsa key: %s", err)
	}

	config := &fosite.Config{
		AccessTokenLifespan:   time.Minute * 30,
		RefreshTokenLifespan:  time.Hour * 24 * 7,
		AuthorizeCodeLifespan: time.Minute * 5,
		GlobalSecret:          secret,
	}

	OAuth2 = compose.ComposeAllEnabled(config, store, privateKey)

	log.Printf("Successfully composed Fosite provider !")
}
