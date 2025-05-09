package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

type Sale struct {
	Product  string
	Count    int
	Price    float64
	Date     string // YYYY-MM-DD
	Category string
}

type SaleData struct {
	Count int
	Price float64
}

const (
	layout = "2006-01-02"
)

func main() {
	sales := []Sale{
		{"яблоко", 5, 10.5, "2025-03-10", "фрукты"},
		{"банан", 3, 15.0, "2025-03-12", "фрукты"},
		{"яблоко", 2, 10.5, "2025-04-05", "фрукты"}, // не попадает
		{"апельсин", 4, 8.0, "2025-03-18", "фрукты"},
		{"банан", 1, 15.0, "2025-04-02", "фрукты"}, // не попадает
		{"картошка", 10, 5.0, "2025-03-20", "овощи"},
		{"морковь", 7, 4.0, "2025-04-01", "овощи"},
	}

	startDate := ParseTimeSafe("2025-03-01")
	endDate := ParseTimeSafe("2025-04-01")

	data := make(map[string]SaleData)
	dataCat := make(map[string]SaleData)
	sliceProducts := []string{}
	sliceCat := []string{}

	for _, saleItem := range sales {

		d := ParseTimeSafe(saleItem.Date)

		if d.Before(startDate) || d.After(endDate) {
			continue
		}

		if _, e := data[saleItem.Product]; !e {
			sliceProducts = append(sliceProducts, saleItem.Product)
		}

		if _, e := dataCat[saleItem.Category]; !e {
			sliceCat = append(sliceCat, saleItem.Category)
		}

		data[saleItem.Product] = SaleData{
			Count: data[saleItem.Product].Count + saleItem.Count,
			Price: data[saleItem.Product].Price + saleItem.Price,
		}

		dataCat[saleItem.Category] = SaleData{
			Count: dataCat[saleItem.Category].Count + saleItem.Count,
			Price: dataCat[saleItem.Category].Price + saleItem.Price,
		}
	}

	fmt.Printf("Статистика продаж по товарам:\n")
	fmt.Printf("-----------------------------------\n")

	totalNumber := 0
	totalPrice := float64(0.0)
	maxPrice := float64(0.0)
	bestID := ""
	sort.Strings(sliceProducts)
	sort.Strings(sliceCat)

	for _, itemName := range sliceProducts {
		item := data[itemName]
		fmt.Printf("%s — количество: %d, общая сумма: %.2f\n", itemName, item.Count, item.Price)

		totalNumber += item.Count
		totalPrice += item.Price
		if item.Price > maxPrice {
			maxPrice = item.Price
			bestID = itemName
		}
	}

	fmt.Printf("\n")
	fmt.Printf("Статистика продаж по категориям:\n")
	fmt.Printf("-----------------------------------\n")

	for _, itemName := range sliceCat {
		item := dataCat[itemName]
		fmt.Printf("%s — количество: %d, общая сумма: %.2f\n", itemName, item.Count, item.Price)
	}

	fmt.Printf("\n")
	fmt.Printf("Общее количество проданных товаров: %d\n", totalNumber)
	fmt.Printf("Общая сумма всех продаж: %.2f\n", totalPrice)
	fmt.Printf("Товар с наибольшей суммой продаж: %s\n", bestID)

}

func ParseTimeSafe(s string) time.Time {
	d, err := time.Parse(layout, s)
	if err != nil {
		fmt.Println("Ошибка при обработке даты!")
		os.Exit(1)
	}
	return d
}
