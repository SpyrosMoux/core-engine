package common

import "spyrosmoux/core-engine/internal/helpers"

var (
	GhSecret = helpers.LoadEnvVariable("GH_SECRET")
	GhToken  = helpers.LoadEnvVariable("GH_TOKEN")
)
