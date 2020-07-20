package letter

import (
	"reflect"
	"testing"
)

// In the separate file frequency.go, you are given a function, Frequency(),
// to sequentially count letter frequencies in a single text.
//
// Perform this exercise on parallelism using Go concurrency features.
// Make concurrent calls to Frequency and combine results to obtain the answer.

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`

	newsMk = `До 1 ноември во Македонија ќе починат 1.629 луѓе од
КОВИД-19 ако продолжи олеснувањето на мерките за заштита сè
додека дневната бројка на мртви не достигне 8 на 1.000.000 жители,
кога повторно би се вовеле построги мерки, според проекцијата на
американскиот Институт за здравствени мерења и проценки при
Универзитетот Вашингтон од Сиетл.
Покрај ова, институтот од Сиетл предвидува уште две сценарија.
Најлошото сценарио е ако продолжи олеснувањето на мерките и никогаш
не бидат вратени, и во тој случај проекцијата е 1.805 мртви до 1 ноември.
Според најповолното сценарио, ако 95 отсто од граѓаните носат маски на
јавни места, бројот на починати до 1 ноември би изнесувал 752.
Според проекцијата за најлошото сценарио, бројот на починати би пораснал
до 32 дневно до 1 ноември, додека во најповолното сценарио, таа бројка е 3 дневно.
Проекциите на Институтот за здравствени мерења и проценки се врз основа на
бројките до 11 јули, кога имаше 380 починати во Македонија.
Заклучно со вчера, вкупно дијагностицирани со КОВИД-19 во земјава има 8.786,
од кои оздравени се 4.676, починати се 406, а активни се 3.704.`

	newsUs = `Protests in Portland, Ore., continued through early Sunday morning,
following the Oregon Department of Justice's announcement it would be suing several
federal agencies for civil rights abuses in the state. Demonstrations have taken
place in the city for weeks following the police killing of George Floyd in May.
Protesters dismantled fences in front of the Multnomah County Justice Center and
Hatfield Federal Courthouse. The U.S. Attorney's Office in Oregon had previously
said the fences were in place to "de-escalate tensions between protesters and
federal law enforcement officers" and allow repairs on the buildings to begin.
Demonstrators also set fire to the Portland Police Association building,
according to Portland police. The fire was put out later in the evening,
and police declared a riot in the area.
Tear gas and flash bangs were used on protesters and arrests were made,
according to videos and photos from the scene posted on social media.
The Oregon Department of Justice announced Saturday it would be suing several
federal agencies for civil rights abuses, and state prosecutors will potentially
pursue criminal charges against a federal officer who seriously injured a protester.
The federal lawsuit names the U.S. Department of Homeland Security, the
U.S. Marshals Service, the United States Customs and Border Protection,
and the Federal Protective Service, agencies that have had a role in stepped-up
force used against protesters since early July.
The state filed the lawsuit late Friday night.`
)

func OriginalFrequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func TestConcurrentFrequency(t *testing.T) {
	seq := OriginalFrequency(euro + dutch + us + newsMk + newsUs)
	con := ConcurrentFrequency([]string{euro, dutch, us, newsMk, newsUs})
	if !reflect.DeepEqual(con, seq) {
		t.Fatal("ConcurrentFrequency wrong result")
	}
}

func TestSequentialFrequency(t *testing.T) {
	oSeq := OriginalFrequency(euro + dutch + us)
	seq := Frequency(euro + dutch + us)
	if !reflect.DeepEqual(oSeq, seq) {
		t.Fatal("Frequency wrong result")
	}
}

func BenchmarkSequentialFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Frequency(euro + dutch + us)
	}
}

func BenchmarkConcurrentFrequency(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConcurrentFrequency([]string{euro, dutch, us})
	}
}
