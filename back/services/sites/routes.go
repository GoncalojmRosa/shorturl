package sites

import (
	"log"
	"net/http"
	"time"

	"github.com/goncalojmrosa/shorturl/types"
	"github.com/goncalojmrosa/shorturl/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.SiteStore
}

func NewHandler(store types.SiteStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/sites", h.createSite).Methods("POST")
	mux.HandleFunc("/sites", h.getSites).Methods("GET")
}

func (h *Handler) createSite(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterSitePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//create site
	site, err := h.store.Insert(r.Context(), &types.Site{
		Id: uuid.New().String(),
		ShortUrl: payload.ShortUrl,
		UrlCode: utils.GenerateShortUrl(),
		CreateAt: time.Now(),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, site)
}

func (h *Handler) getSites(w http.ResponseWriter, r *http.Request) {
	log.Println("Getting all sites")
	//get all sites
	sites, err := h.store.FindAll(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, sites)
}