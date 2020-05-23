package handlers

import (
	"TestMS/product-api/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products  {
	return  &Products{l}
}
/*
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}  else if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}else if r.Method == http.MethodPut {
		rege := regexp.MustCompile("/([0-9]+)")
		g := rege.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			http.Error(rw, "Invalid url", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			http.Error(rw, "Invalid url", http.StatusBadRequest)
			return
		}
		idSting := g[0][1]

		id , err := strconv.Atoi(idSting)
		if err != nil {
			http.Error(rw, "Invalid url", http.StatusBadRequest)
			return
		}

		p.updateProduct(id, rw, r)
	}
}
 */

func (p *Products)GetProducts(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Handle get method")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Handle post method")

	prod := &data.Product{}
	e := prod.FromJSON(r.Body)
	if e != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}
	data.AddPoduct(prod)
}

func (p *Products)UpdateProduct(rw http.ResponseWriter, r* http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest);
		return
	}

	prod := &data.Product{}
	err = prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest);
		return
	}
	err = data.UpdatePoduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusBadRequest);
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError);
		return
	}
}
