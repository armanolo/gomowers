package main

import (
	"bytes"
	"mowers/internal/application"
	"mowers/internal/domain"
	"mowers/internal/infrastructure"
	"net/http"
)

// ProcessInstructions Controller to manage
//
//	@Summary	Manage instructions
//	@Description Data input with plateau upper-right coordinate and then mower coordinates and movements
//	@Tags		Mowers operation
//	@Param requestBody body string true "Data input with plateau upper-right coordinate and then mower coordinates and movements" example('55\n12N\nLMLMLMLMM\n33E\nMMRMMRMRRM')
//	@Accept		text/plain
//	@Produce	text/plain
//	@Success	200	{string}	string	"Success script of operations"
//	@Failure	400	{object}	string
//	@Router		/manage [POST]
func (app *Application) Manage(w http.ResponseWriter, r *http.Request) {
	buf := make([]byte, 512)
	r.Body.Read(buf)
	content := bytes.Trim(buf, "\x00")
	var p = new(domain.Plateau)
	var ml []*domain.Mower
	ml, err := infrastructure.ValidateInput(string(content), p)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}
	out, err := application.ManageMowers(ml)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	app.writeJSON(w, http.StatusOK, out)
}
