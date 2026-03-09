package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    TaskName  string             `bson:"taskName" json:"taskName"`
    AssignedTo primitive.ObjectID `bson:"assignedTo" json:"assignedTo"`
    Remarks   string             `bson:"remarks" json:"remarks"`
    Priority  string             `bson:"priority" json:"priority"`
    Status    string             `bson:"status" json:"status"`
}