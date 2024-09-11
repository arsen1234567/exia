package models

type ReservesOilNgs struct {
	Year                                          string  `json:"Год"`
	Region                                        string  `json:"Регион"`
	Nedropolzhovatel                              string  `json:"Недропользователь"`
	NGSRegion                                     string  `json:"Месторождения в балансе НГС"`
	AlphRegion                                    string  `json:"Месторождения в алфавитном справ."`
	DegreeOfDevelopment                           string  `json:"Степень освоения месторождения"`
	ContractID                                    string  `json:"Номер контракта"`
	Type                                          string  `json:"Тип"`
	DataOfIssue                                   string  `json:"Дата выдачи"`
	Production_since_the_beginning_of_development float64 `json:"Добыча с начала разработки"`
	DateForGKZ                                    string  `json:"Добыча на дату утверждения ГКЗ"`
	ABC1                                          string  `json:"Балансовые запасы на начало(А+В+С1)"`
	C2                                            string  `json:"Балансовые запасы на начало(С2)"`
	Lost_ABC1                                     string  `json:"Добычи, потери(А+В+С1)"`
	IBZDP_ABC1                                    string  `json:"Изм. Бал. Зап. добычи, потери(А+В+С1)"`
	IBZR_ABC1                                     string  `json:"Изм. Бал. Зап. разведки(А+В+С1)"`
	IBZPP_ABC1                                    string  `json:"Изм. Бал. Зап. переоц. и перед.(А+В+С1)"`
	IBSNZ_ABC1                                    string  `json:"Изм. Бал. спис. неподт. Зап.(А+В+С1)"`
	A_end                                         string  `json:"Балансовые запасы на конец(А)"`
	B_end                                         string  `json:"Балансовые запасы на конец(B)"`
	AB_end                                        string  `json:"Балансовые запасы на конец(А+В)"`
	C1_end                                        string  `json:"Балансовые запасы на конец(С1)"`
	ABC1_end                                      string  `json:"Балансовые запасы на конец(А+В+С1)"`
	C2_end                                        string  `json:"Балансовые запасы на конец(С2)"`
	Utv_GKZ_ABC1                                  string  `json:"Бал. запасы утв. ГКЗ (А+В+С1)"`
	Utv_GKZ_C2                                    float64 `json:"Бал. запасы утв. ГКЗ (С2)"`
	Com_ID                                        int     `json:"com_id"`
	Bin                                           string  `json:"bin"`
	Dep_id                                        int     `json:"dep_id"`
	Red_id                                        int     `json:"reg_id"`
	AlphRegionCount                               int     `json:"Count"`
}

type NgsReservesOilSummary struct {
	Year            int `json:"year" db:"Год"`
	AlphRegionCount int `json:"Count"`
}

type NgsReservesOilTopCompanies struct {
	CompanyName   string  `json:"Company"`
	TotalReserves float64 `json:"TotalReserves"`
}
