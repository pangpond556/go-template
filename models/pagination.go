package models

// Total - total row
type Total struct {
	Total int64 `bson:"total" json:"total"`
}

// Pagination - pagination
type Pagination struct {
	Data       []map[string]interface{} `bson:"data" json:"data"`
	Pagination []Total                  `bson:"pagination" json:"pagination"`
}

// GetTotalAndPagination - calculate total rows and last page from pagination
func GetTotalAndPagination(p Pagination) (int64, int64) {
	var lastPage int64
	var total int64
	if len(p.Pagination) > 0 {
		lastPage = p.Pagination[0].Total / 15
		if lastPage == 0 {
			lastPage = 1
		}
		total = p.Pagination[0].Total
	} else {
		lastPage = 1
		total = 0
	}
	return total, lastPage
}
