package input

import (
	domainClient "influ-dojo/api/domain/client"
	"influ-dojo/api/domain/repository"
	"log"
)

type RecordUsers struct {
	FollowerClient    domainClient.Follower `json:"-"`
	UserRepo          repository.User       `json:"-"`
	DailyWorkRepo     repository.Work       `json:"-"`
	DailyResultRepo   repository.Result     `json:"-"`
	WeeklyWorkRepo    repository.Work       `json:"-"`
	WeeklyResultRepo  repository.Result     `json:"-"`
	MonthlyWorkRepo   repository.Work       `json:"-"`
	MonthlyResultRepo repository.Result     `json:"-"`
}

func (in *RecordUsers) RecordUsers() error {
	latestFollowers, err := in.FollowerClient.GetFollowers()
	if err != nil {
		return err
	}

	IDs, err := in.UserRepo.LoadIDs()
	if err != nil {
		return err
	}

	for _, f := range latestFollowers {
		if contains(IDs, f.UserID) {
			IDs = remove(IDs, f.UserID)
			continue
		}

		if err := in.UserRepo.Save(f.User); err != nil {
			return err
		}

		if err := in.DailyWorkRepo.Save(f.Work); err != nil {
			return err
		}

		if err := in.DailyResultRepo.Save(f.Result); err != nil {
			return err
		}

		if err := in.WeeklyWorkRepo.Save(f.Work); err != nil {
			return err
		}

		if err := in.WeeklyResultRepo.Save(f.Result); err != nil {
			return err
		}

		if err := in.MonthlyWorkRepo.Save(f.Work); err != nil {
			return err
		}

		if err := in.MonthlyResultRepo.Save(f.Result); err != nil {
			return err
		}
	}

	log.Printf("削除対象IDs: %v", IDs)
	for _, userID := range IDs {
		if err := in.UserRepo.Delete(userID); err != nil {
			return err
		}
	}

	return nil
}

func contains(slice []string, object string) bool {
	for _, element := range slice {
		if element == object {
			return true
		}
	}
	return false
}

func remove(elements []string, search string) []string {
	var result []string
	for _, element := range elements {
		if element != search {
			result = append(result, element)
		}
	}

	return result
}
