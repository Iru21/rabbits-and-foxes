package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

func DrawChart(rabbits, foxes []int) {
	p := plot.New()

	p.Title.Text = "Population of Rabbits and Foxes"
	p.X.Label.Text = "Tiicks"
	p.Y.Label.Text = "Population"

	rabbitPoints := make(plotter.XYs, len(rabbits))
	foxPoints := make(plotter.XYs, len(foxes))

	for i := range rabbits {
		rabbitPoints[i].X = float64(i * 10)
		rabbitPoints[i].Y = float64(rabbits[i])
		foxPoints[i].X = float64(i * 10)
		foxPoints[i].Y = float64(foxes[i])
	}

	//err := plotutil.AddLinePoints(p,
	//	"Rabbits", rabbitPoints,
	//	"Foxes", foxPoints,
	//)
	//if err != nil {
	//	panic(err)
	//}

	rabbitsLine, err := plotter.NewLine(rabbitPoints)
	if err != nil {
		panic(err)
	}
	rabbitsLine.LineStyle.Width = vg.Points(1)
	rabbitsLine.LineStyle.Color = plotutil.Color(0)
	foxesLine, err := plotter.NewLine(foxPoints)
	if err != nil {
		panic(err)
	}
	foxesLine.LineStyle.Width = vg.Points(1)
	foxesLine.LineStyle.Color = plotutil.Color(1)

	p.Add(rabbitsLine, foxesLine)
	p.Legend.Add("Rabbits", rabbitsLine)
	p.Legend.Add("Foxes", foxesLine)

	if err := p.Save(8*vg.Inch, 4*vg.Inch, "population.png"); err != nil {
		panic(err)
	}
}
