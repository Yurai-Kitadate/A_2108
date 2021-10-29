package repository

import (
	"sync"

	"github.com/jphacks/A_2108/src/domain"
	"gorm.io/gorm"
)

type PlanRepository struct {
	db *gorm.DB
}

func NewPlanRepository(db *gorm.DB) *PlanRepository {
	return &PlanRepository{
		db: db,
	}
}

type PlanError struct {
	s string
}

func (e PlanError) Error() string {
	return e.s
}

type dictionary_i2s struct {
	sync.Mutex
	dict map[int]string
}

func NewDict_i2s() *dictionary_i2s {
	return &dictionary_i2s{
		dict: make(map[int]string),
	}
}

func (d *dictionary_i2s) Set(key int, value string) {
	d.Lock()
	d.dict[key] = value
	d.Unlock()
}

func (d *dictionary_i2s) Get(key int) string {
	d.Lock()
	defer d.Unlock()
	return d.dict[key]
}

type dictionary_s2i struct {
	sync.Mutex
	dict map[string]int
}

func NewDict_s2i() *dictionary_s2i {
	return &dictionary_s2i{
		dict: make(map[string]int),
	}
}

func (d *dictionary_s2i) Set(key string, value int) {
	d.Lock()
	d.dict[key] = value
	d.Unlock()
}

func (d *dictionary_s2i) Get(key string) int {
	d.Lock()
	defer d.Unlock()
	return d.dict[key]
}

var (
	seasonKey2def   *dictionary_i2s
	categoryKey2def *dictionary_i2s
	timespanKey2def *dictionary_i2s

	seasonDef2key   *dictionary_s2i
	categoryDef2key *dictionary_s2i
	timespanDef2key *dictionary_s2i
)

func initPlan(db *gorm.DB) {
	// Prefetch Definitions
	{
		seasonKey2def = NewDict_i2s()
		categoryKey2def = NewDict_i2s()
		timespanKey2def = NewDict_i2s()

		seasonDef2key = NewDict_s2i()
		categoryDef2key = NewDict_s2i()
		timespanDef2key = NewDict_s2i()

		seasonDefinition := []domain.DBSeasonDefinition{}
		timeSpanDefinition := []domain.DBTimeSpanDefinition{}
		categoryDefinition := []domain.DBCategoryDefinition{}

		if err := db.Find(&seasonDefinition).Error; err != nil {
			panic("DB Load Error")
		}
		for _, v := range seasonDefinition {
			seasonKey2def.Set(v.ID, v.Description)
			seasonDef2key.Set(v.Description, v.ID)
		}

		if err := db.Find(&timeSpanDefinition).Error; err != nil {
			panic("DB Load Error")
		}
		for _, v := range timeSpanDefinition {
			timespanKey2def.Set(v.ID, v.Description)
			timespanDef2key.Set(v.Description, v.ID)
		}

		if err := db.Find(&categoryDefinition).Error; err != nil {
			panic("DB Load Error")
		}
		for _, v := range categoryDefinition {
			categoryKey2def.Set(v.ID, v.Description)
			categoryDef2key.Set(v.Description, v.ID)
		}
	}

}

func (pr PlanRepository) GetPlanByID(planID int) (domain.Plan, error) {
	db := pr.db
	plan := domain.Plan{}

	{
		db_plan := domain.DBPlan{}

		err := db.First(&db_plan, planID).Error
		if err != nil {
			return domain.Plan{}, errHandling(err)
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
		if err != nil {
			return domain.Plan{}, errHandling(err)
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
			if err != nil {
				return domain.Plan{}, errHandling(err)
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
			if err != nil {
				return domain.Plan{}, errHandling(err)
			}

			plan.Days[i].Schedule = make(domain.Schedule, len(db_schedules))
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
					schedule.Address = nil
				} else if err != nil {
					return domain.Plan{}, err
				} else {
					schedule.Address = &domain.Address{
						ID:       db_address.ID,
						PlusCode: db_address.PlusCode,
					}
				}
			}
		}
	}

	{
		conditions := domain.DBCondition{}
		err := db.Where("plan_id = ?", plan.PlanId).First(&conditions).Error
		if err != nil {
			return domain.Plan{}, errHandling(err)
		}

		plan.Conditions.ID = conditions.ID
		plan.Conditions.EstimatedCost = conditions.EstimatedCharge

		db_seasons := []domain.DBSeason{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_seasons).Error
		if err != nil {
			return domain.Plan{}, errHandling(err)
		}

		plan.Conditions.Season = make([]domain.Season, len(db_seasons))
		for i, v := range db_seasons {
			season := &plan.Conditions.Season[i]
			season.ID = v.ID
			season.Text = seasonKey2def.Get(v.SeasonDefinitionID)
		}

		db_categories := []domain.DBCategory{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_categories).Error
		if err != nil {
			return domain.Plan{}, errHandling(err)
		}

		plan.Conditions.Category = make([]domain.Category, len(db_categories))
		for i, v := range db_categories {
			category := &plan.Conditions.Category[i]
			category.ID = v.ID
			category.Text = categoryKey2def.Get(v.CategoryDefinitionID)
		}

		db_timeSpan := []domain.DBTimeSpan{}
		err = db.Where("condition_id = ?", conditions.ID).Find(&db_timeSpan).Error
		if err != nil {
			return domain.Plan{}, errHandling(err)
		}

		plan.Conditions.TimeSpan = make([]domain.TimeSpan, len(db_timeSpan))
		for i, v := range db_timeSpan {
			timespan := &plan.Conditions.TimeSpan[i]
			timespan.ID = v.ID
			timespan.Text = timespanKey2def.Get(v.TimeSpanDefinitionID)
		}
	}

	return plan, nil
}

