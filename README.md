# Basic-Go-WebServer
This is a single routed web server written in go, which on visiting [https://basic-go-webserver.onrender.com/] will increment the counter and display the number of times the website has been visited.
This code locks the go rountines for concurrency safety, blocking the threads spawned by multiple requests at the same time from fetching the counter value avoiding uncertainity and releases it sequentially.
Incrementing the counter, it writes the value onto a text file (which i thought would be solution for persistent storage) and this idea turn of writing it onto a text file has become a jk as however when render(deployed cloud) restarts the server the text file is lost and it fetches the text file from the github repo which contains the value of counter when the code was pushed. 
So it has become meaningless as it is the same as storing the value in the heap.
however im recording this.
