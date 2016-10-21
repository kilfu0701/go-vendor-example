package config

func Load() map[string]interface{} {
	config := map[string]interface{}{
		"Website": "This is my website",
		"host": "localhost",
		"port": "8081",
	}

	return config
}
