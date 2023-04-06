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

- Run all the .go files in the main directory to start the web app. (By default, the program will take the gopher.json to build the story tree. You can also use the f flag to specify the path of the file)<br/><br/>
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

- Execute the main.go file in the main folder with exactly two arguments. (The first argument specifies the text that will be encrypted with the Caesar-chipper encryption, and the second one is for the offset value for the encryption. As an example, we will run the program with "testTheProgram" as the target text and number one as the offset)<br/><br/>
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

- Check if the program is installed correctly by running the task command with no additional subcommands or arguments. (The output of this command will also show the documentation of the command, including all its available subcommands)<br/><br/>
![task4](readme_images/task4.png)

- Run the task command followed by one of its subcommands and arguments for that subcommand. Let's say we're going to add "cleaning rooms", "studying math", and "fixing the car" to our to-do list. We can do these as follow.<br/><br/>
![task5](readme_images/task5.png)

- Let's check our active to-do list by using the list subcommand. (As you will see below, the data are stored in alphabetical order)<br/><br/>
![task6](readme_images/task6.png)

- You can mark an active task as done by using the do subcommand and list all tasks marked as done with the completed subcommand. (As an example, we will mark the "studying math" task and then check the completed and active task list)<br/><br/>
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

Solution for the [9th](https://github.com/gophercises/deck) exercise. Activities involved: creating a package representing a card in a card game, specifically a suit card game. There's no real program in this directory, but there is one printing the contents of a newly created deck.

To run the program:

- Open a terminal session and go to the deck directory.<br/><br/>
![deck1](readme_images/deck1.png)

- Run the main.go file in the main directory. (By default, the program will create a deck of suit cards with no joker included and sort the deck based on the card IDs ascendingly. You can customize this by using the j flag to specify the number of jokers and the s flag to point out the sorting algorithm)<br/><br/>
![deck2](readme_images/deck2.png)


**10. blackjacks**

Solution for the [10th](https://github.com/gophercises/blackjack) and [11th](https://github.com/gophercises/blackjack_ai) exercises. Activities involved: creating a CLI program that simulates the blackjack card game.

To run the program:

- Open a terminal session and go to the blackjack directory.<br/><br/>
![blackjack1](readme_images/blackjack1.png)

- Run the main.go file in the main folder with any number of arguments specifying the active players included in the game. (By default, other than the players, the program will initiate the game with one dealer and one AI. You can customize the number of AIs by using the c flag. In this example, we will create the game with two computers and one player named Farhan)<br/><br/>
![blackjack2](readme_images/blackjack2.png)

- Let's interact more with the program by playing one round of the game.<br/><br/>
![blackjack](readme_images/blackjack3.gif)


**11. renamer**

Solution for the [12th](https://github.com/gophercises/renamer) exercise. Activities involved: Creating a program that will rename files in the specified directory that match the provided pattern.

To run the program:

- Open a terminal session and go to the renamer directory.<br/><br/>
![renamer1](readme_images/renamer1.png)

- Before we run the program, let's take a look at the folder that we will use as the target directory in our program.<br/><br/>
![renamer2](readme_images/renamer2.png)

- Now, run the main.go file in the main folder with exactly three arguments and check again the previous folder. (The first argument specifies the pattern used to find the target files and parts of the file names that will be altered, the second one defines the replacement words for altering the file names, and the third one points out the path of the folder that will be scanned by the program)<br/><br/>
![renamer3](readme_images/renamer3.png)<br/><br/>
![renamer4](readme_images/renamer4.png)

- Additionally, you can also use regex patterns in the first and second arguments. (In the following example, we will use regex patterns to match sequences of digits and append the N letter at the front of the sequences)<br/><br/>
![renamer5](readme_images/renamer5.png)
![renamer6](readme_images/renamer6.png)


**12. quiet_hn**

Solution for the [13th](https://github.com/gophercises/quiet_hn) exercise. Activities involved: Creating a web application that retrieves stories from one of the Hackernews' APIs and shows several top stories based on the query of the path.

To run the program:

- Open a terminal session and go to the quiet_hn directory.<br/><br/>
![quiet_hn1](readme_images/quiet_hn1.png)

- Run the main.go file in the main folder to start running the web application. (By default, the application will create variables in the memory for caching that will be refreshed every ten seconds. You can use the r flag to change this number)<br/><br/>
![quiet_hn2](readme_images/quiet_hn2.png)

- Let's visit our running web application on a browser. (The application will return the top 30 stories from the API if the user visits its root path with no additional query parameter. You can set the n query parameter to specify the number of stories shown)<br/><br/>
![quiet_hn3](readme_images/quiet_hn3.png)<br/><br/>
![quiet_hn4](readme_images/quiet_hn4.png)


**13. recover**

Solution for the [14th](https://github.com/gophercises/recover) and [15th](https://github.com/gophercises/recover_chroma) exercises. Activities involved: Creating a web application simulating a panic-recovery event and printing the stack trace and source codes related to that event.

To run the program:

- Open a terminal session and go to the recovery directory.<br/><br/>
![recover1](readme_images/recover1.png)

- Run the main.go file in the main folder to start the web server. (The stack trace and source code printing feature is only available for the dev mode and is the default mode. You can change this by using the d flag)<br/><br/>
![recover2](readme_images/recover2.png)

- Let's visit the root path of our web.<br/><br/>
![recover3](readme_images/recover3.png)

- Though there are not many things there, the web application is running normally. Now, let's visit the /panic path of the web application. <br/><br/>
![recover4](readme_images/recover4.png)

- The above output indicates the web application has recovered from a panic situation, and the stack traces printed are the traces for the corresponding panic situation. Click one of the links provided to show the source code associated with the stack trace. (Below, we will see the source code that is lying at the very top of the stack traces)<br/><br/>
![recover5](readme_images/recover5.png)


**14. secret**

Solution for the [17th](https://github.com/gophercises/secret) exercise. Activities involved: Making a CLI program that stores and retrieves key-value pair data from encrypted storage that lives in a file.

To run the program:

- Open a terminal session and go to the secret directory.<br/><br/>
![secret1](readme_images/secret1.png)

- Install the entry of the program beneath the inner secret folder. (Make sure you have the GOBIN environment variable set correctly. Use the echo command to check the variable)<br/><br/>
![secret2](readme_images/secret2.png)

- Run the secret command with no additional subcommands or arguments to test if the program is installed correctly. You should see the documentation of the program as follow.<br/><br/>
![secret3](readme_images/secret3.png)

- Let's populate some data into our storage by using the set subcommand. (By default, all of the commands executed will use the default encoding key and storage file path. Use the k flag to change the encoding key and the f flag to point out the path of the file that will be used. In this example, we will store the following three "key:value" pair data into our encrypted storage, namely "fb-password:123456", "ig-password:aabbccdd", and "gmail-password:uncrackable_pw")<br/><br/>
![secret4](readme_images/secret4.png)

- Use the list subcommand to show all key-value pairs stored in our storage and the get subcommand for an individual value corresponding to the specified key. <br/><br/>
![secret5](readme_images/secret5.png)
![secret6](readme_images/secret6.png)

- Below is the content of the file used for the encrypted storage, which is at the default ./test path. As you will see, it is encrypted and unreadable.<br/><br/>
![secret7](readme_images/secret7.png)


**15. transform**

Solution for the [18th](https://github.com/gophercises/transform) exercise. Activities involved: Making a web application that generates primitive-shape brushed images from an image uploaded by the user.

To run the program:

- Open a terminal session and go to the transform directory.<br/><br/>
![transform1](readme_images/transform1.png)

- Run the .go files beneath the main folder to run the web application.<br/><br/>
![transform2](readme_images/transform2.png)

- Let's visit the root path of our application on a browser to see our web application.<br/><br/>
![transform3](readme_images/transform3.png)

- To generate the images, select an image in your local workspace that will be used as the base, and then click the submit button to upload the image to the backend.<br/><br/>
![transform4](readme_images/transform4.png)<br/><br/>
![transform5](readme_images/transform5.png)

- After several seconds, the backend will return us with eight generated images with different modes. Select one of the modes you like and click the submit button. (In this example, we will select the triangle mode to proceed with this step)<br/><br/>
![transform6](readme_images/transform6.png)<br/><br/>
![transform7](readme_images/transform7.png)

- After a while, the final results of the generated images are returned. The generator used different numbers to draw the shape we have chosen. You can download the results of the images by clicking the download buttons.<br/><br/>
![transform8](readme_images/transform8.png)


**16. img**

Solution for the [19th](https://github.com/gophercises/image) exercise. Activities involved: Making a program that draws a chart on a PNG or SVG file.

To run the program:

- Open a terminal session and go to the img directory.<br/><br/>
![img1](readme_images/img1.png)

- Run the main.go file in the main folder with exactly two arguments. The first argument specifies the path of the CSV file containing the chart data and the second one defines the path for the output file. (In this example, We will draw the chart on the two of the possible output formats, which are PNG and SVG. We will use the data from the demoData.csv file and defines the demoPNG.png and demoSVG.svg as the output paths)<br/><br/>
![img2](readme_images/img2.png)

- Now, let's take a look at our created files on a browser. (As you will see below, the SVG version is made more sophisticated than the PNG one. This is due to the robustness of the svggo library that supports the limited time, whereas the PNG version uses the basic built-in image library with finite functionalities)
![img3](readme_images/img3.png)<br/><br/>
![img4](readme_images/img4.png)

**17. pdf**

Solution for the [20th](https://github.com/gophercises/pdf) exercise. Activities involved: making functions to create an invoice and certificate pdf, using the gofpdf library. 