package main

import (
	"fmt"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type ImageRef struct {
	Name     string
	X        float64
	Y        float64
	Markdown string
}

type pdfMatrix struct {
	a float64
	b float64
	c float64
	d float64
	e float64
	f float64
}

func identityMatrix() pdfMatrix {
	return pdfMatrix{a: 1, d: 1}
}

func multiplyMatrix(left, right pdfMatrix) pdfMatrix {
	return pdfMatrix{
		a: left.a*right.a + left.c*right.b,
		b: left.b*right.a + left.d*right.b,
		c: left.a*right.c + left.c*right.d,
		d: left.b*right.c + left.d*right.d,
		e: left.a*right.e + left.c*right.f + left.e,
		f: left.b*right.e + left.d*right.f + left.f,
	}
}

var (
	reCMLine = regexp.MustCompile(`^\s*([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+([+-]?\d*\.?\d+)\s+cm\s*$`)
	reDoLine = regexp.MustCompile(`^\s*/([A-Za-z0-9]+)\s+Do\s*$`)
)

func parsePageImageRefs(content string) []ImageRef {
	refs := []ImageRef{}
	ctm := identityMatrix()
	stack := []pdfMatrix{}

	for _, raw := range strings.Split(content, "\n") {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		switch line {
		case "q":
			stack = append(stack, ctm)
			continue
		case "Q":
			if len(stack) > 0 {
				ctm = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			continue
		}

		if m := reCMLine.FindStringSubmatch(line); len(m) == 7 {
			a, _ := strconv.ParseFloat(m[1], 64)
			b, _ := strconv.ParseFloat(m[2], 64)
			c, _ := strconv.ParseFloat(m[3], 64)
			d, _ := strconv.ParseFloat(m[4], 64)
			e, _ := strconv.ParseFloat(m[5], 64)
			f, _ := strconv.ParseFloat(m[6], 64)
			ctm = multiplyMatrix(ctm, pdfMatrix{a: a, b: b, c: c, d: d, e: e, f: f})
			continue
		}

		if m := reDoLine.FindStringSubmatch(line); len(m) == 2 {
			refs = append(refs, ImageRef{
				Name: m[1],
				X:    ctm.e,
				Y:    ctm.f,
			})
		}
	}

	sort.SliceStable(refs, func(i, j int) bool {
		if refs[i].Y == refs[j].Y {
			return refs[i].X < refs[j].X
		}
		return refs[i].Y > refs[j].Y
	})

	return refs
}

func resolvePageImageRefs(content, imagesDir, baseName string, page, maxPageDigits int) ([]ImageRef, error) {
	refs := parsePageImageRefs(content)
	result := make([]ImageRef, 0, len(refs))

	for _, ref := range refs {
		pattern := fmt.Sprintf("%s_%0*d_%s.*", baseName, maxPageDigits, page, ref.Name)
		matches, err := filepath.Glob(filepath.Join(imagesDir, pattern))
		if err != nil {
			return nil, err
		}
		if len(matches) == 0 {
			continue
		}
		sort.Strings(matches)

		relPath, err := filepath.Rel("mds", matches[0])
		if err != nil {
			return nil, err
		}
		ref.Markdown = fmt.Sprintf("![%s](%s)", ref.Name, filepath.ToSlash(relPath))
		result = append(result, ref)
	}

	return result, nil
}
