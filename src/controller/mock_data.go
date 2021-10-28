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
	{
		MockPlan.PlanId = 1
		MockPlan.Title = "のんびり会津旅"
		MockPlan.Description = "のんびりと会津の旅行を楽しむ"
		MockPlan.Image = "https://avatars.githubusercontent.com/u/34409044?s=40&v=4"
		MockPlan.CreatedAt = time.Now()
		MockPlan.CreatorUser = MockUser2
		hplink := "https://twitter.com/home"
		reservelink := "https://twitter.com/home"
		MockPlan.Days = domain.Days{
			{
				NthDay: 1,
				Headings: domain.Headings{
					{
						ID:    1,
						Text:  "今日一番目の予定",
						Order: 1,
					},
					{
						ID:    2,
						Text:  "今日2番目の予定",
						Order: 3,
					},
					{
						ID:    3,
						Text:  "今日参番目の予定",
						Order: 5,
					},
				},
				Schedule: domain.Schedule{
					{
						ID:          1,
						Title:       "電車に乗る",
						Description: "遠くに移動するので電車に乗る",
						StartTime:   420,
						EndTime:     435,
						Addresss: &domain.Address{
							ID:       1,
							PlusCode: "MP43+V5 渋谷区",
						},
						HpLink:          &hplink,
						ReservationLink: &reservelink,
						Order:           4,
					},
				},
			},
		}
		MockPlan.Conditions = domain.Conditions{
			ID: 123,
			Place: []domain.Place{
				{
					ID:         1,
					Area:       4,
					Prefecture: 3,
					City:       2,
					Name:       "doko?koko",
				},
			},
			Season: []domain.Season{
				{
					ID:   5,
					Text: "夏",
				},
			},
			TimeSpan: []domain.TimeSpan{
				{
					ID:   1,
					Text: "1月",
				},
			},
			Category: []domain.Category{
				{
					ID:   1,
					Text: "娯楽",
				},
			},
		}
		MockPlans = make(domain.Plans, 0)
		for i := 1; i <= 10; i++ {
			MockPlans = append(MockPlans, MockPlan)
		}
	}
}
