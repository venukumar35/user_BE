package request

type UserCreationRequest struct {
	Name       string `json:"name" binding:"required,max=150"`
	Email      string `json:"email" binding:"required,email"`
	Mobile     string `json:"mobile" binding:"required,min=7,max=14"`
	DoorNumber string `json:"doorNumber" binding:"required"`
	Street     string `json:"street" binding:"required"`
	Pincode    string `json:"pincode" binding:"required"`
	StateId    int    `json:"stateId" binding:"required"`
}
