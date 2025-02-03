package parsing

import (
	"AISale/database/models/repos/phone_repos"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func parseBool(cell string) bool {
	return cell == "1" || cell == "true" || cell == "TRUE"
}

func parseFloat(cell string) float64 {
	val, err := strconv.ParseFloat(cell, 64)
	if err != nil {
		return 0
	}
	return val
}

func parseInt(cell string) int {
	val, err := strconv.Atoi(cell)
	if err != nil {
		return 0
	}
	return val
}

func readCSV(filePath string) ([]Phone, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV header: %w", err)
	}

	var phones []Phone
	for {
		record, err := reader.Read()
		if err != nil {
			break
		}

		if len(record) < 22 { // Проверка на наличие всех колонок
			continue
		}

		phone := Phone{
			PhoneName:        record[0],
			Brand:            record[1],
			OS:               record[2],
			Inches:           parseFloat(record[3]),
			Resolution:       record[4],
			Battery:          parseInt(record[5]),
			BatteryType:      record[6],
			RAM:              parseInt(record[7]),
			AnnouncementDate: record[8],
			Weight:           parseInt(record[9]),
			Storage:          parseInt(record[10]),
			Video720p:        parseBool(record[11]),
			Video1080p:       parseBool(record[12]),
			Video4K:          parseBool(record[13]),
			Video8K:          parseBool(record[14]),
			Video30fps:       parseBool(record[15]),
			Video60fps:       parseBool(record[16]),
			Video120fps:      parseBool(record[17]),
			Video240fps:      parseBool(record[18]),
			Video480fps:      parseBool(record[19]),
			Video960fps:      parseBool(record[20]),
			PriceUSD:         parseFloat(record[21]),
		}
		phones = append(phones, phone)
	}
	return phones, nil
}

func ParsePhonesCSV(filePath string) error {
	phones, err := readCSV(filePath)
	if err != nil {
		return errors.New("failed to read CSV file: " + err.Error())
	}

	for _, phone := range phones {
		err = phone_repos.Create(phone)
	}
	if err != nil {
		return errors.New("failed to insert phones into database: " + err.Error())
	}

	return nil
}
