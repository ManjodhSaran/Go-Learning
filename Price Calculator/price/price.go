package price

import (
	"fmt"

	"pricecalculator.com/conversion"
	"pricecalculator.com/filemanager"
)

type TaxIncludedPriceJob struct {
	TaxRate        float64
	InputPrices    []float64
	TaxIncludedJob map[string]string
	IOManage       filemanager.FileManager
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOManage.ReadLines()

	if err != nil {
		fmt.Println(err)
		return
	}

	prices, err := conversion.StringsToFloats(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices

}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		tax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", tax)
	}

	job.TaxIncludedJob = result

	job.IOManage.WriteResult(job)
}

func NewTaxIncludedPriceJob(fm filemanager.FileManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManage:    fm,
	}
}
