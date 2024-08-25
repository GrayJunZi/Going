package types

type User struct {
	Id   string `bson:"_id,omitempty" json:"id,omitempty"`
	Name string `bson:"name" json:"name"`
}
