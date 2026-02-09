package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

func (h *WalletHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	log := zerolog.New(os.Stderr)

	id, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error().Err(errors.New("id empty")).Stack().Send()

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	balance, err := h.uc.GetBalance(r.Context(), id)
	if err != nil {
		log.Error().Err(err).Stack().Msg("failed to get balance")

		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]int64{
		"balance": balance,
	})
}
