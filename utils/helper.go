package utils

import "os"

// env utk local bisa disimpan di .env
// sedangkan di online server biasanya di simpan di OSnya

// func untuk mendapatkan value env,
// jika value tidak di temukan, maka value dari fallback menjadi default valuenya
// GetEnv(keyname, defaultvalue)
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
