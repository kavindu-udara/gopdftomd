package controllers

import (
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func GetPageCount(context *model.Context) int {
	return context.PageCount
}


