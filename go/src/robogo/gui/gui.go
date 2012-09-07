package gui

import (
	"fmt"
	"strings"
	"net/http"
	"html/template"
	"robogo/core"
	"appengine"
)

const MAX_DEPTH = 15

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
	fmt.Fprintf(w, solve(p, robots, appengine.NewContext(r)))
}

func solve(p *core.Position, robots *[4][2]uint, ctx appengine.Context) string {
	for depth := uint(2); depth <= MAX_DEPTH; depth++ {
		ctx.Infof("Solving with max depth %v...", depth);
		if p.Solve(depth) {
			ctx.Infof("Solved!");
			return strings.Join(p.Move(), ", ")
		}
		p.Reset(robots)
	}
	ctx.Infof("Search aborted!");
	return fmt.Sprintf("No solution found at depth %v!", MAX_DEPTH)
}

var mainTemplate = template.Must(template.New("test").Parse(mainTemplateHTML))

const mainTemplateHTML = `
<!DOCTYPE html>
<html lang="en">
<head>
<meta charset="utf-8"/>
<script src="static/scripts/jquery-1.7.2.min.js">
</script>
<script src="static/scripts/jquery-ui-1.8.18.custom.min.js">
</script>
<script>
$(function(){
  var robogo = {
	  positions: {}
	},
	solutionDiv = $("#solution"),
	positionsComplete = function() {
	  var p = robogo.positions;
	  if (!p["target"]) return false;
	  for (var i = 1; i <= 4; i++) {
		if (!p["robot" + i]) return false;
	  }
	  return true;
	},
	showSolving = function() {
	  solutionDiv.text("Solving...");
	},
	solve = function() {
	  $.get("solve", robogo.positions, function(data) {
		solutionDiv.text(data);
	  });
	};
  solutionDiv.ajaxError(function() {
	$(this).text( "AJAX error!" );
  });
  $("img").draggable({revert: "invalid"});
  $(".box").droppable({drop: function(event, ui) {
	  var target = $(this),
		id = ui.draggable.attr("id");
	  if (robogo.positions[id]) {
		$("#" + robogo.positions[id]).droppable("option", "disabled", false);
	  }
	  robogo.positions[id] = this.id;
	  ui.draggable.detach().css({top: 0, left: 0}).appendTo(target);
	  $("#" + this.id).droppable("option", "disabled", true);
	},
	hoverClass: "ui-state-hover"});
  $("button").click(function(){
	if (positionsComplete()) {
	  showSolving();
	  solve();
	} else {
	  alert("Please place target and all robots!");
	}
  });
});
</script>
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
  vertical-align: bottom;
  width: 35px;
  height: 35px;
  border: 5px solid white;
  background-color: #EEE;
}
div.ui-state-hover { background-color: #CFC; }
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
button { margin: 2ex 0px 2ex 0px; }
</style>
</head>
<body>
<div class="main">
<h1>Robogo</h1>
<h2>A “Ricochet Robots” solver in Go</h2>
<div>
{{ range . }}
  <div><!--
  {{ range . }}
	--><div class="box box{{.Walls}}" id="{{.Id}}"></div><!--
  {{ end }}
  --></div>
{{ end }}
</div>
<div>
  <img id="target" src="static/images/s.png" alt="T" title="drag me">
</div>
<div>
  <img id="robot1" src="static/images/r1.png" alt="1" title="drag me">
  <img id="robot2" src="static/images/r2.png" alt="2" title="drag me">
  <img id="robot3" src="static/images/r3.png" alt="3" title="drag me">
  <img id="robot4" src="static/images/r4.png" alt="4" title="drag me">
</div>
<button>Solve!</button>
<div id="solution"/>
</div>
</body>
</html>
`
