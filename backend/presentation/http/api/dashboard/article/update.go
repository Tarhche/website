package article

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/khanzadimahdi/testproject/application/auth"
	updatearticle "github.com/khanzadimahdi/testproject/application/dashboard/article/updateArticle"
	"github.com/khanzadimahdi/testproject/domain"
	"github.com/khanzadimahdi/testproject/domain/permission"
)

type updateHandler struct {
	useCase    *updatearticle.UseCase
	authorizer domain.Authorizer
}

func NewUpdateHandler(useCase *updatearticle.UseCase, a domain.Authorizer) *updateHandler {
	return &updateHandler{
		useCase:    useCase,
		authorizer: a,
	}
}

func (h *updateHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	userUUID := auth.FromContext(r.Context()).UUID
	if ok, err := h.authorizer.Authorize(userUUID, permission.ArticlesUpdate); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	} else if !ok {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	var request updatearticle.Request
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	request.AuthorUUID = userUUID

	response, err := h.useCase.Execute(request)
	switch {
	case errors.Is(err, domain.ErrNotExists):
		rw.WriteHeader(http.StatusNotFound)
	case err != nil:
		rw.WriteHeader(http.StatusInternalServerError)
	case len(response.ValidationErrors) > 0:
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(response)
	default:
		rw.WriteHeader(http.StatusNoContent)
	}
}
