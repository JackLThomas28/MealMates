package databaseoperations

type OperationResponse struct {
	Success bool `json:"success"`
	Body []ReqItem `json:"body"`
}

type OperationError struct {
	ErrorMessage string `json:"errorMessage"`
	ErrorType string `json:"errorType"`
}

type ReqItem interface {
	GetRecipeIds() []int
}

