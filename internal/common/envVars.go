package common

import "spyrosmoux/core-engine/internal/helpers"

var (
	ApiPort  = helpers.LoadEnvVariable("API_PORT")
	GhSecret = helpers.LoadEnvVariable("GH_SECRET")
	GhToken  = helpers.LoadEnvVariable("GH_TOKEN")
)
