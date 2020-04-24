package main

import (
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api"
	"github.com/nafisfaysal/erply/cfg"
	"github.com/nafisfaysal/erply/errors"
	"strconv"
)

type erplyApiService struct {
	api.IClient
}

func NewErplyApiService(sessionKey string, clientCode string) *erplyApiService {
	return &erplyApiService{api.NewClient(sessionKey, clientCode, nil)}
}

func (s *erplyApiService) getPointsOfSale(posID string) (string, error) {
	res, err := s.GetPointsOfSaleByID(posID)
	if err != nil {
		return "", err
	}
	return strconv.Itoa(res.WarehouseID), nil
}

type CompanyInfo struct {
	ID      string
	Name    string
	Email   string
	Web     string
	Country string
}

func (s *erplyApiService) getCompanyInfo() (CompanyInfo, error) {
	res, err := s.GetCompanyInfo()
	if err != nil {
		fmt.Println("got an  error: ", err)
	}

	comInfo := CompanyInfo{
		ID:      res.ID,
		Name:    res.Name,
		Email:   res.Email,
		Web:     res.Web,
		Country: res.Country,
	}

	return comInfo, nil
}

func (s *erplyApiService) getUserName() (string, error) {
	res, err := s.GetUserName()
	if err != nil {
		fmt.Println(err)
	}

	return res, nil
}

type Customer struct {
	ID          int
	CustomerID  int
	TypeID      string
	FullName    string
	CompanyName string
	FirstName   string
	LastName    string
	GroupID     int
}

func (s *erplyApiService) getSupplierByName(name string) (Customer, error) {
	res, err := s.GetSupplierByName(name)
	if err != nil {
		fmt.Println(err)
	}

	ctmr := Customer{
		ID:          res.ID,
		CustomerID:  res.CustomerID,
		TypeID:      res.TypeID,
		FullName:    res.FullName,
		CompanyName: res.CompanyName,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		GroupID:     res.GroupID,
	}

	return ctmr, nil
}

func main() {
	err := cfg.Load()
	if err != nil {
		errors.CheckError(err)
	}

	fmt.Println(cfg.UserName, "clinet", cfg.ClientCode, cfg.Password)

	ses, err := api.VerifyUser(cfg.UserName, cfg.Password, cfg.ClientCode)
	if err != nil {
		errors.CheckError(err)
	}

	s := NewErplyApiService(ses, cfg.ClientCode)

	userName, err := s.getUserName()
	if err != nil {
		errors.CheckError(err)
	}

	fmt.Println(userName)

	ids := []string{"1", "2", "3"}
	customerList, err := s.GetCustomersByIDs(ids)
	if err != nil {
		errors.CheckError(err)
	}
	fmt.Println(customerList)

	PointOfSale, err := s.GetPointsOfSaleByID("1")
	if err != nil {
		errors.CheckError(err)
	}
	fmt.Println(PointOfSale)
}
