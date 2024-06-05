package sites

import (
	"fmt"
	"net/http"
	"time"

	"github.com/goncalojmrosa/shorturl/types"
	"github.com/goncalojmrosa/shorturl/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const baseURL = "http://localhost:8080/api/v1/s"

type Handler struct {
	store types.SiteStore
}

func NewHandler(store types.SiteStore) *Handler {
	return &Handler{store}
}

func (h *Handler) RegisterRoutes(mux *mux.Router) {
	mux.HandleFunc("/s/{urlCode}", h.redirectToSite).Methods("GET")

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

	urlCode := utils.GenerateShortUrl()
	//create site
	newSite, err := h.store.Insert(r.Context(), &types.Site{
		Id:       primitive.NewObjectID(),
		Url:      payload.Url,
		UrlCode:  urlCode,
		ShortUrl: fmt.Sprintf("%s/%s", baseURL, urlCode), //generate short url
		CreateAt: time.Now(),
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, newSite)
}

func (h *Handler) getSites(w http.ResponseWriter, r *http.Request) {
	//get all sites
	sites, err := h.store.FindAll(r.Context())
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, sites)
}

func (h *Handler) redirectToSite(w http.ResponseWriter, r *http.Request) {
	c := mux.Vars(r)["urlCode"]
	site, err := h.store.FindByUrlCode(r.Context(), c)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, site.Url, http.StatusMovedPermanently)

	//utils.WriteJSON(w, http.StatusOK, site)
}
