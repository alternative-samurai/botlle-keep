package main

import (
	"context"
	"example/src/controllers"
	"example/src/infrastructures/generated"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

// const a =

type contextKey struct {
	reqDate string
}

func main() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgresp@localhost:5432/bottle"))
	if err != nil {
		panic(err)
	}
	if db == nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &controllers.Resolver{}}))

	// ここでコンテキストを追加できる!!
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		// ここでgraphqlのコンテキストを取得している!!
		// oc := graphql.GetOperationContext(ctx)
		// リクエスト時刻を取得!!
		ctxs := context.WithValue(ctx, "reqDate", time.Now())
		// fmt.Printf("around: %s %s", oc.OperationName, oc.RawQuery)
		return next(ctxs)
	})
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", baseHandler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func baseHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// common

		// bufbody := new(bytes.Buffer)
		// defer r.Body.Close()
		// bufbody.ReadFrom(r.Body)
		// body := bufbody.String()
		// e := r.ParseForm()
		// log.Println(e)
		// defer r.Body.Close()
		// body, err := ioutil.ReadAll(r.Body)
		// if err != nil {
		// 	panic(err)
		// }
		// body, _ := ioutil.ReadAll(r.Body)
		// bodyString := string(body)
		log.Println(r.URL, r.Method)
		// log.Println(bodyString)
		// log.Println(r.Form.Get("query"), "aaa")
		handler.ServeHTTP(w, r)
	})
}
