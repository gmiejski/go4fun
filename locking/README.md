### Locking benchmark

Imagine that you have to implement service instances lock mechanism.

There are 3 possible lock operations:

1) lock given N service instances (by id) for given duration
2) mark instance as removed 
3) add instance to available instances


Implement mechanism for efficient locking. Requests contains:
* Lock instances request, containing:
    * service instances ids 
    * time to lock them for
* Remove instances request, containing:
    * instances ids  
* add instances request, containing:
    * instances ids    

Rules:
* service cannot be locked twice in the same time.
* request is handled when:
    * all services have been locked for proper time (simultaneously)
    * OR lock cannot be applied because instance is removed (return error which wil be checked)

GO!

