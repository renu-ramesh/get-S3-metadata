package cmd

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"aws-lambda-in-go-lang/helpers"
	"aws-lambda-in-go-lang/models"
	"aws-lambda-in-go-lang/pkg/contenttype"

	"go.uber.org/zap"
)

var (
	Logger *zap.Logger
)

type Handler struct {
	Ctx         context.Context
	Evt         json.RawMessage
	Logger      *zap.Logger
	Helpers     helpers.Service
	ContentType contenttype.Service
}

func Lambdahandler(ctx context.Context, evt json.RawMessage) (models.LambdaResponse, helpers.Error) {
	logger, _ := zap.NewProduction()
	handler := &Handler{
		Ctx:         ctx,
		Evt:         evt,
		Logger:      logger,
		Helpers:     helpers.NewHelperService(),
		ContentType: contenttype.NewContentType(logger),
	}
	response, err := handler.HandlerExecution()
	if err != nil {
		return models.LambdaResponse{
			Message: "error occured",
			Success: false,
			Error: map[string]string{
				"err": err.Err().Error(),
				"msg": err.Msg(),
			},
		}, err
	}
	return response, err
}

func (handler *Handler) HandlerExecution() (models.LambdaResponse, helpers.Error) {

	var eventParams models.InputPayLoadParams
	json.Unmarshal(handler.Evt, &eventParams)

	// load events params
	payLoad := eventParams.PayLoad
	var inputPayloadMap string
	if payLoad != "" {
		inputPayloadMap = payLoad
	}

	// parse and load configuration
	data, err := base64.StdEncoding.DecodeString(inputPayloadMap)
	if err != nil {
		handler.Logger.Error("processing payload",
			zap.Any("err", err),
			zap.Any("input", inputPayloadMap),
		)
		return models.LambdaResponse{}, helpers.NewError(
			helpers.ErrBadRequest,
			fmt.Sprintf("error happened while process your request error : %v", err),
		)
	}

	// check data exists
	if len(data) == 0 {
		handler.Logger.Error("processing payload",
			zap.Any("err", err),
			zap.Any("input", inputPayloadMap),
		)
		return models.LambdaResponse{}, helpers.NewError(
			helpers.ErrBadRequest, "No request payload found",
		)
	}

	// marshal into json
	var inputParams models.RequestParams
	err = json.Unmarshal(data, &inputParams)
	if err != nil {
		handler.Logger.Error("json marshal",
			zap.Any("err", err),
			zap.Any("input", inputPayloadMap),
		)
		return models.LambdaResponse{}, helpers.NewError(
			helpers.ErrBadRequest, "json marshal",
		)
	}

	handler.Logger.Info("Requested with inputs", zap.Any("input", inputParams))

	// validate the inputs
	// validationError := handler.Helpers.ValidateInput(gojsonschema.NewStringLoader(string(data)), "RequestParams")
	// if validationError != nil {
	// 	handler.Logger.Error("validation error",
	// 		zap.Any("err", err),
	// 		zap.Any("input", inputPayloadMap),
	// 	)
	// 	return models.LambdaResponse{}, helpers.NewError(
	// 		helpers.ErrBadRequest, fmt.Sprintf("validation error %v", validationError),
	// 	)
	// }

	// if the input validation passes
	mimeData, err := handler.ContentType.DetectContentType(inputParams)
	if err != nil {
		handler.Logger.Error("error occured when detecting MIME Type",
			zap.Any("err", err),
			zap.Any("input", inputPayloadMap),
		)
		return models.LambdaResponse{}, helpers.NewError(
			helpers.ErrBadRequest, fmt.Sprintf("MIME Type detection error %v", err),
		)
	}

	if mimeData == nil {
		mimeData = &models.LambdaResponse{
			Message: "unexpected error occured, please see the log",
			Success: false,
		}
	}
	return *mimeData, nil
}
