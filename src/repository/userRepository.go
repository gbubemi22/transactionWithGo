package repository



import (
	"github.com/jinzhu/gorm"
	
	"github.com/gbubemi22/transaction/src/model"
 )


 type UserRepository struct {
	db *gorm.DB
 }
 func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
 }


 
 

//  func NewUserRepository(db *gorm.DB, firstName string, lastName string, email string, phone string, password string) *UserRepository {

// 	user := &model.User{
// 		First_name: firstName,
// 		Last_name:  lastName,
// 		Email:      email,
// 		Phone:      phone,
// 		Password: password,
// 	 }
// 	 db.Create(user)

// 	return &UserRepository{db}
//  }
 
 // CreateUser creates a new user record in the database.
 func (ur *UserRepository) CreateUser(user *model.User) error {
	return ur.db.Create(user).Error
 }
 
   

 // GetUserByID retrieves a user record by ID from the database.
 func (ur *UserRepository) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	err := ur.db.First(user, id).Error
	return user, err
 }
 
 func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.Where("email = ?", email).First(user).Error
	return user, err
 }

 func (ur *UserRepository) GetUserByPhone(phone string) (*model.User, error) {
	user := &model.User{}
	err := ur.db.Where("phone = ?", phone).First(user).Error
	return user, err
 }
 
 // UpdateUser updates an existing user record in the database.
 func (ur *UserRepository) UpdateUser(user *model.User) error {
	return ur.db.Save(user).Error
 }
 
 // DeleteUser deletes a user record from the database.
 func (ur *UserRepository) DeleteUser(user *model.User) error {
	return ur.db.Delete(user).Error
 }


 // GetAllUsers retrieves a list of all users from the database.
func (ur *UserRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := ur.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}



