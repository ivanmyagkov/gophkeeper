package v2

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"gophkeeper/internal/domain"
)

func (h *Handler) initMaterialsRoutes(gr *echo.Group) {
	materialsGr := gr.Group("/materials")

	authGr := materialsGr.Group("", h.checkUserIdentity)
	authGr.GET("/text", h.getAllTextData)
	authGr.POST("/text", h.CreateNewTextData)
	authGr.PUT("/text", h.UpdateTextDataByID)
	authGr.GET("/cred", h.getAllCredData)
	authGr.POST("/cred", h.CreateNewCredData)
	authGr.PUT("/cred", h.UpdateCredDataByID)
	authGr.GET("/card", h.getAllCardData)
	authGr.POST("/card", h.CreateNewCardData)
	authGr.PUT("/card", h.UpdateCardDataByID)
	authGr.GET("/blob", h.getAllBlobData)
	authGr.POST("/blob", h.CreateNewBlobData)
	authGr.PUT("/blob", h.UpdateBlobDataByID)
}

func (h Handler) getAllTextData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	dataArray, err := h.services.Materials.GetAllTextData(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDataNotFound):
			return c.NoContent(http.StatusNoContent)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}
	if len(dataArray) == 0 {
		return c.NoContent(http.StatusNoContent)
	}

	return c.JSON(http.StatusOK, dataArray)
}

func (h Handler) UpdateTextDataByID(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp domain.TextData
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.UpdateTextDataByID(c.Request().Context(), userID, inp)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type newTextDataInput struct {
	Text     string `json:"text"`
	Metadata string `json:"metadata"`
}

func (h Handler) CreateNewTextData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp newTextDataInput
	if err := c.Bind(&inp); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	log.Println(inp.Text)

	err := h.services.Materials.CreateNewTextData(c.Request().Context(), userID, domain.TextData{
		ID:       -1, //this field fill be ignored
		Text:     inp.Text,
		Metadata: inp.Metadata,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) getAllCredData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	dataArray, err := h.services.Materials.GetAllCredData(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDataNotFound):
			return c.NoContent(http.StatusNoContent)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, dataArray)
}

func (h Handler) UpdateCredDataByID(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp domain.CredData
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.UpdateCredDataByID(c.Request().Context(), userID, inp)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type newCredDataInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Metadata string `json:"metadata"`
}

func (h Handler) CreateNewCredData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp newCredDataInput
	if err := c.Bind(&inp); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.CreateNewCredData(c.Request().Context(), userID, domain.CredData{
		ID:       -1, //this field fill be ignored
		Login:    inp.Login,
		Password: inp.Password,
		Metadata: inp.Metadata,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) getAllCardData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	dataArray, err := h.services.Materials.GetAllCardData(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDataNotFound):
			return c.NoContent(http.StatusNoContent)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, dataArray)
}

func (h Handler) UpdateCardDataByID(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp domain.CardData
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.UpdateCardDataByID(c.Request().Context(), userID, inp)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type newCardDataInput struct {
	CardNumber string    `json:"card_number"`
	ExpDate    time.Time `json:"exp_date"`
	CVC        string    `json:"cvc"`
	Name       string    `json:"name"`
	Surname    string    `json:"surname"`
	Metadata   string    `json:"metadata"`
}

func (h Handler) CreateNewCardData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp newCardDataInput
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.CreateNewCardData(c.Request().Context(), userID, domain.CardData{
		ID:         -1, //this field fill be ignored
		CardNumber: inp.CardNumber,
		ExpDate:    inp.ExpDate,
		CVC:        inp.CVC,
		Name:       inp.Name,
		Surname:    inp.Surname,
		Metadata:   inp.Metadata,
	})
	log.Println(parseExpireDate(inp.ExpDate.String()))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func (h Handler) getAllBlobData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	dataArray, err := h.services.Materials.GetAllBlobData(c.Request().Context(), userID)
	if err != nil {
		switch {
		case errors.Is(err, domain.ErrDataNotFound):
			return c.NoContent(http.StatusNoContent)
		default:
			return echo.NewHTTPError(http.StatusInternalServerError)
		}
	}

	return c.JSON(http.StatusOK, dataArray)
}

func (h Handler) UpdateBlobDataByID(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	var inp domain.BlobData
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err := h.services.Materials.UpdateBlobDataByID(c.Request().Context(), userID, inp)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

type newBlobDataInput struct {
	Data     []byte `json:"data"`
	Metadata string `json:"metadata"`
}

func (h Handler) CreateNewBlobData(c echo.Context) error {
	userID := c.Get(UserIDCtxName.String()).(int)

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, src); err != nil {
		return err
	}

	var inp newBlobDataInput
	inp.Data = buf.Bytes()
	inp.Metadata = c.FormValue("metadata")
	if err := c.Bind(&inp); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = h.services.Materials.CreateNewBlobData(c.Request().Context(), userID, domain.BlobData{
		ID:       -1, //this field fill be ignored
		Data:     inp.Data,
		Metadata: inp.Metadata,
	})

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusOK)
}

func parseExpireDate(exp string) time.Time {
	date, err := time.Parse("01/06", exp)
	if err != nil {
		log.Panicln(err.Error())
	}
	return date
}
