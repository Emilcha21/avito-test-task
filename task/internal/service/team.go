package service

import (
	"avito-test/internal/dto"
	apperrors "avito-test/internal/errors"
	"avito-test/internal/models"
	"errors"

	"gorm.io/gorm"
)

func (s *Service) CreateTeam(req *dto.TeamReq) error {
	s.Logger.Info("Starting function: CreateTeam",
		"team_name", req.TeamName)
	teamModel := &models.Team{
		TeamName: req.TeamName,
	}

	existsName, err := s.ITeamRepo.GetTeamByName(req.TeamName)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		s.Logger.Error("Failed to get team",
			"error", err)
		return err
	}

	if existsName == req.TeamName {
		s.Logger.Warn("Team exists",
			"team_name", req.TeamName)
		return apperrors.ErrTeamExists
	}

	var userModels []models.User
	for _, user := range req.Members {
		s.Logger.Debug("User data", "user value", user)
		userModels = append(userModels, models.User{
			UserId:   user.UserId,
			Username: user.Username,
			IsActive: *user.IsActive,
			TeamName: req.TeamName,
		})
	}

	err = s.ITeamRepo.CreateTeam(teamModel, userModels)
	if err != nil {
		s.Logger.Error("Failed to create team",
			"error", err)
		return err
	}

	s.Logger.Info("Successfully create team",
		"team_name", req.TeamName)

	return nil
}

func (s *Service) GetTeamByName(name string) ([]dto.UserResponse, error) {
	s.Logger.Info("Starting function: GetTeamByName",
		"name", name)
	_, err := s.ITeamRepo.GetTeamByName(name)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			s.Logger.Warn("Team not found",
				"team_name", name)
			return nil, apperrors.ErrTeamNotFound
		}
		s.Logger.Error("Failed to get team",
			"error", err)
		return nil, err
	}
	s.Logger.Debug("Get team")

	users, err := s.GetUsersByTeam(name)
	if err != nil {
		s.Logger.Error("Failed to get users by team",
			"error", err)
		return nil, err
	}

	s.Logger.Info("Successfully get team by name")

	return users, nil
}
