package player

type BM struct {
	SerpentsSwiftness    int
	FerociousInspiration int
}

type MM struct {
	TrueshotAura int
	LethalShots  int
	MortalShots  int
	CarefulAim   int
}

type SV struct {
	KillerInstinct int
}

type Talents struct {
	BM BM
	MM MM
	SV SV
}

const (
	LETHAL_SHOTS_MODIFIER    = 0.01
	KILLER_INSTINCT_MODIFIER = 0.01
)
