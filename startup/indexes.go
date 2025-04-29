package startup

import (
	auth "go-tutorial/api/auth/model"
	blog "go-tutorial/api/blog/model"
	contact "go-tutorial/api/contact/model"
	user "go-tutorial/api/user/model"
	"go-tutorial/arch/mongo"
)

func EnsureDbIndexes(db mongo.Database) {
	go mongo.Document[auth.Keystore](&auth.Keystore{}).EnsureIndexes(db)
	go mongo.Document[auth.ApiKey](&auth.ApiKey{}).EnsureIndexes(db)
	go mongo.Document[user.User](&user.User{}).EnsureIndexes(db)
	go mongo.Document[user.Role](&user.Role{}).EnsureIndexes(db)
	go mongo.Document[blog.Blog](&blog.Blog{}).EnsureIndexes(db)
	go mongo.Document[contact.Message](&contact.Message{}).EnsureIndexes(db)
}
