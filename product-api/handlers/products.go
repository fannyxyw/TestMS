package handlers

import (
	"TestMS/product-api/data"
	"context"
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

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request)  {
	p.l.Println("Handle get method")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request)  {
	// TODO(shidf): understand
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddPoduct(&prod)
}

func (p *Products) UpdateProduct(rw http.ResponseWriter, r* http.Request)  {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest);
		return
	}

	// TODO(shidf): understand
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	err = data.UpdatePoduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusBadRequest);
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError);
		return
	}
}

type KeyProduct struct {}

func (p *Products)MiddlewarValidProduct(next http.Handler) http.Handler  {
	return http.HandlerFunc(func(rw http.ResponseWriter, r* http.Request) {
			prod := data.Product{}
			err := prod.FromJSON(r.Body)
			if err != nil {
				http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest);
				return
			}

			// TODO(shidf): understand
			ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)
	})
}
