package domain

type CustomerRepositoryStub struct {
	Customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customers, nil
}
func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Jose", City: "Nairobi", Zipcode: "254", DateofBirth: "01-01-1990", Status: "Active"},
		{Id: "2", Name: "Vee", City: "Nairobi", Zipcode: "254", DateofBirth: "01-01-1990", Status: "Active"},
	}
	return CustomerRepositoryStub{customers}
}
