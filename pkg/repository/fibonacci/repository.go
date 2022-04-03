package fibonacci

import "github.com/Kuwerin/fibonacci/pkg/domain"

type Repository interface {
	Save(domain.Fibonacci) error
	Find(key uint64) (domain.Fibonacci, error)
}
