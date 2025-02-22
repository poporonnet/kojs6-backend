package controller

import (
	"fmt"
	"time"

	"github.com/mct-joken/kojs5-backend/pkg/application/problem"
	"github.com/mct-joken/kojs5-backend/pkg/repository"
	"github.com/mct-joken/kojs5-backend/pkg/server/controller/model"
	"github.com/mct-joken/kojs5-backend/pkg/utils/id"
)

type ProblemController struct {
	repository    repository.ProblemRepository
	createService problem.CreateProblemService
	findService   problem.FindProblemService
}

func NewProblemController(
	repository repository.ProblemRepository,
	createService problem.CreateProblemService,
	findService problem.FindProblemService,
) *ProblemController {
	return &ProblemController{
		repository:    repository,
		createService: createService,
		findService:   findService,
	}
}

func (c *ProblemController) CreateProblem(req model.CreateProblemRequestJSON) (model.CreateProblemResponseJSON, error) {
	res, err := c.createService.Handle(id.SnowFlakeID(req.ContestID), string(req.Title[0]), req.Title, req.Text, req.Points, req.Limits.Time)
	if err != nil {
		return model.CreateProblemResponseJSON{}, fmt.Errorf("failed to create problem: %w", err)
	}
	return model.CreateProblemResponseJSON{
		ID:     string(res.GetID()),
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Points: res.GetPoint(),
		Limits: struct {
			Memory int `json:"memory"`
			Time   int `json:"time"`
		}(struct {
			Memory int
			Time   int
		}{
			Memory: res.GetMemoryLimit(),
			Time:   res.GetTimeLimit(),
		}),
	}, nil
}

func (c *ProblemController) FindByID(i string) (model.FindProblemResponseJSON, error) {
	// ToDo: 検索を実行したユーザーを渡す
	res, err := c.findService.FindByID(id.SnowFlakeID(i), time.Now(), "")
	if err != nil {
		return model.FindProblemResponseJSON{}, fmt.Errorf("failed to find problem: %w", err)
	}
	return model.CreateProblemResponseJSON{
		ID:     string(res.GetID()),
		Title:  res.GetTitle(),
		Text:   res.GetText(),
		Points: res.GetPoint(),
		Limits: struct {
			Memory int `json:"memory"`
			Time   int `json:"time"`
		}(struct {
			Memory int
			Time   int
		}{
			Memory: res.GetMemoryLimit(),
			Time:   res.GetTimeLimit(),
		}),
	}, nil
}

func (c *ProblemController) FindByContestID(id id.SnowFlakeID) ([]model.FindProblemResponseJSON, error) {
	res, err := c.findService.FindByContestID(id)
	if err != nil {
		return []model.FindProblemResponseJSON{}, fmt.Errorf("failed to find problem: %w", err)
	}

	response := make([]model.FindProblemResponseJSON, len(res))
	for i, v := range res {
		response[i] = model.CreateProblemResponseJSON{
			ID:     string(v.GetID()),
			Title:  v.GetTitle(),
			Text:   v.GetText(),
			Points: v.GetPoint(),
			Limits: struct {
				Memory int `json:"memory"`
				Time   int `json:"time"`
			}(struct {
				Memory int
				Time   int
			}{
				Memory: v.GetMemoryLimit(),
				Time:   v.GetTimeLimit(),
			}),
		}
	}

	return response, nil
}
