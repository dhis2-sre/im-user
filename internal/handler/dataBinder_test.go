package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestDataBinder(t *testing.T) {
	id := uint(1)

	type request struct {
		Id uint
	}

	body, err := json.Marshal(&request{id})
	assert.NoError(t, err)
	httpRequest := httptest.NewRequest(http.MethodPost, "/whatever", bytes.NewBuffer(body))
	httpRequest.Header.Set("Content-Type", "application/json; charset=UTF-8")
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httpRequest
	req := &request{}

	err = DataBinder(c, req)

	assert.NoError(t, err)
	assert.NoError(t, err)
	assert.Equal(t, id, req.Id)
}
