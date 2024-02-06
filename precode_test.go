package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	// Проверка, если в параметре `count` указано больше, чем есть всего, должны вернуться все доступные кафе.
	totalCount := 4
	req := httptest.NewRequest("GET", "/cafe?count=10&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)
	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)
	body := responseRecorder.Body.String()
	assert.NotEqual(t, "wrong city value", body)
	list := strings.Split(body, ",")
	assert.Equal(t, len(list), totalCount)

}
func TestMainHandlerSuccess(t *testing.T) {
	// Проверка, запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое
	req := httptest.NewRequest("GET", "/cafe?count=4&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)
	status := responseRecorder.Code
	require.Equal(t, status, http.StatusOK)
}

func TestMainHandWrongCity(t *testing.T) {
	// Проверка, корректность введенного названия города
	req := httptest.NewRequest("GET", "/cafe?count=10&city=mocsow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	body := responseRecorder.Body.String()
	assert.Equal(t, "wrong city value", body)

}
