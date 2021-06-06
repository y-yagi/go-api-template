package routes

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/y-yagi/go-api-template/database"
	"github.com/y-yagi/go-api-template/ent"
)

func removeTestData() {
	database.Client.Author.Delete().ExecX(context.Background())
	database.Client.Book.Delete().ExecX(context.Background())
}

func TestMain(m *testing.M) {
	err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Client.Close()

	removeTestData()

	code := m.Run()

	defer os.Exit(code)
}

func TestGetBooks(t *testing.T) {
	author := database.Client.Author.Create().SetName("Bob").SaveX(context.Background())
	database.Client.Book.Create().SetName("GoBook").AddAuthor(author).SaveX(context.Background())
	database.Client.Book.Create().SetName("JSBook").AddAuthor(author).SaveX(context.Background())

	addr := startServer(t)

	reqURL := &url.URL{Scheme: "http", Host: addr.String(), Path: "/api/books"}
	req, _ := http.NewRequest("GET", reqURL.String(), nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	var books []ent.Book
	json.Unmarshal([]byte(string(body)), &books)

	if len(books) != 2 {
		t.Fatalf("got: %d\nwont: %d", len(books), 2)
	}

	expected := "GoBook"
	actual := books[0].Name
	if actual != expected {
		t.Errorf("got: %s\nwont: %s", actual, expected)
	}
}

func startServer(t *testing.T) net.Addr {
	addr := randomAddress(t)
	app := New()
	go func() {
		app.Listen(addr.String())
	}()

	return addr
}

func randomAddress(t *testing.T) net.Addr {
	t.Helper()

	listener, err := net.Listen("tcp", ":0")
	listener.Close()

	if err != nil {
		t.Fatal(err)
	}
	return listener.Addr()
}
