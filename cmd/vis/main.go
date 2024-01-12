package main

import (
	"github.com/leoffx/bandit-algorithms/lib/arm"
	"github.com/leoffx/bandit-algorithms/lib/bandit"
	"github.com/leoffx/bandit-algorithms/lib/database"
	"github.com/leoffx/bandit-algorithms/lib/simulation"
	"github.com/leoffx/bandit-algorithms/lib/strategy"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

const numRounds = 1000
const numArms = 3

func main() {
	strategy := strategy.NewEpsilonGreedy(0.3)
	// strategy := strategy.NewRecoveringDifferenceSoftmax(0.3,)
	bandit, db := simulation.Run(numArms, numRounds, strategy)
	createPlot(bandit, db)

}

func createPlot(bandit *bandit.Bandit, db *database.DatabaseAggregator) {
	p := plot.New()

	p.Title.Text = "Average Reward over Time"
	p.X.Label.Text = "Round"
	p.Y.Label.Text = "Average Reward"

	var bestArm arm.Arm
	for _, arm := range bandit.Arms {
		if bestArm == nil || arm.Mean() > bestArm.Mean() {
			bestArm = arm
		}
	}

	pts := getPoints(db.Entries)
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	p.Add(line)

	hline, err := plotter.NewLine(plotter.XYs{
		{X: 0, Y: bestArm.Mean()},
		{X: float64(numRounds), Y: bestArm.Mean()},
	})
	if err != nil {
		panic(err)
	}
	hline.Color = plotutil.Color(0)
	p.Add(hline)

	if err := p.Save(400, 400, "points.png"); err != nil {
		panic(err)
	}
}

func getPoints(entries []*database.Entry) plotter.XYs {
	pts := make(plotter.XYs, len(entries))
	avgReward := 0.0
	for i, entry := range entries {
		pts[i].X = float64(entry.Round)
		avgReward = (avgReward*float64(i) + entry.Reward) / float64(i+1)
		pts[i].Y = avgReward
	}
	return pts

}
