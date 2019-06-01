package lib

import (
	cloudLib "CloudStorage/server/lib"
	"encoding/json"
	"fmt"
	"github.com/dcbCIn/MidCloud/lib"
	"io/ioutil"
	"os"
)

type Regions struct {
	Br_sp Region `json:"South America (Sao Paulo)"`
}

type Region struct {
	Tag_storage ServiceDetail `json:"Tags Storage per Tag Mo"`
}

type ServiceDetail struct {
	Price float64
}

type AwsFunctions struct {
}

func (AwsFunctions) SendFile(file *os.File) (createdFile cloudLib.CloudFile, err error) {
	panic("implement me")
}

func (AwsFunctions) GetFile() (file *os.File, err error) {
	panic("implement me")
}

func (AwsFunctions) List(path string) (files []cloudLib.CloudFile, err error) {
	panic("implement me")
}

func (AwsFunctions) Price(size float64) float64 {

	jsonFile, err := os.Open("data.json")

	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		lib.PrintlnError("Erro ao abrir json. Erro", err)
	}

	defer jsonFile.Close()

	byteValueJSON, _ := ioutil.ReadAll(jsonFile)

	//s := make([]string, 3)
	regions := Regions{}

	err = json.Unmarshal(byteValueJSON, &regions)
	if err != nil {
		lib.PrintlnError("Erro ao realizar unmarshal. Erro:", err)
	}

	//fmt.Println(regions.Name)
	price := size * regions.Br_sp.Tag_storage.Price

	fmt.Println(price)

	return price
}

func (AwsFunctions) Availability() bool {
	return false
}
