import os

service_code = """
func (s *MasterService) UpdateDepartmentFloorPlan(ctx context.Context, id uuid.UUID, floorPlanUrl string) (repository.Department, error) {
	urlArg := pgtype.Text{String: floorPlanUrl, Valid: floorPlanUrl != ""}
	return s.queries.UpdateDepartmentFloorPlan(ctx, repository.UpdateDepartmentFloorPlanParams{
		ID:           id,
		FloorPlanUrl: urlArg,
	})
}

func (s *MasterService) UpdateRackMapCoordinates(ctx context.Context, id uuid.UUID, posX, posY float64) (repository.Rack, error) {
	var px, py pgtype.Numeric
	px.Scan(fmt.Sprintf("%.2f", posX))
	py.Scan(fmt.Sprintf("%.2f", posY))

	return s.queries.UpdateRackMapCoordinates(ctx, repository.UpdateRackMapCoordinatesParams{
		ID:      id,
		MapPosX: px,
		MapPosY: py,
	})
}
"""

handler_code = """
type UpdateDepartmentFloorPlanRequest struct {
	FloorPlanUrl string `json:"floor_plan_url"`
}

// UpdateDepartmentFloorPlan updates the floor plan for a department
// @Summary Update Department Floor Plan
// @Description Update the floor plan URL for a department (indoor mapping)
// @Param id path string true "Department ID"
// @Param request body UpdateDepartmentFloorPlanRequest true "Floor Plan URL"
// @Router /master/departments/{id}/floor-plan [put]
func (h *MasterHandler) UpdateDepartmentFloorPlan(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid department ID", err.Error())
	}

	req := new(UpdateDepartmentFloorPlanRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	dept, err := h.svc.UpdateDepartmentFloorPlan(c.Context(), id, req.FloorPlanUrl)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update department floor plan", err.Error())
	}

	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "department_floor_plan", &id, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Department floor plan updated", dept)
}

type UpdateRackCoordinatesRequest struct {
	PosX float64 `json:"pos_x"`
	PosY float64 `json:"pos_y"`
}

// UpdateRackCoordinates updates the rack coordinates
// @Summary Update Rack Coordinates
// @Description Update the indoor map coordinates (posX, posY) for a rack
// @Param id path string true "Rack ID"
// @Param request body UpdateRackCoordinatesRequest true "Coordinates"
// @Router /master/racks/{id}/coordinates [put]
func (h *MasterHandler) UpdateRackCoordinates(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}

	req := new(UpdateRackCoordinatesRequest)
	if err := c.BodyParser(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	rack, err := h.svc.UpdateRackMapCoordinates(c.Context(), id, req.PosX, req.PosY)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update rack coordinates", err.Error())
	}

	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "rack_coordinates", &id, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Rack coordinates updated", rack)
}
"""

with open("c:/lara/www/kreatif/kreatif-dms/backend/internal/service/master_service.go", "a") as f:
    f.write("\n" + service_code + "\n")

with open("c:/lara/www/kreatif/kreatif-dms/backend/internal/handler/master_handler.go", "a") as f:
    f.write("\n" + handler_code + "\n")

"""
