package service

type orderStatusGetter struct {
}

func GetOrderStatusGetter() orderStatusGetter {
	return orderStatusGetter{}
}

func (service orderStatusGetter) GetStatuses() map[int]string {
	return map[int]string{
		1: "NEW",
		2: "INVALID",
		3: "PROCESSING",
		4: "PROCESSED",
	}
}

func (service orderStatusGetter) GetRegisteredStatusID() int {
	return 1
}

func (service orderStatusGetter) GetInvalidStatusID() int {
	return 2
}

func (service orderStatusGetter) GetProcessingStatusID() int {
	return 3
}

func (service orderStatusGetter) GetProcessedStatusID() int {
	return 4
}
