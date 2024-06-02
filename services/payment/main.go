package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shopito/pkg/types/errors"
	jsonutil "shopito/pkg/util/json"
	cfg "shopito/services/notifier/config"
	"shopito/services/payment/models"
	"time"

	"github.com/go-chi/chi"
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

const (
	ADDR = ":3010"
)

func main() {

	router := chi.NewRouter()
	router.Post("/cart", func(w http.ResponseWriter, r *http.Request) {
		jsonutil.EncodeJson(w, 200, "item added to cart")
	})
	router.Post("/payment", func(w http.ResponseWriter, r *http.Request) {
		Cart, err := jsonutil.DecodeJson[models.Cart](r)
		if err != nil {
			errors.SendErr(w, errors.ErrUnpocessableEntity.SetMessage("Wrong Credentials"))
			return
		}
		ProcessPayment(models.ReceiptData{CompanyName: "Shopito", Items: Cart.Items, Customer: Cart.Name, Date: time.Now()})
		if err := sendEmailWithAttachment(Cart.Email); err != nil {
			jsonutil.EncodeJson(w, 400, "Couldn't send email")
			return
		}
		jsonutil.EncodeJson(w, 200, "Success")
	})
	logrus.WithField("addr", ADDR).Info("server started")
	srv := http.Server{
		Addr:    ADDR,
		Handler: router,
	}

	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}

func sendEmailWithAttachment(to string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", cfg.EMAIL.USERNAME)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Your Receipt")
	m.SetBody("text/plain", "Thank you for your purchase. Please find your receipt attached.")

	m.Attach("docs/receipt.pdf")

	d := gomail.NewDialer(cfg.EMAIL.HOST, 587, cfg.EMAIL.USERNAME, cfg.EMAIL.PASSWORD)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}

func ProcessPayment(receiptData models.ReceiptData) {
	fmt.Println("Processing payment...")
	m := maroto.New(config.NewBuilder().
		WithPageNumber().
		WithLeftMargin(10).
		WithTopMargin(15).
		WithRightMargin(10).
		Build(),
	)

	// Adding the header
	m.AddRow(20,
		col.New(6).Add(
			text.New(receiptData.CompanyName,
				props.Text{
					Top:   5,
					Size:  20,
					Align: align.Center,
					Style: fontstyle.Bold,
				})))

	m.AddRow(10,
		col.New(6).Add(text.New("Invoice #: 1234567890", props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
		})),
		col.New(6).Add(text.New(fmt.Sprintf("Date: %v", receiptData.Date.Format("2006-01-02 15:04:05")), props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
			Align: align.Right,
		})),
	)

	// Adding the customer information
	m.AddRow(10,
		col.New(12).Add(text.New(fmt.Sprintf("Customer: %v", receiptData.Customer), props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
		})))

	// Adding the table header
	m.AddRow(10,
		col.New(6).Add(text.New("Item", props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
			// Align: align.Right,
		})),
		col.New(3).Add(text.New("Quantity", props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
			Align: align.Center,
		})),
		col.New(3).Add(text.New("Price", props.Text{
			Top:   3,
			Size:  12,
			Style: fontstyle.Bold,
			Align: align.Right,
		})),
	)

	// Adding the table content
	total := 0
	for _, item := range receiptData.Items {
		total += item.Price
		m.AddRow(10,
			col.New(6).Add(text.New(item.Name, props.Text{
				Top:   3,
				Size:  12,
				Align: align.Left,
			})),
			col.New(3).Add(text.New(fmt.Sprintf("%v", item.Quantity), props.Text{
				Top:   3,
				Size:  12,
				Align: align.Center,
			})),
			col.New(3).Add(text.New(fmt.Sprintf("$%v", item.Price), props.Text{
				Top:   3,
				Size:  12,
				Align: align.Right,
			})),
		)
	}

	// Adding the total
	m.AddRow(10,
		col.New(6).Add(text.New("")),
		col.New(3).Add(text.New("Total", props.Text{
			Top:   3,
			Size:  12,
			Align: align.Center,
			Style: fontstyle.Bold,
		})),
		col.New(3).Add(text.New(fmt.Sprintf("$%v", total), props.Text{
			Top:   3,
			Size:  12,
			Align: align.Right,
			Style: fontstyle.Bold,
		})),
	)

	m.AddRow(20, code.NewQrCol(10, "barcode"))
	doc, err := m.Generate()
	if err != nil {
		log.Fatal(err)
	}
	doc.Save("docs/receipt.pdf")
}
