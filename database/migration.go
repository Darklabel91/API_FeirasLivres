package database

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/Darklabel91/API_FeirasLivres/models"
	"log"
	"os"
	"strconv"
)

//migrateCSV migrate the csv file to the database
func migrateCSV(path string) error {
	err := DB.AutoMigrate(&models.Fair{})
	if err != nil {
		return err
	}

	var count int64
	DB.Model(&models.Fair{}).Count(&count)

	if count == 0 {
		fairs, err := readCsvFile(path)
		if err != nil {
			return err
		}

		result := DB.Create(fairs)
		if result.Error != nil {
			return err
		}

		log.Println("INFO: CSV imported on database")

		return nil
	}

	return errors.New("ERROR: csv already migrated")

}

//readCsvFile reads a csv file from a given path returning it as models.Fair
func readCsvFile(filePath string) ([]models.Fair, error) {
	var data2 []models.Fair

	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = ','

	csvData, err := csvR.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	for i, line := range csvData {
		//ignore the header
		if i != 0 {
			id, err := strconv.Atoi(line[0])
			if err != nil {
				return nil, err
			}
			t := models.Fair{
				Id:            id,
				Longitude:     line[1],
				Latitude:      line[2],
				SetCen:        line[3],
				AreaP:         line[4],
				CodDist:       line[5],
				District:      line[6],
				CodSubPref:    line[7],
				SubPref:       line[8],
				RegionFive:    line[9],
				RegionEight:   line[10],
				NameFair:      line[11],
				Record:        line[12],
				Street:        line[13],
				Number:        line[14],
				Neighbourhood: line[15],
				Reference:     line[16],
			}
			data2 = append(data2, t)
		}
	}
	return data2, nil
}
