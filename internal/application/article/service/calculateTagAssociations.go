package service

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/radityacandra/go-cms/internal/application/article/model"
	"github.com/radityacandra/go-cms/internal/application/article/types"
)

func (s *Service) CalculateTagAssociations(ctx context.Context, input model.Article) error {
	article, err := s.Repository.FindArticleByIdAndOptionalStatus(ctx, input.Id, "")
	if err != nil {
		return err
	}

	pairList := article.TagCombinationPairs()

	articles, err := s.Repository.ListArticle(ctx, types.ListArticleInput{
		Page:     1,
		PageSize: 1,
		Status:   "all",
	})
	if err != nil {
		return err
	}
	totalArticle := articles.Pagination.TotalData

	for _, pair := range pairList {
		// association_score = log2((count(tagA AND tagB) * N) / (count(tagA) * count(tagB)))
		countPair, err := s.Repository.CountArticleContainingTags(ctx, pair)
		countTag1, err1 := s.Repository.CountArticleContainingTags(ctx, []string{pair[0]})
		countTag2, err2 := s.Repository.CountArticleContainingTags(ctx, []string{pair[1]})
		if err != nil || err1 != nil || err2 != nil {
			return errors.Join(err, err1, err2)
		}

		score := math.Log2(float64(countPair*totalArticle) / float64(countTag1*countTag2))

		err = s.Repository.UpsertTagAssociation(ctx, types.UpsertTagAssociationInput{
			Id:        uuid.NewString(),
			Tag1Id:    pair[0],
			Tag2Id:    pair[1],
			Score:     score,
			CreatedBy: article.CreatedBy,
			CreatedAt: time.Now().UnixMilli(),
		})
		if err != nil {
			return err
		}
	}

	return nil
}
