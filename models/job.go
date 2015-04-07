package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Job struct {
	Id      bson.ObjectId `json:"Id,omitempty" bson:"_id,omitempty"`
	Name    string        `json:"name,omitempty"`
	Created time.Time     `json:"created_at,omitempty"`
	Updated time.Time     `json:"updated_at,omitempty"`
}

// creates a new job for the give Job
// returns an error if the job creation fails
func (ds *DataStore) CreateJob(job *Job) error {
	jobCollection := ds.session.DB(DbName).C(JobsCollection)

	//set the time
	job.Id = bson.NewObjectId()
	job.Created = time.Now()
	job.Updated = time.Now()

	err := jobCollection.Insert(job)
	return err
}

// updates a job
func (ds *DataStore) UpdateJob(job *Job) error {
	jobCollection := ds.session.DB(DbName).C(JobsCollection)

	//set the time
	job.Updated = time.Now()

	err := jobCollection.Update(bson.M{"_id": job.Id}, job)
	return err
}

func (ds *DataStore) FindJobs() []Job {
	jobCollection := ds.session.DB(DbName).C(JobsCollection)
	jobs := []Job{}

	query := jobCollection.Find(bson.M{}).Sort("-created").Limit(50)
	query.All(&jobs)

	return jobs
}
