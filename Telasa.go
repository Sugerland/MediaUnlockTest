package mediaunlocktest

import (
	"encoding/json"
	"io"
	"net/http"
)

func Telasa(c http.Client) Result {
	req, err := http.NewRequest("GET", "https://api-videopass-anon.kddi-video.com/v1/playback/system_status", nil)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	req.Header.Set("X-Device-ID", "d36f8e6b-e344-4f5e-9a55-90aeb3403799")

	resp, err := cdo(c, req)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return Result{Status: StatusNetworkErr, Err: err}
	}
	// log.Println(string(b))
	var res struct {
		Status struct {
			Type    string
			Subtype string
		}
	}
	if err := json.Unmarshal(b, &res); err != nil {
		return Result{Status: StatusErr, Err: err}
	}
	if res.Status.Subtype == "IPLocationNotAllowed" {
		return Result{Status: StatusNo}
	}
	if res.Status.Type != "" {
		return Result{Status: StatusOK}
	}
	return Result{Status: StatusUnexpected}
}
