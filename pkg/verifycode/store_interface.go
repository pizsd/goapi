package verifycode

type Store interface {
	Set(id string, value string) bool
	Get(id string, clear bool) string
	Verify(id string, answer string, clear bool) bool
}
