# GEOPOLITICS

Web API to store and analyse geopolitics/diplomatics relationships between countries.

Linked to a graph-oriented database ArangoDB.

## Global view
```
GET /countries?region=${region}
```
Needs:
- the region (world, europe, africa...)

Returns:
- the list of country belonging to that region
- the score edges connecting those countries

The score edges are summaries of relationship connecting 2 countries. It summarizes all the edges concerning those 2 countries with a score value and other data.

## Relationship
```
GET /relationships?country=${key}
``` 
Needs:
- the 3-alpha code of the country

Returns:
- all the countries which does have at least one relationship stored connected those and the asked country.
- all the score edges

## Relationship details
```
GET /relationships/details?countryAKey=${keyA}&countryBKey=${keyB}
```
Needs:
- the 3-alpha code of the country A
- the 3-alpha code of the country B

The order doesn't matter

Returns:
- the 2 countries asked
- the score edge connecting those 2 countries
- all the relationships edges connecting those 2 countries, direction-independent
 
