package main

import (
	"fmt"
	"os"

	// "strconv"

	api "github.com/pdfcpu/pdfcpu/pkg/api"
	// reader "github.com/kavindu-udara/PDFtoMD/reader"
)

func main() {

	file, err := os.Open("test.pdf")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	context, err := api.ReadContextFile("test.pdf")
	if err != nil {
		panic(err)
	}

	dic, types, model, err := context.PageDict(1, true)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Context : ", context)
	fmt.Println("Page Dict : ", dic)
	fmt.Println("Types : ", types)
	fmt.Println("Model : ", model)

	// get pages count
	fmt.Printf("Pages count : %d\n", context.PageCount)

	// extract first page
	modelContext, err := api.ReadContext(file, nil)
	if err != nil {
		panic(err)
	}

	fmt.Println(modelContext.FindObject(11))
	// read object 11 stream
	obj, err := modelContext.FindObject(11)
	if err != nil {
		panic(err)
	}
	fmt.Println("Object 11 : ", obj.PDFString())

	// err = api.ExtractContentFile("test.pdf", "content", nil, nil)
	// if err != nil {
	// 	panic(err)
	// }

	// err = api.ExportFormFile("test.pdf", "form.json", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Extracted Page : ", modelContext)

	// readable

	raw, err := readFile("content/test_Content_page_153.txt")
	if err != nil {
		panic(err)
	}

	runs := parseContentStream(raw)

	for _, run := range runs {
		fmt.Printf("Text: %q, Font: %s, Size: %.2f, Position: (%.2f, %.2f)\n",
			run.Text, run.Font, run.Size, run.X, run.Y)
	}

	rules := DefaultMarkdownSizeRules
	md := runsToMarkdownWithRules(runs, rules)

	if err := os.WriteFile("output.md", []byte(md), 0644); err != nil {
		panic(err)
	}

	fmt.Println("Markdown output written to output.md")

}
