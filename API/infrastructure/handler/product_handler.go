package handler

import (
	"ApiShortLong/domain/entities"
	"ApiShortLong/domain/repo"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ProductHandler struct {
	service repo.ProductService
}

func NewProductHandler(service repo.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) SetupRoutes(router *mux.Router) {
	router.HandleFunc("/addProduct", h.addProduct).Methods("POST")
	router.HandleFunc("/isNewProductAdded", h.isNewProductAdded).Methods("GET")
	router.HandleFunc("/CountProductIsInDiscount", h.countProductsInDiscount).Methods("GET")
}

func (h *ProductHandler) addProduct(w http.ResponseWriter, r *http.Request) {
	var product entities.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.AddProduct(product); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) isNewProductAdded(w http.ResponseWriter, r *http.Request) {
    // Timeout más corto para long polling (2 segundos)
    timeout := time.After(2 * time.Second)
    ticker := time.NewTicker(500 * time.Millisecond) // Verificación cada 500ms
    defer ticker.Stop()

    // Obtener el último producto (limit = 1)
    lastProducts, err := h.service.GetLastAddedProducts(1)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var lastProduct *entities.Product
    if len(lastProducts) > 0 {
        lastProduct = &lastProducts[0]
    }

    for {
        select {
        case <-timeout:
            json.NewEncoder(w).Encode(map[string]interface{}{
                "hasNewProduct": false,
                "product":       nil,
            })
            return
        case <-ticker.C:
            currentProducts, err := h.service.GetLastAddedProducts(1)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            var currentProduct *entities.Product
            if len(currentProducts) > 0 {
                currentProduct = &currentProducts[0]
            }

            // Comparar IDs para detectar cambios reales
            if currentProduct != nil && (lastProduct == nil || currentProduct.ID != lastProduct.ID) {
                json.NewEncoder(w).Encode(map[string]interface{}{
                    "hasNewProduct": true,
                    "product":       currentProduct,
                })
                return
            }
        }
    }
}

func (h *ProductHandler) countProductsInDiscount(w http.ResponseWriter, r *http.Request) {
	// Long polling implementation
	timeout := time.After(30 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	initialCount, err := h.service.CountProductsInDiscount()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for {
		select {
		case <-timeout:
			json.NewEncoder(w).Encode(map[string]int{"count": initialCount})
			return
		case <-ticker.C:
			currentCount, err := h.service.CountProductsInDiscount()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if currentCount != initialCount {
				json.NewEncoder(w).Encode(map[string]int{"count": currentCount})
				return
			}
		}
	}
}