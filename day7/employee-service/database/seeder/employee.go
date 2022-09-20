package seeder

import (
	"log"
	"time"

	"github.com/EvanKA/alterra-agmc/day6/employee-service/internal/model"
	"gorm.io/gorm"
)

func employeeSeeder(db *gorm.DB) {
	now := time.Now()
	var employees = []model.Employee{
		{
			Fullname:   "Evan Kurnia Alim",
			Email:      "e0337884@u.nus.edu",
			Password:   "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:     1,
			DivisionID: 1,
			Common:     model.Common{ID: 1, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname:   "Claudia Citra",
			Email:      "claudiacitra@gmail.com",
			Password:   "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:     2,
			DivisionID: 1,
			Common:     model.Common{ID: 2, CreatedAt: now, UpdatedAt: now},
		},
		{
			Fullname:   "Natural Language Processing",
			Email:      "nlp@gmail.com",
			Password:   "$2a$10$rfpS/jJ.a5J9seBM5sNPTeMQ0iVcAjoox3TDZqLE7omptkVQfaRwW", // 123abcABC!
			RoleID:     2,
			DivisionID: 2,
			Common:     model.Common{ID: 3, CreatedAt: now, UpdatedAt: now},
		},
	}
	if err := db.Create(&employees).Error; err != nil {
		log.Printf("cannot seed data employees, with error %v\n", err)
	}
	log.Println("success seed data employees")
}
