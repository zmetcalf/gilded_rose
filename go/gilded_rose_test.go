package gilded_rose

import (
	"testing"
)

func TestNormalItemBeforeSellDate(t *testing.T) {
	items := []*Item{&Item{"randomstring", 5, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 9 {
		t.Error("Expected quality to be 9, got ", items[0].quality)
	}
}

func TestNormalItemWithMinQuality(t *testing.T) {
	items := []*Item{&Item{"randomstring", 5, 0}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestNormalItemOnSellDate(t *testing.T) {
	items := []*Item{&Item{"randomstring", 0, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 8 {
		t.Error("Expected quality to be 8, got ", items[0].quality)
	}
}

func TestNormalItemOnSellDateWithMinQuality(t *testing.T) {
	items := []*Item{&Item{"randomstring", 0, 0}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestNormalItemOnSellDateNearMinQuality(t *testing.T) {
	items := []*Item{&Item{"randomstring", 0, 1}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestNormalItemAfterSellDate(t *testing.T) {
	items := []*Item{&Item{"randomstring", -10, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 8 {
		t.Error("Expected quality to be 8, got ", items[0].quality)
	}
}

func TestNormalItemAfterSellDateWithMinQuality(t *testing.T) {
	items := []*Item{&Item{"randomstring", -10, 0}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestNormalItemAfterSellDateNearMinQuality(t *testing.T) {
	items := []*Item{&Item{"randomstring", -10, 1}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestAgedCheddarBeforeSellDate(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", 5, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 11 {
		t.Error("Expected quality to be 11, got ", items[0].quality)
	}
}

func TestAgedCheddarWithMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", 5, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestAgedCheddarOnSellDate(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", 0, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 12 {
		t.Error("Expected quality to be 12, got ", items[0].quality)
	}
}

func TestAgedCheddarOnSellDateWithMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", 0, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestAgedCheddarOnSellDateNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", 0, 49}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestAgedCheddarAfterSellDate(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", -10, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 12 {
		t.Error("Expected quality to be 12, got ", items[0].quality)
	}
}

func TestAgedCheddarAfterSellDateWithMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", -10, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestAgedCheddarAfterSellDateNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Aged Cheddar", -10, 49}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestHammerBeforeSellDate(t *testing.T) {
	items := []*Item{&Item{"Hammer", 5, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 10 {
		t.Error("Expected quality to be 10, got ", items[0].quality)
	}
}

func TestHammerOnSellDate(t *testing.T) {
	items := []*Item{&Item{"Hammer", 0, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 10 {
		t.Error("Expected quality to be 10, got ", items[0].quality)
	}
}

func TestHammerAfterSellDate(t *testing.T) {
	items := []*Item{&Item{"Hammer", -10, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 10 {
		t.Error("Expected quality to be 10, got ", items[0].quality)
	}
}

func TestConcertTicketsLongBeforeSellDate(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 11, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 19 {
		t.Error("Expected daysRemaining to be 10, got ", items[0].daysRemaining)
	}
	if items[0].quality != 11 {
		t.Error("Expected quality to be 11, got ", items[0].quality)
	}
}

func TestConcertTicketsLongBeforeSellDateAtMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 11, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 10 {
		t.Error("Expected daysRemaining to be 10, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateUpperBound(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 10, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 9 {
		t.Error("Expected daysRemaining to be 9, got ", items[0].daysRemaining)
	}
	if items[0].quality != 12 {
		t.Error("Expected quality to be 12, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateUpperBoundAtMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 10, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 9 {
		t.Error("Expected daysRemaining to be 9, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateUpperBoundNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 10, 49}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 9 {
		t.Error("Expected daysRemaining to be 9, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateLowerBound(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 6, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 5 {
		t.Error("Expected daysRemaining to be 5, got ", items[0].daysRemaining)
	}
	if items[0].quality != 12 {
		t.Error("Expected quality to be 12, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateLowerBoundAtMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 6, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 5 {
		t.Error("Expected daysRemaining to be 5, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsMediumCloseToSellDateLowerBoundNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 6, 49}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 5 {
		t.Error("Expected daysRemaining to be 5, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateUpperBound(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 5, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 13 {
		t.Error("Expected quality to be 13, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateUpperBoundAtMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 5, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateUpperBoundNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 5, 49}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 4 {
		t.Error("Expected daysRemaining to be 4, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateLowerBound(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 1, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 0 {
		t.Error("Expected daysRemaining to be 0, got ", items[0].daysRemaining)
	}
	if items[0].quality != 13 {
		t.Error("Expected quality to be 13, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateLowerBoundAtMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 1, 50}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 0 {
		t.Error("Expected daysRemaining to be 0, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsVeryCloseToSellDateLowerBoundNearMaxQuality(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 1, 48}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != 0 {
		t.Error("Expected daysRemaining to be 0, got ", items[0].daysRemaining)
	}
	if items[0].quality != 50 {
		t.Error("Expected quality to be 50, got ", items[0].quality)
	}
}

func TestConcertTicketsOnSellDate(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", 0, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -1 {
		t.Error("Expected daysRemaining to be -1, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}

func TestConcertTicketsAfterSellDate(t *testing.T) {
	items := []*Item{&Item{"Concert Tickets", -10, 10}}
	items[0].ProcessItemEndOfDay()
	if items[0].daysRemaining != -11 {
		t.Error("Expected daysRemaining to be -11, got ", items[0].daysRemaining)
	}
	if items[0].quality != 0 {
		t.Error("Expected quality to be 0, got ", items[0].quality)
	}
}
