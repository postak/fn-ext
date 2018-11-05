package proxima_auth

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"

	"net/http"
	"os"
)

var fnToken = ""
var fnUsername = ""
var fnPassword = ""

func init() {
	server.RegisterExtension(&proximaPlugin{})
	fnToken = os.Getenv("FN_TOKEN")
	fnUsername = os.Getenv("FN_USERNAME")
	fnPassword = os.Getenv("FN_PASSWORD")

	fmt.Println("fnToken=" + fnToken)
	fmt.Println("fnUsername=" + fnUsername)
	fmt.Println("fnPassword=" + fnPassword)

}

type proximaPlugin struct {
}

func (e *proximaPlugin) Name() string {
	return "github.com/postak/fn-ext/proxima_auth"
}

func (e *proximaPlugin) Setup(s fnext.ExtServer) error {
	s.AddCallListener(&Proxima{})
	s.AddRootMiddleware(&ProximaCheckHeader{})

	return nil
}

type Proxima struct {
}

func (l *Proxima) BeforeCall(ctx context.Context, call *models.Call) error {
	fmt.Println("Proxima BeforeCall.")
	return nil
}

func (l *Proxima) AfterCall(ctx context.Context, call *models.Call) error {
	fmt.Println("Proxima AfterCall.")
	return nil
}

// struct ProximaCheckHeader
type ProximaCheckHeader struct{}

func (h *ProximaCheckHeader) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("Proxima Middleware called")
		// fmt.Println("fnToken=", fnToken)
		// fmt.Println("URL path=", r.URL.Path)

		// if r.URL.Path == "/version" {
		// 	fmt.Println("skip authentication for this URL")
		// 	next.ServeHTTP(w, r)
		// 	return
		// }

		// // check which auth header to use

		// var authHeader = ""

		// if r.Header.Get("Authorization") != "" {
		// 	authHeader = r.Header.Get("Authorization")
		// 	fmt.Printf("use header Authorization=%s\n", authHeader)
		// } else if r.Header.Get("fn_authorization") != "" {
		// 	// some service filter Authorization header so accept
		// 	// also custom header fn_authorization
		// 	fmt.Printf("use header fn_authorization=%s\n", authHeader)
		// } else {
		// 	fmt.Printf("authorization failed")
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	m2 := map[string]string{"message": "Invalid token use either Authorization or fn_authorization."}
		// 	m := map[string]map[string]string{"error": m2}
		// 	json.NewEncoder(w).Encode(m)
		// 	fmt.Println("auth NOT succeeded!")
		// 	return
		// }

		// tokenFromHeader := strings.SplitN(authHeader, " ", 3)

		// // check auth method

		// if tokenFromHeader[0] == "Basic" {
		// 	fmt.Printf("Check Basic auth\n")
		// 	payload, _ := base64.StdEncoding.DecodeString(tokenFromHeader[1])

		// 	pair := strings.SplitN(string(payload), ":", 2)
		// 	if len(pair) != 2 || pair[0] != fnUsername || pair[1] != fnPassword || len(fnUsername) == 0 || len(fnPassword) == 0 {
		// 		fmt.Printf("authorization failed")
		// 		w.Header().Set("Content-Type", "application/json")
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		m2 := map[string]string{"message": "Invalid username/password."}
		// 		m := map[string]map[string]string{"error": m2}
		// 		json.NewEncoder(w).Encode(m)
		// 		fmt.Println("auth NOT succeeded!")
		// 		return
		// 	}
		// 	fmt.Printf("authorization ok")
		// 	//
		// } else if tokenFromHeader[0] == "Bearer" {
		// 	fmt.Printf("Check Bearer token %s\n", tokenFromHeader)
		// 	if len(tokenFromHeader) < 2 || tokenFromHeader[1] != fnToken || len(fnToken) == 0 {
		// 		w.Header().Set("Content-Type", "application/json")
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		m2 := map[string]string{"message": "Invalid Authorization token."}
		// 		m := map[string]map[string]string{"error": m2}
		// 		json.NewEncoder(w).Encode(m)
		// 		fmt.Println("auth NOT succeeded!")
		// 		return
		// 	}
		// } else {
		// 	w.Header().Set("Content-Type", "application/json")
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	m2 := map[string]string{"message": "Invalid Authorization method. Use either Basic or Bearer."}
		// 	m := map[string]map[string]string{"error": m2}
		// 	json.NewEncoder(w).Encode(m)
		// 	fmt.Println("auth NOT succeeded!")
		// 	return
		// }

		fmt.Println("auth succeeded!")
		r = r.WithContext(context.WithValue(r.Context(), contextKey("user"), "I'm in!"))
		next.ServeHTTP(w, r)
	})
}

type contextKey string
