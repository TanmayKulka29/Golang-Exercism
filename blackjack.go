package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch card {
    case "ace":
        return 11
    case "eight":
        return 8
    case "two":
        return 2
    case "ten":
        return 10
    case "nine":
        return 9
    case "three":
        return 3
    case "four":
        return 4
    case "five":
        return 5
    case "six":
        return 6
    case "seven":
        return 7
    case "jack":
        return 10
    case "queen":
        return 10
    case "king":
        return 10
    case "other":
        return 0
    default:
        return 0
        panic("Unknown card")
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    if card1=="ace" && card2=="ace" {
        return "P"
    }
    if ParseCard(card1)+ParseCard(card2)>=21 {
        if ParseCard(dealerCard)<10 {
        	return "W"
        } else {
            return "S"
        }
    }
    if ParseCard(card1)+ParseCard(card2)>=17 && ParseCard(card1)+ParseCard(card2)<=20 {
        return "S"
    }
    if ParseCard(card1)+ParseCard(card2)>=12 && ParseCard(card1)+ParseCard(card2)<=16 {
        if ParseCard(dealerCard)>=7 {
            return "H"
        } else {
            return "S"
        }
    }
    if ParseCard(card1)+ParseCard(card2)<=11 {
        return "H"
    }
	panic("Please implement the FirstTurn function")
}
