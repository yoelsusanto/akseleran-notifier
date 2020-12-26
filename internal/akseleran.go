package internal

import (
	"bytes"
	"encoding/json"
	"net/http"
	"text/template"
	"time"

	"github.com/yoelsusanto/akseleran-notifier/models"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func GetAkseleranCampaign() (*models.AkseleranResponse, error) {
	req, err := http.NewRequest("GET", "https://core.akseleran.com/api/v1/campaigns", nil)
	if err != nil {
		return nil, err
	}

	query := req.URL.Query()
	query.Add("campaign_status", "OPEN_CAMPAIGN")
	query.Add("campaign_type", "P2P_BUSINESS,P2P_ONLINE_MERCHANT,P2P_LOAN_PURCHASE,P2P_EMPLOYEE_LOAN,P2P_INVOICE_FINANCING,P2P_RECEIVEABLE_FINANCING,P2P_CAPEX_FINANCING,P2P_INVENTORY_FINANCING,P2P_PORTFOLIO_FINANCING")
	query.Add("location", "front")
	query.Add("min_installment_range", "0")
	query.Add("max_installment_range", "99")
	query.Add("min_interest_range", "0")
	query.Add("max_interest_range", "99")
	query.Add("sort", "campaigndate_desc")
	query.Add("limit", "6")

	req.URL.RawQuery = query.Encode()

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	akseleranResponse := models.AkseleranResponse{}

	err = json.NewDecoder(resp.Body).Decode(&akseleranResponse)
	if err != nil {
		return nil, err
	}
	return &akseleranResponse, nil
}

func CampaignToMessage(campaign models.Campaign) (string, error) {
	campaignMessageTemplate :=
		`
Hi investor, ada campaign baru nih!
Nama: {{ .CampaignName }}
Tipe: {{ .CampaignType }}
Rating: {{ .LoanCreditRating }}
Interest: {{ .RoundedEffectiveInterest }}
Agunan: {{ if .HaveCollateral }}Ada{{ else }}Tidak{{ end }}
Frekuensi Pembayaran Pokok: {{ .InstallmentPaymentFreq }}
Tenor: {{ .InstallmentLength }} {{ .InstallmentLengthUnit.Value }}
{{ if .CampaignStart -}}
Campaign Start: {{ convertAkseleranDate .CampaignStart }}
{{ end -}}
{{ if .CampaignEnd -}}
Campaign End: {{ convertAkseleranDate .CampaignEnd }}
{{ end -}}
Remaining Time: {{ .CampaignTimeRemaining }}
Total Invested: Rp. {{ floatToRupiah .TotalInvestment }}
Funded Percentange: {{ .FundedPercentage }}%
Status: {{ .Status }}
`

	campaignTemplate, err := template.New("").Funcs(template.FuncMap{
		"convertAkseleranDate": reformatAkseleranDate,
		"floatToRupiah":        floatToRupiah,
	}).Parse(campaignMessageTemplate)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = campaignTemplate.Execute(&buf, campaign)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func reformatAkseleranDate(akseleranDate string) (string, error) {
	akseleranDateFormat := "2006-01-02 15:04:05"
	parsedDate, err := time.Parse(akseleranDateFormat, akseleranDate)

	if err != nil {
		return "", err
	}

	result := parsedDate.Format("02 January 2006 15:04:05")
	return result, nil
}

func floatToRupiah(f float64) string {
	p := message.NewPrinter(language.Indonesian)
	return p.Sprintf("%.0f", f)
}
