package medods

type User struct {
	GUID    string `bson:"guid, omitempty"`
	Refresh string `bson:"refresh, omitempty"`
}
