package context

import (
	"context"
	"errors"
	"net/url"
)

//KeyString is context key
type KeyString string

const (
	// FBAccessToken is url key context
	FBAccessToken KeyString = "facebook_access_token"

	// HostContextKey is url key context
	HostContextKey KeyString = "host"

	// ManagixToken is url key context
	ManagixToken KeyString = "access_token"

	// RoleContextKey is url key context
	RoleContextKey KeyString = "role"

	// URLContextKey is url key context
	URLContextKey KeyString = "url"

	// UsernameContextKey is url key context
	UsernameContextKey KeyString = "username"

	// CreatedAtKey is createed at key
	CreatedAtKey KeyString = "createdAt"
)

// SetContext is set context value
func SetContext(ctx context.Context, key KeyString, value string) context.Context {
	return context.WithValue(ctx, key, value)
}

// GetContextString is get url context
func GetContextString(ctx context.Context, key KeyString) (string, error) {
	value := ctx.Value(key)
	if value != nil {
		return value.(string), nil
	}
	return "", errors.New("no context key not found")
}

// GetHostContext is get url context
func GetHostContext(ctx context.Context) (string, error) {
	host := ctx.Value(HostContextKey)
	if host != nil {
		return host.(string), nil
	}
	return "", errors.New("host context ket not found")
}

// GetURLContext is get url context
func GetURLContext(ctx context.Context) (*url.URL, error) {
	u := ctx.Value(URLContextKey)
	if u != nil {
		return u.(*url.URL), nil

	}
	return nil, errors.New("URL context ket not found")
}
