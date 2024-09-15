// RAW SQL

type Result struct {
	Name  string
	Total int
}

var result Result
db.Raw("SELECT name, SUM(sales) as total FROM sales GROUP BY name").Scan(&result)

// JOIN

var users []User
db.Joins("Company").Find(&users)



// SCOPES

func OlderThan(age int) func(db *gorm.DB) *gorm.DB {
    return func(db *gorm.DB) *gorm.DB {
        return db.Where("age > ?", age)
    }
}

db.Scopes(OlderThan(30)).Find(&users)


// HOOKS
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
    u.UUID = uuid.New().String()
    return
}

// BEST PRATICES FOR GORM
// Contextual Logging: Integrate logging with your database operations.

// Connection Pooling: Configure max open connections and idle connections.