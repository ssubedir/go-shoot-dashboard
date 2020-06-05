package handlers

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EnqueuedTask struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	TaskID     string             `bson:"task_id,omitempty" json:"task_id"`
	TaskName   string             `bson:"task_name,omitempty" json:"task_name"`
	TaskStatus string             `bson:"task_status,omitempty" json:"task_status"`
	TaskTime   time.Time          `bson:"task_enqueued_time,omitempty" json:"task_enqueued_time"`
}

type ScheduledTask struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	TaskID            string             `bson:"task_id,omitempty" json:"task_id"`
	TaskName          string             `bson:"task_name,omitempty" json:"task_name"`
	TaskStatus        string             `bson:"task_status,omitempty" json:"task_status"`
	TaskCreatedTime   time.Time          `bson:"task_created_time,omitempty" json:"task_created_time"`
	TaskScheduledTime time.Time          `bson:"task_scheduled_time,omitempty" json:"task_scheduled_time"`
}

type SucceededTask struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	TaskID            string             `bson:"task_id,omitempty" json:"task_id"`
	TaskName          string             `bson:"task_name,omitempty" json:"task_name"`
	TaskStatus        string             `bson:"task_status,omitempty" json:"task_status"`
	TaskDuration      int                `bson:"task_duration,omitempty" json:"task_duration"`
	TaskType          string             `bson:"task_type,omitempty" json:"task_type"`
	TaskSucceededTime time.Time          `bson:"task_succeeded_time,omitempty" json:"task_succeeded_time"`
}
