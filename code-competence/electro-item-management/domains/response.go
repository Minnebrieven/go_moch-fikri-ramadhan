package domains

type (
	ResponseMetadata struct {
		StatusCode int    `json:"status_code"`
		Message    string `json:"message"`
	}

	Pagination struct {
		Page      int `json:"page"`
		DataShown int `json:"data_shown"`
		TotalData int `json:"total_data"`
	}

	GeneralListResponse struct {
		ResponseMetadata `json:"metadata"`
		Pagination       `json:"pagination"`
		Data             interface{} `json:"data"`
	}

	GeneralResponse struct {
		ResponseMetadata `json:"metadata"`
		Data             interface{} `json:"data"`
	}

	LoginResponse struct {
		ResponseMetadata `json:"metadata"`
		Data             interface{} `json:"data"`
		Token            string      `json:"token"`
	}
)
