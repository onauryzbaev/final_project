package comment

import (
	"encoding/json"
	"net/http"
)

// Comment - структура комментария
type Comment struct {
	PublicationID int    `json:"publication_id"`
	User          string `json:"user"`
	Text          string `json:"text"`
}

// CommentResponse - структура ответа на комментарий
type CommentResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// HandleComment - обработчик для добавления комментария
func HandleComment(w http.ResponseWriter, r *http.Request) {
	var comment Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Invalid comment data", http.StatusBadRequest)
		return
	}

	if IsCensored(comment.Text) {
		response := CommentResponse{Status: "rejected", Message: "Комментарий содержит нецензурные слова"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Если комментарий проходит проверку, сохраняем его в базе данных (здесь просто пример)
	response := CommentResponse{Status: "accepted", Message: "Комментарий принят"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
