package pages

import (
	"github.com/gofiber/fiber/v2"
	"github.com/maddalax/mhtml/framework/h"
	"starter-template/pages/base"
	"starter-template/partials/patient"
)

func PatientsIndex(ctx *fiber.Ctx) *h.Page {
	return h.NewPage(base.RootPage(
		h.Div(
			h.Class("flex flex-col p-4 w-full"),
			h.Div(
				h.Div(
					h.Class("flex justify-between items-center"),
					h.P(h.Text("Manage Patients"), h.Class("text-lg font-bold")),
					patient.AddPatientButton(),
				),
				h.View(patient.List, h.ReloadParams{
					Triggers: h.CreateTriggers("load", "path-deps"),
					Children: h.Children(h.Attribute("path-deps", h.GetPartialPath(patient.Create))),
				}),
			),
		),
	))
}
