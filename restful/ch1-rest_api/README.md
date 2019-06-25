# Getting Started with REST API Development

A web service is a communication mechanism defined between different computer systems. Is a road between two endpoints where messages are transferred smoothly. Here, this transfer is usually one way. Two additional programmable entities can also communicate with each other through their own APIs.

## Types of Web Services
There are many types of web services which have evolved over time. 

- [SOAP](https://en.wikipedia.org/wiki/SOAP)

- [REST](https://en.wikipedia.org/wiki/Representational_state_transfer)
    - This type of web service it is very simplicity and lightweight compared to SOAP, Performance, portability, and modifiability are the main principles behind the REST design. This architecture allows different systems to communicate and send/recive data in a very simple way. Also all request calls has a relation between an HTTP verb adn the URL. The resources in the database in an application acan be mapped with an API endpoint in the REST. REST is a stateless, cacheable, and simple architecture that is not a protocol but a pattern.

    - # Characteristics of REST services
        - **Client-server based Architecture**: This architecture is most essential for the modern web to communicate over HTTP. A single client-server may look naive initially, but many hybrid architectures are evolving.
        - **Stateless**: A REST HTTP request consists of all the data needed by the server to understand and give back the response. Once a request is served, the server doesn't remember if the request has arrived after a while.
        - **Cacheable**: In order to sacale an application well, we need to cache content and deliver it as a response. If the cache is not valid, it is our responsability to bust it. REST services should be properly cached for scaling.
        - **Scripts on demand**: REST can provide. It is more common to request scripts and data from the server.
        - **Multiple layered system**: The REST API can be served from multiple servers. One server can request the other and so forth. This easy implementation is always a good stategy for keeping the web application loosely coupled.

- [UDDI](https://en.wikipedia.org/wiki/Web_Services_Discovery)

- [WSDL](https://en.wikipedia.org/wiki/Web_Services_Description_Language)

## REST verbs and Status Codes
Rest verbs specify an action to be performed on a specific resource or collection of resources. When a request is made by the client, it should send this information in the HTTP request:
- Rest verb
- Header information
- Body (optional)

REST uses the URI to decode its resource to be handled. There are quite a few REST verbs available, but this are the most used.
- GET: Fetches a record of set of resources from the server. 200 or 404
- POST: Creates a new set of resources or just one. 201 or 404, 409
- PUT: Updates or replaces the given record. 200, 204 or 404
- PATCH: Modifies the given record. 200, 204 or 404
- DELETE: Delete the given resource. 200 or 404
- OPTIONS: Fetches all available REST operations. 200

The numbers from above, are *Success* and *Failure*, reprecents the HTTP status code. Whenever a client initiates a REST operation, since REST is stateless, the client should know a way to find out whether the operation was successful or not.
