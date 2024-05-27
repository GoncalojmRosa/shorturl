package sites

import (
	"net/http"

	"github.com/goncalojmrosa/shorturl/types"
	"github.com/goncalojmrosa/shorturl/utils"
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
}

func (h *Handler) createSite(w http.ResponseWriter, r *http.Request) {
	//get JSON payload
	var payload types.RegisterSitePayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}
}