# Skills for gorm/02-crud

## What You'll Learn

**Previous:** [gorm/01-connect-models](../01-connect-models/skills.md) | **Next:** [gorm/03-associations](../03-associations/skills.md)

**Challenge:** Master the full GORM CRUD API and the repository pattern.

## Create

```go
// Create one
user := &User{Name: "Alice", Email: "alice@example.com"}
db.Create(user)        // user.ID populated
db.Create(user)        // error: duplicate email

// Create in batch
users := []User{{Name: "A"}, {Name: "B"}}
db.Create(&users)

// Create with selected fields only
db.Select("Name", "Email").Create(user)
```

## Read

```go
// Find by primary key
var user User
db.First(&user, 1)          // SELECT * FROM users WHERE id = 1 LIMIT 1
db.First(&user, "id = ?", 1)

// Find all
var users []User
db.Find(&users)

// With conditions
db.Where("name = ?", "Alice").Find(&users)
db.Where("age > ? AND role = ?", 18, "admin").Find(&users)

// With limit and offset
db.Limit(10).Offset(20).Find(&users)

// Count
var count int64
db.Model(&User{}).Count(&count)

// First vs Find
db.First(&user)  // returns error if not found (gorm.ErrRecordNotFound)
db.Find(&users)  // returns empty slice if none found, no error
```

## Update

```go
// Save (updates all fields)
user.Name = "Alice Updated"
db.Save(&user)

// Updates (only specified fields)
db.Model(&user).Updates(User{Name: "Bob"})           // from struct (ignores zero values)
db.Model(&user).Updates(map[string]interface{}{      // from map (updates even zero values)
    "name": "Bob",
    "age": 0,
})

// Update single field
db.Model(&user).Update("name", "Carol")

// Update with condition (no model — updates all matching rows)
db.Model(User{}).Where("role = ?", "guest").Update("role", "user")
```

## Delete

```go
// Soft delete (sets DeletedAt)
db.Delete(&user, 1)

// Hard delete (if using soft delete, use Unscoped)
db.Unscoped().Delete(&user, 1)

// Bulk delete
db.Where("created_at < ?", cutoff).Delete(&User{})
```

## Error Handling

```go
result := db.First(&user, id)
if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return nil, fmt.Errorf("user %d not found", id)
}
if result.Error != nil {
    return nil, result.Error
}
```

## Repository Pattern

```go
type UserRepository struct{ db *gorm.DB }

func (r *UserRepository) Create(name, email string) (*User, error) {
    u := &User{Name: name, Email: email}
    return u, r.db.Create(u).Error
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
    var u User
    err := r.db.First(&u, id).Error
    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, fmt.Errorf("user not found")
    }
    return &u, err
}

func (r *UserRepository) Search(query string) ([]User, error) {
    var users []User
    like := "%" + query + "%"
    err := r.db.Where("name LIKE ? OR email LIKE ?", like, like).Find(&users).Error
    return users, err
}
```

## Documentation

- [GORM CRUD interface](https://gorm.io/docs/crud_interface.html)
- [GORM error handling](https://gorm.io/docs/error_handling.html)

**Next:** [gorm/03-associations](../03-associations/README.md)
