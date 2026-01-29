package reports

import (
	"encoding/json"
	"fmt"
	"log"
	"online-bookstore-api/interfaces"
	"online-bookstore-api/models"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const (
	topSellingLimit = 10
	reportPrefix    = "report_"
	reportSuffix    = ".json"
)

// GenerateSalesReport fetches orders in the last 24 hours, aggregates data,
// and returns a SalesReport. Caller is responsible for saving if needed.
func GenerateSalesReport(orderStore interfaces.OrderStore) (*models.SalesReport, error) {
	end := time.Now()
	start := end.Add(-24 * time.Hour)

	orders, err := orderStore.GetOrdersInTimeRange(start, end)
	if err != nil {
		return nil, fmt.Errorf("get orders in time range: %w", err)
	}

	var totalRevenue float64
	var totalBooksSold int
	bookQuantities := make(map[int]models.BookSales) // book ID -> BookSales (book + quantity)

	for _, order := range orders {
		totalRevenue += order.TotalPrice
		for _, item := range order.Items {
			totalBooksSold += item.Quantity
			key := item.Book.ID
			existing, ok := bookQuantities[key]
			if !ok {
				bookQuantities[key] = models.BookSales{Book: item.Book, Quantity: item.Quantity}
			} else {
				existing.Quantity += item.Quantity
				bookQuantities[key] = existing
			}
		}
	}

	// Top-selling books: sort by quantity descending, take top N
	type pair struct {
		bookID int
		bs     models.BookSales
	}
	var pairs []pair
	for id, bs := range bookQuantities {
		pairs = append(pairs, pair{id, bs})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].bs.Quantity > pairs[j].bs.Quantity
	})
	topN := topSellingLimit
	if len(pairs) < topN {
		topN = len(pairs)
	}
	topSelling := make([]models.BookSales, topN)
	for i := 0; i < topN; i++ {
		topSelling[i] = pairs[i].bs
	}

	report := &models.SalesReport{
		Timestamp:       end,
		TotalRevenue:    totalRevenue,
		TotalOrders:     len(orders),
		TotalBooksSold:  totalBooksSold,
		TopSellingBooks: topSelling,
	}
	return report, nil
}

// ReportFilename returns the filename for a report at the given time.
// Format: report_DDMMYYYYHHmm.json (e.g. report_250120261430.json).
func ReportFilename(t time.Time) string {
	return reportPrefix + t.Format("020120061504") + reportSuffix
}

// SaveReport writes the report to outputDir with filename report_DDMMYYYYHHmm.json.
func SaveReport(report *models.SalesReport, outputDir string) (string, error) {
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return "", fmt.Errorf("create output dir: %w", err)
	}
	filename := ReportFilename(report.Timestamp)
	path := filepath.Join(outputDir, filename)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("create report file: %w", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(report); err != nil {
		_ = os.Remove(path)
		return "", fmt.Errorf("encode report: %w", err)
	}
	return path, nil
}

// GenerateAndSave generates a sales report for the last 24 hours and saves it to outputDir.
func GenerateAndSave(orderStore interfaces.OrderStore, outputDir string) (*models.SalesReport, string, error) {
	report, err := GenerateSalesReport(orderStore)
	if err != nil {
		return nil, "", err
	}
	path, err := SaveReport(report, outputDir)
	if err != nil {
		return nil, "", err
	}
	log.Printf("Sales report generated: %s (revenue=%.2f, orders=%d, books_sold=%d)", path, report.TotalRevenue, report.TotalOrders, report.TotalBooksSold)
	return report, path, nil
}
