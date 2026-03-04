package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type TextRun struct {
	Text string
	Font string
	Size float64
	X, Y float64
}

var (
	reTf      = regexp.MustCompile(`/([A-Za-z0-9]+)\s+([+-]?\d*\.?\d+)\s+Tf`)
	reTd      = regexp.MustCompile(`([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+Td`)
	reTj      = regexp.MustCompile(`\(((?:\\.|[^\\()])*)\)\s*Tj`)
	reTJ      = regexp.MustCompile(`\[(.*?)\]\s*TJ`)
	reTJToken = regexp.MustCompile(`\(((?:\\.|[^\\()])*)\)|([+-]?\d*\.?\d+)`)
	reOp      = regexp.MustCompile(`/[A-Za-z0-9]+\s+[+-]?\d*\.?\d+\s+Tf|[+-]?\d*\.?\d+\s+[+-]?\d*\.?\d+\s+Td|\[(?:.|\n)*?\]\s*TJ|\((?:\\.|[^\\()])*\)\s*Tj|BT|ET`)
)

func parseContentStream(s string) []TextRun {
	var runs []TextRun
	font := ""
	size := 0.0
	x, y := 0.0, 0.0

	ops := reOp.FindAllString(s, -1)
	for _, op := range ops {
		op = strings.TrimSpace(op)
		if op == "" || op == "BT" || op == "ET" {
			continue
		}

		if m := reTf.FindStringSubmatch(op); len(m) == 3 {
			font = m[1]
			size, _ = strconv.ParseFloat(m[2], 64)
			continue
		}
		if m := reTd.FindStringSubmatch(op); len(m) == 3 {
			dx, _ := strconv.ParseFloat(m[1], 64)
			dy, _ := strconv.ParseFloat(m[2], 64)
			x += dx
			y += dy
			continue
		}
		if m := reTj.FindStringSubmatch(op); len(m) == 2 {
			txt := cleanText(decodePDFLiteral(m[1]))
			if txt != "" {
				runs = append(runs, TextRun{Text: txt, Font: font, Size: size, X: x, Y: y})
			}
			continue
		}
		if m := reTJ.FindStringSubmatch(op); len(m) == 2 {
			var b strings.Builder
			for _, t := range reTJToken.FindAllStringSubmatch(m[1], -1) {
				if t[1] != "" {
					b.WriteString(decodePDFLiteral(t[1]))
					continue
				}
				// PDF TJ: strong negative adjustment usually means visible gap -> word space.
				if t[2] != "" {
					n, _ := strconv.ParseFloat(t[2], 64)
					if n < -120 {
						if b.Len() > 0 && !strings.HasSuffix(b.String(), " ") {
							b.WriteByte(' ')
						}
					}
				}
			}
			txt := cleanText(b.String())
			if txt != "" {
				runs = append(runs, TextRun{Text: txt, Font: font, Size: size, X: x, Y: y})
			}
		}
	}
	return runs
}

func decodePDFLiteral(s string) string {
	s = strings.ReplaceAll(s, `\\`, `\`)
	s = strings.ReplaceAll(s, `\(`, `(`)
	s = strings.ReplaceAll(s, `\)`, `)`)
	s = strings.ReplaceAll(s, `\n`, "\n")
	s = strings.ReplaceAll(s, `\r`, "\r")
	s = strings.ReplaceAll(s, `\t`, "\t")

	reOct := regexp.MustCompile(`\\([0-7]{3})`)
	s = reOct.ReplaceAllStringFunc(s, func(m string) string {
		v, err := strconv.ParseInt(m[1:], 8, 32)
		if err != nil {
			return m
		}
		return string(rune(v))
	})
	return s
}

func cleanText(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsPrint(r) {
			b.WriteRune(r)
		} else {
			b.WriteByte(' ')
		}
	}
	return strings.Join(strings.Fields(b.String()), " ")
}

func readFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
