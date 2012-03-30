package service_test

import (
	. "github.com/timeredbull/tsuru/api/app"
	. "github.com/timeredbull/tsuru/api/service"
	. "github.com/timeredbull/tsuru/database"
	. "launchpad.net/gocheck"
	"fmt"
)

func (s *ServiceSuite) createServiceApp() {
	st := ServiceType{Name: "mysql", Charm: "mysql"}
	st.Create()
	fmt.Println()
	fmt.Println(st.Id)
	fmt.Println()
	s.serviceApp = &ServiceApp{
		ServiceId: st.Id,
		AppId:     1,
	}
	s.serviceApp.Create()
}

func (s *ServiceSuite) TestCreateServiceApp(c *C) {
	s.createServiceApp()
	rows, err := Db.Query("SELECT service_id, app_id FROM service_apps WHERE service_id = 2 AND app_id = 1")
	c.Check(err, IsNil)

	var serviceId int
	var appId int

	for rows.Next() {
		rows.Scan(&serviceId, &appId)
	}

	c.Assert(s.serviceApp.Id, Not(Equals), 0)
	c.Assert(serviceId, Equals, 2)
	c.Assert(appId, Equals, 1)
}

func (s *ServiceSuite) TestDeleteServiceApp(c *C) {
	s.createServiceApp()
	s.serviceApp.Delete()

	rows, err := Db.Query("SELECT count(*) FROM service_apps WHERE service_id = 2 AND app_id = 1")
	c.Assert(err, IsNil)

	var qtd int
	for rows.Next() {
		rows.Scan(&qtd)
	}

	c.Assert(qtd, Equals, 0)
}

func (s *ServiceSuite) TestRetrieveAssociatedService(c *C) {
	st := ServiceType{Name: "mysql", Charm: "mysql"}
	st.Create()
	service := Service{Name: "my_service", ServiceTypeId: st.Id}
	service.Create()

	s.serviceApp = &ServiceApp{
		ServiceId: service.Id,
		AppId:     1,
	}
	s.serviceApp.Create()

	retrievedService := s.serviceApp.Service()

	c.Assert(service.Name, Equals, retrievedService.Name)
	c.Assert(service.Id, Equals, retrievedService.Id)
	c.Assert(service.ServiceTypeId, Equals, retrievedService.ServiceTypeId)
}

func (s *ServiceSuite) TestRetrieveAssociatedApp(c *C) {
	app := App{Name: "my_app", Framework: "django"}
	app.Create()

	st := ServiceType{Name: "mysql", Charm: "mysql"}
	st.Create()

	s.serviceApp = &ServiceApp{
		ServiceId: st.Id,
		AppId:     app.Id,
	}
	s.serviceApp.Create()

	retrievedApp := s.serviceApp.App()

	c.Assert(app.Name, Equals, retrievedApp.Name)
	c.Assert(app.Framework, Equals, retrievedApp.Framework)
}
