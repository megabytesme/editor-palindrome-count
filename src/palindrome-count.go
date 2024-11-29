package main

import (
    "net/http"
    "strings"
    "unicode"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/count_palindromes", countPalindromes)
    router.Run(":8080")
}

func countPalindromes(c *gin.Context) {
    text := c.Query("text")
    words := strings.Fields(text)
    palindromeCount := 0
    for _, word := range words {
        if isPalindrome(word) {
            palindromeCount++
        }
    }
    c.JSON(http.StatusOK, gin.H{"palindrome_count": palindromeCount})
}

func isPalindrome(s string) bool {
    if len(s) <= 1 {
        return false
    }
    runes := []rune(s)
    var cleaned []rune
    for _, r := range runes {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            cleaned = append(cleaned, unicode.ToLower(r))
        }
    }
    for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
        if cleaned[i] != cleaned[j] {
            return false
        }
    }
    return true
}
