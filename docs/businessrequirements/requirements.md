
# Business requirements

Build and API for bike rental service.

1. Manage bikes
    1. Add bikes
    1. Update bikes
    1. Remove bikes
    1. Delete bikes
1. Manage bike rentals
    1. Check bike availability
        API user provides:
        - bike id,
        - time from,
        - time to,

        API returns information whether the bike is available or not.
    1. Rent a bike for some amount of time
        API user provides:
        - customer info,
        - bike id,
        - time from,
        - time to,
        - rental location coordinates,

        API returns:
        - On success: full reservation information with it's unique id, applied discount amount.
        - On failure: reason of failure (i.e. bike is not available at this time range).
        1. API should calculate discount for reservation. rules for discount described later in specs.
    1. Check possible discount
        API user provides:
        - customer info,
        - bike id,
        - time from,
        - time to,
        - rental location coordinates,

        API returns:
        - Value of possible discount

        It shouldn't matter if the bike is available at that time.
1. Discount rules
    1. Discounts for individual customers
        1. Discount on bike weight
            If bike weight >= 15kg, apply 1% discount on each additional kg up to 20%
        1. Discount on weather conditions
            If temperature at rental location is less than 10C, apply 5% discount.
        1. discount on number of bike incidents in rental location neighborhood
            If there are 3-4 incidents around rental location, apply 5% discount.
            If there are 5 or more incidents around rental location, apply 10% discount.
    1. Discounts for business customers
        1. Discount on reservation value
            If reservation value is >= 100€, apply 5% discount
        1. Discount on reservation time
            If reservation time is >= 24h, apply 15% discount
    1. Combining discounts
        If customer applies to more than 1 discount, choose only one with highest value.
1. Authentication/Authorization
   1. Simple authentication (TODO: specify)
   2. No special authorization.
