select
    id,
    property_id,
    `url`
from
    property_images
order by
    random()
limit
    5;