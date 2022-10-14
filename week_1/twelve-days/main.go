package twelve

var verse = map[string]string{
	"first":    "a Partridge in a Pear Tree.",
	"twelfth":  "twelve Drummers Drumming",
	"eleventh": "eleven Pipers Piping",
	"tenth":    "ten Lords-a-Leaping",
	"ninth":    "nine Ladies Dancing",
	"eighth":   "eight Maids-a-Milking",
	"seventh":  "seven Swans-a-Swimming",
	"sixth":    "six Geese-a-Laying",
	"fifth":    "five Gold Rings",
	"fourth":   "four Calling Birds",
	"third":    "three French Hens",
	"second":   "two Turtle Doves",
}

var wording = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

func Verse(j int) string {
	accum := ""

	for i := 1; i <= j; i++ {
		if i == 1 {
			accum = verse[wording[i]]
		} else if i == 2 {
			accum = verse[wording[i]] + ", and " + accum
		} else {
			accum = verse[wording[i]] + ", " + accum
		}
	}

	return "On the " + wording[j] + " day of Christmas my true love gave to me, " + accum
}

func Song() string {
	song := ""

	for i := 1; i <= len(wording); i++ {
		song = song + Verse(i)+ "\n"
	}
	
	return song
}

func main() {
	Song()
}
