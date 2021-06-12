package index

import (
	"encoding/json"
	"net/http"
)

func Welcome(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
