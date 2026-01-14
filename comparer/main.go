package main

import (
	"fmt"
	"reflect"
)

/*
{"__faust": {"ns": "scoring_kontur_sanctions.agents.Sanction.model.MessageOut"},

	"sanctions": {
	"primary": [
	{"name": "Санкционный список Канады SEMA", "programs": ["Russia / Russie"], "source_url": "https://www.international.gc.ca/", "external_id": "Russia / Russie-1, Part 2-416", "legal_basis": "Нет данных", "source_name": "Government of Canada", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": ["Schedule 1, Part 2"], "actual_list_date": "15.07.2024"}, {"name": "Санкционный список Украины МЭРТ", "programs": ["МЭРТ"], "source_url": "https://www.president.gov.ua/", "external_id": "1027700229193", "legal_basis": "14.05.2020", "source_name": "Министерство экономического развития и торговли", "exclude_date": "Бессрочно", "include_date": "14.05.2020", "restrictions": [], "actual_list_date": "24.06.2024"}]}}

'{"__faust": {"ns": "scoring_kontur_sanctions.agents.Sanction.model.MessageOut"},

	"sanctions": {"related": [{"name": "Санкционный список Украины МЭРТ", "programs": [], "source_url": "https://www.president.gov.ua/", "external_id": "AHfpnZ7vFEKdJXjIC/1vEA==", "legal_basis": "Нет данных", "source_name": "Министерство экономического развития и торговли", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.06.2024"}, {"name": "Санкционный список США OFAC", "programs": [], "source_url": "https://home.treasury.gov/", "external_id": "Nr9W/uC7h0KEkJcQxPCmYQ==", "legal_basis": "Нет данных", "source_name": "U.S. Department of the Treasury - Office of Foreign Assets Control", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список США OFAC", "programs": [], "source_url": "https://home.treasury.gov/", "external_id": "Nr9W/uC7h0KEkJcQxPCmYQ==", "legal_basis": "Нет данных", "source_name": "U.S. Department of the Treasury - Office of Foreign Assets Control", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Великобритании HM Treasury", "programs": [], "source_url": "https://www.gov.uk/government/organisations/hm-treasury", "external_id": "RkY/uByCsUuq4MkgfEUDXA==", "legal_basis": "Нет данных", "source_name": "Office of Financial Sanctions Implementation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Великобритании HM Treasury", "programs": [], "source_url": "https://www.gov.uk/government/organisations/hm-treasury", "external_id": "RkY/uByCsUuq4MkgfEUDXA==", "legal_basis": "Нет данных", "source_name": "Office of Financial Sanctions Implementation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Франции GEL", "programs": [], "source_url": "https://www.tresor.economie.gouv.fr/", "external_id": "3Kd57XeBqU2xqLBVP1zjhg==", "legal_basis": "Нет данных", "source_name": "Ministry of Finance", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Японии JMETI", "programs": [], "source_url": "https://www.meti.go.jp/", "external_id": "eWP8Sn/IXUeLZL9YJ0M9Qg==", "legal_basis": "Нет данных", "source_name": "Ministry of Economy, Trade and Industry", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "06.12.2023"}, {"name": "Санкционный список Швейцарии SECO", "programs": [], "source_url": "https://www.seco.admin.ch/seco/de/home.html", "external_id": "0e8LL3Fa5kGnkw+YOZwx4A==", "legal_basis": "Нет данных", "source_name": "Swiss State Secretariat for Economic Affairs", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "09.07.2024"}, {"name": "Санкционный список Польши MWSIA", "programs": [], "source_url": "https://www.gov.pl/web/mswia/", "external_id": "TB6pGpFh8kyilnzpTyV/vw==", "legal_basis": "Нет данных", "source_name": "Ministry of Internal Affairs and Administration", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Великобритании UKSL", "programs": [], "source_url": "https://www.gov.uk/", "external_id": "bZ4x/0iVTU6zCoEEmC4/iQ==", "legal_basis": "Нет данных", "source_name": "Government of the United Kingdom", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "11.03.2024"}, {"name": "Санкционный список Гибралтар GFIU", "programs": [], "source_url": "https://www.gfiu.gov.gi/", "external_id": "aSg7/95AQk2O1mTdx8CvaA==", "legal_basis": "Нет данных", "source_name": "Financial Intelligence Unit", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Монако SICCFIN National", "programs": [], "source_url": "https://geldefonds.gouv.mc/", "external_id": "srZ4hgW7y0Gh5Kd8AIaRdg==", "legal_basis": "Нет данных", "source_name": "Service d`Information et de Controle sur les Circuits Financiers", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "23.11.2023"}, {"name": "Санкционный список Джерси JFSC HM Treasury", "programs": [], "source_url": "https://www.gov.je/", "external_id": "uLPi2yCxYECgDaxkXqPISA==", "legal_basis": "Нет данных", "source_name": "Financial Services Comission", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Бельгии FOJ", "programs": [], "source_url": "https://finance.belgium.be/", "external_id": "jqe9jvDPG0CxcuBE+dPJ+w==", "legal_basis": "Нет данных", "source_name": "Federal Public Service for Justice", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "19.07.2024"}, {"name": "Санкционный список Японии JMOF", "programs": [], "source_url": "https://www.mof.go.jp/", "external_id": "KMI2lkE05E+l5juaBR1+zA==", "legal_basis": "Нет данных", "source_name": "Ministry of Finance", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "23.07.2024"}], "fifty_percent": [{"name": "Санкционный список США BIS", "programs": [], "source_url": "https://www.bis.doc.gov", "external_id": "D4vtWsrD6ECeCOXFHDj+rQ==", "legal_basis": "Нет данных", "source_name": "U.S. Department of Commerce - Bureau of Industry and Security", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "17.07.2024"}, {"name": "Санкционный список США BEBA OFAC", "programs": [], "source_url": "https://home.treasury.gov/", "external_id": "f4v3BBlMbEm7NfWmgG1tcQ==", "legal_basis": "Нет данных", "source_name": "U.S. Deparment of State - Bureau of Economic and Business Administration", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список США OFAC", "programs": [], "source_url": "https://home.treasury.gov/", "external_id": "Nr9W/uC7h0KEkJcQxPCmYQ==", "legal_basis": "Нет данных", "source_name": "U.S. Department of the Treasury - Office of Foreign Assets Control", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список США OFAC", "programs": [], "source_url": "https://home.treasury.gov/", "external_id": "Nr9W/uC7h0KEkJcQxPCmYQ==", "legal_basis": "Нет данных", "source_name": "U.S. Department of the Treasury - Office of Foreign Assets Control", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Великобритании HM Treasury", "programs": [], "source_url": "https://www.gov.uk/government/organisations/hm-treasury", "external_id": "RkY/uByCsUuq4MkgfEUDXA==", "legal_basis": "Нет данных", "source_name": "Office of Financial Sanctions Implementation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Великобритании HM Treasury", "programs": [], "source_url": "https://www.gov.uk/government/organisations/hm-treasury", "external_id": "RkY/uByCsUuq4MkgfEUDXA==", "legal_basis": "Нет данных", "source_name": "Office of Financial Sanctions Implementation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "26.06.2024"}, {"name": "Санкционный список Евросоюза EU Manual (ручной)", "programs": [], "source_url": "https://eur-lex.europa.eu/", "external_id": "jt29owa7ykiIy6ludI4h+Q==", "legal_basis": "Нет данных", "source_name": "European Commission", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.06.2024"}, {"name": "Санкционный список Новой Зеландии MFAT (Russia sanctions register)", "programs": [], "source_url": "https://www.mfat.govt.nz/", "external_id": "hGqR/ily3kex5NQkd0Xzjw==", "legal_basis": "Нет данных", "source_name": "New Zealand Foreign Affairs & Trade", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "12.07.2024"}, {"name": "Санкционный список Европейского Союза EU", "programs": [], "source_url": "https://eeas.europa.eu/headquarters/headquarters-homepage_en", "external_id": "uSF1OkTDqUaEeOJJFFGhiA==", "legal_basis": "Нет данных", "source_name": "European Commission", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Нидерландов NNTS EU", "programs": [], "source_url": "https://www.government.nl/", "external_id": "CY3pXL6iF0+Yq8qjgzNoYg==", "legal_basis": "Нет данных", "source_name": "Government of the Netherlands", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Мальты MFSA EU", "programs": [], "source_url": "https://www.mfsa.mt/", "external_id": "aNK3cSylr0yLG5SmiFD/NQ==", "legal_basis": "Нет данных", "source_name": "Financial Services Authority", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Латвии LVSCS EU", "programs": [], "source_url": "https://www.fid.gov.lv/", "external_id": "P9eF/mVcpEW1Wod6dbqhOA==", "legal_basis": "Нет данных", "source_name": "Sanctions Control Service", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Германии DBB EU", "programs": [], "source_url": "https://www.bundesbank.de/", "external_id": "DEJeQa/k8UCsZ2CiIhfpbA==", "legal_basis": "Нет данных", "source_name": "Deutche Bundesbank", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Испании SMFA EU", "programs": [], "source_url": "http://www.exteriores.gob.es/", "external_id": "qb/8DD9yFkWtRcoshE1ktQ==", "legal_basis": "Нет данных", "source_name": "Ministry of Foreign Affairs and Cooperation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Джерси JFSC EU", "programs": [], "source_url": "https://www.gov.je/", "external_id": "KYXxdNYkL0GhhVDhNwiHCg==", "legal_basis": "Нет данных", "source_name": "Financial Services Comission", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Люксембурга CSSF EU", "programs": [], "source_url": "https://www.cssf.lu/", "external_id": "OKO0kwc3J0qTF9+2z+7odA==", "legal_basis": "Нет данных", "source_name": "Commission de Surveilance du Secteur Financier", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список Мальты MFSA OFAC", "programs": [], "source_url": "https://www.mfsa.mt/", "external_id": "6D1LarRjG0CImsSONQSHAQ==", "legal_basis": "Нет данных", "source_name": "Financial Services Authority", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "24.07.2024"}, {"name": "Санкционный список США BEBA BIS", "programs": [], "source_url": "https://www.bis.doc.gov/", "external_id": "1v/Wi0shv0KfkRH2ZJn4gg==", "legal_basis": "Нет данных", "source_name": "U.S. Deparment of State - Bureau of Economic and Business Administration", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "17.07.2024"}, {"name": "Санкционный список Канады SEMA", "programs": [], "source_url": "https://www.international.gc.ca/", "external_id": "Fm7qVo3t/02BV/8DEHCYsQ==", "legal_basis": "Нет данных", "source_name": "Government of Canada", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "15.07.2024"}, {"name": "Санкционный список Австралии DFAT", "programs": [], "source_url": "https://www.dfat.gov.au/", "external_id": "WgIxkzGzzUaTWDNC+iE/Vw==", "legal_basis": "Нет данных", "source_name": "Australian Government", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "25.07.2024"}], "twenty_five_percent": [{"name": "Санкционный список Японии MOFA", "programs": [], "source_url": "https://www.mofa.go.jp", "external_id": "J0zq7iTl6EyzXZk187EHkg==", "legal_basis": "Нет данных", "source_name": "Ministry of Foreign Affairs", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "21.06.2024"}, {"name": "Санкционный список США BISN", "programs": [], "source_url": "https://www.state.gov/", "external_id": "BkbR+psDjUyy0GW2h5HMjg==", "legal_basis": "Нет данных", "source_name": "U.S. Department of State - Bureau of International Security and Non-Proliferation", "exclude_date": "Бессрочно", "include_date": "Нет данных", "restrictions": [], "actual_list_date": "15.08.2023"}]}}'
*/

