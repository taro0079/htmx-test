package repository

type IUserRepository interface {
    Save() 
    Find(userName string)
}
