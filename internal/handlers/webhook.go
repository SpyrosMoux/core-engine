package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
	"io"
	"net/http"
	"spyrosmoux/core-engine/internal/common"
	"spyrosmoux/core-engine/internal/logger"
	"spyrosmoux/core-engine/internal/pipelines"
)

func HandleWebhook(c *gin.Context) {
	hook, _ := github.New(github.Options.Secret(common.GhSecret))

	payload, err := hook.Parse(c.Request, github.PushEvent)
	if err != nil {
		logger.Log(logger.ErrorLevel, "Error parsing webhook payload: "+err.Error())
	}

	push := payload.(github.PushPayload)

	pipeline, err := fetchPipelineConfig(push.Repository.FullName, push.Ref)
	if err != nil {
		logger.Log(logger.ErrorLevel, "Error fetching pipeline config: "+err.Error())
	}

	pipelines.RunJob(string(pipeline))
}

func fetchPipelineConfig(repoFullName, branchName string) ([]byte, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/contents/sample-pipeline.yaml?ref=%s", repoFullName, branchName)
	logger.Log(logger.InfoLevel, "Fetching pipeline config from "+url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Accept", "application/vnd.github.v3.raw")
	req.Header.Set("Authorization", "Bearer "+common.GhToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pipeline config: %w", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Log(logger.FatalLevel, "Failed to close response body: "+err.Error())
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch pipeline config: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	return body, nil
}
