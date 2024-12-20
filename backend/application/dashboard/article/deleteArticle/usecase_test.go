package deletearticle

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khanzadimahdi/testproject/infrastructure/repository/mocks/articles"
)

func TestUseCase_Execute(t *testing.T) {
	t.Parallel()

	t.Run("deleting an article succeeds", func(t *testing.T) {
		t.Parallel()

		var (
			articleRepository articles.MockArticlesRepository

			r = Request{ArticleUUID: "article-uuid"}
		)

		articleRepository.On("Delete", r.ArticleUUID).Return(nil)
		defer articleRepository.AssertExpectations(t)

		err := NewUseCase(&articleRepository).Execute(&r)

		assert.NoError(t, err)
	})

	t.Run("deleting an article fails", func(t *testing.T) {
		t.Parallel()

		var (
			articleRepository articles.MockArticlesRepository

			r             = Request{ArticleUUID: "article-uuid"}
			expectedError = errors.New("article deletion failed")
		)

		articleRepository.On("Delete", r.ArticleUUID).Return(expectedError)
		defer articleRepository.AssertExpectations(t)

		err := NewUseCase(&articleRepository).Execute(&r)

		assert.ErrorIs(t, err, expectedError)
	})
}
