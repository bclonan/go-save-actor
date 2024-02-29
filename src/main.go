package main

import (
	"./actor"
	"./promises"
	"./githubapi"
	"./adapter"
	"./facade"
)

func main() {
	// Create a new actor
	myActor := actor.NewActor()

	// Create a new promise
	myPromise := promises.NewPromise()

	// Create a new GitHub API client
	myGithubAPI := githubapi.NewGithubAPI()

	// Create a new adapter
	myAdapter := adapter.NewAdapter(myGithubAPI)

	// Create a new facade
	myFacade := facade.NewFacade(myActor, myPromise, myAdapter)

	// Use the facade to interact with the subsystems
	myFacade.DoSomething()
}