type Sanctions struct {
	Primary           []SanctionItem `json:"primary"`
	Related           []SanctionItem `json:"related"`
	FiftyPercent      []SanctionItem `json:"fifty_percent"`
	TwentyFivePercent []SanctionItem `json:"twenty_five_percent"`
}

type SanctionItem struct {
	Name           string   `json:"name"`
	Programs       []string `json:"programs"`
	SourceUrl      string   `json:"source_url"`
	ExternalId     string   `json:"external_id"`
	LegalBasis     string   `json:"legal_basis"`
	SourceName     string   `json:"source_name"`
	ExcludeDate    string   `json:"exclude_date"`
	IncludeDate    string   `json:"include_date"`
	Restrictions   []string `json:"restrictions"`
	ActualListDate string   `json:"actual_list_date"`
}

type SanctionDiffResult struct {
	Primary           *DiffResult
	Related           *DiffResult
	FiftyPercent      *DiffResult
	TwentyFivePercent *DiffResult
}

type DiffResult struct {
	Added    []SanctionItem
	Removed  []SanctionItem
	Modified []struct{ Old, New SanctionItem }
}

func CompareSanctionItems(a, b []SanctionItem) (*DiffResult, error) {
	if !hasUniqueIDs(a) || !hasUniqueIDs(b) {
		return nil, fmt.Errorf("дубликаты ID в одном из контейнеров")
	}

	mapA := createItemMap(a)
	mapB := createItemMap(b)

	result := &DiffResult{}

	for id, itemA := range mapA {
		if itemB, exists := mapB[id]; !exists {
			result.Removed = append(result.Removed, itemA)
		} else if !reflect.DeepEqual(itemA, itemB) {
			result.Modified = append(result.Modified, struct{ Old, New SanctionItem }{itemA, itemB})
		}
	}

	for id, itemB := range mapB {
		if _, exists := mapA[id]; !exists {
			result.Added = append(result.Added, itemB)
		}
	}

	return result, nil
}

