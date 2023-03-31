# Gophercises

![Completion status: on-going](https://img.shields.io/badge/COMPLETION%20STATUS-ON--GOING-informational?style=for-the-badge)

## Description

This repository contains my solutions for the [gophercises](https://gophercises.com/) exercises. As a challenge, all codes are written with my own ideas without viewing the tutorial videos.


### Folder details

**1. quiz**

Solution for the 1st exercise. Activities involved: making several functions to read CSV files & flags, listen to user inputs, and simulate a timer. [Learn more](https://github.com/gophercises/quiz)


**2. urlshort**

Solution for the 2nd exercise. Activities involved: making a simple HTTP listener & its handlers, which can also establish a Postgres database connection, and creating functions to read & parse JSON/YAML files. [Learn more](https://github.com/gophercises/urlshort)


**3. cyoa**

Solution for the 3rd exercise. Activities involved: creating go template files & the functions to parse them and an HTTP handler that renders the parsed go templates. [Learn more](https://github.com/gophercises/cyoa)


**4. link**

Solution for the 4th exercise. Activities involved: creating functions to parse an HTML file and collect the href values & inner texts of all "a" tags within the specified HTML file and starting to write test functions (which will keep doing in the upcoming exercise). [Learn more](https://github.com/gophercises/link)


**5. sitemap**

Solution for the 5th exercise. Activities involved: creating functions to crawl all reachable same-domain URLs from a specific URL, which runs asynchronously with the help of mutex and wait-group structs, and turn them into a sitemap XML, implementing regex match & searching. [Learn more](https://github.com/gophercises/link)


**6. hr1**

Solution for the 6th exercise. Activities involved: creating solutions for the caesarchiper and camelcase hacker rank problems, which play with the string and rune data types. [Learn more](https://github.com/gophercises/hr1)


**7. task**

Solution for the 7th exercise. Activities involved: creating a CLI to-do task manager program with cobra library that stores and retrieve data in a boltDB database. [Learn more](https://github.com/gophercises/task)


**8. phone**

Solution for the 8th exercise. Activities involved: creating more robust functions to interact with a PostgreSQL database, which does queries and alters a table, and a program to normalize phone number data. [Learn more](https://github.com/gophercises/phone)


**9. deck**

Solution for the 9th exercise. Activities involved: creating structs and functions to simulate items in a card game, implementing the functinal-ops coding pattern in a function, and using stringer with go-generate to work with enum-like objects. [Learn more](https://github.com/gophercises/phone)


**10. blackjacks**

Solution for the 10th and 11th exercises. Activities involved: creating a CLI program simulating a blackjack game and making use of the interface data structure to generalize the AI behavior algorithms. [Learn more about the 10th exercise](https://github.com/gophercises/blackjack) [Learn more about the 11th exercise](https://github.com/gophercises/blackjack_ai)


**11. renamer**

Solution for the 12th exercise. Activities involved: Creating functions that will traverse a directory recursively and rename files in that directory (and its subdirectory) that match a pattern specified to follow a given naming pattern. [Learn more](https://github.com/gophercises/renamer)


**12. quiet_hn**

Solution for the 13th exercise. Activities involved: Creating functions that concurrently retrieve data from an API and order back the returned data to follow the original ordering positions, using a wait-group and channel. [Learn more](https://github.com/gophercises/quiet_hn)


**13. recover**

Solution for the 14th and 15th exercises. Activities involved: Creating functions that simulate a panic-recovery event in a web server and show its stack tracks and making use of the Chroma syntax-highlighting library to show a syntax-highlighted source code on the browser. [Learn more about the 14th exercise](https://github.com/gophercises/recover) [Learn more about the 15th exercise](https://github.com/gophercises/recover_chroma)


**14. secret**

Solution for the 17th exercise. Activities involved: Making a CLI program to store and retrieve data from a file that is encrypted, which uses the stream reader and writer from the go's cipher library. [Learn more](https://github.com/gophercises/secret)


**15. transform**

Solution for the 18th exercise. Activities involved: Making a web app generating images from an image uploaded by the user, specifically, the APIs that process the uploaded image, which makes use of the exec library, and the front end that fetches the APIs. [Learn more](https://github.com/gophercises/transform)