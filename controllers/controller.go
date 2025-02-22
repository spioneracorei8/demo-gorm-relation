package controllers

import "github.com/gofiber/fiber/v2"

type Exam interface {
	FetchListQuizChoics(f *fiber.Ctx) error
	FetchQuizChoicesById(f *fiber.Ctx) error
	CreateQuizChoices(f *fiber.Ctx) error
}
