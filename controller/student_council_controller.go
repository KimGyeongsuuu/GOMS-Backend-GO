package controller

import (
	"GOMS-BACKEND-GO/model"
	"GOMS-BACKEND-GO/model/data/constant"
	"GOMS-BACKEND-GO/model/data/input"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StudentCouncilController struct {
	studentCouncilUseCase model.StudentCouncilUseCase
}

func NewStudentCouncilController(studentCouncilUseCase model.StudentCouncilUseCase) *StudentCouncilController {
	return &StudentCouncilController{
		studentCouncilUseCase: studentCouncilUseCase,
	}
}

func (controller *StudentCouncilController) CreateOuting(ctx *gin.Context) {

	outingUUID, err := controller.studentCouncilUseCase.CreateOuting(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"outingUUID": outingUUID.String()})

}

func (controller *StudentCouncilController) FindAccountList(ctx *gin.Context) {

	accounts, err := controller.studentCouncilUseCase.FindAllAccount(ctx)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (controller *StudentCouncilController) SearchAccountByInfo(ctx *gin.Context) {

	grade := ctx.Query("grade")
	gender := ctx.Query("gender")
	name := ctx.Query("name")
	isBlackList := ctx.Query("isBlackList")
	authority := ctx.Query("authority")
	major := ctx.Query("major")

	var input input.SearchAccountInput

	if grade != "" {
		grade, err := strconv.Atoi(grade)
		if err == nil {
			input.Grade = &grade
		}
	}

	if gender != "" {
		gender := constant.Gender(gender)
		input.Gender = &gender
	}

	if name != "" {
		input.Name = &name
	}

	if isBlackList != "" {
		isBlackList, err := strconv.ParseBool(isBlackList)
		if err == nil {
			input.IsBlackList = &isBlackList
		}
	}

	if authority != "" {
		authority := constant.Authority(authority)
		input.Authority = &authority
	}

	if major != "" {
		major := constant.Major(major)
		input.Major = &major
	}

	accounts, err := controller.studentCouncilUseCase.SearchAccount(ctx, &input)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func (controller *StudentCouncilController) UpdateAuthority(ctx *gin.Context) {
	var input input.UpdateAccountAuthorityInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.Error(err)
		return
	}

	err := controller.studentCouncilUseCase.UpdateAccountAuthority(ctx, &input)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *StudentCouncilController) AddBlackList(ctx *gin.Context) {
	accountIDParam := ctx.Param("accountID")

	accountID, err := primitive.ObjectIDFromHex(accountIDParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = controller.studentCouncilUseCase.AddBlackList(ctx, accountID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *StudentCouncilController) DeleteBlackList(ctx *gin.Context) {
	accountIDParam := ctx.Param("accountID")

	accountID, err := primitive.ObjectIDFromHex(accountIDParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = controller.studentCouncilUseCase.ExcludeBlackList(ctx, accountID)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)

}

func (controller *StudentCouncilController) DeleteOutingStudent(ctx *gin.Context) {
	accountIDParam := ctx.Param("accountID")

	accountID, err := primitive.ObjectIDFromHex(accountIDParam)
	if err != nil {
		ctx.Error(err)
		return
	}

	err = controller.studentCouncilUseCase.DeleteOutingStudent(ctx, accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (controller *StudentCouncilController) FindLateList(ctx *gin.Context) {
	dateStr := ctx.Query("date")

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	lateStudents, err := controller.studentCouncilUseCase.FindLateStudentByDate(ctx, date)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch late students."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"late-stduents": lateStudents})
}
