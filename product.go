package go_wix_store

import "fmt"

type ProductService interface {
	List(key string, query string) ([]string, error)
	ListAll() ([]string, error)
	Get(id string) (string, error)
}

type ProductServiceOp struct {
	client *Client
}

var _ ProductService = &ProductServiceOp{}

func (p ProductServiceOp) List(key string, query string) ([]string, error) {
	res, _ := p.client.handler.R().
		SetBody(fmt.Sprintf(`{"query":{"filter":"{\"%s\": \"%s\"}"}}`, key, query)).
		Post("/query")
	fmt.Println(res.Body())

	return []string{"Hello", "Kitty"}, nil
}

func (p ProductServiceOp) ListAll() ([]string, error) {
	res, _ := p.client.handler.R().
		Post("/query")
	fmt.Println(res.Body())
	return []string{"Hello", "Kitty"}, nil
}

func (p ProductServiceOp) Get(id string) (string, error) {
	res, _ := p.client.handler.R().
		Get(fmt.Sprintf("/%s", id))
	fmt.Println(res.Body())
	return "hello", nil
}
