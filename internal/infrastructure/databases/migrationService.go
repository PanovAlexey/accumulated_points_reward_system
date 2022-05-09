package databases

import (
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"log"
	"os"
)

type migrationService struct {
	db *sqlx.DB
}

const migrationUpFilesPath = `internal/infrastructure/databases/migrations/up`
const migrationDownFilesPath = `internal/infrastructure/databases/migrations/down`
const migrationFilePermissions = 0777

func GetMigrationService(db *sqlx.DB) *migrationService {
	migrationService := migrationService{}
	migrationService.db = db

	return &migrationService
}

func (service migrationService) MigrateUp() {
	filesPath := service.getMigrationFilesPath(migrationUpFilesPath)

	for _, filePath := range filesPath {
		migrationFileContent, err := service.getMigrationFileContent(filePath)

		if err != nil {
			log.Println(`impossible to get content of migration file by path: ` + filePath)
		}

		isApplied := service.applyMigrate(migrationFileContent)

		if isApplied {
			log.Println(`migration applied successfully (` + filePath + `)`)
		}
	}
}

func (service migrationService) MigrateDown() {
	filesPath := service.getMigrationFilesPath(migrationDownFilesPath)

	for _, filePath := range filesPath {
		migrationFileContent, err := service.getMigrationFileContent(filePath)

		if err != nil {
			log.Println(`impossible to get content of migration file by path: ` + filePath)
		}

		service.applyMigrate(migrationFileContent)
	}
}

func (service migrationService) getMigrationFileContent(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, migrationFilePermissions)

	if err != nil {
		log.Println(`error with opening migration file ` + filePath + ` ` + err.Error())
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Println(`error with closing migrate file: ` + filePath + ` ` + err.Error())
		}
	}()

	contentString := ``
	content, err := ioutil.ReadAll(file)

	if err == nil {
		contentString = string(content)
	}

	return contentString, err
}

func (service migrationService) applyMigrate(sqlQuery string) bool {
	_, err := service.db.Exec(sqlQuery)

	if err != nil {
		log.Println(`error while applying migration. ` + err.Error())
		return false
	}

	return true
}

func (service migrationService) getMigrationFileFullName(path string, fileName string) string {
	return path + `/` + fileName
}

func (service migrationService) getMigrationFilesPath(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)
	filesPath := make([]string, len(files))

	if err != nil {
		log.Println("error while getting migration files by dir: " + dirPath + ". " + err.Error())
	}

	for key, file := range files {
		if !file.IsDir() {
			filesPath[key] = service.getMigrationFileFullName(dirPath, file.Name())
		}
	}

	return filesPath
}
