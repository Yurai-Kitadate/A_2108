package repository

import (
	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

type PlanError struct {
	s string
}

func (e PlanError) Error() string {
	return e.s
}

func (pr PlanRepository) GetPlanByID(planID int) (domain.Plan, error) {
	db := pr.db
	plan := domain.Plan{}

	{
		db_plan := domain.DBPlan{}

		err := db.First(&db_plan, planID).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		plan.PlanId = db_plan.ID
		plan.Title = db_plan.Title
		plan.Description = db_plan.Description
		plan.Image = db_plan.Image
		plan.CreatedAt = db_plan.CreatedAt
	}

	{
		db_days := []domain.DBDay{}

		err := db.Where("plan_id = ?", planID).Find(&db_days).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}

		plan.Days = make(domain.Days, len(db_days))
		for i, v := range db_days {
			plan.Days[i].ID = v.ID
			plan.Days[i].NthDay = v.NthDay
		}
	}

	{
		for i, day := range plan.Days {
			db_headings := []domain.DBHeading{}

			err := db.Where("day_id = ?", day.ID).Find(&db_headings).Error
			if err == gorm.ErrRecordNotFound {
				return domain.Plan{}, &PlanError{"Record Not Found"}
			} else if err != nil {
				return domain.Plan{}, &PlanError{"Other Error"}
			}

			plan.Days[i].Headings = make(domain.Headings, len(db_headings))
			for j, v := range db_headings {
				heading := &plan.Days[i].Headings[j]
				heading.ID = v.ID
				heading.Text = v.HeadingText
				heading.Order = v.Order
			}

			db_schedules := []domain.DBSchedule{}
			err = db.Where("day_id = ?", day.ID).Find(&db_schedules).Error
			if err == gorm.ErrRecordNotFound {
				return domain.Plan{}, &PlanError{"Record Not Found"}
			} else if err != nil {
				return domain.Plan{}, &PlanError{"Other Error"}
			}
			for j, v := range db_schedules {
				schedule := &plan.Days[i].Schedule[j]
				schedule.ID = v.ID
				schedule.Title = v.Title
				schedule.Description = v.Description
				schedule.StartTime = v.StartTime
				schedule.EndTime = v.EndTime
				schedule.HpLink = v.HPLink
				schedule.ReservationLink = v.ReservationLink
				schedule.Order = v.Order

				db_address := domain.DBAddress{}
				err = db.Where("schedule_id = ?", plan.Days[i].Schedule[j].ID).First(&db_address).Error
				if err == gorm.ErrRecordNotFound {
					return domain.Plan{}, &PlanError{"Record Not Found"}
				} else if err != nil {
					return domain.Plan{}, &PlanError{"Other Error"}
				}
				schedule.Address = &domain.Address{
					ID:       db_address.ID,
					PlusCode: db_address.PlusCode,
				}
			}
		}
	}

	seasonDictionary := map[int]string{}
	timeSpanDictionary := map[int]string{}
	categoryDictionary := map[int]string{}
	// Prefetch Definitions
	{
		seasonDefinition := []domain.DBSeasonDefinition{}
		timeSpanDefinition := []domain.DBTimeSpanDefinition{}
		categoryDefinition := []domain.DBCategoryDefinition{}

		err := db.Find(&seasonDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for _, v := range seasonDefinition {
			seasonDictionary[v.ID] = v.Description
		}

		err = db.Find(&timeSpanDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for _, v := range timeSpanDefinition {
			timeSpanDictionary[v.ID] = v.Description
		}

		err = db.Find(&categoryDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for _, v := range categoryDefinition {
			categoryDictionary[v.ID] = v.Description
		}
	}

	{
		conditions := domain.DBCondition{}
		err := db.Where("plan_id = ?", plan.PlanId).First(&conditions).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}

		plan.Conditions.ID = conditions.ID
		plan.Conditions.EstimatedCost = conditions.EstimatedCharge

		db_seasons := []domain.DBSeason{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_seasons).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for i, v := range db_seasons {
			season := &plan.Conditions.Season[i]
			season.ID = v.ID
			season.Text = seasonDictionary[v.SeasonDefinitionID]
		}

		db_categories := []domain.DBCategory{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_categories).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for i, v := range db_categories {
			category := &plan.Conditions.Category[i]
			category.ID = v.ID
			category.Text = categoryDictionary[v.CategoryDefinitionID]
		}

		db_timeSpan := []domain.DBTimeSpan{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_timeSpan).Error
		if err == gorm.ErrRecordNotFound {
			return domain.Plan{}, &PlanError{"Record Not Found"}
		} else if err != nil {
			return domain.Plan{}, &PlanError{"Other Error"}
		}
		for i, v := range db_timeSpan {
			timespan := &plan.Conditions.TimeSpan[i]
			timespan.ID = v.ID
			timespan.Text = timeSpanDictionary[v.TimeSpanDefinitionID]
		}
	}

	return plan, nil
}

func (pr PlanRepository) GetPlansOrderedbyTime(limit int) ([]domain.DBPlan, error) {
	return nil, nil
}

func (pr PlanRepository) DeletePlanByID(planID int) error {
	return nil
}

func (pr PlanRepository) PostPlan(plan domain.Plan) (int, error) {
	db := pr.db

	userID, err := func() (int, error) {
		if user, ok := plan.CreatorUser.(domain.User); ok {
			return user.ID, nil
		}
		if user, ok := plan.CreatorUser.(domain.MaskedUser); ok {
			return user.ID, nil
		}
		return 0, &PlanError{"Recived type can not be converted into MaskedUser"}
	}()
	if err != nil {
		return 0, err
	}
	var planID int
	{
		db_plan := domain.DBPlan{
			Title:       plan.Title,
			Description: plan.Description,
			Image:       plan.Image,
			UserID:      userID,
			CreatedAt:   plan.CreatedAt,
		}

		err := db.Create(&db_plan).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}
		planID = db_plan.ID
	}

	seasonDictionary := map[string]int{}
	timeSpanDictionary := map[string]int{}
	categoryDictionary := map[string]int{}
	// Prefetch Definitions
	{
		seasonDefinition := []domain.DBSeasonDefinition{}
		timeSpanDefinition := []domain.DBTimeSpanDefinition{}
		categoryDefinition := []domain.DBCategoryDefinition{}
		err := db.Find(&seasonDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return 0, &PlanError{"Record Not Found"}
		} else if err != nil {
			return 0, &PlanError{"Other Error"}
		}
		for _, v := range seasonDefinition {
			seasonDictionary[v.Description] = v.ID
		}

		err = db.Find(&timeSpanDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return 0, &PlanError{"Record Not Found"}
		} else if err != nil {
			return 0, &PlanError{"Other Error"}
		}
		for _, v := range timeSpanDefinition {
			timeSpanDictionary[v.Description] = v.ID
		}

		err = db.Find(&categoryDefinition).Error
		if err == gorm.ErrRecordNotFound {
			return 0, &PlanError{"Record Not Found"}
		} else if err != nil {
			return 0, &PlanError{"Other Error"}
		}
		for _, v := range categoryDefinition {
			categoryDictionary[v.Description] = v.ID
		}
	}

	{
		conditions := domain.DBCondition{
			PlanID:          planID,
			EstimatedCharge: plan.Conditions.EstimatedCost,
		}
		err := db.Create(&conditions).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}

		db_seasons := make([]domain.DBSeason, len(plan.Conditions.Season))
		for i, v := range plan.Conditions.Season {
			season := &db_seasons[i]
			season.ConditionID = conditions.ID
			season.SeasonDefinitionID = seasonDictionary[v.Text]
		}
		err = db.Create(&db_seasons).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}

		db_timeSpan := make([]domain.DBTimeSpan, len(plan.Conditions.TimeSpan))
		for i, v := range plan.Conditions.TimeSpan {
			timespan := &db_timeSpan[i]
			timespan.ConditionID = conditions.ID
			timespan.TimeSpanDefinitionID = timeSpanDictionary[v.Text]
		}
		err = db.Create(&db_timeSpan).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}

		db_categories := make([]domain.DBCategory, len(plan.Conditions.Category))
		for i, v := range plan.Conditions.Category {
			category := &db_categories[i]
			category.ConditionID = conditions.ID
			category.CategoryDefinitionID = categoryDictionary[v.Text]
		}
		err = db.Create(&db_seasons).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}
	}

	{
		db_days := make([]domain.DBDay, len(plan.Days))
		for i, v := range plan.Days {
			db_days[i].ID = v.ID
			db_days[i].PlanID = planID
			db_days[i].NthDay = v.NthDay
		}

		err := db.Create(&db_days).Error
		if err != nil {
			return 0, &PlanError{"Other Error"}
		}
	}

	{
		for _, day := range plan.Days {
			db_headings := make([]domain.DBHeading, len(day.Headings))

			for j, v := range day.Headings {
				db_heading := &db_headings[j]
				db_heading.ID = v.ID
				db_heading.HeadingText = v.Text
				db_heading.Order = v.Order
			}

			err := db.Create(&db_headings).Error
			if err != nil {
				return 0, &PlanError{"Other Error"}
			}

			db_schedules := make([]domain.DBSchedule, len(day.Schedule))
			for j, v := range day.Schedule {
				schedule := &db_schedules[j]
				schedule.DayID = day.ID
				schedule.Title = v.Title
				schedule.Description = v.Description
				schedule.StartTime = v.StartTime
				schedule.EndTime = v.EndTime
				schedule.HPLink = v.HpLink
				schedule.ReservationLink = v.ReservationLink
				schedule.Order = v.Order

				db_address := domain.DBAddress{
					PlusCode: v.Address.PlusCode,
				}
				if err = db.Create(&db_address).Error; err != nil {
					return 0, &PlanError{"Other Error"}
				}
			}

			if err = db.Create(&db_schedules).Error; err != nil {
				return 0, &PlanError{"Other Error"}
			}
		}
	}

	return planID, nil
}

func (pr PlanRepository) PutPlan(plan domain.Plan) error {
	return nil
}
