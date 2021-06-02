package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/revel/revel"
	"github.com/spf13/viper"
)

type App struct {
	*revel.Controller
}
type Message struct {
	Count  int `json:"count"`
	Result []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"result"`
}

type SubCategory struct {
	Count  int `json:"count"`
	Result []struct {
		ID         int    `json:"Id"`
		Categoryid int    `json:"Categoryid"`
		Name       string `json:"Name"`
	} `json:"result"`
}

type Brand struct {
	Count  int `json:"count"`
	Result []struct {
		ID   int    `json:"Id"`
		Name string `json:"Name"`
	} `json:"result"`
}

type ViewProduct struct {
	Count  int `json:"count"`
	Result []struct {
		CategoryID    int    `json:"CategoryId"`
		SubCategoryID int    `json:"SubCategoryId"`
		BrandID       int    `json:"BrandId"`
		SupplierID    int    `json:"Supplier_id"`
		Name          string `json:"Name"`
		Headder       string `json:"Headder"`
		Description   string `json:"Description"`
	} `json:"result"`
	Images []struct {
		Id        int    `json:"Id"`
		ProductId int    `json:"ProductId"`
		Link      string `json:"Link"`
	}
}

type AllProduct struct {
	Count       int    `json:"count"`
	CurrentPage int    `json:"current_page"`
	Links       string `json:"links"`
	Offset      int    `json:"offset"`
	Pages       int    `json:"pages"`
	PerPage     int    `json:"per_page"`
	Result      []struct {
		CategoryID    int    `json:"CategoryId"`
		SubCategoryID int    `json:"SubCategoryId"`
		BrandID       int    `json:"BrandId"`
		SupplierID    int    `json:"Supplier_id"`
		Name          string `json:"Name"`
		Headder       string `json:"Headder"`
		Description   string `json:"Description"`
	} `json:"result"`
	Images []struct {
		Id        int    `json:"Id"`
		ProductId int    `json:"ProductId"`
		Link      string `json:"Link"`
	} `json:"productImages"`
	TotalRecords int `json:"total_records"`
}

type User struct {
	Count  int `json:"count"`
	Result []struct {
		ID        int    `json:"Id"`
		FirstName string `json:"First_Name"`
		LastName  string `json:"Last_Name"`
		Email     string `json:"Email"`
	} `json:"result"`
}

type session struct {
	Result string `json:"result"`
	Token  string `json:"token"`
	UserID string `json:"userID"`
}

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Service RUN on DEBUG mode")
	}

}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Product() revel.Result {
	return c.Render()
}
func (c App) Rate() revel.Result {
	return c.Render()
}
func (c App) History() revel.Result {
	return c.Render()
}
func (c App) Report() revel.Result {
	return c.Render()
}
func (c App) Payment() revel.Result {
	return c.Render()
}
func (c App) Good() revel.Result {
	return c.Render()
}
func (c App) AddProduct() revel.Result {
	return c.Render()
}
func (c App) ConsumerLoged() revel.Result {
	return c.Render()
}
func (c App) Profile() revel.Result {
	return c.Render()
}

func (c App) Category() revel.Result {

	url := viper.GetString(`api`) + "product/category"

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var m Message
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	return c.RenderJSON(m)
}

func (c App) ProductSubCategory() revel.Result {

	var page string = c.Params.Get("id")

	url := viper.GetString(`api`) + "product/subcategoryByCategory/"+page

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var s SubCategory
	err := json.Unmarshal(body, &s)

	if err != nil {
		panic(err)
	}

	return c.RenderJSON(s)
}

func (c App) Brand() revel.Result {
	url := viper.GetString(`api`) + "product/brand"

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var b Brand
	err := json.Unmarshal(body, &b)

	if err != nil {
		panic(err)
	}

	return c.RenderJSON(b)
}

func (c App) ViewProduct() revel.Result {

	url := viper.GetString(`api`) + "product/get/10"

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var v ViewProduct
	err := json.Unmarshal(body, &v)

	if err != nil {
		panic(err)
	}

	return c.RenderJSON(v)
}

func (c App) AllProduct() revel.Result {

	var page string = c.Params.Get("id")

	// var page string = "1"

	url := viper.GetString(`api`) + "product/page/" + page

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var p AllProduct
	err := json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	return c.RenderJSON(p)
}

func (c App) AddProfile() revel.Result {

	var userID = c.Session["userID"]
	fmt.Println("############################################")
	fmt.Println(userID)
	fmt.Println("############################################")
	url := viper.GetString(`api`) + "consumer/get/" + userID

	var bearer = "Bearer " + c.Session["token"]

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var u User
	err := json.Unmarshal(body, &u)

	if err != nil {
		panic(err)
	}

	return c.RenderJSON(u)
}

func (c App) FbLogin() revel.Result {

	accessToken := c.Params.Get("accessToken")

	fmt.Println(string(accessToken))

	url := viper.GetString(`api`) + "auth/consumer/fb/"

	var bearer = accessToken

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var m session
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	fmt.Println("hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh")

	fmt.Println(string(body))

	resp := string(m.Token)
	reso := string(m.UserID)

	c.Session["token"] = resp
	c.Session["id"] = reso

	fmt.Println(c.Session["token"])

	return c.RenderJSON(m)

	// return c.RenderText("accessToken: " + accessToken + " Body: " + resp)

}

func (c App) GoogleLogin() revel.Result {

	accessToken := c.Params.Get("accessToken")

	fmt.Println(string(accessToken))

	url := viper.GetString(`api`) + "auth/consumer/go/"

	var bearer = accessToken

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var m session
	err := json.Unmarshal(body, &m)

	if err != nil {
		panic(err)
	}

	fmt.Println("1111111111111111111111111111111111111111111111111111111111111")

	fmt.Println(string(body))

	resp := string(m.Token)
	reso := string(m.UserID)

	c.Session["token"] = resp
	c.Session["userID"] = reso

	fmt.Println(c.Session["token"])
	fmt.Println(c.Session["userID"])

	// return c.Redirect(App.Product)
	// // c.Redirect("/product")
	// return c.RenderText("accessToken: " + accessToken + " Body: " + resp)
	return c.RenderJSON(m)

}
