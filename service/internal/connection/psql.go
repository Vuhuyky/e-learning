package connection

import (
	"app/internal/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectPostgresql(migrate bool) error {
	var err error
	dns := fmt.Sprintf(
		`
			host=%s
			user=%s
			password=%s
			dbname=%s
			port=%s
			sslmode=disable`,
		conn.Psql.Host,
		conn.Psql.User,
		conn.Psql.Password,
		conn.Psql.Name,
		conn.Psql.Port,
	)

	dbPsql, err = gorm.Open(postgres.Open(dns), &gorm.Config{})

	if migrate {
		errMigrate := dbPsql.AutoMigrate(
			&entity.Budget{},
			&entity.Category{},
			&entity.Chapter{},
			&entity.Course{},
			&entity.CourseRegister{},
			&entity.DocumentLession{},
			&entity.Lession{},
			&entity.Organization{},
			&entity.Profile{},
			&entity.Quizz{},
			&entity.Role{},
			&entity.SaleCourse{},
			&entity.TextNote{},
			&entity.VideoLession{},
			&entity.RolePermission{},
			&entity.Permission{},
			&entity.CourseCategory{},
			&entity.ProcessStream{},
			&entity.Service{},
		)

		if errMigrate != nil {
			log.Println(errMigrate)
			return errMigrate
		}
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
