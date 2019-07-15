package main
import "encoding/json"
import "net/http"
import "io/ioutil"
import "fmt"
import "strconv"
import "time"

func creatingSampleHttpRequest() {

	for i := 1; i <= 10; i++ {
		response, err := http.Get("http://5a530e1477e1d20012fa066a.mockapi.io/employeedata/" + strconv.Itoa(i))
		if err == nil {
			responseBody, readErr := ioutil.ReadAll(response.Body) 
			
			if readErr == nil {
				fmt.Println(responseBody)
				empList := new(Employee)
				json.Unmarshal(responseBody, &empList)

				fmt.Println(empList.Id)
			}
		}
	}
 
	time.Sleep(10000 * time.Millisecond)
}

type Employee struct {
    Id   		string `json:"id"`
    Name    	string    `json:"name"`
	Avatar 		string `json:"avatar"`
	CreatedAt 	time.Time `json:"createdAt"`
}