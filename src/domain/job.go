package domain

import "time"

type Job struct {
	ID             int       `json:"id"`
	Jobname        string    `json:"jobname"`
	Dateoffirstjob time.Time `json:"dateoffirstjob"`
}
