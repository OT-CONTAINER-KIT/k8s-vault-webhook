package version

// version is a private field and should be set when compiling with --ldflags="-X github.com/innovia/secrets-consumer-env/pkg/version.version=X.Y.Z"
var version = "1.0"

// GetVersion returns the current version
func GetVersion() string {
	return version
}
