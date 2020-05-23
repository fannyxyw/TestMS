package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32  `json:"price"`
	SKU string `json:"sku"`
	CreatedOn string `json:"-"`
	UpdateOn string `json:"-"`
	DeleteOn string `json:"-"`
}

// Products is a collections of product
type Products []*Product

func (p *Product) FromJSON(r io.Reader) error  {
	de := json.NewDecoder(r)
	return  de.Decode(p)
}

func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func GetProducts() Products  {
	return prodctList
}

func AddPoduct(p *Product)  {
	p.ID = getNextID()
	prodctList = append(prodctList, p)
}

var  ErrProductNotFound  = fmt.Errorf("Product not found")

func findProduct(id int)(*Product, int, error)  {
	for i, p := range prodctList{
		if p.ID == id {
			return  p, i, nil
		}
	}

	return  nil, -1, ErrProductNotFound
}

func UpdatePoduct(id int, p *Product) error {
	_, pos, err := findProduct (id)
	if err != nil {
		return err
	}
	p.ID = id
	prodctList[pos] = p

	return nil
}

func getNextID() int  {
	lp := prodctList[len(prodctList) - 1]
	return lp.ID + 1
}

var prodctList  = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Front milky coffee",
		Price:       2.45,
		SKU:         "abc123",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeleteOn:    "",
	},
	&Product{
		ID:          2,
		Name:        "Espressp",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "fds22",
		CreatedOn:   time.Now().UTC().String(),
		UpdateOn:    time.Now().UTC().String(),
		DeleteOn:    "",
	},
}