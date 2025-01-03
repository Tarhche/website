package createarticle

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/article"
	"github.com/khanzadimahdi/testproject/domain/author"
	"github.com/khanzadimahdi/testproject/infrastructure/repository/mocks/articles"
	"github.com/khanzadimahdi/testproject/infrastructure/validator"
)

func TestUseCase_Execute(t *testing.T) {
	t.Parallel()

	t.Run("creates an article", func(t *testing.T) {
		t.Parallel()

		var (
			articleRepository articles.MockArticlesRepository
			validator         validator.MockValidator

			r = Request{
				Title:      "test title",
				Excerpt:    "test excerpt",
				Body:       "test body",
				AuthorUUID: "test-author-uuid",
				Tags:       []string{"tag1", "tag2"},
			}
			a = article.Article{
				Cover:       r.Cover,
				Video:       r.Video,
				Title:       r.Title,
				Excerpt:     r.Excerpt,
				Body:        r.Body,
				PublishedAt: r.PublishedAt,
				Author: author.Author{
					UUID: r.AuthorUUID,
				},
				Tags: r.Tags,
			}

			u                = "article-uuid"
			expectedResponse = Response{UUID: u}
		)

		validator.On("Validate", &r).Once().Return(nil)
		defer validator.AssertExpectations(t)

		articleRepository.On("Save", &a).Once().Return(u, nil)
		defer articleRepository.AssertExpectations(t)

		response, err := NewUseCase(&articleRepository, &validator).Execute(&r)

		assert.NoError(t, err)
		assert.Equal(t, &expectedResponse, response)
	})

	t.Run("validation fails", func(t *testing.T) {
		t.Parallel()

		var (
			articleRepository articles.MockArticlesRepository
			validator         validator.MockValidator

			r                = Request{}
			expectedResponse = Response{
				ValidationErrors: domain.ValidationErrors{
					"title":   "title is required",
					"excerpt": "excerpt is required",
					"body":    "body is required",
					"author":  "author is required",
				},
			}
		)

		validator.On("Validate", &r).Once().Return(expectedResponse.ValidationErrors)
		defer validator.AssertExpectations(t)

		response, err := NewUseCase(&articleRepository, &validator).Execute(&r)

		articleRepository.AssertNotCalled(t, "Save")

		assert.NoError(t, err)
		assert.Equal(t, &expectedResponse, response)
	})

	t.Run("saving the article fails", func(t *testing.T) {
		t.Parallel()

		var (
			articleRepository articles.MockArticlesRepository
			validator         validator.MockValidator

			r = Request{
				Title:      "test title",
				Excerpt:    "test excerpt",
				Body:       "test body",
				AuthorUUID: "test-author-uuid",
				Tags:       []string{"tag1", "tag2"},
			}
			a = article.Article{
				Cover:       r.Cover,
				Video:       r.Video,
				Title:       r.Title,
				Excerpt:     r.Excerpt,
				Body:        r.Body,
				PublishedAt: r.PublishedAt,
				Author: author.Author{
					UUID: r.AuthorUUID,
				},
				Tags: r.Tags,
			}

			expectedErr = errors.New("error happened")
		)

		validator.On("Validate", &r).Once().Return(nil)
		defer validator.AssertExpectations(t)

		articleRepository.On("Save", &a).Once().Return("", expectedErr)
		defer articleRepository.AssertExpectations(t)

		response, err := NewUseCase(&articleRepository, &validator).Execute(&r)

		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, response)
	})
}
