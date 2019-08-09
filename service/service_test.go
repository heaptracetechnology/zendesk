package service

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

var (
	apiKey     = os.Getenv("GODADDY_API_KEY")
	apiSecret  = os.Getenv("GODADDY_API_SECRET")
	domainName = os.Getenv("GODADDY_DOMAIN_NAME")
)

var _ = Describe("List all tickets without environment variables ", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := Ticket{SortBy: "created_at"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/listTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTicket)
	handler.ServeHTTP(recorder, request)

	Describe("List all tickets", func() {
		Context("List tickets", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List all tickets with invalid param", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/listTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTicket)
	handler.ServeHTTP(recorder, request)

	Describe("List all tickets", func() {
		Context("List tickets", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("List all tickets with all required param ", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := Ticket{SortBy: "created_at"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/listTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(ListTicket)
	handler.ServeHTTP(recorder, request)

	Describe("List all tickets", func() {
		Context("List tickets", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete ticket without environment variables ", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := Ticket{ID: 82}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Delete tickets", func() {
		Context("delete tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete ticket with invalid param", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Delete tickets", func() {
		Context("delete tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Delete ticket with all required param ", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := Ticket{ID: 82}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/deleteTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(DeleteTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Delete tickets", func() {
		Context("delete tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create ticket with invalid environment variables", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := Ticket{RequesterID: 387586085654, Subject: "TestCase Ticket", Description: "Ticket from Test Case"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create tickets", func() {
		Context("create tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create ticket with invalid param", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create tickets", func() {
		Context("create tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create ticket with invalid requ param", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := Ticket{RequesterID: 0000, Subject: "TestCase Ticket", Description: "Ticket from Test Case"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create tickets", func() {
		Context("create tickets", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create ticket with all required param ", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := Ticket{RequesterID: 387586085654, Subject: "TestCase Ticket", Description: "Ticket from Test Case"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createTicket", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateTicket)
	handler.ServeHTTP(recorder, request)

	Describe("Create tickets", func() {
		Context("create tickets", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create User without environment variables ", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := User{Name: "TestCase User"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createUser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user", func() {
		Context("create user", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create User with invalid param ", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := User{}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createUser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user", func() {
		Context("create user", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create User with invalid param ", func() {

	os.Setenv("DOMAIN_NAME", "mockDomain")
	os.Setenv("EMAIL", "mockEmail")
	os.Setenv("API_TOKEN", "mockToken")

	zendesk := []byte(`{"status":false}`)
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createUser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user", func() {
		Context("create user", func() {
			It("Should result http.StatusBadRequest", func() {
				Expect(http.StatusBadRequest).To(Equal(recorder.Code))
			})
		})
	})
})

var _ = Describe("Create User with all required param ", func() {

	os.Setenv("DOMAIN_NAME", "heaptrace")
	os.Setenv("EMAIL", "rohits@heaptrace.com")
	os.Setenv("API_TOKEN", "U69Kxjo3FTxfRLT1qWew5JVTBPWCuCFIGubfGGTR")

	zendesk := User{Name: "TestCase User"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(zendesk)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createUser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user", func() {
		Context("create user", func() {
			It("Should result http.StatusOK", func() {
				Expect(http.StatusOK).To(Equal(recorder.Code))
			})
		})
	})
})
