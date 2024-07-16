package main

import (
	"net/http"
)

// 加入前方application當作receiver可以使這個func能夠調用他裡面所有的內容
// 當有人一開始進如我們的網頁 "/" 我們要回傳他的狀態
func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Go Movies up and running",
		Version: "1.0.0",
	}
	// 藉由寫util writeJSON可以把以下輝自轉化成下面那一行app.writeJSON
	// out, err := json.Marshal(payload)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// w.Write(out)
	_ = app.writeJSON(w, http.StatusOK, payload)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	_ = app.writeJSON(w, http.StatusOK, movies) //movies為Payload
}
