# Gophercises

![Completion status: on-going](https://img.shields.io/badge/COMPLETION%20STATUS-ON--GOING-informational?style=for-the-badge)

## Description

This repository contains my solutions for the [gophercises](https://gophercises.com/) exercises. As a challenge, all codes are written with my own ideas without viewing the tutorial videos.


### Folder details

**1. quiz**

Solution for the [1st](https://github.com/gophercises/quiz) exercise. Activities involved: making a CLI program simulating a quiz, which reads user inputs and counts down a timer. 

To run the program:

- Open a terminal session and go to the quiz directory.<br/><br/>
![quiz1](readme_images/quiz1.png)

- Run the main.go file in the main folder to start the program. (By default, the program will take the quiz question in the problem.csv file sequentially and have a duration of 30 seconds. You can use the s flag to shuffle the quiz questions and the t flag to set the duration)<br/><br/>
![quiz2](readme_images/quiz2.png)

- Answer the questions by passing numbers in the terminal.<br/><br/>
![quiz3](readme_images/quiz3.png)

- Here's the final output of the program. (In this example, the program exited due to the timer timeout)<br/><br/>
![quiz4](readme_images/quiz4.png)


**2. urlshort**

Solution for the [2nd](https://github.com/gophercises/urlshort) exercise. Activities involved: making a simple web server that maps its paths into other URLs. The mapped paths can be specified in a JSON/YAML file or database. 

To run the program:

- Open a terminal session and go to the urlshort directory.<br/><br/>
![urlshort1](readme_images/urlshort1.png)

- Run the main.go file in the main folder to start the web server. (By default, the program will use the map on the pathyaml.yaml file to map its paths. You can use the f flag to specify the file or the d flag to specify a database address and use that database instead)<br/><br/>
![urlshort2](readme_images/urlshort2.png)

- Visit the web server by using a browser. (For example, we will visit the /urlshort path, and the program will redirect us to https://github.com/gophercises/urlshort as specified in the path map file)<br/><br/>
![urlshort3](readme_images/urlshort3.png)<br/><br/>
![urlshort4](readme_images/urlshort4.png)


**3. cyoa**

Solution for the [3rd](https://github.com/gophercises/cyoa) exercise. Activities involved: creating a web app simulating a CYOA (Choose Your Own Adventure) experience that makes a story based on options the user chooses.

To run the program:

- Open a terminal session and go to the cyoa directory.<br/><br/>
![cyoa1](readme_images/cyoa1.png)

- Run the main.go file in the main directory to start the web app. (By default, the program will take the gopher.json to build the story tree. You can also use the f flag to specify the path of the file)<br/><br/>
![cyoa2](readme_images/cyoa2.png)

- Visit the root path of the web app on a browser (it will be redirected to the /intro path). (In this example, we will choose some options (the ones in a darker color) to make our story)<br/><br/>
![cyoa3](readme_images/cyoa3.png)<br/><br/>
![cyoa4](readme_images/cyoa4.png)<br/><br/>
![cyoa5](readme_images/cyoa5.png)<br/><br/>
![cyoa6](readme_images/cyoa6.png)<br/><br/>
![cyoa7](readme_images/cyoa7.png)<br/><br/>


**4. link**

Solution for the [4th](https://github.com/gophercises/link) exercise. Activities involved: creating a program that parses an HTML file to collect links and texts from all of the "a" tags.

 To run the program:

- Open a terminal session and go to the link directory.<br/><br/>
![link1](readme_images/link1.png)

- Run the main.go file in the main folder with one argument specifying the HTML file that will be parsed. (In this example, we will use the demo.html file as the target HTML file)<br/><br/>
![link2](readme_images/link2.png)


**5. sitemap**

Solution for the [5th](https://github.com/gophercises/link) exercise. Activities involved: creating a program that crawls all reachable same-domain URLs from a given URL and maps the returned URLs to create a sitemap XML.

To run the program:

- Open a terminal session and go to the sitemap directory.<br/><br/>
![sitemap1](readme_images/sitemap1.png)

- Run the main.go file in the main folder with exactly two arguments, one for the target URL that will be crawled and the other one for the path of the output file. (By default, this program will crawl the target URL for a maximum depth of 2. You can customize this value by using the d flag) <br/><br/>
![sitemap2](readme_images/sitemap2.png)

- In the above example, we crawled one of the main pages of Wikipedia and stored the resulting sitemap at demo.xml. Here's what the result looks like.<br/><br/>
![sitemap3](readme_images/sitemap3.png)


**6. hr1**

Solution for the [6th](https://github.com/gophercises/hr1) exercise. Activities involved: creating solutions for Hackerrank's caesarchiper and camelcase problems and a program making use of those algorithms.

To run the program:

- Open a terminal session and go to the hr1 directory.<br/><br/>
![hr1](readme_images/hr1.png)

- Execute the main.go file in the main folder with exactly two arguments. The first argument specifies the text that will be encrypted with the Caesar-chipper encryption, and the second one is for the offset value for the encryption. (As an example, we will run the program with "testTheProgram" as the target text and number one as the offset)<br/><br/>
![hr2](readme_images/hr2.png)

- And here's the Hackerrank result of the solution.<br/><br/>
![hr3](readme_images/hr3.png)<br/><br/>
![hr4](readme_images/hr4.png)


**7. task**

Solution for the [7th](https://github.com/gophercises/task) exercise. Activities involved: creating a CLI program that can be used to manage a list of to-do tasks and store the list in a local boldDB database.

To run the program:

- Open a terminal session and go to the task directory.<br/><br/>
![task1](readme_images/task1.png)

- Before we install the CLI program, make sure we have the correct GOBIN environment variable pointing to our go bin folder. Below is the sample correct output for echoing the GOBIN variable.<br/><br/>
![task2](readme_images/task2.png)

- Install the entry of the CLI program that is beneath the inner task folder.<br/><br/>
![task3](readme_images/task3.png)

- Check if the program is installed correctly by running an empty task command. (The output of this command will also show the documentation of the command, including all its available subcommands)<br/><br/>
![task4](readme_images/task4.png)

- Run the task command followed by one of its subcommands and arguments for that subcommand. Let's say we're going to add "cleaning rooms", "studying math", and "fixing the car" to our to-do list. We can do these as follow.<br/><br/>
![task5](readme_images/task5.png)

- Let's check our active to-do list by using the list subcommand. (As you will see below, the data are stored in alphabetical order)<br/><br/>
![task6](readme_images/task6.png)

- You can mark an active task as done by using the do subcommand and list all tasks marked as done with the completed subcommand. (As an example, we will mark the "studying math" task before as done and then check the completed and active task list)<br/><br/>
![task7](readme_images/task7.png)


**8. phone**

Solution for the [8th](https://github.com/gophercises/phone) exercise. Activities involved: creating a program that stores normalized (in the same format) and no duplicate phone numbers in a PostgreSQL database.

To run the program:

- Open a terminal session and go to the phone directory.<br/><br/>
![phone1](readme_images/phone1.png)

- Before we run the program, make sure we have the database running. (In here, we will use a local database named Custom and store the data in the phone_numbers table, which is currently empty) <br/><br/>
![phone2](readme_images/phone2.png)

- Now, run the main.go file in the main folder followed by any number of arguments specifying the phone numbers will be stored. (By default, the program will use the default configuration (which includes the address of the database, etc.) to initiate the database connection. You can use the d flag to custom this configuration)<br/><br/>
![phone3](readme_images/phone3.png)

- In the above example, we stored three phone numbers, and the program printed the inputted and resulting phone numbers. And as you will see below, our previously empty table should now have been populated with the new data.<br/><br/>
![phone4](readme_images/phone4.png)


**9. deck**

Solution for the [9th](https://github.com/gophercises/deck) exercise. Activities involved: creating structs and functions to simulate items in a card game, implementing the functinal-ops coding pattern in a function, and using stringer with go-generate to work with enum-like objects.


**10. blackjacks**

Solution for the [10th](https://github.com/gophercises/blackjack) and [11th](https://github.com/gophercises/blackjack_ai) exercises. Activities involved: creating a CLI program simulating a blackjack game and making use of the interface data structure to generalize the AI behavior algorithms. 


**11. renamer**

Solution for the [12th](https://github.com/gophercises/renamer) exercise. Activities involved: Creating functions that will traverse a directory recursively and rename files in that directory (and its subdirectory) that match a pattern specified to follow a given naming pattern. 


**12. quiet_hn**

Solution for the [13th](https://github.com/gophercises/quiet_hn) exercise. Activities involved: Creating functions that concurrently retrieve data from an API and order back the returned data to follow the original ordering positions, using a wait-group and channel.


**13. recover**

Solution for the [14th](https://github.com/gophercises/recover) and [15th](https://github.com/gophercises/recover_chroma) exercises. Activities involved: Creating functions that simulate a panic-recovery event in a web server and show its stack tracks and making use of the Chroma syntax-highlighting library to show a syntax-highlighted source code on the browser. 


**14. secret**

Solution for the [17th](https://github.com/gophercises/secret) exercise. Activities involved: Making a CLI program to store and retrieve data from a file that is encrypted, which uses the stream reader and writer from the go's cipher library.


**15. transform**

Solution for the [18th](https://github.com/gophercises/transform) exercise. Activities involved: Making a web app generating images from an image uploaded by the user, specifically, the APIs that process the uploaded image, which makes use of the exec library, and the front end that fetches the APIs. 


**16. img**

Solution for the [19th](https://github.com/gophercises/image) exercise. Activities involved: Making functions to draw chart bars on a png and SVG file, using the built-in image library and the svggo library. 


**17. pdf**

Solution for the [20th](https://github.com/gophercises/pdf) exercise. Activities involved: making functions to create an invoice and certificate pdf, using the gofpdf library. 