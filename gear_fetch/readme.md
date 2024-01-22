Small set of go files to help pull character gear requirements for Lord of the Rings: Heroes of Middle Earth from lotr.gg to a spreadsheet.

Modules:
- `inputs` processes local files `farming_locations.csv` for a list of known gear and `input_characters.csv` to know which characters to scrape for gear
- `lotr_gg_service` scrapes HTML pages from `lotr.gg` and crawls the HTML to parse characters, gear list urls, and amount of gear required per level
- `gearlist` processes output from `lotr_gg_service` into an intermediate representation (`GearList`), and has utilities for outputting to a local CSV or uploading directly into a google sheet named `HoME`

Code might be a bit rough, used this as a test project while learning Go