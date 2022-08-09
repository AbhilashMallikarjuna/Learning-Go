## Memory Management
* Memory allocation and de-allocation happens automatically
* _new()_
  * Allocates memory but not initialized
  * You will get a memory address
  * Zeroed storage
* _make()_
  * Allocates memory and initialized
  * You will get a memory address
  * Non-zeroed storage
* We cannot put the data initially in the zeroed storage

* Q. Difference between initialization and zeroing, and new() and make() in GoLang:
https://stackoverflow.com/a/37222760

* Garbage collection happens automatically
* Anything out of scope or becomes nil is eligible for GC
* The GOGC variable sets the initial garbage collection target percentage. A collection is triggered when the ratio of freshly allocated data to live data remaining after the previous collection reaches this percentage. The default is GOGC=100. Setting GOGC=off disables the garbage collector entirely. runtime/debug.SetGCPercent allows changing this percentage at run time.