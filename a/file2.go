// // How to make get request in golang
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// func main() {
// 	fmt.Printf("How to make get request in golang\n")
// 	PerformGetRequest()
// }
// func PerformGetRequest() {
// 	const url string = "https://www.themealdb.com/api/json/v1/1/filter.php?c=Chicken"
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	content, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Printf("response is %s", string(content))

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// )

// // Define the structure of the JSON response
// type Meal struct {
// 	StrMeal      string `json:"strMeal"`
// 	StrMealThumb string `json:"strMealThumb"`
// 	IdMeal       string `json:"idMeal"`
// }

// type MealsResponse struct {
// 	Meals []Meal `json:"meals"`
// }

// func main() {
// 	// fmt.Println("Making a GET request in Go")
// 	// PerformGetRequest()
// }

// func PerformGetRequest() {
// 	const url string = "https://www.themealdb.com/api/json/v1/1/search.php?s="

// 	// Perform the GET request
// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Read the response body
// 	// content, err := ioutil.ReadAll(response.Body)
// 	// var responseString strings.Builder
// 	content, _ := ioutil.ReadAll(response.Body)
// 	// byteCount, _ := responseString.WriteString(content)

// 	if err != nil {
// 		fmt.Println("Error reading response body:", err)
// 		return
// 	}
// 	// fmt.Println("Response Body:", byteCount)
// 	// fmt.Println(responseString.String())

// 	// Parse the JSON response
// 	var mealsResponse MealsResponse
// 	err = json.Unmarshal(content, &mealsResponse)
// 	if err != nil {
// 		fmt.Println("Error parsing JSON:", err)
// 		return
// 	}

//		// Print the meals data
//		for _, meal := range mealsResponse.Meals {
//			fmt.Printf("Meal: %s\nImage URL: %s\nID: %s\n\n", meal.StrMeal, meal.StrMealThumb, meal.IdMeal)
//		}
//	}
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// func main() {
// 	fmt.Println("How to make post request in golang")
// 	PerformPostRequest()
// }
// func PerformPostRequest() {
// 	const myurl string = "https://reqres.in/api/users"
// 	responsebody := strings.NewReader(`
// {
// "name":"bharti",
// "job":"developer"
// }
// `)
// 	response, err := http.Post(myurl, "application/json", responsebody)

// 	if err != nil {
// 		panic(err)

// 	}
// 	defer response.Body.Close()
// 	content, err := ioutil.ReadAll(response.Body)
// 	fmt.Printf("response is %s", string(content))

// }
// how to send form data in golang(work on post api )
// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"net/url"
// )

// func main() {
// 	fmt.Println("How to send form data in golang")
// 	PerformPostFormRequest()
// }
// func PerformPostFormRequest() {
// 	const myurl string = "https://reqres.in/api/users"
// 	data := url.Values{}
// 	data.Add(
// 		"name", "bharti")
// 	data.Add("job", "developer")
// 	data.Add("id", "1")
// 	data.Add("email", "bharti@123gmail.com")
// 	response, err := http.PostForm(myurl, data)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer response.Body.Close()
// 	content, err := ioutil.ReadAll(response.Body)
// 	fmt.Printf("response is %s", string(content))

// }
// How to create json data in go lang
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	fmt.Println("How to create json data in go lang")
	Encodejsondata()
	Decodejsondata()
}

type course struct {
	Name     string `json:"CourseName"`
	Platform string
	Price    int      `json:_`
	Tag      []string `json:"tags,omitempty"`
}

func Encodejsondata() {
	listofcourses := []course{
		{
			"React", "Udemy",
			299, []string{"web development", "programming", "react"},
		},
		{
			"Pandas", "Datacamp", 999, []string{"data science", "python", "pandas"},
		},
		{
			"Angular", "Coursera", 199, []string{"web development", "programming", "angular"}},
		{
			"Spring Boot", "Udemy", 249, nil,
		},
	}
	jsondata, err := json.MarshalIndent(listofcourses, " ", "\t")
	if err != nil {
		fmt.Println("Error in encoding json data:", err)
		return
	}
	fmt.Println(string(jsondata))

}

// How to consume json data in go lang
func Decodejsondata() {
	jsondatafromweb := []byte(`{
                "CourseName": "Angular",
                "Platform": "Coursera",
                "Price": 199,
                "tags": ["web development","programming","angular"]
        }`)
	var mycourse course
	checkvalid := json.Valid(jsondatafromweb)
	if checkvalid {
		fmt.Println("json data is valid")
		json.Unmarshal(jsondatafromweb, &mycourse)
		fmt.Printf("%#v\n", mycourse)
	} else {
		fmt.Println("json data is invalid")
	}

}
