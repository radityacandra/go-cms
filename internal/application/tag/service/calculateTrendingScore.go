package service

import "context"

func (s *Service) CalculateTrendingScore(ctx context.Context, tagId string) error {
	tag, err := s.Repository.FindTagById(ctx, tagId)
	if err != nil {
		return err
	}

	score := tag.CalculateTrendingScore()
	tag.UpdateTrendingScore(score, "system")

	return s.Repository.UpdateTag(ctx, *tag)
}
