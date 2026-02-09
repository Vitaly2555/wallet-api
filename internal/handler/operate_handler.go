package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rs/zerolog"

	"wallet/internal/domain"
)

func (h *WalletHandler) Operate(w http.ResponseWriter, r *http.Request) {
	log := zerolog.New(os.Stderr)

	var req OperationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Error().Err(err).Stack().Msg("failed to decode request")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	operation, err := domain.NewOperation(req.Operation)
	if err != nil {
		log.Error().Err(err).Stack().Str("wallet_id", req.WalletID).Msg("failed to create operation")

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err = h.uc.Operate(r.Context(), req.WalletID, operation, req.Amount); err != nil {
		log.Error().Err(err).Stack().Str("wallet_id", req.WalletID).Msg("failed to operate")

		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
