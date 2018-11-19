package main

import (
	"fmt"
)

func getSuggestedItems(userItems []int, customerIds []int, customerOrders []int) []int {
	userHash := map[int]bool{}
	customerItemLists := map[int][]int{}
	customersWithSharedItems := map[int]bool{}

	for _, item := range userItems {
		userHash[item] = true
	}

	for index, customerId := range customerIds {
		customerItemLists[customerId] = append(customerItemLists[customerId], customerOrders[index])
		if userHash[customerOrders[index]] == true {
			customersWithSharedItems[customerId] = true
		}
	}

	suggestedMap := map[int]bool{}
	for custId, _ := range customersWithSharedItems {
		for _, itemId := range customerItemLists[custId] {
			suggestedMap[itemId] = true
		}
	}

	suggestedItems := []int{}
	for itemId, _ := range suggestedMap {
		if userHash[itemId] != true {
			suggestedItems = append(suggestedItems, itemId)
		}
	}

	return suggestedItems
}

func main() {
	userItems := []int{1, 3}
	otherCustomers := []int{1, 2, 3, 2, 3, 2, 3}
	customerItems := []int{6, 1, 3, 3, 4, 5, 1}

	suggestions := getSuggestedItems(userItems, otherCustomers, customerItems)

	fmt.Println(suggestions)
}
