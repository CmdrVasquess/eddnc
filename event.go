package eddnc

type Event struct {
	SchemaRef string `json:"$schemaRef"`
	ID        string `json:"$id"`
	Header    struct {
		UploaderID      string `json:"uploaderID"`
		SoftwareName    string `json:"softwareName"`
		SoftwareVersion string `json:"softwareVersion"`
	}
}
