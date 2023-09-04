# Food Store Calculator

The food store menu calculator

# Spec

- This food store only have 7 items in menu:
  - Red set 50 THB/set.
  - Green set 40 THB/set.
  - Blue set 30 THB/set.
  - Yellow set 50 THB/set.
  - Pink set 80 THB/set.
  - Purple set 90 THB/set.
  - Orange set 120 THB/set.
- Customers can order multiple items.
- Write a function that receives these 7 items, calculate and return price.
- Conditions:
  - Customers can get 10% if customers have a member card.
  - Order doubles of Orange, Pink or Green sets will get a 5% discount.
- If you provide unit-tests; you will get an extra score.
- Example:
  - Desk#1: Customers order Red set and Green set; price from calculation is
	90.
  - Customers can use a 10% discount card, then price should be deducted by
	discount amount.
	- For Orange sets, if customers order more than 2 items per bill. customers
	will get a 5% discount.
