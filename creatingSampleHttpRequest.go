package main
import "encoding/json"
import "net/http"
import "io/ioutil"
import "fmt"
import "strconv"
import "time"
import "runtime"

func main() {

	runtime.GOMAXPROCS(8) 
	
	empList := []Employee{}
	response, err := http.Get("http://5c055de56b84ee00137d25a0.mockapi.io/api/v1/employees")
	if err == nil {
		responseBody, readErr := ioutil.ReadAll(response.Body) 
		if readErr == nil {
			json.Unmarshal(responseBody, &empList)

			for index, emp := range(empList) {
				go fmt.Println("Value at " + strconv.Itoa(index) + " is " + emp.Name)
			}
		}
	}

	time.Sleep(1000 * time.Millisecond)
}

type Employee struct {
    Id   string `json:"id"`
    Name    string    `json:"name"`
	Avatar string `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}