package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func TestCountPalindromes(t *testing.T) {
    router := setupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/count_palindromes?text=racecar%20noon%20level%20toilet%20plunger%20skibidi", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Contains(t, w.Body.String(), `"palindrome_count":3`)
}

func setupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/count_palindromes", countPalindromes)
    return router
}
