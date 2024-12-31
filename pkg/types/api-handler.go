package pkg_types

import "net/http"

type ApiHttpHandler func(w http.ResponseWriter, r *http.Request) error
