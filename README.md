# a genetic algorithm
> written in Go

I recently watched some videos on youtube on the topic "genetic algorithm" and decided to
implement the described solution for fun in Go.

It's the **Shakespeare Monkey Example**

If you are interested in this topic I recommend the following youtube Playlist:  
Jump on **The Coding Train**: https://youtu.be/9zfeTw-uFCw     
This is awesome work, I really appreciate that this man offers open-source, code and free video tutorials.
 
---
## implementation in go
* [population.go](population.go) contains the implementation of the genetic algorithm.
* [cmd/main.go](cmd/main.go) contains the main function and some utilization of the population.

Basically that's all. There's a Population interface used by the main function so there is already
a seam for extending, playing around or testing different strategies.
