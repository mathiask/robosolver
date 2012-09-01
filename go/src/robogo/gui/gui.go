package gui

import (
//    "fmt"
    "net/http"
    "robogo/core"
    "html/template"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    testTemplate.Execute(w, board(core.StandardBoard()))
}

func board(b *core.Board) [][]core.Square {
	n := b.Size()
    result := make([][]core.Square, n)
    for y := uint(0); y < n; y++ {
        result[y] = make([]core.Square, n)
        for x := uint(0); x < n; x++ {
            result[y][x] = b.WallsAt(x, y)
        }
    }
    return result
}


var testTemplate = template.Must(template.New("test").Parse(testTemplateHTML))

const testTemplateHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8"/>
<!-- <script src="scripts/some-script.js"> -->
<!-- </script> -->
<style type="text/css">
body {
  font-family: Trebuchet MS, sans-serif;
  margin-left: auto; margin-right: auto;
  min-width: 100px; max-width: 800px;
  background-color: white;
}
div {
  margin: 0px;
  padding: 0px;
}
div.main {
  padding: 2ex;
  background-color: #ccccff;
  color: #27408b;
  border: 2px solid gray;
  border-radius: 8px; -moz-border-radius: 8px; -webkit-border-radius: 8px;
}
div.box {
  display: inline-block;
  width: 35px;
  height: 35px;
  border: 5px solid white;
  background-color: #EEE;
}
div.box0 {}
div.box1 { border-top: 5px solid blue; }
div.box2 { border-right: 5px solid blue; }
div.box3 {
  border-top: 5px solid blue;
  border-right: 5px solid blue;
}
div.box4 { border-bottom: 5px solid blue; }
div.box5 {
  border-top: 5px solid blue;
  border-bottom: 5px solid blue;
}
div.box6 {
  border-right: 5px solid blue;
  border-bottom: 5px solid blue;
}
div.box7 {
  border-top: 5px solid blue;
  border-right: 5px solid blue;
  border-bottom: 5px solid blue;
}
div.box8 { border-left: 5px solid blue; }
div.box9 {
  border-top: 5px solid blue;
  border-left: 5px solid blue;
}
div.box10 {
  border-right: 5px solid blue;
  border-left: 5px solid blue;
}
div.box11 {
  border-top: 5px solid blue;
  border-right: 5px solid blue;
  border-left: 5px solid blue;
}
div.box12 {
  border-bottom: 5px solid blue;
  border-left: 5px solid blue;
}
div.box13 {
  border-top: 5px solid blue;
  border-bottom: 5px solid blue;
  border-left: 5px solid blue;
}
div.box14 {
  border-right: 5px solid blue;
  border-bottom: 5px solid blue;
  border-left: 5px solid blue;
}
div.box15 {
  border-top: 5px solid blue;
  border-right: 5px solid blue;
  border-bottom: 5px solid blue;
  border-left: 5px solid blue;
}
</style>
</head>
<body>
<div class="main">
<h1>Robogo</h1>
<h2>A “Ricochet Robots” solver in Go</h2>
<h3>Kai Tomerius &amp; Mathias Kegelmann</h3>
{{ range . }}
  <div><!--
  {{ range . }}
    --><div class="box box{{.}}">&nbsp;</div><!--
  {{ end }}
  --></div>
{{ end }}
</div>
</body>
</html>
`