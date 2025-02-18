# Basic-Go-WebServer
This is a single routed web-server written in go. When you visit [https://basic-go-webserver.onrender.com/] will increment the counter and display the number of times the website has been visited.

## Concurrency Safety:
To ensure safe concurrent access. This blocks multiple threads spawned by simeltaneuos requests from accessing and modifying the counter, ensuring a sequential access to the counter.

## Issue With Persistent Storage:
Initially, i tried writing the value of a counter into a file thinking the state would be not be lost. However, this approach has turned out to a joke :sweat_smile:. When the render restarts the server, the text file is lost and fetches the code from this github repo, which contains the counter value when the code was last pushed. 

Again this is same as storing it in the memory(heap).

Learnt file handling and the ues-case of sync.Mutex.
