package model

type Entity struct {
	ID   string `json:"id" firestore:"-"`
	Name string `json:"name" firestore:"name"`

	// Deleted is used for logical deletion
	Deleted bool `json:"-" firestore:"deleted"`
}
