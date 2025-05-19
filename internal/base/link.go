package base

import (
	"gorm.io/gorm"
	"judo/configs"
	"judo/internal/link"
	typesimpo "judo/internal/types"
	"judo/pkg/di"
	"judo/pkg/event"
	"judo/pkg/handlerset"
	"judo/pkg/request"
	"net/http"
	"strconv"
)

type LinkHandler struct {
	LinkRepository di.ILinkRepository
	EventBus       *event.EventBus
	Config         *configs.Config
}

func NewLinkHandler(deps link.LinkRepository, config *configs.Config, event *event.EventBus) *LinkHandler {
	return &LinkHandler{
		LinkRepository: &deps,
		Config:         config,
		EventBus:       event,
	}
}

func (h *LinkHandler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[typesimpo.LinkCreateRequest](w, r)
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
		go h.EventBus.Publish(event.Event{
			Type: event.LinkVisitedEvent,
			Data: link.ID,
		})
		handlerset.HandlerSet(w, link, http.StatusOK)
		http.Redirect(w, r, link.URL, http.StatusTemporaryRedirect)
	}
}

func (h *LinkHandler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		body, err := request.HandleBody[typesimpo.LinkUpdateRequest](w, r)
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
		limit, err := strconv.Atoi(r.URL.Query().Get("limit"))
		if err != nil {
			http.Error(w, "invalid limit ", http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(r.URL.Query().Get("offset"))
		if err != nil {
			http.Error(w, "invalid offset ", http.StatusBadRequest)
			return
		}

		result, count, err := h.LinkRepository.GetLinks(uint(limit), uint(offset))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, typesimpo.AllLinksResponse{
			Links: result,
			Count: count,
		}, http.StatusOK)
	}
}
