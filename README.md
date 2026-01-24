# WillFahren - Apartment search on steroids

WillFahren is an ambitious project to build a comprehensive apartment search engine that aggregates listings from various sources. The goal is to provide a more detailed and user-friendly search experience for users.

Currently, the project is still in its early stages of development and is not yet functional. There is no glue code to bring the different components together, no external API to provide apartment listings, and no frontend to provide a user interface.

The project does include a feature to search for apartment listings along public transit connections. This feature is powered by the **wlclient** component, which provides a client to retrieve public transit information from Wiener Linien. The feature allows users to search for apartment listings within a certain distance of a public transit stop, or to filter the search results by a specific public transit line. This feature is still under development and is not yet fully functional.

The long-term goal is to expand the apartment search engine to include more data sources and cities. Once the initial prototype is functional, the plan is to add more data sources such as Immoscout24, as well as to expand the search functionality to include other cities. This will require adding more clients to the project, as well as integrating the new data sources into the existing caching and search mechanisms.

The project is built using Go as the primary programming language, with the following packages:

* **whclient**: provides a client to retrieve apartment listings from Willhaben, a popular Austrian apartment search platform.
* **wlclient**: provides a client to retrieve public transit information from Wiener Linien.
* **cache**: provides a simple caching mechanism to store and retrieve data from APIs.
* **dto**: provides data transfer objects (DTOs) to represent apartment listings and their associated data.

The project is still under active development, and new features and technical components will be added in the future.

