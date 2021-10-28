package controller

import (
	"time"

	"github.com/jphacks/A_2108/src/domain"
)

var (
	MockUser1 domain.User
	MockUser2 domain.User
	MockPlan  domain.Plan
	MockPlans domain.Plans
)

func init() {
	{
		MockUser1.ID = 1
		MockUser1.UserName = "Mock User 1"
		MockUser1.Email = "mail dayo"
		MockUser1.Password = ""
		MockUser1.Image = "https://storage.googleapis.com/studio-design-asset-files/projects/4BqNmnYVOr/s-1200x1697_v-fms_webp_4e5dd39f-c6ed-48bc-bd0b-e350b05d5425_small.webp"
		MockUser1.DisplayName = "mieru namae"
		MockUser1.DateOfBirth = time.Now()
		MockUser1.Sex = 0
		hp := "google.com"
		instagram := "https://twitter.com/home"
		twitter := "https://github.com/"
		MockUser1.Contacts = domain.Contacts{
			ID:        101,
			Hp:        &hp,
			Instagram: &instagram,
			Twitter:   &twitter,
			Facebook:  nil,
			Tiktok:    nil,
			Biography: nil,
		}
		MockUser1.Creator = nil
		MockUser1.Place = domain.Place{
			ID:         1,
			Area:       2,
			Prefecture: 3,
			City:       4,
			Name:       "siranaibasyo",
		}
	}
	{
		MockUser2 = MockUser1
		MockUser2.ID = 2
		MockUser2.Creator = &domain.Creator{
			ID:   1,
			Name: "kuri eita",
			Job: domain.Job{
				ID:             2,
				Jobname:        "sigoto wo suru",
				DateOfFirstJob: time.Now(),
			},
		}
	}
}
