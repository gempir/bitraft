package server

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *Server) handleGetEmoteHistory(c echo.Context) error {
	username := c.Param("username")
	userID := ""

	if username == "" {
		auth, _, err := s.authenticate(c)
		if err != nil {
			return err
		}
		userID = auth.Data.UserID

		if c.QueryParam("managing") != "" {
			userID, err = s.checkEditor(c, s.getUserConfig(userID))
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
		}
	} else {
		user, err := s.helixClient.GetUserByUsername(username)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		userID = user.ID
	}

	page := c.QueryParam("page")
	if page == "" {
		page = "1"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, s.db.GetEmoteHistory(userID, pageNumber, PAGE_SIZE))
}
