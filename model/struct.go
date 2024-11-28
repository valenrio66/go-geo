package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type GeoPoint struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Properties struct {
	OsmID   int    `json:"osm_id"`
	Name    string `json:"name"`
	Highway string `json:"highway"`
}

type Road struct {
	ID         primitive.ObjectID `bson:"_id"`
	Type       string             `bson:"type"`
	Geometry   GeoPoint           `bson:"geometry"`
	Properties Properties         `bson:"properties"`
}
