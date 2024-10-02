package Controllers

import (
	"log"
	"strconv"

	"net/http"
	"wan-api-kol-event/Const"
	"wan-api-kol-event/Logic"
	"wan-api-kol-event/ViewModels"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetKolsController(context *gin.Context) {
	var KolsVM ViewModels.KolViewModel
	var guid = uuid.New().String()

	log.Println("Failed to start server")
	// * Get Kols from the database based on the range of pageIndex and pageSize
	// @params: pageIndex
	pageIndex, err := strconv.Atoi(context.DefaultQuery("pageIndex", "1"))
	if err != nil || pageIndex < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageIndex"})
		log.Printf("Lỗi khi lấy pageIndex: %v", err)
		return
	}

	log.Printf("pageIndex nhận được từ request: %d", pageIndex)

	// @params: pageSize
	pageSize, err := strconv.Atoi(context.DefaultQuery("pageSize", "10"))

	if err != nil || pageSize < 1 {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize"})
		log.Printf("Lỗi khi lấy pageSize: %v", err)
		return
	}

	log.Printf("pageSize nhận được từ request: %d", pageSize)

	// * TODO: Implement the logic to get parameters from the request
	// ? If parameter passed in the request is not valid, return the response with HTTP Status Bad Request (400)
	// * Perform Logic Here
	// ! Pass the parameters to the Logic Layer
	kols, error := Logic.GetKolLogic(pageIndex, pageSize)

	if error != nil {
		KolsVM.Result = Const.UnSuccess
		KolsVM.ErrorMessage = error.Error()
		KolsVM.PageIndex = int64(pageIndex) // * change this to the actual page index from the request
		KolsVM.PageSize = int64(pageSize)   // * change this to the actual page size from the request
		KolsVM.Guid = guid
		context.JSON(http.StatusInternalServerError, KolsVM)
		return
	}

	// * Return the response after the logic is executed
	// ? If the logic is successful, return the response with HTTP Status OK (200)
	KolsVM.Result = Const.Success
	KolsVM.ErrorMessage = ""
	KolsVM.PageIndex = int64(pageIndex) // * change this to the actual page index from the request
	KolsVM.PageSize = int64(pageSize)   // * change this to the actual page size from the request
	KolsVM.Guid = guid
	KolsVM.KOL = kols
	KolsVM.TotalCount = int64(len(kols))
	context.JSON(http.StatusOK, KolsVM)
}
