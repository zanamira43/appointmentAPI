package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
	"time"
	_ "time/tzdata"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"github.com/zanamira43/appointment-api/database"
	"github.com/zanamira43/appointment-api/repository"
)

// SendEmail sends an email
func SendEmail(to, subject, body string) error {
	// loading env
	err := godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}

	from := os.Getenv("SMTP_EMAIL")
	password := os.Getenv("SMTP_PASSWORD")

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" + body + "\r\n")

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}

func StartDailyEmailJob() {
	// Load Baghdad timezone
	loc, err := time.LoadLocation("Asia/Baghdad")
	if err != nil {
		panic("Failed to load timezone: " + err.Error())
	}

	c := cron.New(cron.WithLocation(loc))

	// Runs every day at 4 AM
	c.AddFunc("5 16 * * *", func() {
		today := time.Now().In(loc).Weekday().String() // Get today's day
		NotifyUsersByDay(today)
	})

	c.Start()
}

func NotifyUsersByDay(dayName string) {
	// Fetch records from DB where DayName == dayName
	timeTableRepo := repository.NewGormTimeTableRepository(database.DB)
	records, err := timeTableRepo.GetTimeTableForDay(dayName)
	if err != nil {
		fmt.Println("Error fetching records:", err)
		return
	}

	if len(*records) == 0 {
		log.Printf("No records found for %s", dayName)
		return
	}

	// Build the email content
	subject := fmt.Sprintf("All Scheduled Meetings for %s", dayName)
	var bodyBuilder strings.Builder

	bodyBuilder.WriteString(fmt.Sprintf("Schedule for %s:\n\n", dayName))

	for _, record := range *records {
		bodyBuilder.WriteString(fmt.Sprintf("Patient: %s\nStart: %s\nEnd: %s\n\n",
			record.PatientName,
			record.StartTime,
			record.EndTime,
		))
	}

	// loading env
	err = godotenv.Load()
	if err != nil {
		log.Printf("Unable to load .env file %v", err)
	}

	// Send the single aggregated email
	targetEmail := os.Getenv("SMTP_TARGET_EMAIL")
	body := bodyBuilder.String()

	err = SendEmail(targetEmail, subject, body)
	if err != nil {
		fmt.Println("Failed to send email:", err)
	} else {
		fmt.Println("Email sent to:", os.Getenv("SMTP_TARGET_EMAIL"))
	}
}
