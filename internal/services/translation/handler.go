package translation

import (
	"net/http"

	"github.com/ducminhgd/go-language-sample/internal/require/response"
	"github.com/ducminhgd/go-language-sample/internal/require/translation"
)

// Handler handles translation HTTP requests
type Handler struct {
	service translation.Service
}

// NewHandler creates a new translation handler
func NewHandler(service translation.Service) *Handler {
	return &Handler{
		service: service,
	}
}

// Translate handles the /translate endpoint
func (h *Handler) Translate(w http.ResponseWriter, r *http.Request) {
	// Get query parameters
	key := r.URL.Query().Get("key")
	lang := r.URL.Query().Get("lang")

	if key == "" {
		resp := response.NewErrorResponse(response.TranslationResponse{
			Message: "Key parameter is required",
		})
		response.JSON(w, http.StatusBadRequest, resp)
		return
	}

	// Get translation
	translatedText, err := h.service.Translate(key, lang)
	if err != nil {
		resp := response.NewErrorResponse(response.TranslationResponse{
			Message: "Failed to translate: " + err.Error(),
		})
		response.JSON(w, http.StatusInternalServerError, resp)
		return
	}

	// Return response
	resp := response.NewSuccessResponse(response.TranslationResponse{
		Message: translatedText,
	})
	response.JSON(w, http.StatusOK, resp)
}
