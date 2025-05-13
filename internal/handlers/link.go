package handlers

import (
	"gorm.io/gorm"
	"judo/internal/link"
	"judo/pkg/handlerset"
	"judo/pkg/request"
	"net/http"
	"strconv"
)

type LinkHandler struct {
	LinkRepository *link.LinkRepository
}

func NewLinkHandler(deps link.LinkRepository) *LinkHandler {
	return &LinkHandler{
		LinkRepository: &deps,
	}
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[link.LinkCreateRequest](w, r)
		if err != nil {
			return
		}
		link := link.NewLink(body.URL)
		for {
			existed, _ := h.LinkRepository.GetByHash(link.Hash)
			if existed == nil {
				break
			}
			link.GenerateHash()
		}

		createdLink, err := h.LinkRepository.Create(link)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handlerset.HandlerSet(w, createdLink, http.StatusCreated)

	}
}

func (h *LinkHandler) Read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := h.LinkRepository.GetByHash(hash)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[link.LinkUpdateRequest](w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			handlerset.HandlerSet(w, err.Error(), http.StatusBadRequest)
			return
		}
		link, err := h.LinkRepository.UpdateLink(&link.Link{
			Model: gorm.Model{ID: uint(id)},
			URL:   body.URL,
			Hash:  body.Hash,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		handlerset.HandlerSet(w, link, http.StatusCreated)
	}
}

func (h *LinkHandler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			handlerset.HandlerSet(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err = h.LinkRepository.DeleteLink(uint(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, "", http.StatusOK)
	}
}

func (h *LinkHandler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var allLinks []*link.LinksResponse

		links, err := h.LinkRepository.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		for _, l := range links {
			allLinks = append(allLinks, &link.LinksResponse{
				ID:   int(l.ID),
				URL:  l.URL,
				Hash: l.Hash,
			})
		}
		handlerset.HandlerSet(w, allLinks, http.StatusOK)

	}
}
