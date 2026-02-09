package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"wallet/internal/usecase"
)

type WalletHandler struct {
	uc *usecase.WalletUsecase
}

func NewWalletHandler(uc *usecase.WalletUsecase) *WalletHandler {
	return &WalletHandler{uc: uc}
}

func (h *WalletHandler) MakeMuxRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallet", h.Operate).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/wallets/{id}", h.GetBalance).Methods(http.MethodGet)

	return r
}
