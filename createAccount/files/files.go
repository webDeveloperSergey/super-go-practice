package files

import (
	"fmt"
	"os"
)

type DbJson struct {
	filename string
}

// Обращаемся к структуре DbJson и создаем функцию конструктор newDbJson, которая инитит новый DbJson с именем
func NewDbJson (name string) *DbJson {
	return &DbJson{
		filename: name,
	}
}

func (db *DbJson) Read() ([]byte, error) {
	data, err := os.ReadFile(db.filename)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil
}

func (db *DbJson) Write(content []byte) {
	file, err := os.Create(db.filename)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Запись успешна")
}