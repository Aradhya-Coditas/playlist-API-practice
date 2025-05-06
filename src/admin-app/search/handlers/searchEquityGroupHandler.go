package handlers

import (
	"admin-app/search/business"
	"admin-app/search/models"
	"encoding/json"
	genericConstants "omnenest-backend/src/constants"
	genericModel "omnenest-backend/src/models"
	"omnenest-backend/src/utils/logger"
	"omnenest-backend/src/utils/responseConversion"
	"omnenest-backend/src/utils/responseUtils"
	"omnenest-backend/src/utils/tracer"
	"omnenest-backend/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type SearchEquityGroupController struct {
	service *business.SearchEquityGroupService
}

func NewSearchEquityGroupController(service *business.SearchEquityGroupService) *SearchEquityGroupController {
	return &SearchEquityGroupController{
		service: service,
	}
}

// @Summary This API is used to get equity group by exchange name.
// @Description This API is used to get equity group by exchange name.
// @Tags Equity
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param xRequestId header string true "Unique request identifier" default(123456789)
// @Param deviceId header string true "Unique device identifier" default(PKQ1.180904.001)
// @Param appVersion header string true "Current app version" default(1.0.0)
// @Param source header string true "Source (MOB or WEB)" default(MOB)
// @Param bypass header string false "Bypass (AUTOMATION or CHART)"
// @Param appInstallId header string true "Unique appInstall identifier" default(ba6eb330-4f7f-11eb-a2fb-67c34e9ac07c)
// @Param userAgent header  string true "userAgent" default(BrokerAppName/3.3.6 (OnePlus ONEPLUS A6010; Android 11 SDK30))
// @Param timestamp header string true "device current day epoch milliseconds timestamp" default(1701063550000)
// @Param request body models.BFFSearchEquityGroupRequest true "Search equity group request JSON"
// @Success 200 {object} models.BFFSearchEquityGroupResponse "Successful response"
// @Failure 400 {object}  models.ErrorAPIResponse "Bad Request: Invalid input data or validation error"
// @Failure 204 "No Content for the request"
// @Failure 500 {object} models.ErrorAPIResponse "The server encountered an unexplained problem which has prevented it from executing the given request"
// @Router /api/search/equity/group [post]
func (controller *SearchEquityGroupController) HandleSearchEquityGroup(ctx *gin.Context) {
	spanCtx, span := tracer.AddToSpan(ctx.Request.Context(), "HandleSearchEquityGroup")
	defer func() {
		if span != nil {
			span.End()
		}
	}()
	log := logger.GetLogger(ctx)

	var bffSearchEquityGroupRequest models.BFFSearchEquityGroupRequest

	if err := ctx.ShouldBindJSON(&bffSearchEquityGroupRequest); err != nil {
		errorMsgs := genericModel.ErrorMessage{Key: err.(*json.UnmarshalTypeError).Field, ErrorMessage: genericConstants.JsonBindingFieldError}
		log.With(zap.Error(err)).Error(err.Error())
		responseUtils.SendBadRequest(ctx, []genericModel.ErrorMessage{errorMsgs})
		return
	}

	ctx.Set(genericConstants.RequestBody, bffSearchEquityGroupRequest)

	if err := validations.GetBFFValidator(spanCtx).Struct(bffSearchEquityGroupRequest); err != nil {
		validationErrors, validationErrorsStr := validations.FormatValidationErrors(spanCtx, err.(validator.ValidationErrors))
		log.With(zap.Error(err)).Error(validationErrorsStr)
		responseUtils.SendBadRequest(ctx, validationErrors)
		return
	}
	responseConversion.ConvertFieldValues(spanCtx, &bffSearchEquityGroupRequest, true)
	bffSearchEquityGroup, err := controller.service.SearchEquityGroup(ctx, spanCtx, bffSearchEquityGroupRequest.ExchangeName)
	if err != nil {
		log.With(zap.Error(err)).Error(err.Error())
		if strings.Contains(strings.ToLower(err.Error()), genericConstants.NoDataFoundError) {
			responseUtils.SendNoContentFoundError(ctx, err)
			return
		}
		responseUtils.SendInternalServerError(ctx, err)
		return
	}
	responseUtils.SendStatusOK(ctx, genericConstants.BFFResponseSuccessMessage, bffSearchEquityGroup)
}
