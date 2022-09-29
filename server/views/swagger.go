package views

import "assignment2/server/models"

type GetAllPeopleSwagger struct {
	Response
	Payload []models.DataRes
}
