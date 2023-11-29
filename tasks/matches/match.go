package matches

// Match ist ein Datentyp, der das Ergebnis eines Sportspiels speichern soll.
type Match struct {
	home     string
	visitors string

	score Score
}

// NewMatch erzeugt ein neues Match-Objekt mit den übergebenen Werten.
func NewMatch(home, visitors string, score Score) Match {
	return Match{home: home, visitors: visitors, score: score}
}

// HomeName gibt den Namen der Heimmannschaft zurück.
// Falls die Heimmannschaft gewonnen hat, wird der Name in Sternchen gesetzt.
func (m Match) HomeName() string {
	// Verwenden Sie die Methode Result() des Score-Typs, um zu prüfen,
	// ob die Heimmannschaft gewonnen hat.

	// TODO
	return ""
}

// VisitorName gibt den Namen der Auswärtsmannschaft zurück.
// Falls die Auswärtsmannschaft gewonnen hat, wird der Name in Sternchen gesetzt.
func (m Match) VisitorName() string {
	// Gehen Sie analog zu HomeName vor.

	// TODO
	return ""
}

// String gibt das Match als String in der Form "FC Freiburg - *Borussia Bremen*: 1:2" zurück.
// Dabei soll immer der Name der Heimmannschaft zuerst stehen und der Name des
// Gewinners in Sternchen gesetzt werden.
func (m Match) String() string {
	// Verwenden Sie die Methoden HomeName und VisitorName, um die Namen der
	// Mannschaften zu erhalten.
	// Verwenden Sie das Feld m.score, um den Spielstand zu erhalten.

	// TODO
	return ""
}

// Winner gibt den Namen des Gewinners zurück.
// Wenn es keinen Gewinner gibt, wird "unentschieden" zurückgegeben.
func (m Match) Winner() string {
	// Verwenden Sie m.score.Result(), um zu prüfen, ob es einen Gewinner gibt.

	// TODO
	return ""
}
