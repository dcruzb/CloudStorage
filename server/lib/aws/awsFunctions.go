package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

type CloudFunctionsImpl struct {
}

type Regions struct {
	br_sp Region `json:"South America (Sao Paulo)"`
}

type Region struct {
	tag_storage ServiceDetail `json:"Tags Storage per Tag Mo"`
}

type ServiceDetail struct {
	price float64
}

func (CloudFunctionsImpl) Price(size float64) float64 {

	jsonFile, err := os.Open("data.json")

	if err != nil {
		//Caso tenha tido erro, ele é apresentado na tela
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValueJSON, _:= ioutil.ReadAll(jsonFile)

	//s := make([]string, 3)
	regions := Regions{}

	json.Unmarshal(byteValueJSON, &regions)

	//fmt.Println(regions.Name)
	price := size * regions.br_sp.tag_storage.price

	fmt.Println(price)

	return price
}

func (CloudFunctionsImpl) Availability() bool {

	return nil
}


