package web

import (
  "html/template"
  "os"
)

type CssTemplate struct {
  Colors struct {
    Primary string
    Secondary string
    Background string
  }   
}

func NewCssTemplate() *CssTemplate {
  return &CssTemplate{
    Colors: {
      Primary: "cyan",
      Secondary: "hotpink",
      Background: "gray",
    },
}
