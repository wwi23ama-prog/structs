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

	// Verwenden Sie die Methode Winner des Match-Typs, um zu prüfen, ob die Mannschaft gewonnen hat.
	// Falls ja, addieren Sie 3 Punkte.
	// Falls nein, prüfen Sie, ob die Mannschaft überhaupt gespielt hat und ob das Spiel unentschieden ausging.

	// TODO

	return points
}

// GoalDiff gibt die Tordifferenz der Mannschaft mit dem übergebenen Namen zurück.
func (ml MatchList) GoalDiff(name string) int {
	goals := 0

	// Iterieren Sie über alle Matches in der MatchList und prüfen Sie,
	// ob die Mannschaft mit dem übergebenen Namen Heim- oder Auswärtsmannschaft war.
	// Addieren Sie dann die Anzahl der geschossenen Tore und subtrahieren Sie die Anzahl der kassierten Tore.

	// TODO

	return goals
}

// Teams gibt die Namen aller Mannschaften zurück, die in der MatchList vorkommen.
// Dabei sollen die Namen alphabetisch sortiert sein.
// Jeder Name soll nur einmal vorkommen.
func (ml MatchList) Teams() []string {
	teams := make([]string, 0)

	// Iterieren Sie über alle Matches in der MatchList und sammlen Sie die Namen der Mannschaften.
	// Verwenden Sie die Funktion slices.Contains, um zu prüfen, ob ein Name schon in der Liste vorkommt.

	// TODO

	slices.Sort(teams)
	return teams
}

// Compare erwartet zwei Teams und gibt zurück, ob das erste Team in der Tabelle vor dem zweiten Team steht.
// Falls t1 vor t2 steht, soll -1 zurückgegeben werden, falls t1 hinter t2 steht, soll 1 zurückgegeben werden.
// Falls t1 und t2 gleich stehen, soll 0 zurückgegeben werden.
// Die Teams sollen dabei nach den gleichen Regeln sortiert werden, wie in der Table-Methode.
func (ml MatchList) Compare(t1, t2 string) int {

	// Vergleichen Sie die Punkte der beiden Teams.
	// Wenn sie unterschiedlich sind, geben Sie das Ergebnis des Vergleichs zurück.
	// Wenn sie gleich sind, vergleichen Sie die Tordifferenz.
	// Wenn sie unterschiedlich sind, geben Sie das Ergebnis des Vergleichs zurück.
	// Wenn sie gleich sind, geben Sie 0 zurück.

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

	// Verwenden Sie die Methode Teams, um die Namen der Mannschaften zu erhalten.
	// Verwenden Sie die Funktion slices.SortFunc, um die Mannschaften zu sortieren.
	//   Dabei sollte die Compare-Methode der MatchList verwendet werden.
	// Iterieren Sie dann über die sortierten Mannschaften und hängen Sie für jede Mannschaft
	// eine Zeile an die Tabelle an, die die Position, den Namen, die Punkte und die Tordifferenz enthält.

	// TODO

	return table
}
