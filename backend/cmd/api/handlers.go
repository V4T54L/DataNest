package main

import (
	"backend/internals/schemas"
	"backend/internals/utils"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

func (a *application) healthHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}

func (a *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var payload schemas.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// hash password
	payload.Password = utils.Hash(payload.Password)

	var userID int

	// create user
	err = WithTxn(a.db, r.Context(), func(tx *sql.Tx) error {
		userID, err = CreateUser(r.Context(), tx, &payload)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// return reponse
	utils.MessageResponse(w, http.StatusCreated, fmt.Sprintf("user created successfully with id: %d", userID))
}

func (a *application) loginHandler(w http.ResponseWriter, r *http.Request) {
	payload := schemas.LoginRequest{}
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	payload.Password = utils.Hash(payload.Password)

	var userDetail *schemas.UserDetails

	userDetail, err = GetUserByCreds(r.Context(), a.db, payload.Username, payload.Password)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenStr, err := utils.GenerateToken(*userDetail)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tokenCookie := http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    tokenStr,
		Expires:  time.Now().Add(time.Minute * 15),
		HttpOnly: true,                 // Prevent JavaScript access
		Secure:   false,                // Set to true if you're using HTTPS
		SameSite: http.SameSiteLaxMode, // Allow cross-origin cookies
	}

	http.SetCookie(w, &tokenCookie)
	utils.DataResponse(w, http.StatusOK, *userDetail)
}

func (a *application) getDashboardsHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	var dashboards []schemas.DashboardInfo
	dashboards, err := GetAllDashboards(r.Context(), a.db, user.ID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.DataResponse(w, http.StatusOK, dashboards)
}

func (a *application) createDashboardHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	var payload schemas.CreateDashboardRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	dashboardID, err := CreateDashboard(r.Context(), a.db, user.ID, payload.Name)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.MessageResponse(w, http.StatusCreated, fmt.Sprintf("Dashboard created with ID : %d", dashboardID))
}

func (a *application) updateDashboardByID(w http.ResponseWriter, r *http.Request) {
	var payload schemas.CreateDashboardRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	dashboardIDStr := chi.URLParam(r, "dashboardId")
	dashboardID, err := strconv.Atoi(dashboardIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = UpdateDashboardByID(r.Context(), a.db, dashboardID, user.ID, payload.Name)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.MessageResponse(w, http.StatusOK, "dashboard name updated successfully!")
}

func (a *application) getDashboardByIDHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	dashboardIDStr := chi.URLParam(r, "dashboardId")
	dashboardID, err := strconv.Atoi(dashboardIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	dashboardDetails, err := GetDashboardDetailsByID(r.Context(), a.db, dashboardID, user.ID)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.DataResponse(w, http.StatusOK, *dashboardDetails)
}

func (a *application) addChartHandler(w http.ResponseWriter, r *http.Request) {
	var payload schemas.CreateChartRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	dashboardIDStr := chi.URLParam(r, "dashboardId")
	dashboardID, err := strconv.Atoi(dashboardIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	chartID, err := CreateChart(r.Context(), a.db, dashboardID, user.ID, &payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.MessageResponse(w, http.StatusCreated, fmt.Sprintf("Chart added with id: %d", chartID))
}

func (a *application) updateChartHandler(w http.ResponseWriter, r *http.Request) {
	var payload schemas.ChartDetail
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	user, ok := r.Context().Value("userID").(schemas.UserDetails)
	if !ok {
		utils.ErrorResponse(w, http.StatusInternalServerError, "TODO : Handle this; Add auth middleware")
		return
	}

	dashboardIDStr := chi.URLParam(r, "dashboardId")
	dashboardID, err := strconv.Atoi(dashboardIDStr)
	if err != nil {
		utils.ErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = UpdateChartByID(r.Context(), a.db, dashboardID, user.ID, &payload)
	if err != nil {
		utils.ErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.MessageResponse(w, http.StatusOK, "Chart updated successfully!")
}
