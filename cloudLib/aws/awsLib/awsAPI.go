package awsLib

import (
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

type Aws struct {
}

func (Aws) Price(size float64) float64 {
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
