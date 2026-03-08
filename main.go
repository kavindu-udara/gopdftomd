package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	api "github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {

	file := flag.String("f", "", "input file")
	option := flag.String("o", "markdown", "output format (markdown, text, json)")
	output := flag.String("out", "output", "output file name")

	flag.Parse()

	if *file == "" {
		fmt.Println("Please provide an input file using -f flag.")
		return
	}

	// check the file is exists
	if _, err := os.Stat(*file); os.IsNotExist(err) {
		fmt.Printf("File %s does not exist.\n", *file)
		return
	}

	// check the output format
	if *option != "markdown" && *option != "text" && *option != "json" {
		fmt.Printf("Invalid output format: %s. Supported formats are: markdown, text, json.\n", *option)
		return
	}

	// check the output folder is mentioned if not default to current folder
	if *output == "" {
		*output = "output"
	}

	// open the file
	pdf, err := os.Open(*file)
	if err != nil {
		panic(err)
	}
	defer pdf.Close()

	//

	context, err := api.ReadContextFile(*file)
	if err != nil {
		panic(err)
	}

	// get Page count
	fmt.Println("Pages count : ", context.PageCount)

	pgCount := context.PageCount
	baseName := strings.TrimSuffix(filepath.Base(*file), filepath.Ext(*file))
	maxPageDigits := len(fmt.Sprint(pgCount))

	if *option == "markdown" {
		fmt.Printf("Extracting content from %d pages and converting to markdown...\n", pgCount)

		// create the output folder if not exists
		if _, err := os.Stat(*output); os.IsNotExist(err) {
			err = os.Mkdir(*output, 0755)
			if err != nil {
				panic(err)
			}
		}

		if err := os.MkdirAll("mds", 0755); err != nil {
			panic(err)
		}

		imagesDir := filepath.Join(*output, "images")
		if err := os.MkdirAll(imagesDir, 0755); err != nil {
			panic(err)
		}

		// Extract all images once, then page-level links are resolved from names like: <pdf>_<page>_<ImX>.<ext>
		err = api.ExtractImagesFile(*file, imagesDir, nil, nil)
		if err != nil {
			panic(err)
		}

		for i := 1; i <= pgCount; i++ {

			err = api.ExtractContentFile(*file, *output, []string{fmt.Sprint(i)}, nil)
			if err != nil {
				panic(err)
			}

			// read the extracted content
			// file name structure is [filename]_Content_page_[page number].txt
			content, err := os.ReadFile(fmt.Sprintf("%s/%s_Content_page_%d.txt", *output, baseName, i))
			if err != nil {
				panic(err)
			}

			runs := parseContentStream(string(content))
			imageRefs, err := resolvePageImageRefs(string(content), imagesDir, baseName, i, maxPageDigits)
			if err != nil {
				panic(err)
			}

			md := runsToMarkdownWithRulesAndImages(runs, imageRefs, DefaultMarkdownSizeRules)
			err = os.WriteFile(fmt.Sprintf("mds/output_page_%d.md", i), []byte(md), 0644)
			if err != nil {
				panic(err)
			}
			fmt.Printf("Page %d processed and saved as mds/output_page_%d.md\n", i, i)
		}
		fmt.Println("Content extraction completed. Converting to markdown...")
	}

	fmt.Println("Done.")

}
