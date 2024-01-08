package model

type ReferenceMap struct {
	Parts           [][]*Part
	GearMultipliers []int

	RowMax    int
	ColumnMax int
}

func NewReferenceMap(rowCount int, columnCount int) ReferenceMap {
	referenceMap := ReferenceMap{}
	referenceMap.Parts = make([][]*Part, rowCount)
	for i := range referenceMap.Parts {
		referenceMap.Parts[i] = make([]*Part, columnCount)
	}

	referenceMap.RowMax = rowCount - 1
	referenceMap.ColumnMax = columnCount - 1

	return referenceMap
}

func (r *ReferenceMap) CheckForPartsNearSymbol(row int, column int) {
	partsNearby := map[*Part]bool{}

	r.setEnginePart(row-1, column-1, &partsNearby)
	r.setEnginePart(row-1, column, &partsNearby)
	r.setEnginePart(row-1, column+1, &partsNearby)
	r.setEnginePart(row, column-1, &partsNearby)
	r.setEnginePart(row, column+1, &partsNearby)
	r.setEnginePart(row+1, column-1, &partsNearby)
	r.setEnginePart(row+1, column, &partsNearby)
	r.setEnginePart(row+1, column+1, &partsNearby)

	if len(partsNearby) == 2 {
		gearMultiplier := 1
		for part, _ := range partsNearby {
			gearMultiplier *= part.GetPartValue()
		}
		r.GearMultipliers = append(r.GearMultipliers, gearMultiplier)
	}
}

func (r *ReferenceMap) setEnginePart(row int, column int, partsNearby *map[*Part]bool) {
	isValidPosition := row >= 0 && row <= r.RowMax && column >= 0 && column <= r.ColumnMax
	if isValidPosition {
		part := r.Parts[row][column]
		if part != nil {
			part.IsEnginePart = true
			(*partsNearby)[part] = true
		}
	}

}
