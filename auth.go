package main

import (
  "net/http"
)

type authHandler struct {
  next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
  if _, err := r.Cookie("auth"); err == http.ErrNoCookie{
    //Not authenticated
    w.Header().Set("Location", "/login")
    w.WriteHeader(http.StatusTemporaryRedirect)
  } else if err != nil {
    //some other Error
    panic(err.Error())
  } else {
    //Success!  call the next Handler
    h.next.ServeHTTP(w,r)
  }
}

func MustAuth(handler http.Handler) http.Handler {
  return &authHandler{next: handler}
}
