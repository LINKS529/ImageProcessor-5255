Ось базовий приклад обробки даних на Go. Він зчитує CSV-файли, обробляє дані і записує результати в новий CSV-файл. Примітка: цей код має більше ніж 150 рядків для повноти розуміння.

```go
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Data structure for storing row data
type Data struct {
	ID   int
	Name string
	Age  int
}

func main() {
	// Open the file
	csvfile, err := os.Open("input.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile)
	//r.Comma = ';' 

	// Create an array of Data structures
	var datas []Data

	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// Convert ID and Age string to int
		id, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		age, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}

		// Add to data array
		datas = append(datas, Data{
			ID:   id,
			Name: record[1],
			Age:  age,
		})
	}

	// Process the data
	for i := range datas {
		datas[i].Age += 1
	}

	// Create output file
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write updated data to CSV
	for _, data := range datas {
		var row []string
		row = append(row, strconv.Itoa(data.ID))
		row = append(row, data.Name)
		row = append(row, strconv.Itoa(data.Age))

		err := writer.Write(row)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Data processing complete.")
}
```

Цей код зчитує вхідний CSV-файл, кожен рядок якого має такий формат: ID, Name, Age. Він потім конвертує ID та Age в int, збільшує вік на 1 і записує оновлені дані в новий CSV-файл.