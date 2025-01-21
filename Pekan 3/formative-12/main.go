package main
import(
	"net/http"
	"formative-12/web-server"
	"fmt"
	)

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	  uname, pwd, ok := r.BasicAuth()
	  if !ok {
		w.Write([]byte("Username atau Password tidak boleh kosong"))
		return
	  }
  
	  if uname == "admin" && pwd == "admin" {
		next.ServeHTTP(w, r)
		return
	 }
	  w.Write([]byte("Username atau Password tidak sesuai"))
	})
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	  }
	http.Handle("/nilai", Auth(log(http.HandlerFunc(webserver.GetNilai))))
	http.Handle("/post-nilai", Auth(log(http.HandlerFunc(webserver.PostNilai))))
	fmt.Println("server running at http://localhost:8080")
  	server.ListenAndServe()
  
}