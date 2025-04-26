package controller

import (
	"net/http"
	"opdServices/model/web"
	"opdServices/service"

	"github.com/labstack/echo/v4"
)

type OpdControllerImpl struct {
	opdService service.OpdService
}

func NewOpdControllerImpl(opdService service.OpdService) *OpdControllerImpl {
	return &OpdControllerImpl{opdService: opdService}
}

// Create godoc
// @Summary Create Opd
// @Description Create a new opd
// @Tags Opd
// @Accept json
// @Produce json
// @Param opd body web.OpdCreateRequest true "Create Opd"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/create [post]
func (controller *OpdControllerImpl) Create(c echo.Context) error {
	OpdCreateRequest := web.OpdCreateRequest{}
	if err := c.Bind(&OpdCreateRequest); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	opdResponse, err := controller.opdService.Create(c.Request().Context(), OpdCreateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Create 	OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Create OPD",
		Data:   opdResponse,
	})
}

// Update godoc
// @Summary Update Opd
// @Description Update an existing opd
// @Tags Opd
// @Accept json
// @Produce json
// @Param id path string true "Opd ID"
// @Param opd body web.OpdUpdateRequest true "Update Opd"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/update/{id} [put]
func (controller *OpdControllerImpl) Update(c echo.Context) error {
	OpdUpdateRequest := web.OpdUpdateRequest{}
	if err := c.Bind(&OpdUpdateRequest); err != nil {
		return err
	}

	id := c.Param("id")

	OpdUpdateRequest.Id = id

	opdResponse, err := controller.opdService.Update(c.Request().Context(), OpdUpdateRequest)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Update OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Update OPD",
		Data:   opdResponse,
	})
}

// Delete godoc
// @Summary Delete Opd
// @Description Delete an existing opd
// @Tags Opd
// @Accept json
// @Produce json
// @Param id path string true "Opd ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/delete/{id} [delete]
func (controller *OpdControllerImpl) Delete(c echo.Context) error {
	id := c.Param("id")

	err := controller.opdService.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Delete OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Delete OPD",
		Data:   nil,
	})
}

// FindById godoc
// @Summary FindById opd
// @Description Find By Id an existing opd
// @Tags Opd
// @Accept json
// @Produce json
// @Param id path string true "Opd ID"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/detail/{id} [get]
func (controller *OpdControllerImpl) FindById(c echo.Context) error {
	id := c.Param("id")

	opdResponse, err := controller.opdService.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find OPD",
				Data:   err.Error(),
			},
		)
	}
	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find OPD",
		Data:   opdResponse,
	})
}

// FindAll godoc
// @Summary FindAll opd
// @Description FindAll an existing opd
// @Tags Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/findall [get]
func (controller *OpdControllerImpl) FindAll(c echo.Context) error {
	opdResponse, err := controller.opdService.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find All OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All OPD",
		Data:   opdResponse,
	})
}

// FindAllOnlyOpd godoc
// @Summary Find all only opd
// @Description Find all only opd for dropdown
// @Tags Opd
// @Accept json
// @Produce json
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/findallopd [get]
func (controller *OpdControllerImpl) FindAllOnlyOpd(c echo.Context) error {
	opdResponse, err := controller.opdService.FindAllOnlyOpd(c.Request().Context())
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find All OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find All OPD",
		Data:   opdResponse,
	})
}

// @Summary FindByKodeOpd opd
// @Description FindByKodeOpd an existing opd
// @Tags Opd
// @Accept json
// @Produce json
// @Param kode_opd path string true "Kode OPD"
// @Success 200 {object} web.WebResponse
// @Failure 400 {object} web.WebResponse
// @Failure 500 {object} web.WebResponse
// @Router /opd/findbyopd/{kode_opd} [get]
func (controller *OpdControllerImpl) FindByKodeOpd(c echo.Context) error {
	opdResponse, err := controller.opdService.FindByKodeOpd(c.Request().Context(), c.Param("kode_opd"))
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			web.WebResponse{
				Code:   http.StatusInternalServerError,
				Status: "Failed Find OPD by Kode OPD",
				Data:   err.Error(),
			},
		)
	}

	return c.JSON(http.StatusOK, web.WebResponse{
		Code:   http.StatusOK,
		Status: "Success Find OPD by Kode OPD",
		Data:   opdResponse,
	})
}
