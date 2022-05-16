package tools

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"rest/pkg/shared/database"
	"strings"
	"time"
)

type Migrator struct {
	MigrationsDir string
	db            *database.Database
}

func NewMigrator(dir string) *Migrator {
	db := database.GetDatabase()
	m := &Migrator{
		MigrationsDir: dir,
		db:            db,
	}
	return m
}

func (m *Migrator) getMigrationFolder() string {
	files, err := ioutil.ReadDir(m.MigrationsDir)
	if err != nil {
		os.Exit(1)
	}
	var lastFolder string
	for i := len(files) - 1; i >= 0; i-- {
		if files[i].IsDir() {
			lastFolder = files[i].Name()
			break
		}
	}
	if lastFolder == "" {
		lastFolder = "v0"
	}
	folderName := fmt.Sprintf("%s/%s/", m.MigrationsDir, lastFolder)

	os.MkdirAll(folderName, os.ModePerm)

	return folderName
}

func (m *Migrator) Up() error {
	files := m.getNotAppliedMigrations()
	m.runMigrations(files)
	return nil
}

func (m *Migrator) Down() error {
	files := m.listAllMigrationFiles("down")
	m.readFiles(files)
	return nil
}

func (m *Migrator) New(title string) error {
	// migration pattern
	// [timestamp]__[description].[up|down].sql
	timestamp := time.Now().Unix()
	description := strings.ReplaceAll(title, " ", "_")
	description = strings.ToLower(description)
	migrationsFolder := m.getMigrationFolder()
	filenameUp := fmt.Sprintf("%s%d__%s.up.sql", migrationsFolder, timestamp, description)
	filenameDown := fmt.Sprintf("%s%d__%s.down.sql", migrationsFolder, timestamp, description)
	comment := []byte(fmt.Sprintf("-- %s", title))
	err := ioutil.WriteFile(filenameUp, comment, 0755)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filenameDown, comment, 0755)
	if err != nil {
		return err
	}
	return nil
}

func (m *Migrator) listAllMigrationFiles(typ string) []string {
	folders, err := ioutil.ReadDir(m.MigrationsDir)
	if err != nil {
		os.Exit(1)
	}
	var allFiles []string
	for _, folder := range folders {
		if folder.IsDir() {
			folderName := fmt.Sprintf("%s/%s/", m.MigrationsDir, folder.Name())
			files, err := ioutil.ReadDir(folderName)
			if err != nil {
				os.Exit(1)
			}
			for _, file := range files {
				if strings.HasSuffix(file.Name(), typ+".sql") {
					allFiles = append(allFiles, fmt.Sprintf("%s%s", folderName, file.Name()))
				}
			}
		}
	}
	return allFiles
}

func (m *Migrator) getNotAppliedMigrations() []string {
	lastmigration := "1651366443__sixth.up.sql"
	files := m.listAllMigrationFiles("up")
	var notApplied []string
	for _, file := range files {
		if strings.Compare(lastmigration, path.Base(file)) < 0 {
			notApplied = append(notApplied, file)
		}
	}
	return notApplied
}

func (m *Migrator) readFiles(files []string) {
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			os.Exit(1)
		}
		fmt.Println(string(content))
	}
}

func (m *Migrator) runMigrations(files []string) {
	tx, _ := m.db.Begin()
	tx.RollbackUnlessCommitted()
	for _, file := range files {
		content, err := ioutil.ReadFile(file)
		if err != nil {
			os.Exit(1)
		}
		_, err = tx.Exec(string(content))
		if err != nil {
			tx.Rollback()
			os.Exit(1)
		}
	}
	tx.Commit()
}

func (m *Migrator) Status() error {
	return nil
}

func (m *Migrator) Migrate() error {
	return nil
}
