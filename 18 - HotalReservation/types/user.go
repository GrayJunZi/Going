package types

type User struct {
	Id   string `bson:"_id" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
}
