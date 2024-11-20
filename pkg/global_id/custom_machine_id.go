package global_id

var _ machineIDGetter = (*customMachineIDGetter)(nil)

type customMachineIDGetter struct {
	id int
}

func (c *customMachineIDGetter) GetMachineID() (int, error) {
	return c.id, nil
}
