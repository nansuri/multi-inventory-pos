package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BranchMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		branchIDStr := c.GetHeader("X-Branch-ID")
		if branchIDStr == "" {
			// Some routes might not require branch context (e.g., listing branches)
			// But for operational routes, we might want to enforce it.
			// For now, we just pass through.
			c.Next()
			return
		}

		branchID, err := strconv.ParseUint(branchIDStr, 10, 32)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Branch ID header"})
			c.Abort()
			return
		}

		c.Set("branch_id", uint(branchID))
		c.Next()
	}
}
