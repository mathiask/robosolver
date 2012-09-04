package gui

import (
	"fmt"
	"net/http"
	"robogo/core"
	"html/template"
)

type Cell struct {
	Walls core.Square
	Id string
}

func init() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/solve", solveHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	mainTemplate.Execute(w, board(core.StandardBoard()))
}

func board(b *core.Board) [][]Cell {
	n := b.Size()
	result := make([][]Cell, n)
	for y := uint(0); y < n; y++ {
		result[y] = make([]Cell, n)
		for x := uint(0); x < n; x++ {
			result[y][x] = Cell{ b.WallsAt(x, y), fmt.Sprintf("%v_%v", x, y) }
		}
	}
	return result
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	robots := new([4][2]uint)
	for i := 0; i < 4; i++ {
		fmt.Sscanf(r.FormValue(fmt.Sprintf("robot%v", i + 1)),
			"%v_%v", &robots[i][0], &robots[i][1])
	}
	var tx, ty uint
	fmt.Sscanf(r.FormValue("target"), "%v_%v", &tx, &ty)
	b := core.StandardBoard().Reset(robots)
	p := core.NewPosition(b, b.Location(tx, ty))
	fmt.Fprintf(w, solve(p))
}

func solve(p *core.Position) string {
	p.Solve(10)
	return fmt.Sprintf("%v", p.Move())
}

var mainTemplate = template.Must(template.New("test").Parse(mainTemplateHTML))

const mainTemplateHTML = `
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
{{ range . }}
  <div><!--
  {{ range . }}
	--><div class="box box{{.Walls}}" id="{{.Id}}">&nbsp;</div><!--
  {{ end }}
  --></div>
{{ end }}
</div>
<div>
  <img id="target" src="static/images/s.png">
</div>
<div>
  <img id="robot1" src="static/images/r1.png">
  <img id="robot2" src="static/images/r2.png">
  <img id="robot3" src="static/images/r3.png">
  <img id="robot4" src="static/images/r4.png">
</div>
<form action="solve">
<p><input type="text" name="target"></p>
<input type="text" name="robot1">
<input type="text" name="robot2">
<input type="text" name="robot3">
<input type="text" name="robot4">
<input type="submit" value="Solve">
</form>
</body>
</html>
`