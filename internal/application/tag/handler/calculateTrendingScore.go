package handler

import (
	"context"

	"go.uber.org/zap"
)

func (h *Handler) CalculateTrendingScore(ctx context.Context) error {
	h.Logger.Info("executing trending score calculation")

	tags, err := h.Service.ListActiveTag(ctx)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		err = h.Service.CalculateTrendingScore(ctx, tag.Id)
		if err != nil {
			h.Logger.Error("error occured on CalculateTrendingScore", zap.Error(err))
		}
	}

	h.Logger.Info("finished trending score calculation")
	return err
}
