package main

import (
	"regexp"
	"strconv"
)

func parseLengthFromDict(dictBytes []byte) int {
    // Look for /Length NNN pattern
    re := regexp.MustCompile(`/Length\s+(\d+)`)
    match := re.FindSubmatch(dictBytes)
    if len(match) < 2 {
        return 0
    }
    
    length, _ := strconv.Atoi(string(match[1]))
    return length
}

func parseSimpleDict(dictBytes []byte) map[string]interface{} {
    result := make(map[string]interface{})
    
    // Extract /Key /Value or /Key NNN patterns
    re := regexp.MustCompile(`/([A-Za-z0-9]+)\s+(/?[A-Za-z0-9]+|\d+)`)
    matches := re.FindAllSubmatch(dictBytes, -1)
    
    for _, match := range matches {
        key := string(match[1])
        value := string(match[2])
        
        // Try to convert numeric values
        if num, err := strconv.Atoi(value); err == nil {
            result[key] = num
        } else {
            result[key] = value
        }
    }
    
    return result
}

func isLikelyText(data []byte) bool {
    // Heuristic: check if most bytes are printable ASCII
    printable := 0
    for _, b := range data {
        if (b >= 32 && b <= 126) || b == '\n' || b == '\r' || b == '\t' {
            printable++
        }
    }
    return len(data) > 0 && float64(printable)/float64(len(data)) > 0.8
}

func truncateString(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen] + "..."
}