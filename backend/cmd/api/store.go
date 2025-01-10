package main

import (
	"backend/internals/schemas"
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	queryTimeoutDuration = time.Duration(5 * time.Second)
)

func WithTxn(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

func CreateUser(ctx context.Context, tx *sql.Tx, user *schemas.SignupRequest) (int, error) {
	query := `
	Insert into users  (username, password, email, name) VALUES 
    ($1, $2, $3, $4)
    RETURNING id;
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	userID := 0

	err := tx.QueryRowContext(
		ctx, query,
		user.Username, user.Password,
		user.Email, user.Name,
	).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetUserByCreds(ctx context.Context, db *sql.DB, username, hashedPassword string) (*schemas.UserDetails, error) {
	query := `
	Select id, email, name from users
	where username=$1 and password=$2;
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	userDetail := schemas.UserDetails{
		Username: username,
	}

	err := db.QueryRowContext(
		ctx, query,
		username, hashedPassword,
	).Scan(&userDetail.ID, &userDetail.Email, &userDetail.Name)
	if err != nil {
		return nil, err
	}

	return &userDetail, nil
}

func GetAllDashboards(
	ctx context.Context, db *sql.DB, userID int,
) ([]schemas.DashboardInfo, error) {
	query := `
	Select dashboard_id, dashboard_name,
	(Select count(chart_id) from charts group by dashboard_id)
	as chart_count from dashboards where user_id=$1
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	rows, err := db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	dashboards := []schemas.DashboardInfo{}

	for rows.Next() {
		dashboard := schemas.DashboardInfo{}

		err = rows.Scan(&dashboard.ID, &dashboard.Name, &dashboard.ChartsCount)
		if err != nil {
			return nil, err
		}
		dashboards = append(dashboards, dashboard)
	}

	return dashboards, nil
}

// TODO: Convert to txn
func CreateDashboard(
	ctx context.Context, db *sql.DB,
	userID int, dashboardName string,
) (int, error) {
	query := `
	Insert into dashboards (user_id, dashboard_name)
	Values ($1,$2) returning dashboard_id;
	`

	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	dashboardID := 0

	err := db.QueryRowContext(
		ctx, query, userID, dashboardName,
	).Scan(&dashboardID)
	if err != nil {
		return 0, err
	}

	return dashboardID, nil
}

func GetDashboardDetailsByID(ctx context.Context, db *sql.DB, dashboardID, userID int) (*schemas.DashboardDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	var dashboard schemas.DashboardDetail

	query := `SELECT dashboard_id, dashboard_name FROM dashboards WHERE dashboard_id = $1 AND user_id = $2`
	err := db.QueryRowContext(ctx, query, dashboardID, userID).Scan(&dashboard.ID, &dashboard.Name)
	if err != nil {
		return &dashboard, err
	}

	queryCharts := `SELECT chart_id, chart_type, chart_data FROM charts WHERE dashboard_id = $1`
	rows, err := db.QueryContext(ctx, queryCharts, dashboardID)
	if err != nil {
		return &dashboard, err
	}
	defer rows.Close()

	for rows.Next() {
		var chart schemas.ChartDetail
		var chartData []byte

		if err := rows.Scan(&chart.ID, &chart.Type, &chartData); err != nil {
			return &dashboard, err
		}
		chart.Data = string(chartData)

		dashboard.Charts = append(dashboard.Charts, chart)
	}

	if err := rows.Err(); err != nil {
		return &dashboard, err
	}

	return &dashboard, nil

}

func UpdateDashboardByID(ctx context.Context, db *sql.DB, dashBoardID, userID int, newName string) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	query := `
	Update dashboards set dashboard_name=$1 where dashboard_id=$2 and user_id=$3;
	`

	rows, err := db.ExecContext(ctx, query, newName, dashBoardID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("invalid dashboard_id provided")
	}

	return nil
}

// TODO: Convert to txn
func CreateChart(ctx context.Context, db *sql.DB, dashboardID, userID int, chartData *schemas.CreateChartRequest) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	queryOne := `
	Select dashboard_id from dashboards
	where user_id=$1 and dashboard_id=$2;
	`
	temp := 0
	err := db.QueryRowContext(ctx, queryOne, userID, dashboardID).Scan(&temp)
	if err != nil {
		return 0, err
	}

	queryTwo := `
	Insert into charts (dashboard_id,chart_type,chart_data)
	Values ($1,$2,$3) returning id
	`

	chartID := 0
	err = db.QueryRowContext(ctx, queryTwo, dashboardID, chartData.Type, []byte(chartData.Data)).Scan(&chartID)
	if err != nil {
		return 0, err
	}

	return chartID, nil
}

// TODO: Convert to txn
func UpdateChartByID(ctx context.Context, db *sql.DB, dashboardID, userID int, chartDetail *schemas.ChartDetail) error {
	ctx, cancel := context.WithTimeout(ctx, queryTimeoutDuration)
	defer cancel()

	queryOne := `
	Select dashboard_id from dashboards
	where user_id=$1 and dashboard_id=$2;
	`
	temp := 0
	err := db.QueryRowContext(ctx, queryOne, userID, dashboardID).Scan(&temp)
	if err != nil {
		return err
	}

	queryTwo := `
		UPDATE charts
		SET chart_type = $1, chart_data = $2
		WHERE chart_id = $3 and dashboard_id=$4;
	`

	res, err := db.ExecContext(ctx, queryTwo, chartDetail.Type, []byte(chartDetail.Data), chartDetail.ID, dashboardID)
	if err != nil {
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rowCount == 0 {
		return errors.New("invalid row provided")
	}

	return nil
}
