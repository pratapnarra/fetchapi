package handlers

import(
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/pratapnarra/fetchapi/models"
	"unicode"
	"net/http"
	"strconv"
	"strings"
	"math"
	"time"
)

func PostHandler(w http.ResponseWriter, r *http.Request){
	var receipt models.Receipt
	json.NewDecoder(r.Body).Decode(&receipt)
	
	
	newID := uuid.New().String()

	models.MapMutex.Lock()
	models.PointsMap[newID] = CalculatePoints(w,receipt)
	models.MapMutex.Unlock()

	response := models.PostResponse{
		ID: newID,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, string(jsonResponse))
	
}

func CalculatePoints(w http.ResponseWriter,receipt models.Receipt) int{
	var apoints int = 0

	//1 Count alpha numeric characters in the Name of the retailer
	for _, char := range receipt.Retailer {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			apoints++
		}
	}

	//2 Check if total is round dollar amount
	f, err := strconv.ParseFloat(receipt.Total, 64)
	if err!=nil{
				fmt.Println(err)
				http.Error(w, "Total Not Valid", http.StatusBadRequest)
				return 0
			}
	if( float64(int(f)) == f ){
       apoints += 50
	}
	
	//3 Check if multiple of 0.25
	tolerance := 1e-9 // A small tolerance value for comparison
	remainder := math.Mod(f, 0.25)
	if(math.Abs(remainder) < tolerance){
		apoints +=25
	}
	

	//4 5 points for every 2 items
	apoints += (5 * (len(receipt.Items)/2))

	//5 0.2 * price if item desc is multiple of 3
	for _,item := range receipt.Items{
		trimed := strings.TrimSpace(item.ShortDescription)
		if( len(trimed)%3 ==0 ){
			itemp, err := strconv.ParseFloat(item.Price, 64)
			if err!=nil{
				fmt.Println(err)
				http.Error(w, "Item Price Not Valid", http.StatusBadRequest)
				return 0
			}
			itemp *= 0.2
			apoints += int(math.Ceil(itemp))
		}
	}

	//6 6 points if purchase date is odd
	//date is always of length 10
	date,err := strconv.Atoi(receipt.PurchaseDate[8:])
	if err!=nil{
		fmt.Println(err)
		http.Error(w, "Date Not Valid", http.StatusBadRequest)
		return 0
	}
	if (date%2!=0){
      apoints +=6
	}

	//7 Purchase time between 2pm to 4pm
	timeLayout := "2006-01-02T15:04"
	
	t,err := time.Parse(timeLayout,receipt.PurchaseDate+"T"+receipt.PurchaseTime)
	start,_ := time.Parse(timeLayout,receipt.PurchaseDate+"T14:00")
	end,_ := time.Parse(timeLayout,receipt.PurchaseDate+"T16:00")
	
	if err!=nil{
		fmt.Println(err)
		http.Error(w, "Time not valid", http.StatusBadRequest)
		return 0
	}
	
	if (t.After(start) && t.Before(end)){
		apoints +=10
	}

	return apoints
}

