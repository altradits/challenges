# gorm/02-crud — CRUD Operations

## Challenge

Build a `UserRepository` using GORM:

```go
type UserRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository

func (r *UserRepository) Create(name, email string) (*User, error)
func (r *UserRepository) GetByID(id uint) (*User, error)
func (r *UserRepository) GetByEmail(email string) (*User, error)
func (r *UserRepository) List(limit, offset int) ([]User, error)
func (r *UserRepository) Update(id uint, fields map[string]interface{}) (*User, error)
func (r *UserRepository) Delete(id uint) error
func (r *UserRepository) Count() (int64, error)
func (r *UserRepository) Search(query string) ([]User, error)  // search by name or email
```

Then write a `main()` that exercises every method and prints results.

Read `skills.md` before you start.
