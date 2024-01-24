package testseries

// EmpiricalDistribution erwartet eine Liste mit relativen Häufigkeiten einer Messreihe.
// Die Funktion liefert eine Liste, in der für jede Zahl der
// entsprechende Wert der empirischen Verteilungsfunktion steht.
func EmpiricalDistribution(relativeFreqs []float64) []float64 {
	emp := make([]float64, len(relativeFreqs))
	/* Hinweis:
	   Verwenden Sie eine for-Schleife, um die relativen Häufigkeiten aufzusummieren.
	   In jedem Schritt speichern Sie die bisherige Summe an Stelle i in der Liste emp.
	*/
	// TODO
	return emp
}

// Distribution erwartet eine Liste mit ganzzahligen Messwerten.
// Die Funktion liefert eine Liste mit den Werten der empirischen Verteilungsfunktion.
func Distribution(values []int) []float64 {
	/* Hinweis:
	Verwenden Sie die Funktionen AbsoluteFrequencies, RelativeFrequencies und EmpiricalDistribution.
	*/
	// TODO
	return []float64{}
}
