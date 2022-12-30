package tickets

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	filePath = "./tickets.csv"
)

type Ticket struct {
}

func GetTotalTickets(destination string) (total int, err error) {

	total = 0

	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line

		//If the register is equals to the destination, plus one over total
		if rec[3] == destination {
			total++
		}
		//fmt.Printf("%+v\n", rec[3])
	}

	return total, nil
}

// How many persons fly at certain period
func GetCountByPeriod(time string) (int, error) {
	total := 0

	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line

		lim_inf, lim_sup, err := getPeriod(time)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Printf("Mira el valor primero: %+v, Mira el valor después: %+v\n", rec[4], strings.Split(rec[4], ":")[0])

		//If the register is equals to the destination, plus one over total
		hour, err := strconv.Atoi(strings.Split(rec[4], ":")[0])
		if err != nil {
			log.Fatal(err)
		}
		if hour >= lim_inf && hour < lim_sup {
			//fmt.Printf("Hour: %d\n", hour)
			total++
		}
	}
	return total, nil
}

func getPeriod(time string) (lim_inf int, lim_sup int, err error) {
	switch time {
	case "madrugada":
		return 0, 7, nil
	case "mañana":
		return 7, 13, nil
	case "tarde":
		return 13, 20, nil
	case "noche":
		return 20, 24, nil
	default:
		return 0, 0, errors.New("el periodo de tiempo enviado es incorrecto")
	}
}

func AverageDestination(destination string, total int) (average float64, err error) {
	numberPassagers, err := GetTotalTickets(destination)

	if err != nil {
		return 0, err
	}
	average = (float64(numberPassagers) / float64(total)) * 100
	return average, nil
}

func LenData() (int, error) {
	openfile, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	filedata, err := csv.NewReader(openfile).ReadAll()
	if err != nil {
		return 0, err
	}
	totalQuestions := len(filedata)

	return totalQuestions, nil

}
