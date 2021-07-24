pakage models
import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `bson: "_id, omitempty" json:"id"`
	Names string `bson: "names" json:"nombre, omitempty"`
	DateBirth time.Time `bson: "dateBirth" json:"dateBirth, omitempty"`
	Email string `bson: "email" json:"dateBirth"`
	Password string `bson: "password" json:"password, omitempty"`
	Picture string `bson: "picture" json:"picture, omitempty"`
	Banner string `bson: "banner" json:"Banner, omitempty"`
	Residence string `bson: "residence" json:"residence, omitempty"` 
}