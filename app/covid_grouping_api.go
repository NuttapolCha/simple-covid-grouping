package app

import (
	"context"
	"net/http"

	"github.com/NuttapolCha/simple-covid-grouping/custom_error"
	"github.com/spf13/viper"
)

type ageGroup string

const (
	junior  ageGroup = "0-30"
	mid     ageGroup = "31-60"
	senior  ageGroup = "61+"
	unknown ageGroup = "N/A"
)

type CovidCaseSummary struct {
	Province map[string]int   `jsno:"Province"`
	AgeGroup map[ageGroup]int `json:"AgeGroup"`
}

type CovidCaseRespData struct {
	ConfirmDate    string      `json:"ConfirmDate"`
	No             interface{} `json:"No"`
	Age            *int        `json:"Age"`
	Gender         string      `json:"Gender"`
	GenderEn       string      `json:"GenderEn"`
	Nation         string      `json:"Nation"`
	NationEn       string      `json:"NationEn"`
	Province       *string     `json:"Province"`
	ProvinceId     int         `json:"ProvinceId"`
	District       string      `json:"District"`
	ProvinceEn     string      `json:"ProvinceEn"`
	StatQuarantine int         `json:"StatQuarantine"`
}

type CovidCaseResp struct {
	Data []CovidCaseRespData `json:"Data"`
}

func (externalData *CovidCaseResp) Summarize() interface{} {
	ret := &CovidCaseSummary{
		Province: make(map[string]int),
		AgeGroup: make(map[ageGroup]int),
	}
	for _, data := range externalData.Data {
		// group by province
		provinceKey := "N/A"
		if data.Province != nil {
			provinceKey = *data.Province
		}
		ret.Province[provinceKey]++

		// group by age
		ageKey := unknown
		if data.Age != nil {
			age := *data.Age
			switch {
			case age >= 0 && age <= 30:
				ageKey = junior
			case age >= 31 && age <= 60:
				ageKey = mid
			case age >= 61:
				ageKey = senior
			default:
				panic("unrecognized age, this should not be occurred")
			}
		}
		ret.AgeGroup[ageKey]++
	}
	return ret
}

func (app *App) GetCovidCasesSummary(ctx context.Context) (interface{}, error) {
	dataSourceUrl := viper.GetString("ExternalAPIs.CovidCasesSource")

	// get data from external
	externalData := &CovidCaseResp{}
	if err := app.conn.Get(ctx, dataSourceUrl, externalData); err != nil {
		app.Logger.Errorf("could not request COVID data from source because :%v", err)
		return nil, &custom_error.UserError{
			Message:    err.Error(),
			StatusCode: http.StatusBadGateway,
		}
	}

	return externalData.Summarize(), nil
}