func createItemMap(c []SanctionItem) map[string]SanctionItem {
	m := make(map[string]SanctionItem)
	for _, item := range c {
		m[item.ExternalId] = item
	}
	return m
}

func hasUniqueIDs(c []SanctionItem) bool {
	ids := make(map[string]struct{})
	for _, item := range c {
		if _, exists := ids[item.ExternalId]; exists {
			return false
		}
		ids[item.ExternalId] = struct{}{}
	}
	return true
}

func printDiff(diff *DiffResult) {
	fmt.Println("Добавленные элементы:")
	for _, item := range diff.Added {
		fmt.Printf("- ID: %d, Name: %s, Value: %.1f\n", item.ExternalId, item.Name, item.SourceName)
	}

	fmt.Println("\nУдаленные элементы:")
	for _, item := range diff.Removed {
		fmt.Printf("- ID: %d, Name: %s, Value: %.1f\n", item.ExternalId, item.Name, item.SourceName)
	}

	fmt.Println("\nИзмененные элементы:")
	for _, mod := range diff.Modified {
		fmt.Printf("ID: %d\n", mod.Old.ExternalId)
		fmt.Printf("  Было: Name: %s, Value: %.1f\n", mod.Old.Name, mod.Old.SourceName)
		fmt.Printf("  Стало: Name: %s, Value: %.1f\n", mod.New.Name, mod.New.SourceName)
	}
}

