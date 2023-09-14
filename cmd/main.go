package main

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"learn.go/ml/subh/pkg/models"
	"learn.go/ml/subh/pkg/utils"
)

// https://raw.githubusercontent.com/ageron/handson-ml/master/datasets/housing/housing.tgz

var (
	DATA_FILE = "dataset/housing.tgz"
	URL       = "https://raw.githubusercontent.com/ageron/handson-ml/master/datasets/housing/housing.tgz"
)

func ReadCSVAndPrint[T any](file io.Reader, parse func([]string) T) (data []T) {
	log.Info("invoked csv reader")
	csvReader := csv.NewReader(file)
	log.Info("invoked csv reader")

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			log.Debug("eof reached")
			break
		}

		if err != nil {
			log.Errorf("error while reading csv file: %v", err)
			return
		}
		data = append(data, parse(record))
	}

	// records, err := csvReader.ReadAll()
	// if err != nil {
	// 	log.Errorf("error reading csv file: %v", err)
	// }
	// log.Debugf("value : %v", records)
	return
}

// func SimpleCSVRead(file io.Reader) {
// 	// prepare csv reader
// 	csvReader := csv.NewReader(file)

// 	// read record one by one
// 	for {
// 		record, err := csvReader.Read()
// 		if err == io.EOF {
// 			log.Debug()
// 		}
// 	}
// }

func main() {
	ll, _ := log.ParseLevel("debug")
	log.SetLevel(ll)
	if err := utils.DownlowdFile(URL, DATA_FILE); err == nil {
		utils.ExtractFile(DATA_FILE, "dataset")
		file, err := os.Open("dataset/housing.csv")
		if err != nil {
			log.Errorf("error in opening file %v", err)
		}
		defer file.Close()

		records := ReadCSVAndPrint(file, func(record []string) models.House {
			var h models.House = models.House{}
			for id, data := range record {
				if id == 9 {
					h.Ocean_proximity = data
					continue
				}

				data, err := strconv.ParseFloat(data, 64)
				if err != nil {
					h.ParseError = err
					break
				}
				switch id {
				case 0:
					h.Longitude = data
				case 1:
					h.Latitude = data
				case 2:
					h.Housing_median_age = data
				case 3:
					h.Total_rooms = data
				case 4:
					h.Total_bedrooms = data
				case 5:
					h.Population = data
				case 6:
					h.Households = data
				case 7:
					h.Median_income = data
				case 8:
					h.Median_house_value = data

				}
			}
			return h
		})

		for id, record := range records {
			if record.ParseError != nil {
				log.Errorf("[%d]== %v", id, record)
			}
		}
	}
}
