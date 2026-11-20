package handlers

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"

	"portfolio/internal/models"
)

// ContactPage renders the contact form.
func (h *Handler) ContactPage(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{
		"Title": "Contact",
	}, "layout")
}

type contactForm struct {
	Name    string `form:"name" json:"name"`
	Email   string `form:"email" json:"email"`
	Subject string `form:"subject" json:"subject"`
	Message string `form:"message" json:"message"`
}

// SubmitContact handles POST /contact. It always persists the message first
// so a submission is never lost even if outbound email fails, then attempts
// to notify the site owner by email.
func (h *Handler) SubmitContact(c *fiber.Ctx) error {
	var form contactForm
	if err := c.BodyParser(&form); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid form submission")
	}

	form.Name = strings.TrimSpace(form.Name)
	form.Email = strings.TrimSpace(form.Email)
	form.Message = strings.TrimSpace(form.Message)

	if form.Name == "" || form.Email == "" || form.Message == "" {
		return c.Status(fiber.StatusBadRequest).Render("contact", fiber.Map{
			"Title": "Contact",
			"Error": "Name, email, and message are all required.",
			"Form":  form,
		}, "layout")
	}

	msg := models.ContactMessage{
		Name:    form.Name,
		Email:   form.Email,
		Subject: form.Subject,
		Message: form.Message,
	}
	if err := h.DB.Create(&msg).Error; err != nil {
		log.Printf("contact: failed to save message: %v", err)
		return fiber.NewError(fiber.StatusInternalServerError, "failed to save your message, please try again")
	}

	if h.Mailer.Configured() {
		if err := h.Mailer.SendContactNotification(form.Name, form.Email, form.Subject, form.Message); err != nil {
			log.Printf("contact: email notification failed: %v", err)
		} else {
			h.DB.Model(&msg).Update("email_sent", true)
		}
	} else {
		log.Println("contact: SMTP not configured, message saved without email notification")
	}

	return c.Render("contact", fiber.Map{
		"Title":   "Contact",
		"Success": "Thanks for reaching out! I'll get back to you soon.",
	}, "layout")
}
