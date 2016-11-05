package sm

import "testing"

func TestNOOP(t *testing.T) {
	machine := New(
		[]Operation{
			// NOOP
			func(m *Machine) error {
				println("NOOP")
				return nil
			},
		})

	err := machine.FromStr("0000@@", '@', 2, 16, 16)

	if err != nil {
		panic(err)
	}

	err = machine.FromStr("000000", '@', 2, 16, 16)

	if err != nil {
		panic(err)
	}

	err = machine.Run()

	if err != nil {
		panic(err)
	}
}
