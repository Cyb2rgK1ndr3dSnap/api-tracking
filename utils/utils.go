package utils

import (
	"bytes"
	"context"
	"embed"
	"encoding/base64"
	"html/template"
	"net/smtp"
	"os"
	"strconv"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"google.golang.org/api/option"
)

//go:embed templates/*
var templates embed.FS

func RenderTemplate(fs embed.FS, templatePath string, data map[string]string) (string, error) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}
	var buffer bytes.Buffer
	if err := tmpl.Execute(&buffer, data); err != nil {
		return "", err
	}
	return buffer.String(), nil
}

func SendEmail(to, subject, body, username, password, smtpHost, smtpPort string) error {
	from := username
	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
		body)

	auth := smtp.PlainAuth("", username, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
	if err != nil {
		return err
	}
	return nil
}

func ExecuteSend(token int) (chan error, error) {
	toEmail := "rafael135zk@gmail.com"

	// Renderizar la plantilla con el token
	emailBody, err := RenderTemplate(templates, "templates/email.html", map[string]string{
		"Token": strconv.Itoa(token),
	})

	// Configuración de SMTP
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	username := "huniverdad@gmail.com"
	password := "luwj bhzi eolk vdrj"

	errChan := make(chan error, 1)
	// Enviar el correo
	go func() {
		err = SendEmail(toEmail, "Activación de Cuenta", emailBody, username, password, smtpHost, smtpPort)
		if err != nil {
			// Enviar el error a través del canal
			errChan <- err
		} else {
			// Enviar nil si no hay error
			errChan <- nil
		}
	}()

	return errChan, nil
}

func GetDecodedFireBaseKey() ([]byte, error) {

	fireBaseAuthKey := os.Getenv("FIREBASE_AUTH_KEY")

	decodedKey, err := base64.StdEncoding.DecodeString(fireBaseAuthKey)
	if err != nil {
		return nil, err
	}

	return decodedKey, nil
}

func SendNotification(tokens []string, ctx context.Context, data map[string]string) error {
	decodedKey, _ := GetDecodedFireBaseKey()

	opts := []option.ClientOption{option.WithCredentialsJSON(decodedKey)}

	app, err := firebase.NewApp(ctx, nil, opts...)

	if err != nil {
		return err
	}

	fcmClient, err := app.Messaging(ctx)

	if err != nil {
		return err
	}

	if len(tokens) == 1 {
		_, err = fcmClient.Send(ctx, &messaging.Message{
			/*Notification: &messaging.Notification{
				Title: title,
				Body:  body,
			},*/
			Data:  data,
			Token: tokens[0],
		})
		return err
	}

	// Si hay múltiples tokens, usar SendMulticast()
	_, err = fcmClient.SendEachForMulticast(ctx, &messaging.MulticastMessage{
		/*Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},*/
		Data:   data,
		Tokens: tokens,
	})

	return err
}
