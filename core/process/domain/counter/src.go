package counter

type DependencyCounter struct {
	RawDataCounter          int64
	FormulaCounter          int64
	ChangeCounter           int64
	DependentCounters       map[string]int64
	DependentChangeCounters map[string]int64
}
