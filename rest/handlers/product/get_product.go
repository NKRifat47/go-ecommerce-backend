package product

import (
	"ecommerce/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {

		util.SendError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	product, err := h.svc.Get(pId)
	if err != nil {
		util.SendError(w, http.StatusInternalServerError, "Internal server error.")
	}

	if product == nil {
		util.SendError(w, http.StatusNotFound, "Product Not Found")
		return
	}

	util.SendData(w, http.StatusOK, product)
}
