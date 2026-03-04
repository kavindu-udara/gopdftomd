package main

import (
    "math"
    "strings"
)

type MarkdownSizeRules struct {
    H1Min              float64
    H2Min              float64
    H3Min              float64
    BodyMin            float64
    IgnoreBelow        float64
    LineMergeTolerance float64
    ParagraphGap       float64
}

var DefaultMarkdownSizeRules = MarkdownSizeRules{
    H1Min:              20.0,
    H2Min:              16.0,
    H3Min:              13.0, // 14.35 becomes heading
    BodyMin:            9.0,
    IgnoreBelow:        7.5, // skip tiny footer markers
    LineMergeTolerance: 2.0,
    ParagraphGap:       12.0,
}

// ...existing code...

func classifyBySize(size float64, rules MarkdownSizeRules) string {
    switch {
    case size < rules.IgnoreBelow:
        return "ignore"
    case size >= rules.H1Min:
        return "h1"
    case size >= rules.H2Min:
        return "h2"
    case size >= rules.H3Min:
        return "h3"
    case size >= rules.BodyMin:
        return "body"
    default:
        return "small"
    }
}

func runsToMarkdownWithRules(runs []TextRun, rules MarkdownSizeRules) string {
    var out strings.Builder
    var line strings.Builder
    prevY := 0.0
    hasPrev := false

    flushLine := func(paragraphBreak bool) {
        txt := strings.TrimSpace(line.String())
        if txt != "" {
            out.WriteString(txt)
            out.WriteString("\n")
            if paragraphBreak {
                out.WriteString("\n")
            }
        }
        line.Reset()
    }

    for _, r := range runs {
        txt := normalizeSpaces(r.Text)
        if txt == "" {
            continue
        }
        kind := classifyBySize(r.Size, rules)
        if kind == "ignore" {
            continue
        }

        if kind == "h1" || kind == "h2" || kind == "h3" {
            flushLine(true)
            prefix := "### "
            if kind == "h1" {
                prefix = "# "
            } else if kind == "h2" {
                prefix = "## "
            }
            out.WriteString(prefix + txt + "\n\n")
            hasPrev = false
            continue
        }

        if !hasPrev {
            line.WriteString(txt)
            prevY = r.Y
            hasPrev = true
            continue
        }

        dy := math.Abs(r.Y - prevY)
        if dy <= rules.LineMergeTolerance {
            line.WriteString(" ")
            line.WriteString(txt)
        } else {
            flushLine(dy >= rules.ParagraphGap)
            line.WriteString(txt)
        }
        prevY = r.Y
    }

    flushLine(false)
    return strings.TrimSpace(out.String()) + "\n"
}
func normalizeSpaces(s string) string {
    return strings.Join(strings.Fields(s), " ")
}