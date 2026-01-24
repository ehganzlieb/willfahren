# Willfahren - Apartment search on steroids

WillFahren is an ambitious project to build a comprehensive apartment search engine that aggregates listings from various sources. The goal is to provide a more detailed and user-friendly search experience for users.

Currently, the project is in its early stages of development and is focused on the city of Vienna. The project is not yet functional and is missing key components such as glue code to connect the different components, an external API to provide apartment listings, and a frontend to provide a user interface.

The project does include a feature to search for apartment listings along public transit connections. This feature is powered by the **wlclient** component, which provides a client to retrieve public transit information from Wiener Linien. The feature allows users to search for apartment listings within a certain distance of a public transit stop, or to filter the search results by a specific public transit line. This feature is still under development and is not yet fully functional.

The long-term goal is to expand the apartment search engine to include more cities and data sources. Once the initial prototype is functional, the plan is to add more data sources such as Immoscout24 and expand the search functionality to other cities.

The project is built using Go as the primary programming language, with the following packages:

* **whclient**: provides a client to retrieve apartment listings from Willhaben, a popular Austrian apartment search platform.
* **wlclient**: provides a client to retrieve public transit information from Wiener Linien.
* **cache**: provides a simple caching mechanism to store and retrieve data from APIs.
* **dto**: provides data transfer objects (DTOs) to represent apartment listings and their associated data.

