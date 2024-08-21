package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	repositories "tender/internal/repositories/prod"
)

type SpecificTaxesService struct {
	Repo *repositories.SpecificTaxesRepository
}

// Function to get the exchange rate from KZT to USD using the amount
func getExchangeRate(amount float64) (float64, error) {
	apiKey := "uBEwbRsUj1jy3mbTbiK6ta8lNakAgg14" // Your API key
	url := fmt.Sprintf("https://api.apilayer.com/exchangerates_data/convert?to=USD&from=KZT&amount=%f", amount)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}
	req.Header.Set("apikey", apiKey)

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	var data struct {
		Result float64 `json:"result"`
	}
	if err := json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Result, nil
}

func (s *SpecificTaxesService) GetSpecificTaxes(ctx context.Context, productionunit string, year int) (map[string]map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetSpecificTaxes(ctx, productionunit, year)
	if err != nil {
		return nil, err
	}

	var totalTaxAmount float64
	for _, tax := range totalSumSummary {
		totalTaxAmount += tax
	}

	exchangeRate, err := getExchangeRate(totalTaxAmount)
	if err != nil {
		return nil, err
	}

	result := make(map[string]map[string]float64)
	for name, tax := range totalSumSummary {
		result[name] = map[string]float64{
			"KZT": tax,
			"USD": tax * exchangeRate / totalTaxAmount, // Calculate the equivalent in USD
		}
	}

	return result, nil
}
