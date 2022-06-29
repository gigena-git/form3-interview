package delete

/*
 * In this file is specified the struct that one can feed to, or receive from Delete. For now, there
 * is no response from said method.
 */
type AccountDataRequest struct {
	ID      string `json:"id,omitempty"`
	Version *int64 `json:"version,omitempty"`
}
