package handlers

import (
	"context"
	"net/http"
	"online-bookstore-api/models"
	"online-bookstore-api/reports"
	"time"
)

const (
	defaultReportsStartDays = 30
	reportsTimeout          = 5 * time.Second
)

// GetSalesReports handles GET /reports/sales?start_date=YYYY-MM-DD&end_date=YYYY-MM-DD.
// If start_date or end_date are omitted, defaults to last 30 days.
func (h *Handler) GetSalesReports(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), reportsTimeout)
	defer cancel()

	if checkContext(ctx, w) {
		return
	}

	start, end, err := parseReportDateRange(r)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if checkContext(ctx, w) {
		return
	}

	list, err := reports.ListReportsInRange(h.ReportOutputDir, start, end)
	if err != nil {
		LogError("GetSalesReports", "list reports failed", err)
		respondWithError(w, http.StatusInternalServerError, "Failed to load reports")
		return
	}

	if list == nil {
		list = []models.SalesReport{}
	}

	respondWithJSON(w, http.StatusOK, list)
	LogInfo("GetSalesReports", "reports retrieved", map[string]interface{}{
		"start": start.Format("2006-01-02"),
		"end":   end.Format("2006-01-02"),
		"count": len(list),
	})
}

// parseReportDateRange parses start_date and end_date query params (YYYY-MM-DD).
// If missing, uses last defaultReportsStartDays days.
func parseReportDateRange(r *http.Request) (start, end time.Time, err error) {
	now := time.Now()
	end = now
	start = now.AddDate(0, 0, -defaultReportsStartDays)

	if s := r.URL.Query().Get("start_date"); s != "" {
		t, e := time.Parse("2006-01-02", s)
		if e != nil {
			return time.Time{}, time.Time{}, errBadDate("start_date", s)
		}
		start = t
	}
	if s := r.URL.Query().Get("end_date"); s != "" {
		t, e := time.Parse("2006-01-02", s)
		if e != nil {
			return time.Time{}, time.Time{}, errBadDate("end_date", s)
		}
		end = t
	}

	if start.After(end) {
		start, end = end, start
	}
	return start, end, nil
}

func errBadDate(name, value string) error {
	return &badDateError{name: name, value: value}
}

type badDateError struct {
	name  string
	value string
}

func (e *badDateError) Error() string {
	return "invalid " + e.name + " (use YYYY-MM-DD): " + e.value
}
