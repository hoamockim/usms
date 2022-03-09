package migration

import (
	/*	"encodingFile/json"
		"fmt"
		"github.com/go-sql-driver/mysql"
		"go.uber.org/zap"
		"io/ioutil"
		"os"
		"strings"
		"time"
		"usms/app"

		database "usms/db"
		cfg "usms/pkg/configs"
		"usms/pkg/errors"*/
	"fmt"
	"os"
	"strings"
	"usms/app"
)

func main() {
	/*db := database.NewDB()
	if err := db.Open(cfg.GetDb().ConnectionString()); err != nil {
		fatalf("Creating connection to DB: %v", err)
	}

	jobMigrate, errCode := initJob(fmt.Sprintf("./resources/%v.json", cfg.GetDb().MigratePath))
	if errCode != nil {
		fatalf("Cannot init job, %v, %v", errCode.Code(), errCode.Message())
	}

	errCode = runJob(db, jobMigrate)
	if errCode != nil {
		fatalf("run job err %v, %v", errCode.Code(), errCode.Message())
	}
	zap.S().Infof("Job %v is success", jobMigrate.JobId)*/
}

/*func initJob(filePath string) (jM *job.JobMigrate, errCode *errors.ErrorCode) {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		errCode = errors.ErrFile
		errCode.SetMessageError(err.Error())
		return
	}
	err = json.Unmarshal([]byte(file), &jM)
	if err != nil {
		errCode = errors.ErrFile
		errCode.SetMessageError(err.Error())
		return
	}
	return
}*/

/*func runJob(orm database.DBAdapter, jM *job.JobMigrate) (errCode *errors.ErrorCode) {
	if strings.TrimSpace(jM.JobId) == "" {
		errCode = errors.ErrData
		errCode.SetMessageError("Job id is empty")
		return nil
	}
	if jM != nil && len(jM.Migrates) > 0 {
		tx := orm.Begin()
		defer func() {
			tx.RollbackUselessCommitted()
		}()

		var jMEntity job.Migration
		tx.Gormer().Table("migrations").
			Where("job_id = ?", jM.JobId).First(&jMEntity)
		if jMEntity.JobId == jM.JobId {
			errCode = errors.ErrData
			errCode.SetMessageError(fmt.Sprintf("run migrate err, job id is existed %v", jM.JobId))
			return
		}
		for _, migrate := range jM.Migrates {
			if strings.TrimSpace(migrate.FullName) != "" && strings.TrimSpace(migrate.Sql) != "" {
				var entity job.Migration
				err := tx.Gormer().Table("migrations").
					Where("name = ?", migrate.FullName).First(&entity).Error
				if err != nil {
					if sqlError, _ := err.(*mysql.MySQLError);
					//error number 1146: table is not exist
						sqlError != nil && sqlError.Number != 1146 {
						errCode = errors.ErrData
						errCode.SetMessageError(fmt.Sprintf("run migrate err %v", err.Error()))
						return
					}
				}
				if entity.IdMigration > 0 {
					continue
				}
				entity = job.Migration{
					JobId:      jM.JobId,
					FullName:       migrate.FullName,
					Statements: migrate.Sql,
					Status:     "update",
					StartTime:  time.Now().UTC(),
					EndTime:    time.Now().UTC(),
				}

				err = tx.Gormer().Exec(migrate.Sql).Error
				if err != nil {
					errCode = errors.ErrData
					errCode.SetMessageError(fmt.Sprintf("run migrate err %v", err.Error()))
					return
				}

				err = tx.Gormer().Table("migrations").Save(&entity).Error
				if err != nil {
					errCode = errors.ErrData
					errCode.SetMessageError(fmt.Sprintf("run migrate err %v", err.Error()))
					tx.RollbackUselessCommitted()
					return
				}

			} else {
				errCode = errors.ErrData
				errCode.SetMessageError(fmt.Sprintln("Job is not valid :%v, %v", migrate.FullName, migrate.Sql))
				return
			}
			zap.S().Infof("migrate name %v success, detail: %v ", migrate.FullName, migrate.Sql)
		}

		tx.Commit()
	}
	return
}*/

func fatalf(format string, values ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	_, _ = fmt.Fprintf(os.Stderr, "\n")
	_, _ = fmt.Fprintf(os.Stderr, "[GIN-debug] [ERROR] "+format, values...)
	os.Exit(1)
}

func init() {
	app.InitLogger()
}
