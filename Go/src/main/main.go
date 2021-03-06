/*
   MAIN
   Author: Michal Kukowski
   email: michalkukowski10@gmail.com
   LICENCE: GPL3.0
*/

package main

import (
	"configs"
	"graph"
	"myswitch"
	"track"
	"train"
	"sync"
	"client"
	"driver"
)

func main() {

	var wg sync.WaitGroup
	/* Load configs from file */
	configs.Conf.Load()

	/* Create Global Arrays */
	train.Trains.NewTrains(configs.Conf.NumTrains())
	myswitch.Switches.NewSwitches(configs.Conf.NumSwitches())
	track.Tracks.NewTracks(configs.Conf.NumTracks())

	/* Load ALL */
	train.Trains.Load()
	myswitch.Switches.Load()
	track.Tracks.Load()
	graph.Load()


	/* Start Thread if needed */
	if configs.Conf.Mode() == configs.SILENT {
		wg.Add(1)

		go client.Talk()
	}

	drivers := make([]*driver.Driver, configs.Conf.NumTrains())
	for i := 0; i < len(drivers); i++ {
		drivers[i] = driver.New(train.Trains.GetTrainByID(i + 1))
		wg.Add(1)
		go drivers[i].Drive()
	}
	wg.Wait()
}
