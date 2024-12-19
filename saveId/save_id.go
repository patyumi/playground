package SaveId

import "fmt"

type IdDatabase struct {
	Id            int
	IdDescription string
	TotalUpdates  int
}

func Save(idDescription string) bool {
	fmt.Println("Initializing save id ...")

	var selectedDescription IdDatabase

	d := []IdDatabase{
		{Id: 1, IdDescription: "cpf", TotalUpdates: 0},
		{Id: 2, IdDescription: "email", TotalUpdates: 5},
		{Id: 3, IdDescription: "telefone", TotalUpdates: 0},
	}

	for _, v := range d {
		if v.IdDescription == idDescription {
			v.TotalUpdates++
			selectedDescription = v
		}
	}

	if selectedDescription.TotalUpdates == 0 {
		return false
	}

	if selectedDescription.TotalUpdates > 5 {
		selectedDescription.TotalUpdates = 0
		return false
	}

	return true
}
