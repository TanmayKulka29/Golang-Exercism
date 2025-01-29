package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int {
        "quarter_of_a_dozen" : 3,
        "half_of_a_dozen": 6,
        "dozen" :	12,
        "small_gross" : 120,
        "gross" : 144,
        "great_gross" : 1728,
    }
    return units
	panic("Please implement the Units() function")
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    return make(map[string]int)
	panic("Please implement the NewBill() function")
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    myUnits := Units()
    if _,val := myUnits[unit]; !val {
        return false
    } 
    if _,val := bill[item]; !val {
        bill[item] = myUnits[unit]
        return true
    } else {
        bill[item] += myUnits[unit]
    	return true
    }
	panic("Please implement the AddItem() function")
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    if _,val := bill[item]; !val {
        return false
    }
    if _,val := units[unit]; !val {
        return false
    }
    quantity := bill[item] - units[unit]
    if quantity<0 {
        return false
    } else if quantity==0 {
        delete(bill, item)
        return true
    } else {
        bill[item] = quantity
        return true
    }
	panic("Please implement the RemoveItem() function")
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    if key,val := bill[item]; !val {
        return 0, false
    } else {
        return key, true
    }
	panic("Please implement the GetItem() function")
}