func (pr PlanRepository) GetPlansOrderedbyTime(limit int) (domain.Plans, error) {
	db := pr.db
	var plans domain.Plans
	db_plans := []domain.DBPlan{}

	err := db.Limit(limit).Find(&db_plans).Error
	if err != nil {
		return domain.Plans{}, errHandling(err)
	}

	plans = make(domain.Plans, len(db_plans))
	for i, v := range db_plans {
		plans[i], err = pr.GetPlanByID(v.ID)
		if err != nil {
			return plans, errHandling(err)
		}
	}
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
		return 0, &PlanError{CANNOT_CONVERT}
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
			return 0, errHandling(err)
		}
		planID = db_plan.ID
	}

	{
		conditions := domain.DBCondition{
			PlanID:          planID,
			EstimatedCharge: plan.Conditions.EstimatedCost,
		}
		err := db.Create(&conditions).Error
		if err != nil {
			return 0, errHandling(err)
		}

		db_seasons := make([]domain.DBSeason, len(plan.Conditions.Season))
		for i, v := range plan.Conditions.Season {
			season := &db_seasons[i]
			season.ConditionID = conditions.ID
			season.SeasonDefinitionID = seasonDef2key.Get(v.Text)
		}
		err = db.Create(&db_seasons).Error
		if err != nil {
			return 0, errHandling(err)
		}

		db_timeSpan := make([]domain.DBTimeSpan, len(plan.Conditions.TimeSpan))
		for i, v := range plan.Conditions.TimeSpan {
			timespan := &db_timeSpan[i]
			timespan.ConditionID = conditions.ID
			timespan.TimeSpanDefinitionID = timespanDef2key.Get(v.Text)
		}
		err = db.Create(&db_timeSpan).Error
		if err != nil {
			return 0, errHandling(err)
		}

		db_categories := make([]domain.DBCategory, len(plan.Conditions.Category))
		for i, v := range plan.Conditions.Category {
			category := &db_categories[i]
			category.ConditionID = conditions.ID
			category.CategoryDefinitionID = categoryDef2key.Get(v.Text)
		}
		err = db.Create(&db_categories).Error
		if err != nil {
			return 0, errHandling(err)
		}
	}

	{
		db_days := make([]domain.DBDay, len(plan.Days))
		for i, v := range plan.Days {
			db_days[i].PlanID = planID
			db_days[i].NthDay = v.NthDay
		}

		err := db.Create(&db_days).Error
		if err != nil {
			return 0, errHandling(err)
		}
		for i, v := range db_days {
			plan.Days[i].ID = v.ID
		}
	}

	{
		for _, day := range plan.Days {
			db_headings := make([]domain.DBHeading, len(day.Headings))

			for j, v := range day.Headings {
				db_heading := &db_headings[j]
				db_heading.DayID = day.ID
				db_heading.HeadingText = v.Text
				db_heading.Order = v.Order
			}

			err := db.Create(&db_headings).Error
			if err != nil {
				return 0, errHandling(err)
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
					return 0, errHandling(err)
				}
			}
			err = db.Create(&db_schedules).Error
			if err != nil {
				return 0, errHandling(err)
			}
		}
	}

	return planID, nil
}

func (pr PlanRepository) PutPlan(plan domain.Plan) error {
	return nil
}
