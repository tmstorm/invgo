// Package utils contains internal utility functions for invgo
package utils

import (
	"fmt"
	"log"
	"net/url"
	"reflect"
	"strings"
)

func StructToQuery(v any) (url.Values, error) {
	q := make(url.Values)
	if v == nil {
		return q, nil
	}
	if err := addQuery(q, v, ""); err != nil {
		return nil, err
	}
	return q, nil
}

func addQuery(q url.Values, v any, prefix string) error {
	if v == nil {
		return nil
	}

	vals := reflect.ValueOf(v)
	if vals.Kind() == reflect.Pointer {
		if vals.IsNil() {
			return nil
		}
		vals = vals.Elem()
	}

	switch vals.Kind() {
	case reflect.Struct:
		t := vals.Type()
		for i := range vals.NumField() {
			field := vals.Field(i)
			if !field.CanInterface() {
				continue
			}

			if field.Kind() == reflect.Struct {
				if err := addQuery(q, field.Interface(), prefix); err != nil {
					return err
				}
			}

			tag := t.Field(i).Tag.Get("url")
			if tag == "" || tag == "-" {
				continue
			}

			parts := strings.Split(tag, ",")
			key := strings.TrimSpace(parts[0])
			isRequired := len(parts) > 1 && strings.TrimSpace(parts[1]) == "required"

			fullKey := key
			if prefix != "" {
				fullKey = prefix + tag
			}

			if isRequired && isZero(field.Interface()) {
				return fmt.Errorf("field %s is required", key)
			}

			if err := addQuery(q, field.Interface(), fullKey); err != nil {
				return err
			}
		}

	case reflect.Slice, reflect.Array:
		if vals.Len() == 0 {
			return nil
		}

		for i := range vals.Len() {
			elem := vals.Index(i)
			key := fmt.Sprintf("%s[%d]", prefix, i)
			if err := addQuery(q, elem.Interface(), key); err != nil {
				return err
			}
		}
	default:
		if isZero(vals.Interface()) {
			return nil
		}
		q.Add(prefix, fmt.Sprint(vals.Interface()))
	}

	return nil
}

func isZero(v any) bool {
	return reflect.DeepEqual(v, reflect.Zero(reflect.TypeOf(v)).Interface())
}

// ParseURL is used to pre-parse the provided rawURL before a client is created.
// It will attempt to enforce https if allowHTTP = false. This should only be set to true
// in testing, to prevent instances where the server is not configured correctly and data is
// sent in cleartext.
func ParseURL(rawURL string, invgateAPIPath string, allowHTTP bool) (*url.URL, error) {
	base := strings.TrimSuffix(strings.TrimSpace(rawURL), "/")

	path := invgateAPIPath
	if path != "" && !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	fullURL := base + path
	u, err := url.Parse(fullURL)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "":
		u.Scheme = "https"
		log.Printf("[INVGO] No scheme in provided URL, defaulting to https://%s", u.Host)
	case "http":
		if !allowHTTP {
			oldURL := u.String()
			u.Scheme = "https"
			log.Printf("[INVGO] SECURITY: Auto-upgrading insecure provided URL from %s -> %s\n"+
				" This is to prevent accidental cleartext traffic. To disable, set AllowHTTP: true", oldURL, u)
		} else {
			log.Printf("[INVGO] WARNING: Client configured with insecure HTTP: %s\n"+
				" this could be a security risk if the destination server is not configured correctly.\n"+
				" to disable this set AllowHTTP=false or change the URL to https.", u)
		}
	}
	return u, nil
}
