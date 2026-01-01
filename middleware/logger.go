package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ANSI color codes
const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

// colorizeMethod returns colored method string
func colorizeMethod(method string) string {
	switch method {
	case "GET":
		return green + method + reset
	case "POST":
		return blue + method + reset
	case "PUT":
		return yellow + method + reset
	case "DELETE":
		return red + method + reset
	case "PATCH":
		return magenta + method + reset
	default:
		return cyan + method + reset
	}
}

// colorizeStatus returns colored status code string
func colorizeStatus(status int) string {
	statusStr := strconv.Itoa(status)
	switch {
	case status >= 200 && status < 300:
		return green + statusStr + reset
	case status >= 300 && status < 400:
		return yellow + statusStr + reset
	case status >= 400 && status < 500:
		return red + statusStr + reset
	case status >= 500:
		return red + statusStr + reset
	default:
		return white + statusStr + reset
	}
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// Fixed width formatting
		timestamp := params.TimeStamp.Format("2006/01/02 15:04:05")
		method := colorizeMethod(params.Method)
		path := params.Path
		status := colorizeStatus(params.StatusCode)
		latency := params.Latency.String()
		clientIP := params.ClientIP
		requestSize := fmt.Sprintf("%d", params.Request.ContentLength)
		responseSize := fmt.Sprintf("%d", params.BodySize)

		// Format with fixed widths: timestamp(20), method(10), path(40), status(5), latency(15), ip(15), reqSize(10), respSize(10)
		logLine := fmt.Sprintf("[ZUM] %-20s %-10s %-40s %-5s %-15s %-15s %-10s %-10s\n",
			timestamp,
			method,
			path,
			status,
			latency,
			clientIP,
			requestSize,
			responseSize,
		)

		return logLine
	})
}
