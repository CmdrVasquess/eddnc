package eddn

func ExampleUpload() {
	u := Upload{vaildate: true}
	u.header.Uploader = "me"
	u.header.SwName = "test"
	u.header.SwVersion = "0.0.1"
}
