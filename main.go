package main

import (
	"compress/gzip"
	"io"
	"log"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/RomuloDurante/baseNewServer/config"
	"github.com/RomuloDurante/baseNewServer/controller"
)

/*********************************/
/********** GZIP ****************/

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

//GzipHandlerFunc ....
func GzipHandlerFunc(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn(gzr, r)
	}
}

//GzipHandler ...
func GzipHandler(fn http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			fn.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		fn.ServeHTTP(gzr, r)
	}
}

//**          END GZIP           **//
//********************************//

//################################//
//             MAIN              //
//###############################//
func main() {
	//runtime
	runtime.GOMAXPROCS(4)
	//server
	mux := http.NewServeMux()
	mux.HandleFunc("/", GzipHandlerFunc(controller.HandleController))
	// mux.HandleFunc("/api/", GzipHandlerFunc(api.HandleAPI))
	// mux.HandleFunc("/upload", api.HandleFile)
	// mux.Handle("/css/", GzipHandler(http.FileServer(http.Dir("view/public"))))
	// mux.Handle("/script/", GzipHandler(http.FileServer(http.Dir("view/public"))))
	// mux.Handle("/img/", http.FileServer(http.Dir("view/public")))
	// mux.Handle("/video/", http.FileServer(http.Dir("view/public")))
	// mux.Handle("/favicon.ico", http.FileServer(http.Dir("view/public/img/icons")))
	//service worker for progressive web application
	//mux.Handle("/sw.js", GzipHandler(http.FileServer(http.Dir("view/"))))

	//config server securit
	// cfg := &tls.Config{
	// 	MinVersion:               tls.VersionTLS12,
	// 	CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 	PreferServerCipherSuites: true,
	// 	CipherSuites: []uint16{
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
	// 		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
	// 	},
	// }
	//the server
	server := &http.Server{
		Addr:              config.PortToDeploy,
		Handler:           mux,
		ReadHeaderTimeout: 20 * time.Second,
		ReadTimeout:       1 * time.Minute,
		WriteTimeout:      2 * time.Minute,
		//TLSConfig:         cfg,
		//TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}
	// go log.Fatal(server.ListenAndServeTLS("/etc/letsencrypt/live/www.bunkerlibrary.com/fullchain.pem", "/etc/letsencrypt/live/www.bunkerlibrary.com/privkey.pem"))

	go log.Fatal(server.ListenAndServe())
}
