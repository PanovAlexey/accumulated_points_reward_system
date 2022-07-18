package service

// orderStatusGetter service contains all possible order statuses and implements work with them
type orderStatusGetter struct {
}

func GetOrderStatusGetter() orderStatusGetter {
	return orderStatusGetter{}
}

func (service orderStatusGetter) GetRegisteredStatus() string {
	return "NEW"
}

func (service orderStatusGetter) GetInvalidStatus() string {
	return "INVALID"
}

func (service orderStatusGetter) GetProcessingStatus() string {
	return "PROCESSING"
}

func (service orderStatusGetter) GetProcessedStatus() string {
	return "PROCESSED"
}

func (service orderStatusGetter) GetUnfinishedStatuses() []string {
	return []string{service.GetRegisteredStatus(), service.GetProcessingStatus()}
}