func main() {

	sanctions1 := &Sanctions{
		Primary: []SanctionItem{
			{ExternalId: "Russia / Russie-1, Part 2-416", Name: "Санкционный список Канады SEMA", SourceName: "Government of Canada"},
			{ExternalId: "AHfpnZ7vFEKdJXjIC/1vEA==", Name: "Санкционный список Украины МЭРТ", SourceName: "Министерство экономического развития и торговли"},
		},
		Related: []SanctionItem{
			{ExternalId: "Russia / Russie-1, Part 2-416", Name: "Санкционный список Канады SEMA", SourceName: "Government of Canada"},
		},
	}

	sanctions2 := &Sanctions{
		Primary: []SanctionItem{
			{ExternalId: "Russia / Russie-1, Part 2-416", Name: "Санкционный список Канады SEMAz", SourceName: "Government of Canada"},
			{ExternalId: "AHfpnZ7vFEKdJXjIC/1vEA==", Name: "Санкционный список Украины МЭРТ", SourceName: "Министерство экономического развития и торговли"},
			{ExternalId: "Nr9W/uC7h0KEkJcQxPCmYQ==", Name: "Санкционный список США OFAC", SourceName: "U.S. Department of the Treasury - Office of Foreign Assets Control"},
		},
		FiftyPercent: []SanctionItem{
			{ExternalId: "Russia / Russie-1, Part 2-416", Name: "Санкционный список Канады SEMA", SourceName: "Government of Canada"},
		},
	}

	var err error

	sanctionDiffResult := SanctionDiffResult{}

	if sanctions1 == nil {
		sanctions1 = &Sanctions{}
	}
	if sanctions2 == nil {
		sanctions2 = &Sanctions{}
	}

	if sanctions1.Primary != nil || sanctions2.Primary != nil {
		sanctionDiffResult.Primary, err = CompareSanctionItems(sanctions1.Primary, sanctions2.Primary)
		if err != nil {
			//p.logger.Warn("Monitoring verification", zap.Error(err))
		}
		printDiff(sanctionDiffResult.Primary)
	}
	if sanctions1.Related != nil || sanctions2.Related != nil {
		sanctionDiffResult.Related, err = CompareSanctionItems(sanctions1.Related, sanctions2.Related)
		if err != nil {
			//p.logger.Warn("Monitoring verification", zap.Error(err))
		}
		printDiff(sanctionDiffResult.Related)
	}
	if sanctions1.FiftyPercent != nil || sanctions2.FiftyPercent != nil {
		sanctionDiffResult.FiftyPercent, err = CompareSanctionItems(sanctions1.FiftyPercent, sanctions2.FiftyPercent)
		if err != nil {
			//p.logger.Warn("Monitoring verification", zap.Error(err))
		}
		printDiff(sanctionDiffResult.FiftyPercent)
	}
	if sanctions1.TwentyFivePercent != nil || sanctions2.TwentyFivePercent != nil {
		sanctionDiffResult.TwentyFivePercent, err = CompareSanctionItems(sanctions1.TwentyFivePercent, sanctions2.TwentyFivePercent)
		if err != nil {
			//p.logger.Warn("Monitoring verification", zap.Error(err))
		}
		printDiff(sanctionDiffResult.TwentyFivePercent)
	}

	//if sanctions1.Primary != nil || sanctions2.Primary != nil {
	//	sanctionDiffResult.Primary, err = CompareSanctionItems(sanctions1.Primary, sanctions2.Primary)
	//	if err != nil {
	//		fmt.Println("Ошибка:", err)
	//		return
	//	}
	//	printDiff(sanctionDiffResult.Primary)
	//}
	//if sanctions1.Related != nil || sanctions2.Related != nil {
	//	sanctionDiffResult.Related, err = CompareSanctionItems(sanctions1.Related, sanctions2.Related)
	//	if err != nil {
	//		fmt.Println("Ошибка:", err)
	//		return
	//	}
	//	printDiff(sanctionDiffResult.Related)
	//}

}
