package mysql

import (
	"silver.com/internal/entity"
	"silver.com/internal/infra/db"
)

type JobRepository struct {
	db *db.MysqlDB
}

func NewJobRepository(db *db.MysqlDB) *JobRepository {
	return &JobRepository{
		db: db,
	}
}

func (r *JobRepository) FetchJobs() ([]*entity.Job, error) {
	query := "SELECT * FROM Jobs WHERE active = 1 LIMIT 20"
	var jobs []*entity.Job

	rows, err := r.db.Conn.Query(query)
	if err != nil {
		return jobs, err
	}
	job := &entity.Job{}

	for rows.Next() {
		err = rows.Scan(
			&job.ID,
			&job.Name,
			&job.Owner,
			&job.ProducerName,
			&job.Active,
			&job.FunctionID,
			&job.Arguments,
			&job.ExtraArguments,
			&job.Priority,
			&job.RunInterval,
			&job.LastRun,
			&job.LastRunMessage,
			&job.NextRun,
			&job.RunTimeout,
			&job.PossibleRetries,
			&job.LastErrors,
			&job.ConcurrentRun,
			&job.HighPriority,
			&job.DispatchedAt,
			&job.StartedAt,
			&job.LastRunHash,
			&job.Destination,
			&job.SourceName,
			&job.SourceUrl,
			&job.SourceSection,
			&job.SourceChannel,
			&job.SourceProducer,
			&job.RemovalPolicyType,
			&job.RemovalPolicyArgs,
			&job.Comments,
			&job.Deleted,
			&job.AlarmLimit,
			&job.Alarmed,
		)
		if err != nil {
			return jobs, err
		}

		jobs = append(jobs, job)
	}

	return jobs, err
}
