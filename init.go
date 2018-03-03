package ego

// Dependency-order initialization for variables.
// See https://golang.org/ref/spec#Package_initialization

func (s *Scope) initializeVars() error {
	for {
		var (
			anyUninitialized bool
			initializedAny   bool
		)
		for name, obj := range s.objs {
			obj, ok := obj.(uninitialized)
			if !ok {
				continue
			}
			anyUninitialized = true
			if xxx /* can initialize */ {
				val, err := s.Eval1(obj)
				if err != nil {
					return err
				}
				s.objs[name] = val
				initializedAny = true
			}
		}
		if !anyUninitialized {
			return nil
		}
		if !initializedAny {
			// xxx err
		}
	}
}
