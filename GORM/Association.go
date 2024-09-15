package main

// one to many

type Company struct {
	ID        uint
	Name      string
	Employees []Employee
}

type Employee struct {
	ID        uint
	Name      string
	CompanyID uint
}

func main() {
	// Auto migrate
	db.AutoMigrate(&Company{}, &Employee{})
}

// Creating Records with Associations

company := Company{
	Name: "Tech Corp",
	Employees: []Employee{
		{Name: "John Doe"},
		{Name: "Jane Smith"},
	},
}

db.Create(&company)

// Querying with Preload

var company Company
db.Preload("Employees").First(&company, company.ID)
