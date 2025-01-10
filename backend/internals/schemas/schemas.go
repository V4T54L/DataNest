package schemas

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

type UserDetails struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateDashboardRequest struct {
	Name string `json:"name"`
}

type ChartDetail struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Data string `json:"data"`
}

type DashboardDetail struct {
	ID     int           `json:"id"`
	Name   string        `json:"name"`
	Charts []ChartDetail `json:"charts"`
}

type DashboardInfo struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	ChartsCount int    `json:"charts_count"`
}

type CreateChartRequest struct {
	Type string `json:"type"`
	Data string `json:"data"`
}
