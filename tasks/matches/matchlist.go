package matches

import (
	"slices"
)

// MatchList ist ein Datentyp für eine Liste von Spielergebnissen.
type MatchList struct {
	matches []Match
}

// NewMatchList erzeugt eine neue MatchList mit den übergebenen Matches.
func NewMatchList(matches ...Match) MatchList {
	return MatchList{matches: matches}
}

// String gibt die MatchList als String zurück.
// Dabei sollen die einzelnen Matches jeweils in einer eigenen Zeile stehen.
// Für jedes Match soll die String-Methode des Match-Typs verwendet werden.
func (ml MatchList) String() string {
	var s string
	for _, m := range ml.matches {
		s += m.String() + "\n"
	}
	return s
}

// Add fügt ein Match zur MatchList hinzu.
func (ml *MatchList) Add(m Match) {
	ml.matches = append(ml.matches, m)
}

// Points gibt die Punkte der Mannschaft mit dem übergebenen Namen zurück.
// Dabei gibt es für einen Sieg 3 Punkte, für ein Unentschieden 1 Punkt und für eine Niederlage 0 Punkte.
func (ml MatchList) Points(name string) int {
	points := 0

	// TODO

	return points
}

// GoalDiff gibt die Tordifferenz der Mannschaft mit dem übergebenen Namen zurück.
func (ml MatchList) GoalDiff(name string) int {
	goals := 0

	// TODO

	return goals
}

// Teams gibt die Namen aller Mannschaften zurück, die in der MatchList vorkommen.
// Dabei sollen die Namen alphabetisch sortiert sein.
// Jeder Name soll nur einmal vorkommen.
func (ml MatchList) Teams() []string {
	teams := make([]string, 0)

	// TODO

	slices.Sort(teams)
	return teams
}

// Compare erwartet zwei Teams und gibt zurück, ob das erste Team in der Tabelle vor dem zweiten Team steht.
// Falls t1 vor t2 steht, soll -1 zurückgegeben werden, falls t1 hinter t2 steht, soll 1 zurückgegeben werden.
// Falls t1 und t2 gleich stehen, soll 0 zurückgegeben werden.
// Die Teams sollen dabei nach den gleichen Regeln sortiert werden, wie in der Table-Methode.
func (ml MatchList) Compare(t1, t2 string) int {

	// TODO

	return 0
}

// Table gibt die Tabelle der MatchList zurück.
// Dabei soll die Tabelle in jeder Zeile die Position, den Namen einer Mannschaft,
// die Anzahl der Punkte und die Tordifferenz enthalten.
// Die Mannschaften sollen dabei absteigend nach Punkten sortiert werden.
// Bei gleicher Punktzahl soll nach Tordifferenz sortiert werden.
// Bei gleicher Tordifferenz soll alphabetisch nach Namen sortiert werden.
func (ml MatchList) Table() []string {
	table := make([]string, 0)

	// TODO

	return table
}
