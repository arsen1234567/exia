package models

import "time"

// Main model for the NgsReservesGas table
type NgsReservesCond struct {
	Year                        int       `json:"year" db:"Год"`
	Region                      string    `json:"region" db:"Регион"`
	SubsoilUser                 string    `json:"subsoil_user" db:"Недропользователь"`
	FieldBalanceNGS             string    `json:"field_balance_ngs" db:"Месторождения в балансе НГС"`
	FieldAlphabetical           string    `json:"field_alphabetical" db:"Месторождения в алфавитном справ."`
	FieldDevelopmentStage       string    `json:"field_development_stage" db:"Степень освоения месторождения"`
	ContractNumber              string    `json:"contract_number" db:"Номер контракта"`
	IssuanceDate                time.Time `json:"issuance_date" db:"Дата выдачи"`
	GasType                     string    `json:"gas_type" db:"Вид газа"`
	Type                        string    `json:"type" db:"Тип"`
	ProductionSinceStart        float64   `json:"production_since_start" db:"Добыча с начала разработки"`
	ProductionAtGKZApproval     float64   `json:"production_at_gkz_approval" db:"Добыча на дату утверждения ГКЗ"`
	BalanceAtStartABC1          float64   `json:"balance_at_start_abc1" db:"Балансовые запасы на начало(А+В+С1)"`
	BalanceAtStartC2            float64   `json:"balance_at_start_c2" db:"Балансовые запасы на начало(С2)"`
	BalanceChangeExtractionABC1 float64   `json:"balance_change_extraction_abc1" db:"Изм. Бал. Зап. добычи(А+В+С1)"`
	BalanceChangeLossABC1       float64   `json:"balance_change_loss_abc1" db:"Изм. Бал. Зап. потери(А+В+С1)"`
	BalanceChangeExplorationABC float64   `json:"balance_change_exploration_abc" db:"Изм. Бал. Зап. разведки(А+В+С)"`
	BalanceChangeTransferABC    float64   `json:"balance_change_transfer_abc" db:"Изм. Бал. Зап. пересеоц. и перер. (А+В+С)"`
	BalanceEndUnaccountedABC1   float64   `json:"balance_end_unaccounted_abc1" db:"Изм. Бал. спис. неподт. Зап.(А+В+С1)"`
	BalanceEndA                 float64   `json:"balance_end_a" db:"Балансовые запасы на конец(А)"`
	BalanceEndAB                float64   `json:"balance_end_ab" db:"Балансовые запасы на конец(А+В)"`
	BalanceEndABC1              float64   `json:"balance_end_abc1" db:"Балансовые запасы на конец(А+В+С1)"`
	BalanceEndC2                float64   `json:"balance_end_c2" db:"Балансовые запасы на конец(С2)"`
	BalanceGKZAB                float64   `json:"balance_gkz_ab" db:"Бал. запасы утв. ГКЗ (А+В)"`
	BalanceGKZABC1              float64   `json:"balance_gkz_abc1" db:"Бал. запасы утв. ГКЗ (А+В+С1)"`
	BalanceGKZC2                float64   `json:"balance_gkz_c2" db:"Бал. запасы утв. ГКЗ (С2)"`
	RegionID                    int       `json:"region_id" db:"reg_id"`
	CompanyID                   int       `json:"company_id" db:"com_id"`
	Bin                         string    `json:"bin" db:"bin"`
	ProductionABC1              float64   `json:"production_abc1" db:"Добыча(А+В+С1)"`
	LossABC1                    float64   `json:"loss_abc1" db:"Потери(А+В+С1)"`
	DevelopmentStage            string    `json:"development_stage" db:"stepen_osoveyna"`
	GasTypeText                 string    `json:"gas_type_text" db:"Type gas"`
	AlphRegionCount             int       `json:"Count"`
	TotalReservesSum            float64   `json:"total_reserves_sum"`
}

type NgsReservesCondSummary struct {
	Year            int `json:"year" db:"Год"`
	AlphRegionCount int `json:"Count"`
}

type NgsReservesCondTotalReservesSummary struct {
	Year             int     `json:"year" db:"Год"`
	TotalReservesSum float64 `json:"total_reserves_sum"`
}

type RecoverableCondReservesSummary struct {
	Year          int     `json:"year" db:"year"`
	TotalReserves float64 `json:"total_reserves" db:"total_reserves"`
}

type NgsReservesCondTopCompanies struct {
	CompanyName   string  `json:"Company"`
	TotalReserves float64 `json:"TotalReserves"`
}
