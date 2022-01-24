select
    id,
    title,
    substr(description, 1, 9) description,
    substr(location, 1, 9) location,
    price,
    beds,
    bathrooms,
    property_type,
    garden,
    parking,
    new_home,
    current_rating,
    potential_rating
from
    properties
order by
    random()
limit
    5;