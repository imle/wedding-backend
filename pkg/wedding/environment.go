package wedding

type Environment string

const (
	EnvironmentProduction  Environment = "production"
	EnvironmentStaging     Environment = "staging"
	EnvironmentDevelopment Environment = "development"
	EnvironmentLocal       Environment = "local"
	EnvironmentTest        Environment = "test"
)

func (e Environment) Is(env ...Environment) bool {
	for _, environment := range env {
		if e == environment {
			return true
		}
	}

	return false
}
