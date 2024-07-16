package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// omitempty代表如果是空的就不要include在這個json
// interface{} 也可以寫成any, 一樣意思
type JSONResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ...代表可0或多
// write to responseWriter
// payload就是data interface{}
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}
	if len(headers) > 0 { //檢查看我們有沒有要加入ｈｅａｄｅｒ
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// 假設我們預計讀到data是pointer
func (app *application) readJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1024 * 1024 //1mega byte
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	err := dec.Decode(data)
	if err != nil {
		return err //太大 or 包含未知領域or無法decode
	}
	// decode into any struct, a throwaway variable
	// 如果请求体中只有一个合法的 JSON 值，那么在解码第一次后，dec 的缓冲区应该已经耗尽，没有更多的数据可供解码。因此，这次解码应该会返回 io.EOF，表示没有更多的数据了
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only contain a single JSON value") //代表body不只有一份json
	}
	return nil
}

func (app *application) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest // default錯誤
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload JSONResponse
	payload.Error = true
	payload.Message = err.Error()
	// 沒問題的話就會回傳null
	return app.writeJSON(w, statusCode, payload)
}
