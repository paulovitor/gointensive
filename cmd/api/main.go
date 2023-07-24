package main

import (
	// "encoding/json"
	// "net/http"

	// "github.com/go-chi/chi/v5"
	// "github.com/go-chi/chi/v5/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/paulovitor/gointensive/internal/entity"
)

func main() {
	// chi
	// r := chi.NewRouter()
	// r.Use(middleware.Logger)
	// r.Get("/order", Order)
	// http.ListenAndServe(":8888", r)

	// http.HandleFunc("/order", Order)
	// http.ListenAndServe(":8888", nil)

	e := echo.New()
	e.GET("/order", Order)
	e.Logger.Fatal(e.Start(":8888"))
}

// func Order(w http.ResponseWriter, r *http.Request) {
func Order(c echo.Context) error {
	order := entity.Order{
		ID:    "1",
		Price: 10,
		Tax:   1,
	}
	// order, err := entity.NewOrder("1234", 1000, 10)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(w).Encode(err.Error())
	// 	return
	// }
	err := order.CalculateFinalPrice()
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(order)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, order)
}
