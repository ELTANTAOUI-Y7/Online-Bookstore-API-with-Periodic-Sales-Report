package reports

import (
	"encoding/json"
	"fmt"
	"online-bookstore-api/models"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ListReportsInRange returns all sales reports in outputDir whose timestamp
// falls within [start, end] (inclusive). start and end are truncated to date only.
func ListReportsInRange(outputDir string, start, end time.Time) ([]models.SalesReport, error) {
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	end = time.Date(end.Year(), end.Month(), end.Day(), 23, 59, 59, 999999999, end.Location())

	entries, err := os.ReadDir(outputDir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("read report dir: %w", err)
	}

	var reports []models.SalesReport
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasPrefix(name, reportPrefix) || !strings.HasSuffix(name, reportSuffix) {
			continue
		}
		// name = report_DDMMYYYYHHmm.json
		dateStr := strings.TrimPrefix(name, reportPrefix)
		dateStr = strings.TrimSuffix(dateStr, reportSuffix)
		t, err := time.Parse("020120061504", dateStr)
		if err != nil {
			continue
		}
		if t.Before(start) || t.After(end) {
			continue
		}
		path := filepath.Join(outputDir, name)
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		var r models.SalesReport
		if err := json.Unmarshal(data, &r); err != nil {
			continue
		}
		reports = append(reports, r)
	}
	return reports, nil
}
