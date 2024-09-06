package services

import (
	"context" // Ensure this is correct
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	repositories_dmart "tender/internal/repositories/dmart"
)

type InvestmentsDashService struct {
	Repo *repositories_dmart.InvestmentsDashRepository
}

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

func (s *InvestmentsDashService) GetInvestmentsDash(ctx context.Context, companyName, finReportType string, reportYear int) (map[int]float64, error) {
	summary, err := s.Repo.GetInvestmentsDash(ctx, companyName, finReportType, reportYear)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashOilProduction(ctx context.Context, companyname, productionunit, finreporttype string, reportYear int) (map[int]float64, error) {
	summary, err := s.Repo.GetInvestmentsDashOilProduction(ctx, companyname, productionunit, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashSpecificRevenue(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportYear int) (map[int]float64, error) {
	summary, err := s.Repo.GetInvestmentsDashSpecificRevenue(ctx, currencyunit, companyname, productionunit, finreporttype, reportYear)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashROA(ctx context.Context, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	roaSummary, err := s.Repo.GetInvestmentsDashROA(ctx, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return roaSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashNetProfitMargin(ctx context.Context, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	avgNetProfitToRevenueSummary, err := s.Repo.GetInvestmentsDashNetProfitMargin(ctx, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return avgNetProfitToRevenueSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashSpecificNetProfit(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportyear int) (map[int]float64, error) {
	avgNetProfitToRevenueSummary, err := s.Repo.GetInvestmentsDashSpecificNetProfit(ctx, currencyunit, companyname, productionunit, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return avgNetProfitToRevenueSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashRevenue(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	revenueSummary, err := s.Repo.GetInvestmentsDashRevenue(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return revenueSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashOperatingProfit(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	operatingProfitSummary, err := s.Repo.GetInvestmentsDashOperatingProfit(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return operatingProfitSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashEBITDA(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashEBITDA(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashNetProfit(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashNetProfit(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashTotalTaxes(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashTotalTaxes(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashTaxBurden(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashTaxBurden(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashSpecificTaxes(ctx context.Context, currencyunit, companyname, productionunit, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashSpecificTaxes(ctx, currencyunit, companyname, productionunit, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashAssets(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashAssets(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashCapital(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashCapital(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashLiabilities(ctx context.Context, currencyunit, companyname, finreporttype string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashLiabilities(ctx, currencyunit, companyname, finreporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashSpecificNetProfitGraph(ctx context.Context, currencyunit, productionunit, reporttype string, reportyear int) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashSpecificNetProfitGraph(ctx, currencyunit, productionunit, reporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
} // s

func (s *InvestmentsDashService) GetInvestmentsDashROAGraph(ctx context.Context, reporttype string, reportyear int) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashROAGraph(ctx, reporttype, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashCurrentRatio(ctx context.Context, reporttype, company, currency, unit string, reportyear int) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashCurrentRatio(ctx, reporttype, company, currency, unit, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashCF(ctx context.Context, reporttype, company, currency, unit string, reportyear int) (map[string]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashCF(ctx, reporttype, company, currency, unit, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashCapEx(ctx context.Context, reporttype, company, currency, unit string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashCapExSummary(ctx, reporttype, company, currency, unit, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}

func (s *InvestmentsDashService) GetInvestmentsDashCashEndPeriod(ctx context.Context, reporttype, company, currency, unit string, reportyear int) (map[int]float64, error) {
	totalSumSummary, err := s.Repo.GetInvestmentsDashCashEndPeriod(ctx, reporttype, company, currency, unit, reportyear)
	if err != nil {
		return nil, err
	}
	return totalSumSummary, nil
}
