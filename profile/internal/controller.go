package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"social-network/lib/mysql"
	"social-network/profile/internal/types"

	"github.com/labstack/echo/v4"
)

// ShowAccount godoc
// @Summary      Show an account
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (o *App) Register(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.RegisterRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	auth, err := r.NewAuth()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	profile, err := r.NewProfile()

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fn := func(ctx context.Context, tx *sql.Tx) error {

		qr, err := auth.InsertAuth().Exec(ctx, tx)
		if err != nil {
			return err
		}

		profile.UserId = qr.LastInsertId

		qr, err = profile.UpsertProfile().Exec(ctx, tx)
		if err != nil {
			return err
		}

		if qr.RowsAffected == 0 {
			return errors.New("profile was not saved")
		}

		return nil
	}

	err = mysql.BeginTxFunc(ctx, nil, o.Db, fn)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, profile)
}

func (o *App) Profiles(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.ProfilesRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pr := new(types.PageRequest)

	if err := c.Bind(pr); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var profiles = make([]types.Profile, 0)

	err := r.ReadProfilesPage(pr).Query(ctx, o.Db, profiles)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var total int64

	err = r.ReadProfilesTotal().QueryOne(ctx, o.Db, &total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := types.NewPageResponse(pr, profiles, total)

	return c.JSON(http.StatusOK, result)
}

func (o *App) Profile(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.UserIdRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	result := new(types.Profile)

	err := r.ReadProfileByUserId().QueryOne(ctx, o.Db, result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func (o *App) SaveProfile(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Profile)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	qr, err := r.UpsertProfile().Exec(ctx, o.Db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if qr.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, "profile was not saved")
	}

	return c.JSON(http.StatusOK, r)
}

func (o *App) Friends(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.UserIdRequest)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	pr := new(types.PageRequest)

	if err := c.Bind(pr); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	var profiles = make([]types.Profile, 0)

	err := r.ReadUserFriendsProfiles(pr).Query(ctx, o.Db, profiles)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var total int64

	err = r.ReadUserFriendsTotal().QueryOne(ctx, o.Db, &total)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result := types.NewPageResponse(pr, profiles, total)

	return c.JSON(http.StatusOK, result)
}

func (o *App) AddFriend(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Friend)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	qr, err := r.InsertFriend().Exec(ctx, o.Db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if qr.RowsAffected == 0 {
		return c.JSON(http.StatusInternalServerError, "friend was not added")
	}

	return c.JSON(http.StatusOK, r)
}

func (o *App) DeleteFriend(c echo.Context) error {
	ctx := c.Request().Context()
	r := new(types.Friend)

	if err := c.Bind(r); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	_, err := r.DeleteFriend().Exec(ctx, o.Db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, r)
}
