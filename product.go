package go_wix_store

import (
	"encoding/json"
	"fmt"
	. "github.com/marinekev88/go-wix-store/types"
)

type ProductService interface {
	List(key string, query string) (products Products, err error)
	ListAll() (products Products, err error)
	Get(id string) (Product, error)
}

type ProductServiceOp struct {
	client *Client
}

var _ ProductService = &ProductServiceOp{}

func (p ProductServiceOp) List(key string, query string) (products Products, err error) {
	res, _ := p.client.handler.R().
		SetBody(fmt.Sprintf(`{"query":{"filter":"{\"%s\": \"%s\"}"}}`, key, query)).
		Post("/query")

	if err = json.Unmarshal(res.Body(), &products); err != nil {
		return
	}

	return
}

func (p ProductServiceOp) ListAll() (products Products, err error) {
	res, _ := p.client.handler.R().
		Post("/query")

	if err = json.Unmarshal(res.Body(), &products); err != nil {
		return
	}

	return
}

func (p ProductServiceOp) Get(id string) (product Product, err error) {
	res, _ := p.client.handler.R().
		Get(fmt.Sprintf("/%s", id))

	if err = json.Unmarshal(res.Body(), &product); err != nil {
		return
	}

	return
}
