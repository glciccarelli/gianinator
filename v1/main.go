package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fLog, err := os.OpenFile("gianinator.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(fLog)

	var filePath string
	for {
		fmt.Println("Please provide file path")
		_, err := fmt.Scanln(&filePath)
		if err != nil {
			log.Fatal("Failed to read filePath")
		} else {
			break
		}
	}

	var action string
	for {

		fmt.Println("Select an option: ACTAS[A], CAMPUS[C]")
		_, err := fmt.Scanln(&action)
		if err != nil {
			log.Fatal("Failed to read filePath")
		} else if strings.ToUpper(action) == "A" {
			fmt.Println("Reading csv..")
			err = parseCSVActas(filePath)
			if err != nil {
				log.Print("Failed to read csv")
			}

			break
		} else if strings.ToUpper(action) == "C" {
			fmt.Println("Reading csv..")
			err = parseCSVCampus(filePath)
			if err != nil {
				log.Print("Failed to read csv")
			}
		}
		log.Fatal("a not valid option was specified")
	}
}

func parseCSVActas(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed opening file")

		return err
	}
	log.Printf("Parsing csv actas...")
	defer closeFile(f)

	filename := f.Name()
	filename = fmt.Sprintf("./%s.txt", filename)
	file, _ := os.Create(filename)
	defer func() {
		closeFile(file)

	}()

	lines, _ := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i == 0 {
			continue
		}
		name := line[0]
		dni := line[1]
		phone := line[2]
		email := line[3]
		course := line[4]
		course = getCourse(course)
		endDate := line[6] // 2021-12-31
		score := line[7]

		script := fmt.Sprintf("document.forms['form']['dni'].value ='%s'\ndocument.forms['form']['nombre'].value ='%s'\ndocument.forms['form']['telefono'].value ='%s'\ndocument.forms['form']['idcurso'].value = %s\ndocument.forms['form']['calificacion'].value ='%s'\ndocument.forms['form']['fecha_finaliza'].value='%s'\ndocument.forms['form']['mail'].value='%s'\n\n\n", dni, name, phone, course, score, endDate, email)

		_, err := file.Write([]byte(script))
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func parseCSVCampus(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Failed opening file")

		return err
	}
	log.Printf("Parsing csv campus...")
	defer closeFile(f)

	filename := f.Name()
	filename = fmt.Sprintf("./%s.txt", filename)
	file, _ := os.Create(filename)
	defer func() {
		closeFile(file)

	}()

	lines, _ := csv.NewReader(f).ReadAll()
	for i, line := range lines {
		if i == 0 {
			continue
		}

		lastname := line[0]
		firstname := line[1]
		email := line[2]
		username := line[3]
		password := line[3]

		scriptOnClick := `document.getElementsByName("submit_plus")[0].click()`

		script := fmt.Sprintf("document.forms['user_add']['lastname'].value ='%s'\ndocument.forms['user_add']['firstname'].value ='%s'\ndocument.forms['user_add']['email'].value ='%s'\ndocument.forms['user_add']['username'].value ='%s'\ndocument.forms['user_add']['password[password]'].value ='%s'\ndocument.forms['user_add']['password[password_auto]'].value = '0'\n%s\n\n\n\n", lastname, firstname, email, username, password, scriptOnClick)

		_, err := file.Write([]byte(script))
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

/*
EXTRACCIONISTA 13
FARMACIA 31
ENFERMERIA DOMICILIARIA 232
SUELDOS 37
ASISTENTE LABORATORIO 14
CELULARES 319
SECRETARIADO MEDICO 970
PRESTACIONES 1562
*/
func getCourse(course string) string {
	courseValue := ""
	course = strings.ToUpper(course)

	switch course {
	case "EXTRACCIONISTA":
		courseValue = "13"
	case "FARMACIA":
		courseValue = "31"
	case "ENFERMERIA DOMICILIARIA":
		courseValue = "232"
	case "SUELDOS":
		courseValue = "37"
	case "ASISTENTE LABORATORIO":
		courseValue = "14"
	case "CELULARES":
		courseValue = "319"
	case "SECRETARIADO MEDICO":
		courseValue = "970"
	case "PRESTACIONES MEDICAS":
		courseValue = "1562"
	}

	return courseValue
}

func closeFile(f *os.File) {
	fmt.Println("closing", f.Name())
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
