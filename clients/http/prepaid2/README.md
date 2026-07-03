# Prepaid dictionary requests

Open these `.http` files in IntelliJ IDEA to inspect prepaid topup and prepaid card dictionaries.

Use dictionary responses to calculate payment amounts before creating `TOPUP_SERVICE` or `BUY_ITEM` payments:

- `TOPUP_SERVICE`: `amount = quantity * unfixedRate`.
- `BUY_ITEM`: `amount = denomination.price`.

If the rate or price changes before execution, the payment can be declined.
