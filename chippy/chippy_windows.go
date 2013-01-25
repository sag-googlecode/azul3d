package chippy

func backend_Init() error {
	// There is no need to get for minimum Windows version right now, since all the API's we are
	// currently using are 2000 Professional/Server+, which is the same requirement for the version
	// function of the API (GetVersion, GetVersionEx)
	return nil
}

func backend_Destroy() {
}
