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
    result := make([][]core.Square, 16)
    for y := uint(0); y < 16; y++ {
        result[y] = make([]core.Square, 16)
        for x := uint(0); x < 16; x++ {
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
div {
  margin: 0px;
  padding: 0px;
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
<h1>Test</h1>
{{ range . }}
  <div><!--
  {{ range . }}
    --><div class="box box{{.}}">{{ . }}</div><!--
  {{ end }}
  --></div>
{{ end }}
</body>
</html>
`