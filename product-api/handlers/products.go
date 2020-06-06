// Package classification of Product API
//
// Documentation for Product API
//
// Schemes: http
// BasePath: /
// Version: 1.0.1
//
// Consumes:
// - application

package handlers

import (
	protos "TestMS/CURRENCY/protos/currency"
	"TestMS/product-api/data"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Products struct {
	l *log.Logger
	cc protos.CurrencyClient
}

func NewProducts(l *log.Logger, cc protos.CurrencyClient) *Products  {
	return  &Products{l, cc}
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request)  {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

func (p *Products) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	prod, err := data.GetProductByID(id)
	if err != nil {
		http.Error(rw, "Unable to convert id", http.StatusBadRequest);
		return
	}
	rr := &protos.RateRequest{
		Base: protos.Currencies(protos.Currencies_value["BGN"]),
		Destination: protos.Currencies(protos.Currencies_value["BGN"]),
	}

	rp, err := p.cc.GetRate(context.Background(), rr)
	if err != nil {
		http.Error(rw, "Grpc error", http.StatusInternalServerError)
		return
	}

	p.l.Printf("currency Get rate:", rp.Rate)
	prod.Price = prod.Price * rp.Rate
	e := json.NewEncoder(rw)
	e.Encode(prod)
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request)  {
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

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)
		next.ServeHTTP(rw, req)
	})
}


// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that
// this is a valid number
func getProductID(r *http.Request) int {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}