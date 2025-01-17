package api

/* \
**From the Synology API Spec:**


This is a non-blocking API. You need to start to search files with the start method. Then, you should poll
requests with list method to get more information, or make a request with the stop method to cancel the
operation. Otherwise, search results are stored in a search temporary database so you need to call clean
method to delete it at the end of operation.

TODO:
	SYNO.FileStation.Search.start
	SYNO.FileStation.Search.list
	SYNO.FileStation.Search.stop
	SYNO.FileStation.Search.clean

Maybe set up a struct that implements the closer interface meant for
using with Non Blocking APIs like this. We'll see. I'll have to set up the base methods anyway
*/
